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
			|| !validAlias(certificateRequest.getAlias())
			|| !validEndDate(certificateRequest.getEndDate()))
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
		} else if (surname.length() > 100) {
			System.out.println("Your surname shouldn't contain more than 100 characters!");
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
	
	private boolean validEndDate(String endDate) {
		if (endDate.isBlank()) {
			System.out.println("Your end date needs to be inserted!");
			return false;
		}else if (Pattern.compile("[a-zA-Z]+").matcher(endDate).find()) {
			System.out.println("Your end date shouldn't contain letters.");
			return false;
		}else if (Pattern.compile("[!@#$%^&*,:'/.\"]+").matcher(endDate).find()) {
			System.out.println("Your end date shouldn't contain special character other than [-].");
			return false;
		}else if (!Pattern.compile("[2][0-9]{3}-[0-1][0-9]-[0-3][0-9]").matcher(endDate).matches()) {
			System.out.println("Your end date is not set in right format.");
			return false;
		}else if(!validEndDateParts(endDate)) return false;
		
		return true;
	}

	private boolean validEndDateParts(String endDate) {
		String[] endDateSplit = endDate.split("-");
		int eDSYear = Integer.parseInt(endDateSplit[0]);
		int eDSMonth = Integer.parseInt(endDateSplit[1]);
		int eDSDay = Integer.parseInt(endDateSplit[2]);
		
		if (eDSYear > 3000 || eDSYear < 2021){
			 System.out.println("Year of end date isn't valid");
				return false;
		} else if (eDSYear < 2025) {
			System.out.println("End date must me valid 5 years from current date.");
			return false;
		} else if (eDSMonth > 12 || eDSMonth < 0) {
			System.out.println("Month of end date isn't valid");
			return false;
		} else if (eDSDay > 31 || eDSDay < 1) {
			System.out.println("Day of end date isn't valid");
			return false;
		}
		return true;
	}
}
