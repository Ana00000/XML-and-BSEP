package model;

import java.security.PrivateKey;

import javax.persistence.Column;
import javax.persistence.Entity;

import org.bouncycastle.asn1.x500.X500Name;

@Entity
public class Issuer {

	@Column(name = "privateKey", unique = true, nullable = false)
	private PrivateKey privateKey;

	@Column(name = "x500name", unique = true, nullable = false)
	private X500Name x500name;
}
