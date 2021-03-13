package bsep.bsep.model;

import java.security.PublicKey;
import java.time.LocalDateTime;

import org.bouncycastle.asn1.x500.X500Name;

public class Subject {

	private PublicKey publicKey;

	private X500Name x500name;

	private String serialNumber;

	private LocalDateTime start;

	private LocalDateTime end;

	public Subject() {
	}

	public Subject(PublicKey publicKey, X500Name x500name, String serialNumber, LocalDateTime start,
			LocalDateTime end) {
		super();
		this.publicKey = publicKey;
		this.x500name = x500name;
		this.serialNumber = serialNumber;
		this.start = start;
		this.end = end;
	}

	public PublicKey getPublicKey() {
		return publicKey;
	}

	public void setPublicKey(PublicKey publicKey) {
		this.publicKey = publicKey;
	}

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

	public LocalDateTime getStart() {
		return start;
	}

	public void setStart(LocalDateTime start) {
		this.start = start;
	}

	public LocalDateTime getEnd() {
		return end;
	}

	public void setEnd(LocalDateTime end) {
		this.end = end;
	}
}
