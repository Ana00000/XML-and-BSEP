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

	public X509Certificate generateCertificate(Subject subject, Issuer issuer, Boolean isCA) {
		try {

			Security.addProvider(new BouncyCastleProvider());
			// Because the class for generating certificates can not receive the private key directly, we need to make the object builder
			// This object builder contains the private key of the certificate issuer which is used for signing the certificates
			// The parameter is the algorithm used for signing certificates
			JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
			
			// We also need to state the provider (in this case it's Bouncy Castle)
			builder = builder.setProvider("BC");

			// Forming the object which will contain the private key and it will be used for signing certificates
			ContentSigner contentSigner = builder.build(issuer.getPrivateKey());

			Date endDate = Date.valueOf(subject.getEndDate());
		
			// Setting data for generating certificates
			X509v3CertificateBuilder certGen = new JcaX509v3CertificateBuilder(issuer.getX500name(),
					new BigInteger(subject.getSerialNumber().trim()), Date.valueOf(subject.getStartDate()), endDate,
					subject.getX500name(), subject.getPublicKey());

			addOfExtensions(isCA, certGen);

			// Generating certificate
			X509CertificateHolder certHolder = certGen.build(contentSigner);

			// The builder generates a certificate as an X509CertificateHolder object
			// After that it is necessary to use certConverter to convert certHolder into a certificate 
			JcaX509CertificateConverter certConverter = new JcaX509CertificateConverter();
			certConverter = certConverter.setProvider("BC");

			// Converting an object into a certificate
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
