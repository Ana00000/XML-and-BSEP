package bsep.bsep.dto;

import bsep.bsep.model.CertificateStatus;
import bsep.bsep.model.CertificateType;

public class CertificateDTO {

	private Long id;
	private String serialNumber;
	private String commonName;
	private String givenName;
	private String surname;
	private String organization;
	private String organizationalUnitName;
	private String organizationEmail;
	private String countryCode;
	private String alias;
	private String version;
	private String signatureAlgorithmName;
	private CertificateStatus certificateStatus; 
	private CertificateType certificateType; 
	private String subject;
	private String issuer;

	public CertificateDTO() {

	}

	public CertificateDTO(Long id, String serialNumber, String commonName, String givenName, String surname,
			String organization, String organizationalUnitName, String organizationEmail, String countryCode,
			String alias, String version, String signatureAlgorithmName, CertificateStatus certificateStatus,
			CertificateType certificateType, String subject, String issuer) {
		super();
		this.id = id;
		this.serialNumber = serialNumber;
		this.commonName = commonName;
		this.givenName = givenName;
		this.surname = surname;
		this.organization = organization;
		this.organizationalUnitName = organizationalUnitName;
		this.organizationEmail = organizationEmail;
		this.countryCode = countryCode;
		this.alias = alias;
		this.version = version;
		this.signatureAlgorithmName = signatureAlgorithmName;
		this.certificateStatus = certificateStatus;
		this.certificateType = certificateType;
		this.subject = subject;
		this.issuer = issuer;
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

	public String getCommonName() {
		return commonName;
	}

	public void setCommonName(String commonName) {
		this.commonName = commonName;
	}

	public String getGivenName() {
		return givenName;
	}

	public void setGivenName(String givenName) {
		this.givenName = givenName;
	}

	public String getSurname() {
		return surname;
	}

	public void setSurname(String surname) {
		this.surname = surname;
	}

	public String getOrganization() {
		return organization;
	}

	public void setOrganization(String organization) {
		this.organization = organization;
	}

	public String getOrganizationalUnitName() {
		return organizationalUnitName;
	}

	public void setOrganizationalUnitName(String organizationalUnitName) {
		this.organizationalUnitName = organizationalUnitName;
	}

	public String getOrganizationEmail() {
		return organizationEmail;
	}

	public void setOrganizationEmail(String organizationEmail) {
		this.organizationEmail = organizationEmail;
	}

	public String getAlias() {
		return alias;
	}

	public void setAlias(String alias) {
		this.alias = alias;
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

	public String getVersion() {
		return version;
	}

	public void setVersion(String version) {
		this.version = version;
	}

	public String getSignatureAlgorithmName() {
		return signatureAlgorithmName;
	}

	public void setSignatureAlgorithmName(String signatureAlgorithmName) {
		this.signatureAlgorithmName = signatureAlgorithmName;
	}

	public String getCountryCode() {
		return countryCode;
	}

	public void setCountryCode(String countryCode) {
		this.countryCode = countryCode;
	}

}
