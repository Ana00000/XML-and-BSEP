package bsep.bsep.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

import bsep.bsep.dto.UserDTO;

@Entity
@Table(name = "users")
public class Users {

	@Id
	@SequenceGenerator(name = "mySeqGenUsers", sequenceName = "mySeqUsers", initialValue = 1, allocationSize = 1)
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "mySeqGenUsers")
	private Long id;

	@Column(name = "userEmail", unique = true, nullable = false)
	private String userEmail;

	@Column(name = "password", unique = false, nullable = false)
	private String password;

	@Column(name = "firstName", unique = false, nullable = false)
	private String firstName;

	@Column(name = "lastName", unique = false, nullable = false)
	private String lastName;

	@Column(name = "phoneNumber", unique = false, nullable = false)
	private String phoneNumber;

	@Enumerated(EnumType.ORDINAL)
	private UserType typeOfUser;

	public Users() {
	}

	public Users(UserDTO userDTO) {
		this.userEmail = userDTO.getUserEmail();
		this.password = userDTO.getPassword();
		this.firstName = userDTO.getFirstName();
		this.lastName = userDTO.getLastName();
		this.phoneNumber = userDTO.getPhoneNumber();
		
		if(userDTO.getTypeOfUser().equals("ADMIN"))
			this.typeOfUser = UserType.ADMIN;
		else
			this.typeOfUser = UserType.USER;
	}

	public Users(Long id, String userEmail, String password, String firstName, String lastName, String phoneNumber,
			UserType typeOfUser) {
		super();
		this.id = id;
		this.userEmail = userEmail;
		this.password = password;
		this.firstName = firstName;
		this.lastName = lastName;
		this.phoneNumber = phoneNumber;
		this.typeOfUser = typeOfUser;
	}

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
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

	public UserType getTypeOfUser() {
		return typeOfUser;
	}

	public void setTypeOfUser(UserType typeOfUser) {
		this.typeOfUser = typeOfUser;
	}
}
