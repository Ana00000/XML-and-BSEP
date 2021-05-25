package bsep.bsep.model;

import java.security.PublicKey;
import java.time.LocalDate;

import org.bouncycastle.asn1.x500.X500Name;

public class Subject extends Users {

	private PublicKey publicKey;

	private X500Name x500name;

	private String serialNumber;

	private LocalDate startDate;

	private LocalDate endDate;

	public Subject() {
	}

	public Subject(PublicKey publicKey, X500Name x500name, String serialNumber, LocalDate startDate,
			LocalDate endDate) {
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

	public LocalDate getStartDate() {
		return startDate;
	}

	public void setStartDate(LocalDate startDate) {
		this.startDate = startDate;
	}

	public LocalDate getEndDate() {
		return endDate;
	}

	public void setEndDate(LocalDate endDate) {
		this.endDate = endDate;
	}
}
