package bsep.bsep.model;

import java.security.PrivateKey;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

import org.bouncycastle.asn1.x500.X500Name;

@Entity
@Table(name="issuer")
public class Issuer {

	@Id
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "issuerIdSeqGen")
	@SequenceGenerator(name = "issuerIdSeqGen", sequenceName = "issuerIdSeq", initialValue = 1, allocationSize = 1)
	@Column(name = "id", unique = true, nullable = false)
	private Long id;
	
	@Column(name = "privateKey", unique = true, nullable = false)
	private PrivateKey privateKey;

	//@Column(name = "x500name", unique = true, nullable = false)
	//private X500Name x500name;

	public Issuer() {
		// TODO Auto-generated constructor stub
	}

	public Issuer(PrivateKey privateKey, X500Name x500name) {
		super();
		this.privateKey = privateKey;
		//this.x500name = x500name;
	}

	public PrivateKey getPrivateKey() {
		return privateKey;
	}

	public void setPrivateKey(PrivateKey privateKey) {
		this.privateKey = privateKey;
	}
/*
	public X500Name getX500name() {
		return x500name;
	}

	public void setX500name(X500Name x500name) {
		this.x500name = x500name;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((privateKey == null) ? 0 : privateKey.hashCode());
		result = prime * result + ((x500name == null) ? 0 : x500name.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Issuer other = (Issuer) obj;
		if (privateKey == null) {
			if (other.privateKey != null)
				return false;
		} else if (!privateKey.equals(other.privateKey))
			return false;
		if (x500name == null) {
			if (other.x500name != null)
				return false;
		} else if (!x500name.equals(other.x500name))
			return false;
		return true;
	}*/

}
