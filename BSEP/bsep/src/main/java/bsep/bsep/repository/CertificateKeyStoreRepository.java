package bsep.bsep.repository;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.security.KeyStore;
import java.security.KeyStoreException;
import java.security.cert.X509Certificate;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Repository;

@Repository
public class CertificateKeyStoreRepository {
	
	@Autowired
	private Environment env;

	private KeyStore ksRoot;
	private KeyStore ksIntermediate;
	private KeyStore ksEndEntity;
	private final char[] password = env.getProperty("server.ssl.key-store-password").toCharArray();
	
	public CertificateKeyStoreRepository() {
		try {
			ksRoot = KeyStore.getInstance("JKS", "SUN");
			ksIntermediate = KeyStore.getInstance("JKS", "SUN");
			ksEndEntity = KeyStore.getInstance("JKS", "SUN");
			loadKeyStore();
		} catch (Exception e) {
			e.printStackTrace();
		} 
		
	}
	
	private void loadKeyStore() throws Exception {
		
		ksRoot.load(new FileInputStream(env.getProperty("server.ssl.key-store1")), password);
		ksIntermediate.load(new FileInputStream(env.getProperty("server.ssl.key-store2")), password);
		ksEndEntity.load(new FileInputStream(env.getProperty("server.ssl.key-store3")), password);
	}
	
	
	public List<X509Certificate> getCertificates() {
		
		List<X509Certificate> certificatesList = new ArrayList<X509Certificate>();
		 try {
			 	//poziv metode za dodavanje sertifikata u listu u zavisnosti od keystore
			 	//PROVERITI ZA NULL I LOAD KEYSTORE
			 
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

			if(keyStore.isKeyEntry(alias)) {
				certificatesList.add((X509Certificate) keyStore.getCertificate(alias));
			}
		}
	}
		
		
		
	
	

}
