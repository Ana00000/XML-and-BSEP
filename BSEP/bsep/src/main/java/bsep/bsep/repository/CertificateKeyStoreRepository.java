package bsep.bsep.repository;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.security.KeyStore;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.List;

import org.springframework.core.env.Environment;
import org.springframework.stereotype.Repository;

import bsep.bsep.keystores.KeyStoreWriter;

@Repository
public class CertificateKeyStoreRepository {
	
	private Environment env;

	private KeyStore ksRoot;
	private KeyStore ksIntermediate;
	private KeyStore ksEndEntity;
	private char[] password; 
	
	private void createNewKeyStore(KeyStore keyStore1, String filename, String password) {
		
		
		//TODO: Upotrebom klasa iz primeri/pki paketa, implementirati funkciju gde korisnik unosi ime keystore datoteke i ona se kreira
		try {
			
			String s=password;
			char[] pass = s.toCharArray();
			String fileName = filename;
        	File file = new File(fileName);
            if(file.exists())
                keyStore1.load(new FileInputStream(fileName), pass);
            else
                keyStore1.load(null, pass);
			
			FileOutputStream fos = new FileOutputStream(fileName);
			keyStore1.store(fos, pass);
		} catch (KeyStoreException e1) {
			// TODO Auto-generated catch block
			e1.printStackTrace();
		} catch (NoSuchAlgorithmException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (CertificateException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		};
		
	}
	
	public CertificateKeyStoreRepository(Environment env) {
		try {
			this.env = env;
			ksRoot = KeyStore.getInstance("JKS", "SUN");
			ksIntermediate = KeyStore.getInstance("JKS", "SUN");
			ksEndEntity = KeyStore.getInstance("JKS", "SUN");
			password = env.getProperty("server.ssl.key-store-password").toCharArray();
			//createNewKeyStore(ksRoot,env.getProperty("server.ssl.key-store1"), password.toString());
			//createNewKeyStore(ksIntermediate,env.getProperty("server.ssl.key-store2"), password.toString());
			//screateNewKeyStore(ksEndEntity,env.getProperty("server.ssl.key-store"), password.toString());
			loadKeyStore();
		} catch (Exception e) {
			e.printStackTrace();
		} 
		
	}
	
	private void loadKeyStore() throws Exception {
		System.out.println(password);
	//	ksRoot.load(new FileInputStream(env.getProperty("server.ssl.key-store1")), password);
		//ksIntermediate.load(new FileInputStream(env.getProperty("server.ssl.key-store2")), password);
		ksEndEntity.load(new FileInputStream(env.getProperty("server.ssl.key-store")), password);
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
