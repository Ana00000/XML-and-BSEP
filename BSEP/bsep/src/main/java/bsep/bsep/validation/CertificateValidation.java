package bsep.bsep.validation;

import java.util.regex.Pattern;

import bsep.bsep.dto.CertificateInfoDTO;

public class CertificateValidation {
	
	public CertificateValidation() {}
	
	public boolean validCertificate(CertificateInfoDTO certificateRequest) {
		if (!validCommonName(certificateRequest.getCommonName()) || !validGivenName(certificateRequest.getGivenName())
			|| !validSurname(certificateRequest.getSurname()) || !validOrganization(certificateRequest.getOrganization())
			|| !validOrganizationalUnitName(certificateRequest.getOrganizationalUnitName())
			|| !validOrganizationEmail(certificateRequest.getOrganizationEmail()) 
			|| !validCountryCode(certificateRequest.getCountryCode())
			|| !validAlias(certificateRequest.getAlias()))
			return false;
		return true;
	}

	private boolean validCommonName(String commonName) {
		if (commonName.isBlank()) {
			System.out.println("Your common name needs to be inserted!");
			return false;
		} else if (commonName.length() < 2) {
			System.out.println("Your common name should contain at least 2 characters!");
			return false;
		} else if (commonName.length() > 20) {
			System.out.println("Your common name shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}
	
	private boolean validGivenName(String givenName) {
		if (givenName.isBlank()) {
			System.out.println("Your given name needs to be inserted!");
			return false;
		} else if (givenName.length() < 2) {
			System.out.println("Your given name should contain at least 2 characters!");
			return false;
		} else if (givenName.length() > 20) {
			System.out.println("Your given name shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}
	
	private boolean validSurname(String surname) {
		if (surname.isBlank()) {
			System.out.println("Your surname needs to be inserted!");
			return false;
		} else if (surname.length() < 2) {
			System.out.println("Your surname should contain at least 2 characters!");
			return false;
		} else if (surname.length() > 35) {
			System.out.println("Your surname shouldn't contain more than 35 characters!");
			return false;
		}
		return true;
	}
	
	private boolean validOrganization(String organization) {
		if (organization.isBlank()) {
			System.out.println("Your organization needs to be inserted!");
			return false;
		} else if (organization.length() < 2) {
			System.out.println("Your organization should contain at least 2 characters!");
			return false;
		} else if (organization.length() > 20) {
			System.out.println("Your organization shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}
	
	private boolean validOrganizationalUnitName(String organizationalUnitName) {
		if (organizationalUnitName.isBlank()) {
			System.out.println("Your organization unit name needs to be inserted!");
			return false;
		} else if (organizationalUnitName.length() < 1) {
			System.out.println("Your organization unit name should contain at least 1 characters!");
			return false;
		} else if (organizationalUnitName.length() > 20) {
			System.out.println("Your organization unit name shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}
	
	private boolean validOrganizationEmail(String organizationEmail) {
		String ePattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*";
		if (organizationEmail.isBlank()) {
			System.out.println("Organization email needs to be inserted!");
			return false;
		} else if (!Pattern.compile(ePattern).matcher(organizationEmail).matches()) {
			System.out.println("You have entered an invalid organization email address.");
			return false;
		} else if (organizationEmail.length() > 35) {
			System.out.println("Organization email shouldn't contain more than 35 characters!");
			return false;
		}
		return true;
	}

	private boolean validCountryCode(String countryCode) {
		if (countryCode.isBlank()) {
			System.out.println("Your country code needs to be inserted!");
			return false;
		} else if (countryCode.length() < 2) {
			System.out.println("Your country code should contain at least 2 characters!");
			return false;
		} else if (countryCode.length() > 20) {
			System.out.println("Your country code shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}
	
	private boolean validAlias(String alias) {
		if (alias.isBlank()) {
			System.out.println("Your alias needs to be inserted!");
			return false;
		} else if (alias.length() < 2) {
			System.out.println("Your alias should contain at least 2 characters!");
			return false;
		} else if (alias.length() > 20) {
			System.out.println("Your alias shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}
}
