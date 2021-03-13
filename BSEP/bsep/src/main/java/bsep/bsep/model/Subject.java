package bsep.bsep.model;

import java.security.PublicKey;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

import org.bouncycastle.asn1.x500.X500Name;

@Entity
@Table(name="subject")
public class Subject {

	@Id
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "subjectIdSeqGen")
	@SequenceGenerator(name = "subjectIdSeqGen", sequenceName = "subjectIdSeq", initialValue = 1, allocationSize = 1)
	@Column(name = "id", unique = true, nullable = false)
	private Long id;
	
	@Column(name = "publicKey", unique = true, nullable = false)
	private PublicKey publicKey;

	//@Column(name = "x500name", unique = true, nullable = false)
	//private X500Name x500name;

	@Column(name = "serialNumber", unique = true, nullable = false)
	private String serialNumber;

	public Subject() {
		// TODO Auto-generated constructor stub
	}

	public Subject(PublicKey publicKey, X500Name x500name, String serialNumber) {
		super();
		this.publicKey = publicKey;
		//this.x500name = x500name;
		this.serialNumber = serialNumber;
	}

	public PublicKey getPublicKey() {
		return publicKey;
	}

	public void setPublicKey(PublicKey publicKey) {
		this.publicKey = publicKey;
	}
/*
	public X500Name getX500name() {
		return x500name;
	}

	public void setX500name(X500Name x500name) {
		this.x500name = x500name;
	}

	public String getSerialNumber() {
		return serialNumber;
	}

	public void setSerialNumber(String serialNumber) {
		this.serialNumber = serialNumber;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((publicKey == null) ? 0 : publicKey.hashCode());
		result = prime * result + ((serialNumber == null) ? 0 : serialNumber.hashCode());
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
		Subject other = (Subject) obj;
		if (publicKey == null) {
			if (other.publicKey != null)
				return false;
		} else if (!publicKey.equals(other.publicKey))
			return false;
		if (serialNumber == null) {
			if (other.serialNumber != null)
				return false;
		} else if (!serialNumber.equals(other.serialNumber))
			return false;
		if (x500name == null) {
			if (other.x500name != null)
				return false;
		} else if (!x500name.equals(other.x500name))
			return false;
		return true;
	}*/

}
