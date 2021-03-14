package bsep.bsep.controller;

import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.PrivateKey;
import java.security.SecureRandom;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.time.LocalDateTime;
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

	@PostMapping(value = "/create", consumes = "application/json")
	public ResponseEntity<CertificateDTO> createCertificate(@RequestBody CertificateDTO certificateDTO) {
		try {
			certificateService.save(new Certificate(certificateDTO));
			return new ResponseEntity<>(certificateDTO, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
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
