package bsep.bsep.service;

import java.util.ArrayList;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import bsep.bsep.dto.UserDTO;
import bsep.bsep.model.ConfirmationToken;
import bsep.bsep.model.UserType;
import bsep.bsep.model.Users;
import bsep.bsep.repository.IUserRepository;
import bsep.bsep.service.interfaces.IUserService;

@Service
public class UserService implements IUserService {

	private final IUserRepository userRepository;

	@Autowired
	private PasswordEncoder passwordEncoder;
	
	private EmailService emailService;
	
	private final ConfirmationTokenService confirmationTokenService;
	
	private Logger logger = LoggerFactory.getLogger(UserService.class);

	@Autowired
	public UserService(IUserRepository userRepository, ConfirmationTokenService confirmationTokenService,EmailService emailService) {
		this.userRepository = userRepository;
		this.confirmationTokenService = confirmationTokenService;
		this.emailService = emailService;
	}

	public Users findOne(Long id) {
		List<Users> users = findAll();
		for (Users user : users) {
			if (user.getId() == id) {
				return user;
			}
		}
		return null;
	}

	public Users findOneByPassword(String password) {
		List<Users> users = findAll();
		for (Users user : users) {
			if (user.getPassword().equals(password)) {
				return user;
			}
		}
		return null;
	}

	public List<Users> findAll() {
		return userRepository.findAll();
	}

	public List<Users> findAllByFirstName(String firstName) {
		List<Users> users = findAll();
		List<Users> retList = new ArrayList<Users>();
		for (Users user : users) {
			if (user.getFirstName().toLowerCase().equals(firstName.toLowerCase())) {
				retList.add(user);
			}
		}
		return retList;
	}

	public List<Users> findAllByLastName(String lastName) {
		List<Users> users = findAll();
		List<Users> retList = new ArrayList<Users>();
		for (Users user : users) {
			if (user.getLastName().toLowerCase().equals(lastName.toLowerCase())) {
				retList.add(user);
			}
		}
		return retList;
	}

	public Users findByUserEmail(String userEmail) {
		List<Users> users = findAll();
		for (Users user : users) {
			if (user.getUserEmail().equals(userEmail)) {
				return user;
			}
		}
		return null;
	}

	public List<Users> findByFirstNameAndLastNameAllIgnoringCase(String firstName, String lastName) {
		List<Users> users = findAll();
		List<Users> retList = new ArrayList<Users>();
		for (Users user : users) {
			if (user.getLastName().toLowerCase().equals(lastName.toLowerCase())
					&& user.getFirstName().toLowerCase().equals(firstName.toLowerCase())) {
				retList.add(user);
			}
		}
		return retList;
	}

	public void remove(Long id) {
		userRepository.deleteById(id);
	}

	@Override
	public Users login(UserDTO userDTO) {
		List<Users> users = findAll();
		for (Users user : users)
			if (user.getUserEmail().equals(userDTO.getUserEmail()) && user.getPassword().equals(userDTO.getPassword()))
				return user;

		return null;
	}
	
	public Users update(Users user) {
		return userRepository.save(user);
	}
	
	@Override
	public Users save(Users user) {
		user.setPassword(passwordEncoder.encode(user.getPassword()));
		user.setConfirmed(false);
		Users userNew = userRepository.save(user);
		ConfirmationToken confirmationToken = confirmationTokenService.save(userNew);
		sendConfirmationEmail(userNew, confirmationToken);
		return userNew; 
	}
	
	private void sendConfirmationEmail(Users user, ConfirmationToken confirmationToken) {
		try {
			
			String supplierEmail = user.getUserEmail();
			String subject = "Confirm registration";
			String text = "Please confirm your registration by clicking the link below \n\n" + "http://localhost:8081/confirmRegistration/"+ confirmationToken.getConfirmationToken();
			emailService.sendNotificaitionAsync(supplierEmail, subject, text);
			
			System.out.println("Email sent");
	
		}catch(Exception e) {
			logger.info("Error sending email: "+ e.getMessage());
		}
	}

	public List<String> findAllUsersEmails() {
		List<String> returnValues = new ArrayList<String>();
		for (Users user : findAll()) {
			if (user.getTypeOfUser()!=UserType.ADMIN)
				returnValues.add(user.getUserEmail());
		}
		return returnValues;
	}
}