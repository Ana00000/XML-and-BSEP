package bsep.bsep.validation;

import java.util.regex.Pattern;

import bsep.bsep.dto.UserChangePasswordDTO;
import bsep.bsep.dto.UserDTO;

public class UserValidation extends GeneralValidation {

	public UserValidation() {
	}

	public boolean validPasswordAndConfirmPassword(UserChangePasswordDTO userChangePasswordDTO) {
		if (!userChangePasswordDTO.getPassword().equals(userChangePasswordDTO.getConfirmedPassword())
				|| userChangePasswordDTO.getPassword().length() != userChangePasswordDTO.getConfirmedPassword()
						.length()) {
			System.out.println("Your password and confirm password aren't equals!");
			return false;
		}
		return true;
	}

	public boolean validUser(UserDTO userRequest) {
		if (!validUserEmail(userRequest.getUserEmail()) || !validPassword(userRequest.getPassword())
				|| !validFirstName(userRequest.getFirstName()) || !validLastName(userRequest.getLastName())
				|| !validPhoneNumber(userRequest.getPhoneNumber()))
			return false;
		return true;
	}

	public boolean validUserEmail(String userEmail) {
		if (userEmail.isBlank()) {
			System.out.println("Your email needs to be inserted!");
			return false;
		} else if (!IsProperEmail(userEmail)) {
			System.out.println("You have entered an invalid email address.");
			return false;
		} else if (IsTooLong(userEmail, 35)) {
			System.out.println("Your email shouldn't contain more than 35 characters!");
			return false;
		}
		return true;
	}

	public boolean validPassword(String password) {
		if (password.isBlank()) {
			System.out.println("Your password needs to be inserted!");
			return false;
		} else if (!HasLowercaseLetter(password)) {
			System.out.println("Your password should contain at least one lowercase letter.");
			return false;
		} else if (!HasUppercaseLetter(password)) {
			System.out.println("Your password should contain at least one uppercase letter.");
			return false;
		} else if (!HasNumber(password)) {
			System.out.println("Your password should contain at least one number.");
			return false;
		} else if (!HasSpecialCharacter(password)) {
			System.out.println("Your password should contain at least one special character.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(password)) {
			System.out.println("Your password shouldn't contain special character < or >.");
			return false;
		} else if (HasSpace(password)) {
			System.out.println("Your password shouldn't contain spaces!");
			return false;
		} else if (IsTooShort(password, 10)) {
			System.out.println("Your password should contain at least 10 characters!");
			return false;
		} else if (IsTooLong(password, 30)) {
			System.out.println("Your password shouldn't contain more than 30 characters!");
			return false;
		}
		return true;
	}

	private boolean validFirstName(String firstName) {
		if (firstName.isBlank()) {
			System.out.println("Your first name needs to be inserted!");
			return false;
		} else if (HasNumber(firstName)) {
			System.out.println("Your first name shouldn't contain numbers.");
			return false;
		} else if (HasSpecialCharacter(firstName)) {
			System.out.println("Your first name shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(firstName)) {
			System.out.println("Your first name shouldn't contain special character < or >.");
			return false;
		} else if (HasSpace(firstName)) {
			System.out.println("Your first name shouldn't contain spaces!");
			return false;
		} else if (IsTooShort(firstName, 2)) {
			System.out.println("Your first name should contain at least 2 characters!");
			return false;
		} else if (IsTooLong(firstName, 20)) {
			System.out.println("Your first name shouldn't contain more than 20 characters!");
			return false;
		} else if (!HasUppercaseLetterAtStartOnly(firstName)) {
			System.out.println("Your first name needs to have one uppercase letter at the start!");
			return false;
		}
		return true;
	}

	private boolean validLastName(String lastName) {
		if (lastName.isBlank()) {
			System.out.println("Your last name needs to be inserted!");
			return false;
		} else if (HasNumber(lastName)) {
			System.out.println("Your last name shouldn't contain numbers.");
			return false;
		} else if (HasSpecialCharacter(lastName)) {
			System.out.println("Your last name shouldn't contain special characters.");
			return false;
		} else if (HasLessOrGreaterThanCharacter(lastName)) {
			System.out.println("Your last name shouldn't contain special character < or >.");
			return false;
		} else if (HasSpace(lastName)) {
			System.out.println("Your last name shouldn't contain spaces!");
			return false;
		} else if (IsTooShort(lastName, 2)) {
			System.out.println("Your last name should contain at least 2 characters!");
			return false;
		} else if (IsTooLong(lastName, 35)) {
			System.out.println("Your last name shouldn't contain more than 35 characters!");
			return false;
		} else if (!HasUppercaseLetterAtStartOnly(lastName)) {
			System.out.println("Your last name needs to have one uppercase letter at the start!");
			return false;
		}
		return true;
	}

	private boolean validPhoneNumber(String phoneNumber) {
		if (phoneNumber.isBlank()) {
			System.out.println("Your phone number needs to be inserted!");
			return false;
		} else if (HasSpace(phoneNumber)) {
			System.out.println("Your phone number shouldn't contain spaces!");
			return false;
		} else if (HasLowercaseLetter(phoneNumber) || HasUppercaseLetter(phoneNumber)) {
			System.out.println("Your phone number shouldn't contain letters!");
			return false;
		} else if (!Pattern.compile("^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\\s./0-9]*").matcher(phoneNumber).matches()) {
			System.out.println("Your phone number is not in right form!");
			return false;
		}
		return true;
	}
}
