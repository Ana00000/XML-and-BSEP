package bsep.bsep.service;

import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateExpiredException;
import java.security.cert.CertificateNotYetValidException;
import java.security.cert.X509Certificate;
import java.util.ArrayList;
import java.util.List;

import org.bouncycastle.asn1.x500.AttributeTypeAndValue;
import org.bouncycastle.asn1.x500.RDN;
import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.bouncycastle.cert.jcajce.JcaX509CertificateHolder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import bsep.bsep.dto.CertificateDTO;
import bsep.bsep.model.CertificateData;
import bsep.bsep.model.CertificateStatus;
import bsep.bsep.repository.CertificateKeyStoreRepository;
import bsep.bsep.repository.ICertificateRepository;
import bsep.bsep.service.interfaces.ICertificateService;

@Service
public class CertificateService implements ICertificateService {

	private ICertificateRepository certificateRepository;
	private CertificateKeyStoreRepository certificateKeyStoreRepository;

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
	public List<CertificateDTO> findAll() {

		List<CertificateDTO> certificatesDTO = new ArrayList<CertificateDTO>();

		for (X509Certificate certificateX509 : certificateKeyStoreRepository.getCertificates()) {
			try {
				// provera za validnost datuma
				certificateX509.checkValidity();
			} catch (CertificateExpiredException | CertificateNotYetValidException e) {
				continue;
			}

			getAllWithValidStatus(certificatesDTO, certificateX509);

		}

		return certificatesDTO;
	}

	private void getAllWithValidStatus(List<CertificateDTO> certificatesDTO, X509Certificate certificateX509) {

		for (CertificateData certificateData : certificateRepository.findAll()) {
			if ((certificateX509.getSerialNumber().toString().equals(certificateData.getSerialNumber()))
					&& (certificateData.getCertificateStatus() == CertificateStatus.VALID)) {

				// provera da li je povucen/istekao/validan
				certificatesDTO.add(setCertificateData(certificateData, certificateX509));

			}

		}
	}

	private CertificateDTO setCertificateData(CertificateData certificateData, X509Certificate certificateX509) {

		CertificateDTO certificateDTO = convertX509ToCertificateDTO(certificateX509);

		certificateDTO.setCertificateType(certificateData.getCertificateType());
		certificateDTO.setCertificateStatus(certificateData.getCertificateStatus());

		return certificateDTO;
	}

	private CertificateDTO convertX509ToCertificateDTO(X509Certificate certificateX509) {
		CertificateDTO certificateDTO = new CertificateDTO();
		certificateDTO.setSerialNumber(certificateX509.getSerialNumber().toString());

		try {
			X500Name name = new JcaX509CertificateHolder(certificateX509).getSubject();

			for (RDN rdn : name.getRDNs()) {
				// RDN parsira X500Name u niz
				setCertificateDTOFields(certificateDTO, rdn);
			}

			// proba ispisa
			String[] split1 = certificateX509.getIssuerX500Principal().getName().split(",");
			String issuerName = split1[0];
			System.out.println("ISSUER1: " + certificateX509.getIssuerDN().toString());
			System.out.println("ISSUER2: " + issuerName);
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
