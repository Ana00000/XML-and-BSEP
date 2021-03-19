package bsep.bsep.dto;

import bsep.bsep.model.Users;

public class UserDTO {

	private String userEmail;
	private String password;
	private String firstName;
	private String lastName;
	private String phoneNumber;
	private String typeOfUser;

	public UserDTO() {
	}

	public UserDTO(Users user) {
		this(user.getUserEmail(), user.getPassword(), user.getFirstName(), user.getLastName(), user.getPhoneNumber(),
				user.getTypeOfUser().name());
	}

	public UserDTO(String userEmail, String password, String firstName, String lastName, String phoneNumber,
			String typeOfUser) {
		super();
		this.userEmail = userEmail;
		this.password = password;
		this.firstName = firstName;
		this.lastName = lastName;
		this.phoneNumber = phoneNumber;
		this.typeOfUser = typeOfUser;
	}

	public String getUserEmail() {
		return userEmail;
	}

	public void setUserEmail(String userEmail) {
		this.userEmail = userEmail;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public String getFirstName() {
		return firstName;
	}

	public void setFirstName(String firstName) {
		this.firstName = firstName;
	}

	public String getLastName() {
		return lastName;
	}

	public void setLastName(String lastName) {
		this.lastName = lastName;
	}

	public String getPhoneNumber() {
		return phoneNumber;
	}

	public void setPhoneNumber(String phoneNumber) {
		this.phoneNumber = phoneNumber;
	}

	public String getTypeOfUser() {
		return typeOfUser;
	}

	public void setTypeOfUser(String typeOfUser) {
		this.typeOfUser = typeOfUser;
	}
}
