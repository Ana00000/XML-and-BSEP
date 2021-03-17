package bsep.bsep.controller;

import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.PrivateKey;
import java.security.SecureRandom;
import java.security.cert.X509Certificate;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;
import java.util.UUID;

import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

import bsep.bsep.certificates.CertificateGenerator;
import bsep.bsep.dto.CertificateDTO;
import bsep.bsep.dto.CertificateInfoDTO;
import bsep.bsep.model.Issuer;
import bsep.bsep.model.Subject;
import bsep.bsep.repository.CertificateKeyStoreRepository;
import bsep.bsep.service.CertificateService;

@Controller
@CrossOrigin(origins = "", allowedHeaders = "")
@RequestMapping(value = "/certificate")
public class CertificateController {

	private CertificateService certificateService;
	private CertificateKeyStoreRepository certificateKeyStoreRepository;
	private Environment env;

	@Autowired
	public CertificateController(CertificateService certificateService,
			CertificateKeyStoreRepository certificateKeyStoreRepository, Environment env) {
		this.certificateService = certificateService;
		this.certificateKeyStoreRepository = certificateKeyStoreRepository;
		this.env = env;
	}

	@GetMapping(value = "/all")
	public ResponseEntity<List<CertificateDTO>> getAllValidCertificatesDTO() {

		return new ResponseEntity<>(certificateService.findAll(), HttpStatus.OK);
	}

	@PostMapping(value = "/createCertificate", consumes = "application/json")
	public ResponseEntity<CertificateDTO> createRootCertificate(@RequestBody CertificateInfoDTO certificateInfoDTO) {
		try {
			Subject subject = generateSubject(certificateInfoDTO);
			KeyPair keyPairIssuer = generateKeyPair();
			Issuer issuer = generateIssuer(keyPairIssuer.getPrivate(), certificateInfoDTO);

			// Generise se sertifikat za subjekta, potpisan od strane issuer-a
			X509Certificate cert = new CertificateGenerator().generateCertificate(subject, issuer, true, null);
			certificateKeyStoreRepository.saveKSRoot(cert, env.getProperty("server.ssl.key-alias"),
					keyPairIssuer.getPrivate());
			certificateService.save(cert.getSerialNumber().toString());

			return new ResponseEntity<>(new CertificateDTO(), HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
	}

	private Issuer generateIssuer(PrivateKey issuerKey, CertificateInfoDTO certificateInfoDTO) {
		return new Issuer(issuerKey, setBuilder(certificateInfoDTO).build());
	}

	private Subject generateSubject(CertificateInfoDTO certificateInfoDTO) {
		return new Subject(generateKeyPair().getPublic(), setBuilder(certificateInfoDTO).build(),
				String.valueOf(LocalDateTime.now().hashCode()), LocalDate.now(),
				LocalDate.parse(certificateInfoDTO.getEndDate(), DateTimeFormatter.ofPattern("yyyy-MM-dd")));
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
}
