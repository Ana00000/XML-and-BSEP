package bsep.bsep.dto;

import java.time.LocalDate;

import bsep.bsep.model.Certificate;

public class CertificateDTO {

	private Long id;
	private String serialNumber;
	private String signatureAlgorithmName;
	private String version;
	private LocalDate startDate;
	private LocalDate endDate;
	private String subject;
	private String issuer;
	private boolean isExpired;
	private String alias;
	private String type;
	private String keyStoreFileName;

	public CertificateDTO() {

	}

	public CertificateDTO(Long id, String serialNumber, String signatureAlgorithmName, String version,
			LocalDate startDate, LocalDate endDate, String subject, String issuer, boolean isExpired, String alias,
			String type, String keyStoreFileName) {
		super();
		this.id = id;
		this.serialNumber = serialNumber;
		this.signatureAlgorithmName = signatureAlgorithmName;
		this.version = version;
		this.startDate = startDate;
		this.endDate = endDate;
		this.subject = subject;
		this.issuer = issuer;
		this.isExpired = isExpired;
		this.alias = alias;
		this.type = type;
		this.keyStoreFileName = keyStoreFileName;
	}

	public CertificateDTO(Certificate certificate) {
		this.id = certificate.getId();
		this.serialNumber = certificate.getSerialNumber();
		this.signatureAlgorithmName = certificate.getSignatureAlgorithmName();
		this.version = certificate.getVersion();
		this.startDate = certificate.getStartDate();
		this.endDate = certificate.getEndDate();
		this.subject = certificate.getSubject();
		this.issuer = certificate.getIssuer();
		this.isExpired = certificate.isExpired();
		this.type = certificate.getCertificateType();
		this.keyStoreFileName = certificate.getKeyStoreFileName();
	}

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public String getSerialNumber() {
		return serialNumber;
	}

	public void setSerialNumber(String serialNumber) {
		this.serialNumber = serialNumber;
	}

	public String getSignatureAlgorithmName() {
		return signatureAlgorithmName;
	}

	public void setSignatureAlgorithmName(String signatureAlgorithmName) {
		this.signatureAlgorithmName = signatureAlgorithmName;
	}

	public String getVersion() {
		return version;
	}

	public void setVersion(String version) {
		this.version = version;
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

	public String getSubject() {
		return subject;
	}

	public void setSubject(String subject) {
		this.subject = subject;
	}

	public String getIssuer() {
		return issuer;
	}

	public void setIssuer(String issuer) {
		this.issuer = issuer;
	}

	public boolean isExpired() {
		return isExpired;
	}

	public void setExpired(boolean isExpired) {
		this.isExpired = isExpired;
	}

	public String getAlias() {
		return alias;
	}

	public void setAlias(String alias) {
		this.alias = alias;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getKeyStoreFileName() {
		return keyStoreFileName;
	}

	public void setKeyStoreFileName(String keyStoreFileName) {
		this.keyStoreFileName = keyStoreFileName;
	}
	
}
