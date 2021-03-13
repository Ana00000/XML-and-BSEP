package bsep.bsep.model;

import java.time.LocalDateTime;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;

@Entity
@Table(name="certificate")
public class Certificate {

	@Id
	@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "certificateIdSeqGen")
	@SequenceGenerator(name = "certificateIdSeqGen", sequenceName = "certificateIdSeq", initialValue = 1, allocationSize = 1)
	@Column(name = "id", unique = true, nullable = false)
	private Long id;

	@Column(name = "version", unique = false, nullable = false)
	private String version;

	@Column(name = "start", unique = false, nullable = false)
	private LocalDateTime start;

	@Column(name = "end", unique = false, nullable = false)
	private LocalDateTime end;
/*
	@Column(name = "subject", unique = false, nullable = false)
	private Subject subject;

	@Column(name = "issuer", unique = false, nullable = false)
	private Issuer issuer;
*/
	public Certificate() {
		super();
	}

	public Certificate(Long id, String version, LocalDateTime start, LocalDateTime end, Subject subject,
			Issuer issuer) {
		super();
		this.id = id;
		this.version = version;
		this.start = start;
		this.end = end;
		//this.subject = subject;
		//this.issuer = issuer;
	}

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public String getVersion() {
		return version;
	}

	public void setVersion(String version) {
		this.version = version;
	}

	public LocalDateTime getStart() {
		return start;
	}

	public void setStart(LocalDateTime start) {
		this.start = start;
	}

	public LocalDateTime getEnd() {
		return end;
	}

	public void setEnd(LocalDateTime end) {
		this.end = end;
	}
/*
	public Subject getSubject() {
		return subject;
	}

	public void setSubject(Subject subject) {
		this.subject = subject;
	}

	public Issuer getIssuer() {
		return issuer;
	}

	public void setIssuer(Issuer issuer) {
		this.issuer = issuer;
	}
*/
	@Override
	public int hashCode() {
		// TODO Auto-generated method stub
		return super.hashCode();
	}

	@Override
	public String toString() {
		// TODO Auto-generated method stub
		return super.toString();
	}

}
