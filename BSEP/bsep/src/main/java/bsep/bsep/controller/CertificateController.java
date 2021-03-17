package bsep.bsep.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import bsep.bsep.dto.CertificateDTO;
import bsep.bsep.service.CertificateService;

@Controller
@CrossOrigin(origins = "", allowedHeaders = "")
@RequestMapping(value = "/certificate")
public class CertificateController {

	private CertificateService certificateService;

	@Autowired
	public CertificateController(CertificateService certificateService) {
		this.certificateService = certificateService;
	}

	@GetMapping(value = "/all")
	public ResponseEntity<List<CertificateDTO>> getAllValidCertificatesDTO() {

		return new ResponseEntity<>(certificateService.findAll(), HttpStatus.OK);
	}

	/*
	 * @PostMapping(value = "/createCertificate", consumes = "application/json")
	 * public ResponseEntity<CertificateDTO> createCertificate(@RequestBody
	 * CertificateDTO certificateDTO) { try {
	 * 
	 * Subject subject = generateSubject(intermediateCA); KeyPair keyPairIssuer =
	 * generateKeyPair(); Issuer issuer = generateIssuer(keyPairIssuer.getPrivate(),
	 * intermediateCA);
	 * 
	 * // Generise se sertifikat za subjekta, potpisan od strane issuer-a
	 * X509Certificate cert = new
	 * CertificateGenerator().generateCertificate(subject, issuer);
	 * 
	 * writingCertificateInFile(keyPairIssuer, intermediateCA,
	 * KeyStore.getInstance("JKS", "SUN"), cert); //Certificate certificate =
	 * certificateService.save(setCertificate(intermediateCA, cert,
	 * intermediateCA.getKeyStoreName().trim())); //return new ResponseEntity<>(new
	 * CertificateDTO(certificate), HttpStatus.CREATED);
	 * 
	 * return new ResponseEntity<>(new CertificateDTO(), HttpStatus.CREATED); }
	 * catch (Exception e) { return new ResponseEntity<>(null,
	 * HttpStatus.BAD_REQUEST); } }
	 * 
	 * 
	 * private Certificate setCertificate(IntermediateCA intermediateCA,
	 * X509Certificate cert, String fileName) { Certificate certificate = new
	 * Certificate();
	 * 
	 * certificate.setVersion(String.valueOf(cert.getVersion()));
	 * certificate.setAlias(intermediateCA.getAlias());
	 * certificate.setSerialNumber(cert.getSerialNumber().toString());
	 * certificate.setIssuer(cert.getIssuerDN().toString());
	 * certificate.setSubject(cert.getSubjectDN().toString());
	 * certificate.setCertificateType(cert.getType());
	 * certificate.setSignatureAlgorithmName(cert.getSigAlgName());
	 * certificate.setEndDate(LocalDate.ofInstant(cert.getNotAfter().toInstant(),
	 * ZoneId.systemDefault())); certificate.setKeyStoreFileName(fileName);
	 * certificate.setExpired(false);
	 * 
	 * return certificate; }
	 * 
	 * private void writingCertificateInFile(KeyPair keyPairIssuer, IntermediateCA
	 * intermediateCA, KeyStore keyStore, X509Certificate cert) { String password =
	 * intermediateCA.getKeyStorePassword(); String fileName =
	 * intermediateCA.getKeyStoreName().trim(); String alias =
	 * intermediateCA.getAlias(); BufferedInputStream in = null; try { in = new
	 * BufferedInputStream(new FileInputStream(fileName + ".jks")); } catch
	 * (Exception e) { System.out.println(e.getMessage()); } if (in ==null) {
	 * KeyStoreWriter ksw = new KeyStoreWriter(); char[] pass =
	 * password.toCharArray(); ksw.saveKeyStore(fileName, pass); try { in = new
	 * BufferedInputStream(new FileInputStream(fileName + ".jks")); } catch
	 * (FileNotFoundException e) { // TODO Auto-generated catch block
	 * e.printStackTrace(); return; } } try { keyStore.load(in,
	 * password.toCharArray()); keyStore.setCertificateEntry(alias, cert);
	 * keyStore.setKeyEntry(alias, keyPairIssuer.getPrivate(),
	 * password.toCharArray(), new X509Certificate[] { cert }); keyStore.store(new
	 * FileOutputStream(fileName + ".jks"), password.toCharArray()); } catch
	 * (Exception e){ e.printStackTrace(); return; } }
	 */
	/*
	 * private Issuer generateIssuer(PrivateKey issuerKey, IntermediateCA
	 * intermediateCA) { X500NameBuilder builder = new
	 * X500NameBuilder(BCStyle.INSTANCE); builder.addRDN(BCStyle.CN,
	 * intermediateCA.getCommonName()); builder.addRDN(BCStyle.O,
	 * intermediateCA.getOrganization()); builder.addRDN(BCStyle.OU,
	 * intermediateCA.getOrganizationalUnitName()); builder.addRDN(BCStyle.C,
	 * intermediateCA.getIncLocation().getCountry()); builder.addRDN(BCStyle.E,
	 * intermediateCA.getOrganizationEmail());
	 * 
	 * //UID (USER ID) je ID korisnika builder.addRDN(BCStyle.SURNAME, "Luburic");
	 * //builder.addRDN(BCStyle.GIVENNAME, "Nikola"); builder.addRDN(BCStyle.UID,
	 * //"654321");
	 * 
	 * // Kreiraju se podaci za issuer-a, sto u ovom slucaju ukljucuje: // -
	 * privatni kljuc koji ce se koristiti da potpise sertifikat koji se izdaje // -
	 * podatke o vlasniku sertifikata koji izdaje nov sertifikat return new
	 * Issuer(issuerKey, builder.build()); }
	 */
	/*
	 * private Subject generateSubject(IntermediateCA intermediateCA) {
	 * 
	 * KeyPair keyPairSubject = generateKeyPair();
	 * 
	 * LocalDateTime startDate = LocalDateTime.now(); LocalDateTime endDate =
	 * startDate;
	 * 
	 * if (intermediateCA.isValid()) { endDate.plusYears(1); }
	 * 
	 * // klasa X500NameBuilder pravi X500Name objekat koji predstavlja podatke o //
	 * vlasniku X500NameBuilder builder = new X500NameBuilder(BCStyle.INSTANCE);
	 * builder.addRDN(BCStyle.CN, intermediateCA.getCommonName());
	 * builder.addRDN(BCStyle.O, intermediateCA.getOrganization());
	 * builder.addRDN(BCStyle.OU, intermediateCA.getOrganizationalUnitName());
	 * builder.addRDN(BCStyle.C, intermediateCA.getIncLocation().getCountry());
	 * builder.addRDN(BCStyle.E, intermediateCA.getOrganizationEmail());
	 * 
	 * //UID (USER ID) je ID korisnika builder.addRDN(BCStyle.SURNAME, "Luburic");
	 * //builder.addRDN(BCStyle.GIVENNAME, "Nikola"); builder.addRDN(BCStyle.UID,
	 * //"654321");
	 * 
	 * 
	 * // Kreiraju se podaci za sertifikat, sto ukljucuje: // - javni kljuc koji se
	 * vezuje za sertifikat // - podatke o vlasniku // - serijski broj sertifikata
	 * // - od kada do kada vazi sertifikat return new
	 * Subject(keyPairSubject.getPublic(), builder.build(),
	 * intermediateCA.getSerialNumber(), startDate, endDate);
	 * 
	 * }
	 */
	/*
	 * 
	 * private KeyPair generateKeyPair() { try { KeyPairGenerator keyGen =
	 * KeyPairGenerator.getInstance("RSA"); SecureRandom random =
	 * SecureRandom.getInstance("SHA1PRNG", "SUN"); keyGen.initialize(2048, random);
	 * return keyGen.generateKeyPair(); } catch (NoSuchAlgorithmException e) {
	 * e.printStackTrace(); } catch (NoSuchProviderException e) {
	 * e.printStackTrace(); } return null; }
	 */
}
