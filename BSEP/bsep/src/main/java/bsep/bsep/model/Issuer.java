package bsep.bsep.model;

import java.security.PrivateKey;

import org.bouncycastle.asn1.x500.X500Name;

public class Issuer extends Users {

	private static final long serialVersionUID = 1L;

	private PrivateKey privateKey;

	private X500Name x500name;

	public Issuer() {
	}

	public Issuer(PrivateKey privateKey, X500Name x500name) {
		super();
		this.privateKey = privateKey;
		this.x500name = x500name;

	}

	public PrivateKey getPrivateKey() {
		return privateKey;
	}

	public void setPrivateKey(PrivateKey privateKey) {
		this.privateKey = privateKey;
	}

	public X500Name getX500name() {
		return x500name;
	}

	public void setX500name(X500Name x500name) {
		this.x500name = x500name;
	}
}
