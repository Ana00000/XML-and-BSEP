package bsep.bsep.controller;

import java.io.BufferedInputStream;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.KeyStore;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.PrivateKey;
import java.security.SecureRandom;
import java.security.cert.X509Certificate;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.ArrayList;
import java.util.List;

import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.springframework.beans.factory.annotation.Autowired;
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
import bsep.bsep.model.Certificate;
import bsep.bsep.model.IntermediateCA;
import bsep.bsep.model.Issuer;
import bsep.bsep.model.Subject;
import bsep.bsep.service.CertificateService;

@Controller
@CrossOrigin(origins = "", allowedHeaders = "")
@RequestMapping(value = "/certificate")
public class CertificateController {

	private final CertificateService certificateService;

	@Autowired
	public CertificateController(CertificateService certificateService) {
		this.certificateService = certificateService;
	}

	@GetMapping(value = "/all")
	public ResponseEntity<List<CertificateDTO>> getAllCertificates() {

		return new ResponseEntity<>(getAllCertificatesDTO(), HttpStatus.OK);
	}
	
	private List<CertificateDTO> getAllCertificatesDTO() {

		List<CertificateDTO> certificatesDTO = new ArrayList<>();
		for (Certificate certificate : certificateService.findAll()) {
			certificatesDTO.add(new CertificateDTO(certificate));
		}

		return certificatesDTO;
	}

	@PostMapping(value = "/createCertificate", consumes = "application/json")
	public ResponseEntity<CertificateDTO> createCertificate(@RequestBody CertificateDTO certificateDTO) {
		IntermediateCA intermediateCA = new IntermediateCA();
		try {
			Subject subject = generateSubject(intermediateCA);
			KeyPair keyPairIssuer = generateKeyPair();
            Issuer issuer = generateIssuer(keyPairIssuer.getPrivate(),intermediateCA);

            //Generise se sertifikat za subjekta, potpisan od strane issuer-a
            CertificateGenerator cg = new CertificateGenerator();
            X509Certificate cert = cg.generateCertificate(subject, issuer);
			
            KeyStore keyStore = KeyStore.getInstance("JKS", "SUN");
            
            String password = intermediateCA.getKeyStorePassword();
            String fileName = intermediateCA.getKeyStoreName().trim();
            String alias = intermediateCA.getAlias();
            
            writingCertificateInFile(keyPairIssuer,intermediateCA,keyStore ,cert);
            
        	Certificate certificate = new Certificate();
        	certificate.setVersion(String.valueOf(cert.getVersion()));
        	certificate.setAlias(intermediateCA.getAlias());
        	certificate.setSerialNumber(cert.getSerialNumber().toString());
        	certificate.setIssuer(cert.getIssuerDN().toString());
        	certificate.setSubject(cert.getSubjectDN().toString());
        	certificate.setType(cert.getType());
        	certificate.setSignatureAlgorithmName(cert.getSigAlgName());
        	certificate.setEndDate(LocalDate.ofInstant(cert.getNotAfter().toInstant(), ZoneId.systemDefault()));
        	certificate.setKeyStoreFileName(fileName);
        	certificate.setExpired(false);
		
			certificateService.save(certificate);
			return new ResponseEntity<>(certificateDTO, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
	}

	private void writingCertificateInFile(KeyPair keyPair,IntermediateCA intermediateCA, KeyStore keyStore, java.security.cert.Certificate cert) throws FileNotFoundException {
		String password = intermediateCA.getKeyStorePassword();
		String fileName = intermediateCA.getKeyStoreName().trim();
		String alias = intermediateCA.getAlias();
		BufferedInputStream in = new BufferedInputStream(new FileInputStream(fileName+".jks"));
		/*keyStore.load(in, password.toCharArray());
		keyStore.setCertificateEntry(alias, cert);
		keyStore.setKeyEntry(alias, keyPair.getPrivate(), password.toCharArray(), new Certificate[] {cert});
		keyStore.store(new FileOutputStream(fileName+".jks"), password.toCharArray());
		*/
	}

	private Issuer generateIssuer(PrivateKey issuerKey, IntermediateCA intermediateCA) {
		X500NameBuilder builder = new X500NameBuilder(BCStyle.INSTANCE);
	    builder.addRDN(BCStyle.CN, intermediateCA.getCommonName());
	    builder.addRDN(BCStyle.O, intermediateCA.getOrganization());
	    builder.addRDN(BCStyle.OU, intermediateCA.getOrganizationalUnitName());
	    builder.addRDN(BCStyle.C, intermediateCA.getIncLocation().getCountry());
	    builder.addRDN(BCStyle.E, intermediateCA.getOrganizationEmail());
	    /*
	    //UID (USER ID) je ID korisnika
	    builder.addRDN(BCStyle.SURNAME,  "Luburic");
	    builder.addRDN(BCStyle.GIVENNAME, "Nikola");
	    builder.addRDN(BCStyle.UID, "654321");
		*/
		//Kreiraju se podaci za issuer-a, sto u ovom slucaju ukljucuje:
	    // - privatni kljuc koji ce se koristiti da potpise sertifikat koji se izdaje
	    // - podatke o vlasniku sertifikata koji izdaje nov sertifikat
		return new Issuer(issuerKey, builder.build());
	}
	
	private Subject generateSubject(IntermediateCA intermediateCA) {
		
			KeyPair keyPairSubject = generateKeyPair();
			
		
            LocalDateTime startDate = LocalDateTime.now();
            LocalDateTime endDate = startDate;
            
            if(intermediateCA.isValid())
            {
            	endDate.plusYears(1);
            }
           
			
			
			//klasa X500NameBuilder pravi X500Name objekat koji predstavlja podatke o vlasniku
			X500NameBuilder builder = new X500NameBuilder(BCStyle.INSTANCE);
			builder.addRDN(BCStyle.CN, intermediateCA.getCommonName());
			builder.addRDN(BCStyle.O, intermediateCA.getOrganization());
			builder.addRDN(BCStyle.OU, intermediateCA.getOrganizationalUnitName());
			builder.addRDN(BCStyle.C, intermediateCA.getIncLocation().getCountry());
			builder.addRDN(BCStyle.E, intermediateCA.getOrganizationEmail());
			/*
		    //UID (USER ID) je ID korisnika
		    builder.addRDN(BCStyle.SURNAME,  "Luburic");
		    builder.addRDN(BCStyle.GIVENNAME, "Nikola");
		    builder.addRDN(BCStyle.UID, "654321");
			*/
		    
		    //Kreiraju se podaci za sertifikat, sto ukljucuje:
		    // - javni kljuc koji se vezuje za sertifikat
		    // - podatke o vlasniku
		    // - serijski broj sertifikata
		    // - od kada do kada vazi sertifikat
		    return new Subject(keyPairSubject.getPublic(), builder.build(), intermediateCA.getSerialNumber(), startDate, endDate);
		
	}
	
	private KeyPair generateKeyPair() {
        try {
			KeyPairGenerator keyGen = KeyPairGenerator.getInstance("RSA"); 
			SecureRandom random = SecureRandom.getInstance("SHA1PRNG", "SUN");
			keyGen.initialize(2048, random);
			return keyGen.generateKeyPair();
        } catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (NoSuchProviderException e) {
			e.printStackTrace();
		}
        return null;
	}

}
