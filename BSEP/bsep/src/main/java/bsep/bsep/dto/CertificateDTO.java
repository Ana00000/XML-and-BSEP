package bsep.bsep.dto;

import java.time.LocalDateTime;

import bsep.bsep.model.Certificate;

public class CertificateDTO {

	private Long id;
	private String serialNumber;
	private String signatureAlgorithmId;
	private String version;
	private LocalDateTime start;
	private LocalDateTime end;
	private Long subjectId;
	private Long issuerId;
	private boolean isExpired;
	
	public CertificateDTO() {
		
	}

	public CertificateDTO(Long id, String serialNumber, String signatureAlgorithmId, String version,
			LocalDateTime start, LocalDateTime end, Long subjectId, Long issuerId, boolean isExpired) {
		super();
		this.id = id;
		this.serialNumber = serialNumber;
		this.signatureAlgorithmId = signatureAlgorithmId;
		this.version = version;
		this.start = start;
		this.end = end;
		this.subjectId = subjectId;
		this.issuerId = issuerId;
		this.isExpired = isExpired;
	}
	
	public CertificateDTO(Certificate certificate)
	{
		this.id = certificate.getId();
		this.serialNumber = certificate.getSerialNumber();
		this.signatureAlgorithmId = certificate.getSerialNumber();
		this.version = certificate.getVersion();
		this.start = certificate.getStart();
		this.end = certificate.getEnd();
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
	
	
	
	
}
