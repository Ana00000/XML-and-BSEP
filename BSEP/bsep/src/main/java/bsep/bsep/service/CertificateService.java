package bsep.bsep.service;

import java.io.FileOutputStream;
import java.io.IOException;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.PrivateKey;
import java.security.SecureRandom;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateExpiredException;
import java.security.cert.CertificateNotYetValidException;
import java.security.cert.X509Certificate;
import java.sql.Date;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.Base64;
import java.util.List;
import java.util.UUID;

import org.bouncycastle.asn1.x500.AttributeTypeAndValue;
import org.bouncycastle.asn1.x500.RDN;
import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.bouncycastle.cert.jcajce.JcaX509CertificateHolder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import bsep.bsep.certificates.CertificateGenerator;
import bsep.bsep.dto.CertificateDTO;
import bsep.bsep.dto.CertificateInfoDTO;
import bsep.bsep.model.CertificateData;
import bsep.bsep.model.CertificatePurposeType;
import bsep.bsep.model.CertificateStatus;
import bsep.bsep.model.CertificateType;
import bsep.bsep.model.Issuer;
import bsep.bsep.model.Subject;
import bsep.bsep.repository.CertificateKeyStoreRepository;
import bsep.bsep.repository.ICertificateRepository;
import bsep.bsep.service.interfaces.ICertificateService;

@Service
public class CertificateService implements ICertificateService {

	private ICertificateRepository certificateRepository;
	private CertificateKeyStoreRepository certificateKeyStoreRepository;
	private final String BEGIN_CERT = "-----BEGIN CERTIFICATE-----";
	private final String END_CERT = "-----END CERTIFICATE-----";
	private final String LINE_SEPARATOR = System.getProperty("line.separator");

	@Autowired
	public CertificateService(ICertificateRepository certificateRepository,
			CertificateKeyStoreRepository certificateKeyStoreRepository) {
		this.certificateRepository = certificateRepository;
		this.certificateKeyStoreRepository = certificateKeyStoreRepository;
	}

	@Override
	public CertificateData findOne(Long id) {
		return certificateRepository.findById(id).orElseGet(null);
	}

	@Override
	public List<CertificateDTO> findAllRevokedOrExpired() {
		List<CertificateDTO> certificatesDTO = new ArrayList<CertificateDTO>();

		for (X509Certificate certificateX509 : certificateKeyStoreRepository.getCertificates()) {
				getAllRevokedOrExpired(certificatesDTO, certificateX509);
		}
		return certificatesDTO;
	}

	@Override
	public List<CertificateDTO> findAllValid() {
		List<CertificateDTO> certificatesDTO = new ArrayList<CertificateDTO>();

		for (X509Certificate certificateX509 : certificateKeyStoreRepository.getCertificates())
			getAllValid(certificatesDTO, certificateX509);

		return certificatesDTO;
	}

	public CertificateData createCertificate(CertificateInfoDTO certificateInfoDTO) {

		Subject subject = generateSubject(certificateInfoDTO);
		KeyPair keyPairIssuer = generateKeyPair();
		X509Certificate x509certificate = null;
		Issuer issuer = null;

		if (certificateInfoDTO.getCertificateType() == CertificateType.ROOT) {
			issuer = generateIssuer(keyPairIssuer.getPrivate(), certificateInfoDTO);
			x509certificate = new CertificateGenerator().generateCertificate(subject, issuer, true, null);

		} else {
			issuer = certificateKeyStoreRepository.getIssuerBySerialNumber(certificateInfoDTO.getIssuerSerialNumber(),
					certificateInfoDTO.getIssuerAlias());
			// uraditi provjeru da li je sertifikat za issuerSerialNumber validan tj. da
			// nije istekao i da nije povucen
			if (isIssuerInvalid(certificateInfoDTO, issuer)) {
				return null;
			}
			x509certificate = new CertificateGenerator().generateCertificate(subject, issuer,
					isIntermedateCertificate(certificateInfoDTO), convertStringToDate(certificateInfoDTO.getEndDate()));
		}

		certificateKeyStoreRepository.saveKeyStore(certificateInfoDTO.getCertificateType(),
				certificateInfoDTO.getAlias(), x509certificate, keyPairIssuer.getPrivate());
		return save(x509certificate.getSerialNumber().toString(), certificateInfoDTO.getCertificateType(),
				certificateInfoDTO.getCertificatePurposeType());
	}

	private boolean isIssuerInvalid(CertificateInfoDTO certificateInfoDTO, Issuer issuer) {
		return issuer == null || findCertificateDataBySerialNumber(certificateInfoDTO.getIssuerSerialNumber())
				.getCertificateStatus() != CertificateStatus.VALID;
	}

	public void loadCertificateToFile(String serialNumber) throws Exception {

		Base64.Encoder encoder = Base64.getMimeEncoder(64, LINE_SEPARATOR.getBytes());
		if (findCertificateDataBySerialNumber(serialNumber).getCertificateStatus() != CertificateStatus.VALID) {
			throw new Exception();
		}
		byte[] bytes = certificateKeyStoreRepository.findBySerialNumber(serialNumber).getEncoded();

		String certificate = BEGIN_CERT + LINE_SEPARATOR + new String(encoder.encode(bytes)) + LINE_SEPARATOR
				+ END_CERT;

		writeBytesToFile(serialNumber + ".cer", certificate.getBytes());

	}

	private void writeBytesToFile(String fileOutput, byte[] bytes) throws IOException {
		try (FileOutputStream fos = new FileOutputStream(fileOutput)) {
			fos.write(bytes);
		}
	}

	public void revokeCertificate(String serialNumber) {
		CertificateDTO certificateDTO = setCertificateData(findCertificateDataBySerialNumber(serialNumber),
				certificateKeyStoreRepository.findBySerialNumber(serialNumber));
		List<String> certificatesSerialNumbersForRevoke = new ArrayList<String>();
		getSerialNumberOfChildrenCertificate(certificatesSerialNumbersForRevoke, certificateDTO);
		setStatusForChain(certificatesSerialNumbersForRevoke);
	}

	private void setStatusForChain(List<String> certificatesSerialNumbersForRevoke) {
		for (String serialNumberIt : certificatesSerialNumbersForRevoke) {
			setCertificateStatus(findCertificateDataBySerialNumber(serialNumberIt), CertificateStatus.REVOKED);
		}
	}

	public CertificateData findCertificateDataBySerialNumber(String serialNumber) {
		for (CertificateData certificateDataIt : certificateRepository.findAll()) {
			if (certificateDataIt.getSerialNumber().equals(serialNumber))
				return certificateDataIt;
		}
		return null;
	}

	public void getSerialNumberOfChildrenCertificate(List<String> returnValues, CertificateDTO certificateDTO) {
		returnValues.add(certificateDTO.getSerialNumber());
		if (certificateDTO.getCertificateType() == CertificateType.ENDENTITY)
			return;

		for (CertificateDTO certificateDTOIt : findAllValid()) {
			if (isValidChild(certificateDTO.getSubject(), certificateDTOIt))
				getSerialNumberOfChildrenCertificate(returnValues, certificateDTOIt);
		}
	}

	private boolean isValidChild(String issuerInfo, CertificateDTO certificateDTOIt) {
		return certificateDTOIt.getIssuer().equals(issuerInfo)
				&& certificateDTOIt.getCertificateStatus() == CertificateStatus.VALID;
	}

	private Boolean isIntermedateCertificate(CertificateInfoDTO certificateInfoDTO) {
		return certificateInfoDTO.getCertificateType() == CertificateType.INTERMEDIATE;
	}

	public CertificateData save(String serialNumber, CertificateType certificateType,
			CertificatePurposeType certificatePurposeType) {
		return certificateRepository
				.save(convertCertificateInfoDTOToData(serialNumber, certificateType, certificatePurposeType));
	}

	private Issuer generateIssuer(PrivateKey issuerKey, CertificateInfoDTO certificateInfoDTO) {
		return new Issuer(issuerKey, setBuilder(certificateInfoDTO).build());
	}

	private Subject generateSubject(CertificateInfoDTO certificateInfoDTO) {
		return new Subject(generateKeyPair().getPublic(), setBuilder(certificateInfoDTO).build(),
				generateSerialNumber(), LocalDate.now(), convertStringToLocalDate(certificateInfoDTO.getEndDate()));
	}

	private LocalDate convertStringToLocalDate(String date) {
		return LocalDate.parse(date, DateTimeFormatter.ofPattern("yyyy-MM-dd"));
	}

	private Date convertStringToDate(String date) {
		return Date.valueOf(convertStringToLocalDate(date));
	}

	private String generateSerialNumber() {
		String serialNumber = String.valueOf(LocalDateTime.now().hashCode());

		if (serialNumber.contains("-")) {
			return serialNumber.replace("-", "9");

		}

		return serialNumber;
	}

	private X500NameBuilder setBuilder(CertificateInfoDTO certificateInfoDTO) {
		X500NameBuilder builder = new X500NameBuilder(BCStyle.INSTANCE);
		builder.addRDN(BCStyle.UID, UUID.randomUUID().toString());
		builder.addRDN(BCStyle.CN, certificateInfoDTO.getCommonName());
		builder.addRDN(BCStyle.GIVENNAME, certificateInfoDTO.getGivenName());
		builder.addRDN(BCStyle.SURNAME, certificateInfoDTO.getSurname());
		builder.addRDN(BCStyle.O, certificateInfoDTO.getOrganization());
		builder.addRDN(BCStyle.OU, certificateInfoDTO.getOrganizationalUnitName());
		builder.addRDN(BCStyle.E, certificateInfoDTO.getOrganizationEmail());
		builder.addRDN(BCStyle.C, certificateInfoDTO.getCountryCode());
		builder.addRDN(BCStyle.PSEUDONYM, certificateInfoDTO.getAlias());
		return builder;
	}

	private KeyPair generateKeyPair() {
		try {
			KeyPairGenerator keyGen = KeyPairGenerator.getInstance("RSA");
			keyGen.initialize(2048, SecureRandom.getInstance("SHA1PRNG", "SUN"));
			return keyGen.generateKeyPair();
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (NoSuchProviderException e) {
			e.printStackTrace();
		}
		return null;
	}

	private void getAllValid(List<CertificateDTO> certificatesDTO, X509Certificate certificateX509) {
		for (CertificateData certificateData : certificateRepository.findAll()) {
			checkCertificateValidStatus(certificatesDTO, certificateX509, certificateData);
		}
	}
	
	private void getAllRevokedOrExpired(List<CertificateDTO> certificatesDTO, X509Certificate certificateX509) {
		for (CertificateData certificateData : certificateRepository.findAll()) {
			checkCertificateRevokedOrExpiredStatus(certificatesDTO, certificateX509, certificateData);
		}
	}

	private boolean checkCertificateValidStatus(List<CertificateDTO> certificatesDTO, X509Certificate certificateX509,
			CertificateData certificateData) {
		if (certificateX509.getSerialNumber().toString().equals(certificateData.getSerialNumber())) {
			return checkCertificateExpired(certificateX509, certificateData) ? false
					: addValidCertificateToList(certificatesDTO, certificateX509, certificateData);
		}
		return false;
	}
	
	private void checkCertificateRevokedOrExpiredStatus(List<CertificateDTO> certificatesDTO, X509Certificate certificateX509,
			CertificateData certificateData) {
		if (certificateX509.getSerialNumber().toString().equals(certificateData.getSerialNumber())) {
			addRevokedOrExpiredCertificateToList(certificatesDTO, certificateX509, certificateData);
		}
		
	}

	private boolean addValidCertificateToList(List<CertificateDTO> certificatesDTO, X509Certificate certificateX509,
			CertificateData certificateData) {
		if (certificateData.getCertificateStatus() == CertificateStatus.VALID) {
			certificatesDTO.add(setCertificateData(certificateData, certificateX509));
			return true;
		}
		return false;
	}
	
	private boolean addRevokedOrExpiredCertificateToList(List<CertificateDTO> certificatesDTO, X509Certificate certificateX509,
			CertificateData certificateData) {
		if ((certificateData.getCertificateStatus() == CertificateStatus.REVOKED) || (certificateData.getCertificateStatus() == CertificateStatus.EXPIRED) ) {
			certificatesDTO.add(setCertificateData(certificateData, certificateX509));
			return true;
		}
		return false;
	}

	private boolean checkCertificateExpired(X509Certificate certificateX509, CertificateData certificateData) {
		if (isCertificateExpired(certificateX509)) {
			setCertificateStatus(certificateData, CertificateStatus.EXPIRED);
			return true;
		}
		return false;
	}

	private void setCertificateStatus(CertificateData certificateData, CertificateStatus certificateStatus) {
		certificateData.setCertificateStatus(certificateStatus);
		certificateRepository.save(certificateData);
	}

	private boolean isCertificateExpired(X509Certificate certificateX509) {
		boolean invalidDate = false;
		try {
			// provera za validnost datuma
			certificateX509.checkValidity();
		} catch (CertificateExpiredException | CertificateNotYetValidException e) {
			invalidDate = true;
		}
		return invalidDate;
	}

	private CertificateData convertCertificateInfoDTOToData(String serialNumber, CertificateType certificateType,
			CertificatePurposeType certificatePurposeType) {
		CertificateData certificateData = new CertificateData();
		certificateData.setSerialNumber(serialNumber);
		certificateData.setCertificateType(certificateType);
		certificateData.setCertificatePurposeType(certificatePurposeType);
		System.out.println("PURPOSE: " + certificatePurposeType);
		certificateData.setCertificateStatus(CertificateStatus.VALID);

		return certificateData;
	}

	private CertificateDTO setCertificateData(CertificateData certificateData, X509Certificate certificateX509) {
		CertificateDTO certificateDTO = convertX509ToCertificateDTO(certificateX509);
		certificateDTO.setId(certificateData.getId());
		certificateDTO.setCertificateType(certificateData.getCertificateType());
		certificateDTO.setCertificateStatus(certificateData.getCertificateStatus());

		return certificateDTO;
	}

	public CertificateDTO convertX509ToCertificateDTO(X509Certificate certificateX509) {
		CertificateDTO certificateDTO = new CertificateDTO();
		certificateDTO.setSerialNumber(certificateX509.getSerialNumber().toString());

		try {
			X500Name name = new JcaX509CertificateHolder(certificateX509).getSubject();

			for (RDN rdn : name.getRDNs()) {
				// RDN parsira X500Name u niz
				setCertificateDTOFields(certificateDTO, rdn);
			}

			certificateDTO.setIssuer(certificateX509.getIssuerDN().toString());
			certificateDTO.setSubject(certificateX509.getSubjectDN().toString());

			certificateDTO.setVersion(String.valueOf(certificateX509.getVersion()));
			certificateDTO.setSignatureAlgorithmName(certificateX509.getSigAlgName());

			return certificateDTO;

		} catch (CertificateEncodingException e) {
			e.printStackTrace();

			return null;
		}

	}

	private void setCertificateDTOFields(CertificateDTO certificateDTO, RDN rdn) {

		for (AttributeTypeAndValue val : rdn.getTypesAndValues()) {

			if (val.getType().equals(BCStyle.CN)) {
				certificateDTO.setCommonName(val.getValue().toString());
			} else if (val.getType().equals(BCStyle.GIVENNAME)) {
				certificateDTO.setGivenName(val.getValue().toString());
			} else if (val.getType().equals(BCStyle.SURNAME)) {
				certificateDTO.setSurname(val.getValue().toString());
			} else if (val.getType().equals(BCStyle.O)) {
				certificateDTO.setOrganization(val.getValue().toString());
			} else if (val.getType().equals(BCStyle.OU)) {
				certificateDTO.setOrganizationalUnitName(val.getValue().toString());
			} else if (val.getType().equals(BCStyle.C)) {
				certificateDTO.setCountryCode(val.getValue().toString());
			} else if (val.getType().equals(BCStyle.E)) {
				certificateDTO.setOrganizationEmail(val.getValue().toString());
			} else if (val.getType().equals(BCStyle.PSEUDONYM)) {
				certificateDTO.setAlias(val.getValue().toString());
			}
		}
	}

}
