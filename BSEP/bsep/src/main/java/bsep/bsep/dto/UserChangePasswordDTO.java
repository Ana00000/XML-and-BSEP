package bsep.bsep.dto;

public class UserChangePasswordDTO {
	private String emailOfUser;
	private String password;
	private String confirmedPassword;

	public UserChangePasswordDTO() {
	}

	public UserChangePasswordDTO(String emailOfUser, String password, String confirmedPassword) {
		super();
		this.emailOfUser = emailOfUser;
		this.password = password;
		this.confirmedPassword = confirmedPassword;
	}

	public String getEmailOfUser() {
		return emailOfUser;
	}

	public void setEmailOfUser(String emailOfUser) {
		this.emailOfUser = emailOfUser;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public String getConfirmedPassword() {
		return confirmedPassword;
	}

	public void setConfirmedPassword(String confirmedPassword) {
		this.confirmedPassword = confirmedPassword;
	}

}
