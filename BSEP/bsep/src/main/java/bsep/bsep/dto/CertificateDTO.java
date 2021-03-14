package bsep.bsep.dto;

import java.time.LocalDateTime;

import bsep.bsep.model.Certificate;

public class CertificateDTO {

	private Long id;
	private String serialNumber;
	private String signatureAlgorithmId;
	private String version;
	private LocalDateTime startDate;
	private LocalDateTime endDate;
	private Long subjectId;
	private Long issuerId;
	private boolean isExpired;
	private String alias;

	public CertificateDTO() {

	}

	public CertificateDTO(Long id, String serialNumber, String signatureAlgorithmId, String version,
			LocalDateTime startDate, LocalDateTime endDate, Long subjectId, Long issuerId, boolean isExpired, String alias) {
		super();
		this.id = id;
		this.serialNumber = serialNumber;
		this.signatureAlgorithmId = signatureAlgorithmId;
		this.version = version;
		this.startDate = startDate;
		this.endDate = endDate;
		this.subjectId = subjectId;
		this.issuerId = issuerId;
		this.isExpired = isExpired;
		this.alias = alias;
	}

	public CertificateDTO(Certificate certificate) {
		this.id = certificate.getId();
		this.serialNumber = certificate.getSerialNumber();
		this.signatureAlgorithmId = certificate.getSerialNumber();
		this.version = certificate.getVersion();
		this.startDate = certificate.getStartDate();
		this.endDate = certificate.getEndDate();
		this.subjectId = certificate.getSubjectId();
		this.issuerId = certificate.getIssuerId();
		this.isExpired = certificate.isExpired();

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

	public String getSignatureAlgorithmId() {
		return signatureAlgorithmId;
	}

	public void setSignatureAlgorithmId(String signatureAlgorithmId) {
		this.signatureAlgorithmId = signatureAlgorithmId;
	}

	public String getVersion() {
		return version;
	}

	public void setVersion(String version) {
		this.version = version;
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

	public Long getSubjectId() {
		return subjectId;
	}

	public void setSubjectId(Long subjectId) {
		this.subjectId = subjectId;
	}

	public Long getIssuerId() {
		return issuerId;
	}

	public void setIssuerId(Long issuerId) {
		this.issuerId = issuerId;
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

}
