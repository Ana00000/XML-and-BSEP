package bsep.bsep.model;

import java.time.LocalDate;
import java.time.LocalDateTime;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

import bsep.bsep.dto.CertificateDTO;

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

	@Column(name = "signatureAlgorithmName", unique = false, nullable = false)
	private String signatureAlgorithmName;

	@Column(name = "version", unique = false, nullable = false)
	private String version;

	@Column(name = "startDate", unique = false, nullable = false)
	private LocalDate startDate;

	@Column(name = "endDate", unique = false, nullable = false)
	private LocalDate endDate;
	
	@Column(name = "type", unique = false, nullable = false)
	private String type;	

	@Column(name = "subject", unique = false, nullable = false)
	private String subject;

	@Column(name = "issuer", unique = false, nullable = false)
	private String issuer;

	@Column(name = "isExpired", unique = false, nullable = false)
	private boolean isExpired;

	@Column(name = "alias", unique = true, nullable = false)
	private String alias;
	
	@Column(name = "keyStoreFileName", unique = true, nullable = false)
	private String keyStoreFileName;

	public Certificate() {
		super();
	}
	
	public Certificate(Long id, String serialNumber, String signatureAlgorithmName, String version, LocalDate startDate,
			LocalDate endDate, String type, String subject, String issuer, boolean isExpired, String alias,
			String keyStoreFileName) {
		super();
		this.id = id;
		this.serialNumber = serialNumber;
		this.signatureAlgorithmName = signatureAlgorithmName;
		this.version = version;
		this.startDate = startDate;
		this.endDate = endDate;
		this.type = type;
		this.subject = subject;
		this.issuer = issuer;
		this.isExpired = isExpired;
		this.alias = alias;
		this.keyStoreFileName = keyStoreFileName;
	}

	public Certificate(CertificateDTO certificateDTO) {
		this.id = certificateDTO.getId();
		this.serialNumber = certificateDTO.getSerialNumber();
		this.signatureAlgorithmName = certificateDTO.getSignatureAlgorithmName();
		this.version = certificateDTO.getVersion();
		this.startDate = certificateDTO.getStartDate();
		this.endDate = certificateDTO.getEndDate();
		this.subject = certificateDTO.getSubject();
		this.issuer = certificateDTO.getIssuer();
		this.isExpired = certificateDTO.isExpired();
		this.alias = certificateDTO.getAlias();
		this.type = certificateDTO.getType();
		this.keyStoreFileName = certificateDTO.getKeyStoreFileName();
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
