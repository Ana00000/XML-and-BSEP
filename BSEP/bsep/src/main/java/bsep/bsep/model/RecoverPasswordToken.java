package bsep.bsep.model;

import java.time.LocalDate;
import java.util.UUID;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToOne;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

import com.fasterxml.jackson.annotation.JsonFormat;

@Entity
@Table(name = "recovery_password_tokens")
public class RecoverPasswordToken {

	@Id
	@SequenceGenerator(name = "mySeqGenRecoverPasswordToken", sequenceName = "mySeqRecoverPasswordToken", initialValue = 1, allocationSize = 1)
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "mySeqGenRecoverPasswordToken")
	@Column(name = "id", unique = true, nullable = false)
	private Long id;

	@Column(name = "recovery_password_token")
	private String recoveryPasswordToken;

	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd", timezone = "UTC")
	private LocalDate createdDate;

	@OneToOne(targetEntity = Users.class, fetch = FetchType.EAGER)
	@JoinColumn(nullable = false, name = "user_id")
	private Users users;

	public RecoverPasswordToken() {
		// TODO Auto-generated constructor stub
	}

	public RecoverPasswordToken(Long id, String recoveryPasswordToken, LocalDate createdDate, Users users) {
		super();
		this.id = id;
		this.recoveryPasswordToken = recoveryPasswordToken;
		this.createdDate = createdDate;
		this.users = users;
	}

	public RecoverPasswordToken(Users users) {
		this.users = users;
		createdDate = LocalDate.now();
		recoveryPasswordToken = UUID.randomUUID().toString();
	}

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public String getRecoveryPasswordToken() {
		return recoveryPasswordToken;
	}

	public void setRecoveryPasswordToken(String recoveryPasswordToken) {
		this.recoveryPasswordToken = recoveryPasswordToken;
	}

	public LocalDate getCreatedDate() {
		return createdDate;
	}

	public void setCreatedDate(LocalDate createdDate) {
		this.createdDate = createdDate;
	}

	public Users getUsers() {
		return users;
	}

	public void setUsers(Users users) {
		this.users = users;
	}

}
