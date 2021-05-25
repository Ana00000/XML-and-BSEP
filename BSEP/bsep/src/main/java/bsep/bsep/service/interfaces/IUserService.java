package bsep.bsep.service.interfaces;

import java.util.List;

import bsep.bsep.dto.UserDTO;
import bsep.bsep.model.Users;

public interface IUserService {

	Users findOne(Long id);

	Users findOneByPassword(String password);

	List<Users> findAll();

	List<Users> findAllByFirstName(String firstName);

	List<Users> findAllByLastName(String lastName);

	Users findByUserEmail(String email);

	List<Users> findByFirstNameAndLastNameAllIgnoringCase(String firstName, String lastName);

	Users save(Users user);
	
	Users update(Users user);

	Users login(UserDTO userDTO);
}
