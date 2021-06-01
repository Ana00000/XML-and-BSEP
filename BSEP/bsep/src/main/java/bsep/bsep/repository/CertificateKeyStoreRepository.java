package bsep.bsep.repository;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.security.KeyStore;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.Security;
import java.security.UnrecoverableKeyException;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.List;

import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.cert.jcajce.JcaX509CertificateHolder;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import bsep.bsep.model.Issuer;

@Repository
public class CertificateKeyStoreRepository {

	private KeyStore ksRoot;
	private KeyStore ksIntermediate;
	private KeyStore ksEndEntity;

	private final String ksRootPath = "selfsigned.jks";
	private final String ksIntermediatePath = "intermediate.jks";
	private final String ksEndEntityPath = "endEntity.jks";

	private String strPassword;
	private char[] charPassword;
	private String alias;

	@Autowired
	public CertificateKeyStoreRepository() {
		try {
			Security.addProvider(new BouncyCastleProvider());
			strPassword = "password";
			charPassword = "password".toCharArray();
			alias = "1";
			ksRoot = KeyStore.getInstance("JKS");
			ksIntermediate = KeyStore.getInstance("JKS");
			ksEndEntity = KeyStore.getInstance("JKS");
			//createNewKeyStores();
			loadKeyStore();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	public X509Certificate findBySerialNumber(String serialNumber) {
		for (X509Certificate x509Certificate : getCertificates()) 
			if (x509Certificate.getSerialNumber().toString().equals(serialNumber)) 
				return x509Certificate;
		return null;
	}

	public void saveKeyStore(String certificateType, String alias, X509Certificate certificate,
			PrivateKey privateKey) {
		try {
			if (certificateType.equals("ROOT")) {
				ksRoot.setKeyEntry(alias, privateKey, charPassword, new X509Certificate[] { certificate });
				ksRoot.store(new FileOutputStream(ksRootPath), charPassword);
			} else if (certificateType.equals("INTERMEDIATE")) {
				ksIntermediate.setKeyEntry(alias, privateKey, charPassword, new X509Certificate[] { certificate });
				ksIntermediate.store(new FileOutputStream(ksIntermediatePath), charPassword);
			} else {
				ksEndEntity.setKeyEntry(alias, privateKey, charPassword, new X509Certificate[] { certificate });
				ksEndEntity.store(new FileOutputStream(ksEndEntityPath), charPassword);
			}

		} catch (KeyStoreException e) {
			e.printStackTrace();
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (CertificateException e) {
			e.printStackTrace();
		} catch (FileNotFoundException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	private void createNewKeyStores() {
		createNewKeyStore(ksRoot, ksRootPath, strPassword);
		createNewKeyStore(ksIntermediate, ksIntermediatePath, strPassword);
		createNewKeyStore(ksEndEntity, ksEndEntityPath, strPassword);
	}

	private void createNewKeyStore(KeyStore keyStore, String fileName, String password) {
		try {
			if (new File(fileName).exists())
				keyStore.load(new FileInputStream(fileName), password.toCharArray());
			else
				keyStore.load(null, password.toCharArray());

			keyStore.store(new FileOutputStream(fileName), password.toCharArray());
		} catch (KeyStoreException e1) {
			e1.printStackTrace();
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (CertificateException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		}
		;
	}

	private void loadKeyStore() throws Exception {
		ksRoot.load(new FileInputStream(ksRootPath), charPassword);
		ksIntermediate.load(new FileInputStream(ksIntermediatePath), charPassword);
		ksEndEntity.load(new FileInputStream(ksEndEntityPath), charPassword);
	}

	public List<X509Certificate> getCertificates() {
		List<X509Certificate> certificatesList = new ArrayList<X509Certificate>();
		try {
			addCertificatesToList(certificatesList, ksRoot);
			addCertificatesToList(certificatesList, ksIntermediate);
			addCertificatesToList(certificatesList, ksEndEntity);
		} catch (KeyStoreException e) {
			e.printStackTrace();
		} catch (Exception e1) {
			e1.printStackTrace();
		}
		return certificatesList;
	}

	private void addCertificatesToList(List<X509Certificate> certificatesList, KeyStore keyStore)
			throws KeyStoreException {
		Enumeration<String> aliases = keyStore.aliases();
		while (aliases.hasMoreElements()) {
			String alias = aliases.nextElement();
			if (keyStore.isKeyEntry(alias)) {
				System.out.println("This is alias " + alias);
				certificatesList.add((X509Certificate) keyStore.getCertificate(alias));
			}
		}
	}

	public Issuer getIssuerBySerialNumber(String issuerSerialNumber, String issuerAlias) {
		X509Certificate x509Certificate = findBySerialNumber(issuerSerialNumber);
		if (x509Certificate == null) {

			return null;
		}

		return new Issuer(getPrivateKey(issuerAlias), getX500Name(x509Certificate));
	}

	private X500Name getX500Name(X509Certificate x509Certificate) {

		X500Name issuerName = null;
		try {
			issuerName = new JcaX509CertificateHolder(x509Certificate).getSubject();
		} catch (CertificateEncodingException e) {
			e.printStackTrace();
		}
		return issuerName;
	}

	private PrivateKey getPrivateKey(String issuerAlias) {

		try {
			// loadKeyStore();
			PrivateKey privateKey = getPrivateKeyForKeyStore(issuerAlias, ksRoot);

			if (privateKey == null) {

				privateKey = getPrivateKeyForKeyStore(issuerAlias, ksIntermediate);

				if (privateKey == null) {

					return null;
				}
			}

			return privateKey;

		} catch (Exception e) {
			e.printStackTrace();
		}
		return null;
	}

	private PrivateKey getPrivateKeyForKeyStore(String issuerAlias, KeyStore keyStore)
			throws UnrecoverableKeyException, KeyStoreException, NoSuchAlgorithmException {

		return (PrivateKey) keyStore.getKey(issuerAlias, charPassword);
	}
}
