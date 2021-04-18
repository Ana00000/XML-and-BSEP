package bsep.bsep.validation;

import java.util.regex.Pattern;

import bsep.bsep.dto.CertificateInfoDTO;

public class CertificateValidation extends GeneralValidation {

	public CertificateValidation() {
	}

	public boolean validCertificate(CertificateInfoDTO certificateRequest) {
		if (!validCommonName(certificateRequest.getCommonName()) || !validGivenName(certificateRequest.getGivenName())
				|| !validSurname(certificateRequest.getSurname())
				|| !validOrganization(certificateRequest.getOrganization())
				|| !validOrganizationalUnitName(certificateRequest.getOrganizationalUnitName())
				|| !validOrganizationEmail(certificateRequest.getOrganizationEmail())
				|| !validCountryCode(certificateRequest.getCountryCode()) || !validAlias(certificateRequest.getAlias())
				|| !validEndDate(certificateRequest.getEndDate()))
			return false;
		return true;
	}

	private boolean validCommonName(String commonName) {
		if (commonName.isBlank()) {
			System.out.println("Your common name needs to be inserted!");
			return false;
		} else if (HasNumber(commonName)) {
			System.out.println("Your common name shouldn't contain numbers.");
			return false;
		} else if (HasSpecialCharacter(commonName)) {
			System.out.println("Your common name shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(commonName)) {
			System.out.println("Your common name shouldn't contain special character < or >.");
			return false;
		} else if (HasSpace(commonName)) {
			System.out.println("Your common name shouldn't contain spaces!");
			return false;
		} else if (IsTooShort(commonName, 2)) {
			System.out.println("Your common name should contain at least 2 characters!");
			return false;
		} else if (IsTooLong(commonName, 20)) {
			System.out.println("Your common name shouldn't contain more than 20 characters!");
			return false;
		} else if (!HasUppercaseLetterAtStartOnly(commonName)) {
			System.out.println("Your common name needs to have one uppercase letter at the start!");
			return false;
		}
		return true;
	}

	private boolean validGivenName(String givenName) {
		if (givenName.isBlank()) {
			System.out.println("Your given name needs to be inserted!");
			return false;
		} else if (HasNumber(givenName)) {
			System.out.println("Your given name shouldn't contain numbers.");
			return false;
		} else if (HasSpecialCharacter(givenName)) {
			System.out.println("Your given name shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(givenName)) {
			System.out.println("Your given name shouldn't contain special character < or >.");
			return false;
		} else if (HasSpace(givenName)) {
			System.out.println("Your given name shouldn't contain spaces!");
			return false;
		} else if (IsTooShort(givenName, 2)) {
			System.out.println("Your given name should contain at least 2 characters!");
			return false;
		} else if (IsTooLong(givenName, 20)) {
			System.out.println("Your given name shouldn't contain more than 20 characters!");
			return false;
		} else if (!HasUppercaseLetterAtStartOnly(givenName)) {
			System.out.println("Your given name needs to have one uppercase letter at the start!");
			return false;
		}
		return true;
	}

	private boolean validSurname(String surname) {
		if (surname.isBlank()) {
			System.out.println("Your surname needs to be inserted!");
			return false;
		} else if (HasNumber(surname)) {
			System.out.println("Your surname shouldn't contain numbers.");
			return false;
		} else if (HasSpecialCharacter(surname)) {
			System.out.println("Your surname shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(surname)) {
			System.out.println("Your surname shouldn't contain special character < or >.");
			return false;
		} else if (HasSpace(surname)) {
			System.out.println("Your surname shouldn't contain spaces!");
			return false;
		} else if (IsTooShort(surname, 2)) {
			System.out.println("Your surname should contain at least 2 characters!");
			return false;
		} else if (IsTooLong(surname, 35)) {
			System.out.println("Your surname shouldn't contain more than 35 characters!");
			return false;
		} else if (!HasUppercaseLetterAtStartOnly(surname)) {
			System.out.println("Your surname needs to have one uppercase letter at the start!");
			return false;
		}
		return true;
	}

	private boolean validOrganization(String organization) {
		if (organization.isBlank()) {
			System.out.println("Your organization needs to be inserted!");
			return false;
		} else if (HasSpecialCharacter(organization)) {
			System.out.println("Your organization shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(organization)) {
			System.out.println("Your organization shouldn't contain special character < or >.");
			return false;
		} else if (IsTooShort(organization, 2)) {
			System.out.println("Your organization should contain at least 2 characters!");
			return false;
		} else if (IsTooLong(organization, 20)) {
			System.out.println("Your organization shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}

	private boolean validOrganizationalUnitName(String organizationalUnitName) {
		if (organizationalUnitName.isBlank()) {
			System.out.println("Your organization unit name needs to be inserted!");
			return false;
		} else if (HasSpecialCharacter(organizationalUnitName)) {
			System.out.println("Your organization unit name shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(organizationalUnitName)) {
			System.out.println("Your organization unit name shouldn't contain special character < or >.");
			return false;
		} else if (HasNumber(organizationalUnitName)) {
			System.out.println("Your organizationalUnitName shouldn't contain numbers.");
			return false;
		} else if (IsTooShort(organizationalUnitName, 2)) {
			System.out.println("Your organization unit name should contain at least 2 characters!");
			return false;
		} else if (IsTooLong(organizationalUnitName, 20)) {
			System.out.println("Your organization unit name shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}

	private boolean validOrganizationEmail(String organizationEmail) {
		if (organizationEmail.isBlank()) {
			System.out.println("Organization email needs to be inserted!");
			return false;
		} else if (!IsProperEmail(organizationEmail)) {
			System.out.println("You have entered an invalid organization email address.");
			return false;
		} else if (IsTooLong(organizationEmail, 35)) {
			System.out.println("Organization email shouldn't contain more than 35 characters!");
			return false;
		}
		return true;
	}

	private boolean validCountryCode(String countryCode) {
		if (countryCode.isBlank()) {
			System.out.println("Your country code needs to be inserted!");
			return false;
		} else if (!Pattern.compile("^[A-Z]{2,3}").matcher(countryCode).matches()) {
			System.out.println("Your country code is not in right form!");
			return false;
		}
		return true;
	}

	private boolean validAlias(String alias) {
		if (alias.isBlank()) {
			System.out.println("Your alias needs to be inserted!");
			return false;
		} else if (HasSpecialCharacter(alias)) {
			System.out.println("Your alias shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(alias)) {
			System.out.println("Your alias shouldn't contain special character < or >.");
			return false;
		} else if (HasSpace(alias)) {
			System.out.println("Your alias shouldn't contain spaces!");
			return false;
		} else if (IsTooLong(alias, 20)) {
			System.out.println("Your alias shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}

	private boolean validEndDate(String endDate) {
		if (endDate.isBlank()) {
			System.out.println("Your end date needs to be inserted!");
			return false;
		} else if (HasLowercaseLetter(endDate) || HasUppercaseLetter(endDate)) {
			System.out.println("Your end date shouldn't contain letters.");
			return false;
		} else if (Pattern.compile("[!@#$%^&*,:<>+'/.\"]+").matcher(endDate).find()) {
			System.out.println("Your end date shouldn't contain special character other than [-].");
			return false;
		} else if (HasSpace(endDate)) {
			System.out.println("Your end date shouldn't contain spaces!");
			return false;
		} else if (!Pattern.compile("[2][0-9]{3}-[0-1][0-9]-[0-3][0-9]").matcher(endDate).matches()) {
			System.out.println("Your end date is not set in right format.");
			return false;
		} else if (!validEndDateParts(endDate))
			return false;

		return true;
	}

	private boolean validEndDateParts(String endDate) {
		String[] endDateSplit = endDate.split("-");
		int eDSYear = Integer.parseInt(endDateSplit[0]);
		int eDSMonth = Integer.parseInt(endDateSplit[1]);
		int eDSDay = Integer.parseInt(endDateSplit[2]);

		if (eDSYear > 3000 || eDSYear < 2021) {
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
