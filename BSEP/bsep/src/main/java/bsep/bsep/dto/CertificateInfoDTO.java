package bsep.bsep.dto;

import bsep.bsep.model.CertificatePurposeType;
import bsep.bsep.model.CertificateType;

public class CertificateInfoDTO {

	private String commonName;
	private String givenName;
	private String surname;
	private String organization;
	private String organizationalUnitName;
	private String organizationEmail;
	private String countryCode;
	private String alias;
	private String endDate;
	private CertificateType certificateType;
	private CertificatePurposeType certificatePurposeType;
	private String issuerSerialNumber;

	public CertificateInfoDTO() {
	}

	public CertificateInfoDTO(String commonName, String givenName, String surname, String organization,
			String organizationalUnitName, String organizationEmail, String countryCode, String alias, String endDate,
			CertificateType certificateType, CertificatePurposeType certificatePurposeType, String issuerSerialNumber) {
		super();
		this.commonName = commonName;
		this.givenName = givenName;
		this.surname = surname;
		this.organization = organization;
		this.organizationalUnitName = organizationalUnitName;
		this.organizationEmail = organizationEmail;
		this.countryCode = countryCode;
		this.alias = alias;
		this.endDate = endDate;
		this.certificateType = certificateType;
		this.certificatePurposeType = certificatePurposeType;
		this.issuerSerialNumber = issuerSerialNumber;
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

	public String getCountryCode() {
		return countryCode;
	}

	public void setCountryCode(String countryCode) {
		this.countryCode = countryCode;
	}

	public String getAlias() {
		return alias;
	}

	public void setAlias(String alias) {
		this.alias = alias;
	}

	public String getEndDate() {
		return endDate;
	}

	public void setEndDate(String endDate) {
		this.endDate = endDate;
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

	public String getIssuerSerialNumber() {
		return issuerSerialNumber;
	}

	public void setIssuerSerialNumber(String issuerSerialNumber) {
		this.issuerSerialNumber = issuerSerialNumber;
	}

}