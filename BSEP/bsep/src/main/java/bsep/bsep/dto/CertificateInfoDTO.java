package bsep.bsep.dto;

public class CertificateInfoDTO {

	private String commonName;
	private String givenName;
	private String surname;
	private String organization;
	private String organizationalUnitName;
	private String organizationEmail;
	private String countryCode;
	private String endDate;

	public CertificateInfoDTO() {
	}

	public CertificateInfoDTO(String commonName, String givenName, String surname, String organization,
			String organizationalUnitName, String organizationEmail, String countryCode, String endDate) {
		super();
		this.commonName = commonName;
		this.givenName = givenName;
		this.surname = surname;
		this.organization = organization;
		this.organizationalUnitName = organizationalUnitName;
		this.organizationEmail = organizationEmail;
		this.countryCode = countryCode;
		this.endDate = endDate;
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

	public String getEndDate() {
		return endDate;
	}

	public void setEndDate(String endDate) {
		this.endDate = endDate;
	}
}
