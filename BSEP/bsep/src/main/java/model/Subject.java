package model;

import java.security.PublicKey;

import javax.persistence.Column;
import javax.persistence.Entity;

import org.bouncycastle.asn1.x500.X500Name;

@Entity
public class Subject {
	
	@Column(name = "publicKey", unique = true, nullable = false)
	private PublicKey publicKey;

	@Column(name = "x500name", unique = true, nullable = false)
	private X500Name x500name;

	@Column(name = "serialNumber", unique = true, nullable = false)
    private String serialNumber;
}
