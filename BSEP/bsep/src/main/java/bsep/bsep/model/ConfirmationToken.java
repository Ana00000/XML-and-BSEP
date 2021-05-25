package bsep.bsep.model;

import java.time.LocalDate;
import java.util.UUID;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.JoinColumn;
import javax.persistence.OneToOne;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

import javax.persistence.Id;

import com.fasterxml.jackson.annotation.JsonFormat;

@Entity
@Table(name = "confirmation_tokens")
public class ConfirmationToken {
	
	@Id
	@SequenceGenerator(name = "mySeqGenConfirmationToken", sequenceName = "mySeqConfirmationToken", initialValue = 1, allocationSize = 1)
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "mySeqGenConfirmationToken")
	@Column(name = "id", unique = true, nullable = false)
    private Long id;

    @Column(name="confirmation_token")
    private String confirmationToken;

    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern ="yyyy-MM-dd", timezone = "UTC")
    private LocalDate createdDate;

    @OneToOne(targetEntity = Users.class, fetch = FetchType.EAGER)
    @JoinColumn(nullable = false, name = "user_id")
    private Users users;
    
    public ConfirmationToken() {}
    
    

    public ConfirmationToken(Long id, String confirmationToken, LocalDate createdDate, Users users) {
		super();
		this.id = id;
		this.confirmationToken = confirmationToken;
		this.createdDate = createdDate;
		this.users = users;
	}



	public ConfirmationToken(Users users) {
        this.users = users;
        createdDate = LocalDate.now();
        confirmationToken = UUID.randomUUID().toString();
    }

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public String getConfirmationToken() {
		return confirmationToken;
	}

	public void setConfirmationToken(String confirmationToken) {
		this.confirmationToken = confirmationToken;
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
