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
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.List;

import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Repository;

@Repository
public class CertificateKeyStoreRepository {

	private Environment env;

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
	public CertificateKeyStoreRepository(Environment env) {
		try {
			Security.addProvider(new BouncyCastleProvider());
			this.env = env;
			ksRoot = KeyStore.getInstance("JKS");
			ksIntermediate = KeyStore.getInstance("JKS");
			ksEndEntity = KeyStore.getInstance("JKS");
			strPassword = env.getProperty("server.ssl.key-store-password");
			charPassword = env.getProperty("server.ssl.key-store-password").toCharArray();
			alias = env.getProperty("server.ssl.key-alias");
			createNewKeyStores();
			loadKeyStore();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	public void saveKSRoot(String alias, X509Certificate certificate, PrivateKey privateKey) {
		try {
			ksRoot.setKeyEntry(alias, privateKey, charPassword, new X509Certificate[] { certificate });
			ksRoot.store(new FileOutputStream(ksRootPath), charPassword);
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
		//createNewKeyStore(ksRoot, ksRootPath, strPassword);
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
			// poziv metode za dodavanje sertifikata u listu u zavisnosti od keystore
			// PROVERITI ZA NULL I LOAD KEYSTORE
			loadKeyStore();
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
}
