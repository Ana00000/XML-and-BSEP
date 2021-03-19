package bsep.bsep.certificates;

import java.math.BigInteger;
import java.security.Security;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.sql.Date;

import org.bouncycastle.asn1.x509.BasicConstraints;
import org.bouncycastle.asn1.x509.ExtendedKeyUsage;
import org.bouncycastle.asn1.x509.Extension;
import org.bouncycastle.asn1.x509.KeyPurposeId;
import org.bouncycastle.asn1.x509.KeyUsage;
import org.bouncycastle.cert.CertIOException;
import org.bouncycastle.cert.X509CertificateHolder;
import org.bouncycastle.cert.X509v3CertificateBuilder;
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter;
import org.bouncycastle.cert.jcajce.JcaX509v3CertificateBuilder;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.operator.ContentSigner;
import org.bouncycastle.operator.OperatorCreationException;
import org.bouncycastle.operator.jcajce.JcaContentSignerBuilder;

import bsep.bsep.model.Issuer;
import bsep.bsep.model.Subject;

public class CertificateGenerator {
	public CertificateGenerator() {
	}

	public X509Certificate generateCertificate(Subject subject, Issuer issuer, Boolean isCA, Date expirationDate) {
		try {

			Security.addProvider(new BouncyCastleProvider());
			// Posto klasa za generisanje sertifikata ne moze da primi direktno privatni
			// kljuc pravi se builder za objekat
			// Ovaj objekat sadrzi privatni kljuc izdavaoca sertifikata i koristiti se za
			// potpisivanje sertifikata
			// Parametar koji se prosledjuje je algoritam koji se koristi za potpisivanje
			// sertifiakta
			JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
			// Takodje se navodi koji provider se koristi, u ovom slucaju Bouncy Castle
			builder = builder.setProvider("BC");

			// Formira se objekat koji ce sadrzati privatni kljuc i koji ce se koristiti za
			// potpisivanje sertifikata
			ContentSigner contentSigner = builder.build(issuer.getPrivateKey());

			Date endDate = Date.valueOf(subject.getEndDate());
			/*
			if (expirationDate != null && endDate.after(expirationDate))
				endDate = expirationDate;
			*/
			// Postavljaju se podaci za generisanje sertifiakta
			X509v3CertificateBuilder certGen = new JcaX509v3CertificateBuilder(issuer.getX500name(),
					new BigInteger(subject.getSerialNumber().trim()), Date.valueOf(subject.getStartDate()), endDate,
					subject.getX500name(), subject.getPublicKey());

			addOfExtensions(isCA, certGen);

			// Generise se sertifikat
			X509CertificateHolder certHolder = certGen.build(contentSigner);

			// Builder generise sertifikat kao objekat klase X509CertificateHolder
			// Nakon toga je potrebno certHolder konvertovati u sertifikat, za sta se
			// koristi certConverter
			JcaX509CertificateConverter certConverter = new JcaX509CertificateConverter();
			certConverter = certConverter.setProvider("BC");

			// Konvertuje objekat u sertifikat
			return certConverter.getCertificate(certHolder);
		} catch (CertificateEncodingException e) {
			e.printStackTrace();
		} catch (IllegalArgumentException e) {
			e.printStackTrace();
		} catch (IllegalStateException e) {
			e.printStackTrace();
		} catch (OperatorCreationException e) {
			e.printStackTrace();
		} catch (CertificateException e) {
			e.printStackTrace();
		} catch (CertIOException e) {
			e.printStackTrace();
		}
		return null;
	}

	private void addOfExtensions(Boolean isCA, X509v3CertificateBuilder certGen) throws CertIOException {
		if (isCA) {
			// The digitalSignature keyUsage purpose indicating to use the subject public
			// key for verifying digital signatures
			// that have purposes other than non-repudiation, certificate signature, and CRL
			// signature.
			certGen.addExtension(Extension.keyUsage, true, new KeyUsage(KeyUsage.digitalSignature));
		} else {
			// The keyEncipherment keyUsage purpose indicating to use the subject public key
			// for key transport
			certGen.addExtension(Extension.keyUsage, true, new KeyUsage(KeyUsage.dataEncipherment));
			// The keyAgreement keyUsage purpose indicating to use the subject public key
			// for key agreement
			//certGen.addExtension(Extension.keyUsage, true, new KeyUsage(KeyUsage.keyAgreement));
			
			// TLS_WEB_CLIENT_AUTHENTICATION	
			certGen.addExtension(Extension.extendedKeyUsage, false, new ExtendedKeyUsage(KeyPurposeId.id_kp_clientAuth));
			//certGen.addExtension(Extension.extendedKeyUsage, false, new ExtendedKeyUsage(KeyPurposeId.id_kp_emailProtection));
		}
		certGen.addExtension(Extension.basicConstraints, true, new BasicConstraints(isCA));
	}
}
