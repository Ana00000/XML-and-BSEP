package bsep.bsep.service;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import bsep.bsep.dto.UserDTO;
import bsep.bsep.model.UserType;
import bsep.bsep.model.Users;
import bsep.bsep.repository.IUserRepository;
import bsep.bsep.service.interfaces.IUserService;
import ch.qos.logback.core.subst.Token.Type;

@Service
public class UserService implements IUserService {

	private final IUserRepository userRepository;

	@Autowired
	private PasswordEncoder passwordEncoder;

	@Autowired
	public UserService(IUserRepository userRepository) {
		this.userRepository = userRepository;
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
	
	@Override
	public Users save(Users user) {
		user.setPassword(passwordEncoder.encode(user.getPassword()));
		System.out.print("Usao u servis :");
		return userRepository.save(user);
	}

	public List<String> findAllUsersEmails() {
		// TODO Auto-generated method stub
		List<String> returnValues = new ArrayList<String>();
		for (Users user : findAll()) {
			if (user.getTypeOfUser()!=UserType.ADMIN)
				returnValues.add(user.getUserEmail());
		}
		return returnValues;
	}
}