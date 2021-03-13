package model;

import java.time.LocalDateTime;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;

@Entity
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
	
	@Column(name = "subject", unique = false, nullable = false)
	private Subject subject;
	
	@Column(name = "version", unique = false, nullable = false)
	private Issuer issuer;
	
}
