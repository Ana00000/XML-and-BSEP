package bsep.bsep.controller;

import java.util.ArrayList;
import java.util.List;

import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import bsep.bsep.dto.UserDTO;
import bsep.bsep.model.Authority;
import bsep.bsep.model.ConfirmationToken;
import bsep.bsep.model.Users;
import bsep.bsep.security.ResourceConflictException;
import bsep.bsep.security.TokenUtils;
import bsep.bsep.security.UserTokenState;
import bsep.bsep.service.AuthorityService;
import bsep.bsep.service.ConfirmationTokenService;
import bsep.bsep.service.UserService;

@RestController
@CrossOrigin(origins = "http://localhost:8081")
@RequestMapping(value = "/users", produces = MediaType.APPLICATION_JSON_VALUE)
public class UserController {

	@Autowired
	private TokenUtils tokenUtils;

	@Autowired
	private AuthenticationManager authenticationManager;

	private final UserService userService;

	private final AuthorityService authorityService;

	private final ConfirmationTokenService confirmationTokenService;

	@Autowired
	public UserController(UserService userService, AuthorityService authorityService,
			ConfirmationTokenService confirmationTokenService) {
		this.userService = userService;
		this.authorityService = authorityService;
		this.confirmationTokenService = confirmationTokenService;
	}

	@GetMapping("/findAll")
	public ResponseEntity<List<Users>> findAll() {
		return new ResponseEntity<>(userService.findAll(), HttpStatus.OK);
	}

	@GetMapping("/getUsersEmails")
	public ResponseEntity<List<String>> findAllUsersEmails() {
		return new ResponseEntity<>(userService.findAllUsersEmails(), HttpStatus.OK);
	}

	@GetMapping("/redirectMeToMyHomePage")
	public String RedirectionToHome() {
		return "http://localhost:8081/";
	}

	@PostMapping("/login")
	public ResponseEntity<UserTokenState> login(@RequestBody UserDTO authenticationRequest,
			HttpServletResponse response) {

		Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
				authenticationRequest.getUserEmail(), authenticationRequest.getPassword()));

		SecurityContextHolder.getContext().setAuthentication(authentication);
		Users user = (Users) authentication.getPrincipal();
		
		if(user.isConfirmed())
		{
			String jwt = tokenUtils.generateToken(user.getUserEmail());
			return ResponseEntity.ok(new UserTokenState(jwt, tokenUtils.getExpiredIn(), user.getTypeOfUser().name()));
		}
		
		System.out.println("bad request");
		return new ResponseEntity<>(HttpStatus.BAD_REQUEST);

	}

	@PostMapping(value = "/register", consumes = "application/json")
	public ResponseEntity<Users> addUser(@RequestBody UserDTO userRequest) {
		Users existUser;
		if (userRequest.getTypeOfUser().toUpperCase().equals("ADMIN")) {
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
		try {
			existUser = userService.findByUserEmail(userRequest.getUserEmail());

			if (existUser != null) {
				throw new ResourceConflictException(existUser.getId(), "Username already exists");
			}
			Users userWithPermissions = addPermissionsForUser(userRequest);
			Users userRegistered = userService.save(userWithPermissions);
			return new ResponseEntity<>(userRegistered, HttpStatus.CREATED);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
	}

	@PutMapping(value = "/confirm_account/{token}", consumes = "application/json")
	public ResponseEntity<Boolean> confirmAccount(@PathVariable String token) {
		try {

			ConfirmationToken confirmationToken = confirmationTokenService.findByConfirmationToken(token);
			if (confirmationToken != null) {
				setConfirmedAccount(confirmationToken);
				return new ResponseEntity<>(HttpStatus.OK);
				
			} else {
				return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
			}

		} catch (Exception e) {

			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}

	}
	
	private void setConfirmedAccount(ConfirmationToken confirmationToken) {
		Users users = userService.findByUserEmail(confirmationToken.getUsers().getUserEmail());
		users.setConfirmed(true);
		userService.update(users);
	}

	private Users addPermissionsForUser(UserDTO userRequest) {
		Users userNew = new Users(userRequest);
		List<Authority> listOfAuthorities = new ArrayList<Authority>();
		listOfAuthorities.add(authorityService.findByName("USER_GET_CERTIFICATE_DTO_BY_SERIAL_NUMBER_PRIVILEGE"));
		listOfAuthorities.add(authorityService.findByName("USER_ALL_VALID_CERTIFICATES_PRIVILEGE"));
		listOfAuthorities.add(authorityService.findByName("USER_GET_ALL_VALID_CERTIFICATES_DTO_PRIVILEGE"));
		userNew.setAuthorities(listOfAuthorities);
		return userNew;
	}
}
