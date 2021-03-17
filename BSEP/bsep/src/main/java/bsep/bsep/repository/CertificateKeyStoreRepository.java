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

	private char[] password;

	private void createNewKeyStore(KeyStore keyStore, String fileName, String password) {

		// TODO: Upotrebom klasa iz primeri/pki paketa, implementirati funkciju gde
		// korisnik unosi ime keystore datoteke i ona se kreira
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

	@Autowired
	public CertificateKeyStoreRepository(Environment env) {
		try {
			Security.addProvider(new BouncyCastleProvider());
			this.env = env;
			ksRoot = KeyStore.getInstance("JKS");
			ksIntermediate = KeyStore.getInstance("JKS");
			ksEndEntity = KeyStore.getInstance("JKS");
			password = env.getProperty("server.ssl.key-store-password").toCharArray();
			//createNewKeyStores();
			loadKeyStore();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	public void saveKSRoot(X509Certificate certificate, String alias, PrivateKey privateKey) {
		try {
			ksRoot.setKeyEntry(alias, privateKey, password, new X509Certificate[] { certificate });
			ksRoot.store(new FileOutputStream(ksRootPath), password);
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
	
	// Needs to be checked, is it useful
	private void createNewKeyStores() {
		createNewKeyStore(ksRoot, env.getProperty("server.ssl.key-store1"), "password");
		createNewKeyStore(ksIntermediate, env.getProperty("server.ssl.key-store2"), "password");
		createNewKeyStore(ksEndEntity, env.getProperty("server.ssl.key-store3"), "password");
	}

	private void loadKeyStore() throws Exception {
		ksRoot.load(new FileInputStream(ksRootPath), password);
		ksIntermediate.load(new FileInputStream(ksIntermediatePath), password);
		ksEndEntity.load(new FileInputStream(ksEndEntityPath), password);
	}

	public List<X509Certificate> getCertificates() {
		List<X509Certificate> certificatesList = new ArrayList<X509Certificate>();
		try {
			// poziv metode za dodavanje sertifikata u listu u zavisnosti od keystore
			// PROVERITI ZA NULL I LOAD KEYSTORE
			addCertificatesToList(certificatesList, ksRoot);
			addCertificatesToList(certificatesList, ksIntermediate);
			addCertificatesToList(certificatesList, ksEndEntity);
		} catch (KeyStoreException e) {
			e.printStackTrace();
		}
		return certificatesList;
	}

	private void addCertificatesToList(List<X509Certificate> certificatesList, KeyStore keyStore) throws KeyStoreException {
		Enumeration<String> aliases = keyStore.aliases();
		while (aliases.hasMoreElements()) {
			String alias = aliases.nextElement();
			if (keyStore.isKeyEntry(alias)) 
				certificatesList.add((X509Certificate) keyStore.getCertificate(alias));
		}
	}
}
