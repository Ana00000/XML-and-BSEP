package bsep.bsep.repository;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;

import bsep.bsep.model.Users;

public interface IUserRepository extends JpaRepository<Users, Long> {

	Users findOneByPassword(String password);

	List<Users> findAllByFirstName(String firstName);

	List<Users> findAllByLastName(String lastName);

	List<Users> findByFirstNameAndLastNameAllIgnoringCase(String firstName, String lastName);

	Users findByUserEmail(String userEmail);

	Users findByUserEmailAndPassword(String userEmail, String password);
}