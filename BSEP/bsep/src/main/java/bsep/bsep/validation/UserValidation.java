package bsep.bsep.validation;

import java.util.regex.Pattern;

import bsep.bsep.dto.UserChangePasswordDTO;
import bsep.bsep.dto.UserDTO;

public class UserValidation {
	
	public UserValidation() {}
	
	public boolean validUser(UserDTO userRequest) {
		if (!validUserEmail(userRequest.getUserEmail()) || !validPassword(userRequest.getPassword())
				|| !validFirstName(userRequest.getFirstName()) || !validLastName(userRequest.getLastName())
				|| !validPhoneNumber(userRequest.getPhoneNumber()))
			return false;
		return true;
	}

	public boolean validUserEmail(String userEmail) {
		String ePattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*";
		if (userEmail.isBlank()) {
			System.out.println("Your email needs to be inserted!");
			return false;
		} else if (!Pattern.compile(ePattern).matcher(userEmail).matches()) {
			System.out.println("You have entered an invalid email address.");
			return false;
		} else if (userEmail.length() > 35) {
			System.out.println("Your email shouldn't contain more than 35 characters!");
			return false;
		}
		return true;
	}
	
	public boolean validPasswordAndConfirmPassword(UserChangePasswordDTO userChangePasswordDTO) {
		if (!userChangePasswordDTO.getPassword().equals(userChangePasswordDTO.getConfirmedPassword()) || userChangePasswordDTO.getPassword().length()!=userChangePasswordDTO.getConfirmedPassword().length()) {
			System.out.println("Your password and confirm password aren't equals!");
			return false;
		}
		return true;
	}

	public boolean validPassword(String password) {
		if (password.isBlank()) {
			System.out.println("Your password needs to be inserted!");
			return false;
		} else if (password.length() < 10) {
			System.out.println("Your password should contain at least 10 character!");
			return false;
		} else if (password.length() > 30) {
			System.out.println("Your password shouldn't contain more than 30 characters!");
			return false;
		} else if (!Pattern.compile("[a-z]+").matcher(password).find()) {
			System.out.println("Your password should contain at least one small letter.");
			return false;
		} else if (!Pattern.compile("[A-Z]+").matcher(password).find()) {
			System.out.println("Your password should contain at least one big letter.");
			return false;
		} else if (!Pattern.compile("[0-9]+").matcher(password).find()) {
			System.out.println("Your password should contain at least one number.");
			return false;
		} else if (!Pattern.compile("[!@#$%^&*.,:'\\\"]+").matcher(password).find()) {
			System.out.println("Your password should contain at least special character.");
			return false;
		}
		return true;
	}

	private boolean validFirstName(String firstName) {
		if (firstName.isBlank()) {
			System.out.println("Your first name needs to be inserted!");
			return false;
		} else if (firstName.length() < 2) {
			System.out.println("Your first name should contain at least 2 characters!");
			return false;
		} else if (firstName.length() > 20) {
			System.out.println("Your first name shouldn't contain more than 20 characters!");
			return false;
		}
		return true;
	}

	private boolean validLastName(String lastName) {
		if (lastName.isBlank()) {
			System.out.println("Your last name needs to be inserted!");
			return false;
		} else if (lastName.length() < 2) {
			System.out.println("Your last name should contain at least 2 characters!");
			return false;
		} else if (lastName.length() > 25) {
			System.out.println("Your last name shouldn't contain more than 25 characters!");
			return false;
		}
		return true;
	}

	private boolean validPhoneNumber(String phoneNumber) {
		if (phoneNumber.isBlank()) {
			System.out.println("Your phone number needs to be inserted!");
			return false;
		} else if (!phoneNumber.matches("[0-9]{9,12}[+-/]*")) {
			System.out.println(
					"Your phone number needs to contain 9 to 12 numbers and it can contain only [+, -, /] special characters.");
			return false;
		}
		return true;
	}
}
