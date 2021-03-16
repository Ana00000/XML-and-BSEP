package bsep.bsep.model;

import java.security.PublicKey;
import java.time.LocalDateTime;

import org.bouncycastle.asn1.x500.X500Name;

public class Subject extends Users {

	private PublicKey publicKey;

	private X500Name x500name;

	private String serialNumber;

	private LocalDateTime startDate;

	private LocalDateTime endDate;

	public Subject() {
	}

	public Subject(PublicKey publicKey, X500Name x500name, String serialNumber, LocalDateTime startDate,
			LocalDateTime endDate) {
		super();
		this.publicKey = publicKey;
		this.x500name = x500name;
		this.serialNumber = serialNumber;
		this.startDate = startDate;
		this.endDate = endDate;
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

	public LocalDateTime getStartDate() {
		return startDate;
	}

	public void setStartDate(LocalDateTime startDate) {
		this.startDate = startDate;
	}

	public LocalDateTime getEndDate() {
		return endDate;
	}

	public void setEndDate(LocalDateTime endDate) {
		this.endDate = endDate;
	}
}
