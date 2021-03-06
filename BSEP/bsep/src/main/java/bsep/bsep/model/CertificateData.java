package bsep.bsep.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

@Entity
@Table(name = "certificateData")
public class CertificateData {

	@Id
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "certificateIdSeqGen")
	@SequenceGenerator(name = "certificateIdSeqGen", sequenceName = "certificateIdSeq", initialValue = 1, allocationSize = 1)
	@Column(name = "id", unique = true, nullable = false)
	private Long id;

	@Column(name = "serialNumber", unique = true, nullable = false)
	private String serialNumber;
	
	@Column(name = "subjectEmail", unique = false, nullable = false)
	private String subjectEmail;

	@Enumerated(EnumType.ORDINAL)
	private CertificateStatus certificateStatus;

	@Enumerated(EnumType.ORDINAL)
	private CertificateType certificateType;

	@Enumerated(EnumType.ORDINAL)
	private CertificatePurposeType certificatePurposeType;

	public CertificateData() {

	}
	
	public CertificateData(Long id, String serialNumber, String subjectEmail, CertificateStatus certificateStatus,
			CertificateType certificateType, CertificatePurposeType certificatePurposeType) {
		super();
		this.id = id;
		this.serialNumber = serialNumber;
		this.subjectEmail = subjectEmail;
		this.certificateStatus = certificateStatus;
		this.certificateType = certificateType;
		this.certificatePurposeType = certificatePurposeType;
	}
	
	public String getSubjectEmail() {
		return subjectEmail;
	}

	public void setSubjectEmail(String subjectEmail) {
		this.subjectEmail = subjectEmail;
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

	public CertificateStatus getCertificateStatus() {
		return certificateStatus;
	}

	public void setCertificateStatus(CertificateStatus certificateStatus) {
		this.certificateStatus = certificateStatus;
	}

	public CertificateType getCertificateType() {
		return certificateType;
	}

	public void setCertificateType(CertificateType certificateType) {
		this.certificateType = certificateType;
	}

	public CertificatePurposeType getCertificatePurposeType() {
		return certificatePurposeType;
	}

	public void setCertificatePurposeType(CertificatePurposeType certificatePurposeType) {
		this.certificatePurposeType = certificatePurposeType;
	}

}
