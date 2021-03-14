package bsep.bsep.model;

import java.time.LocalDateTime;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

@Entity
@Table(name = "certificate")
public class Certificate {

	@Id
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "certificateIdSeqGen")
	@SequenceGenerator(name = "certificateIdSeqGen", sequenceName = "certificateIdSeq", initialValue = 1, allocationSize = 1)
	@Column(name = "id", unique = true, nullable = false)
	private Long id;

	@Column(name = "serialNumber", unique = true, nullable = false)
	private String serialNumber;

	@Column(name = "signatureAlgorithmId", unique = false, nullable = false)
	private String signatureAlgorithmId;

	@Column(name = "version", unique = false, nullable = false)
	private String version;

	@Column(name = "startDate", unique = false, nullable = false)
	private LocalDateTime startDate;

	@Column(name = "endDate", unique = false, nullable = false)
	private LocalDateTime endDate;

	@Column(name = "subjectId", unique = false, nullable = false)
	private Long subjectId;

	@Column(name = "issuerId", unique = false, nullable = false)
	private Long issuerId;

	@Column(name = "isExpired", unique = false, nullable = false)
	private boolean isExpired;

	@Column(name = "alias", unique = true, nullable = false)
	private String alias;

	public Certificate() {
		super();
	}

	public Certificate(Long id, String serialNumber, String signatureAlgorithmId, String version,
			LocalDateTime startDate, LocalDateTime endDate, Long subjectId, Long issuerId, boolean isExpired,
			String alias) {
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
