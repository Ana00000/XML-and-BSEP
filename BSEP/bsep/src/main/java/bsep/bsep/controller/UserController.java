package bsep.bsep.controller;

import java.util.ArrayList;
import java.util.List;

import javax.servlet.http.HttpServletResponse;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
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

import bsep.bsep.dto.EmailDTO;
import bsep.bsep.dto.TokenDTO;
import bsep.bsep.dto.UserChangePasswordDTO;
import bsep.bsep.dto.UserDTO;
import bsep.bsep.model.Authority;
import bsep.bsep.model.ConfirmationToken;
import bsep.bsep.model.RecoverPasswordToken;
import bsep.bsep.model.Users;
import bsep.bsep.security.TokenUtils;
import bsep.bsep.security.UserTokenState;
import bsep.bsep.service.AuthorityService;
import bsep.bsep.service.ConfirmationTokenService;
import bsep.bsep.service.RecoverPasswordTokenService;
import bsep.bsep.service.UserService;
import bsep.bsep.validation.UserValidation;

@RestController
@CrossOrigin(origins = "https://localhost:8081")
@RequestMapping(value = "/users", produces = MediaType.APPLICATION_JSON_VALUE)
public class UserController {

	@Autowired
	private TokenUtils tokenUtils;

	@Autowired
	private AuthenticationManager authenticationManager;

	private final UserService userService;

	private final AuthorityService authorityService;

	private UserValidation userValidation;

	private final RecoverPasswordTokenService recoverPasswordTokenService;

	private final ConfirmationTokenService confirmationTokenService;

	private Logger logger = LoggerFactory.getLogger(UserService.class);

	@Autowired
	public UserController(UserService userService, AuthorityService authorityService,
			ConfirmationTokenService confirmationTokenService,
			RecoverPasswordTokenService recoverPasswordTokenService) {
		this.userService = userService;
		this.authorityService = authorityService;
		this.recoverPasswordTokenService = recoverPasswordTokenService;
		this.confirmationTokenService = confirmationTokenService;
		this.userValidation = new UserValidation();

	}

	@GetMapping("/findAllUsers")
	public ResponseEntity<List<Users>> findAllUsers() {
		logger.info("action=findAllUsers status=success");
		return new ResponseEntity<>(userService.findAll(), HttpStatus.OK);
	}

	@GetMapping("/getUsersEmails")
	public ResponseEntity<List<String>> findAllUsersEmails() {
		logger.info("action=getUsersEmails status=success");
		return new ResponseEntity<>(userService.findAllUsersEmails(), HttpStatus.OK);
	}

	@GetMapping("/redirectMeToMyHomePage")
	public String RedirectionToHome() {
		logger.info("action=redirectMeToMyHomePage status=success");
		return "https://localhost:8081/";
	}

	@PostMapping("/recoverPasswordWithToken")
	public ResponseEntity<Boolean> recoveringPassword(@RequestBody EmailDTO recoveryPasswordRequestEmail) {
		String userEmail = recoveryPasswordRequestEmail.getEmailOfUser();
		if (!userValidation.validUserEmail(userEmail)) {
			String message = "User email is not valid";
			logger.error("action=recoverPasswordWithToken status=faliure message={} userEmail={}", message, userEmail);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}
		Users user = userService.findByUserEmail(userEmail);
		if (user == null || !user.isConfirmed()) {
			String message = "User not found or not confirmed";
			logger.error("action=recoverPasswordWithToken status=faliure message={} userEmail={}", message, userEmail);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}

		recoverPasswordTokenService.saveTokenAndSendEmailToUser(new RecoverPasswordToken(user));
		logger.info("action=recoverPasswordWithToken status=success userEmail={}", userEmail);
		return new ResponseEntity<>(HttpStatus.OK);
	}

	@PostMapping("/findUserWithToken")
	public ResponseEntity<Users> findUserByToken(@RequestBody TokenDTO token) {
		RecoverPasswordToken recoverPasswordToken = recoverPasswordTokenService
				.findRecoverPasswordTokenByToken(token.getToken());
		if (recoverPasswordToken != null && recoverPasswordToken.getUsers() != null) {
			logger.info("action=findUserByToken status=success userEmail={}",
					recoverPasswordToken.getUsers().getUserEmail());
			return new ResponseEntity<>(recoverPasswordToken.getUsers(), HttpStatus.OK);
		}

		String message = "Token not found";
		logger.error("action=findUserByToken status=failure message={}", message);
		return new ResponseEntity<>(null, HttpStatus.NOT_FOUND);
	}

	@PutMapping(value = "/changePassword", consumes = "application/json")
	public ResponseEntity<Boolean> changePassword(@RequestBody UserChangePasswordDTO userChangePasswordDTO) {
		String userEmail = userChangePasswordDTO.getEmailOfUser();
		if (!userValidation.validUserEmail(userEmail)
				|| !userValidation.validPassword(userChangePasswordDTO.getPassword())
				|| !userValidation.validPassword(userChangePasswordDTO.getConfirmedPassword())
				|| !userValidation.validPasswordAndConfirmPassword(userChangePasswordDTO)) {
			String message = "Bad credentials";
			logger.error("action=changePassword status=failure message={} userEmail={}", message, userEmail);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}
		try {
			Users users = userService.findByUserEmail(userChangePasswordDTO.getEmailOfUser());
			if (users == null) {
				String message = "User not found";
				logger.error("action=changePassword status=failure message={} userEmail={}", message, userEmail);
				return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
			}
			users.setPassword(userChangePasswordDTO.getPassword());
			userService.updatePassword(users);
			logger.info("action=changePassword status=success userEmail={}", userEmail);
			return new ResponseEntity<>(HttpStatus.OK);
		} catch (Exception e) {

			logger.error("action=changePassword status=failure message={} userEmail={}", e.getMessage(), userEmail);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}

	}

	@PostMapping("/login")
	public ResponseEntity<UserTokenState> login(@RequestBody UserDTO authenticationRequest,
			HttpServletResponse response) {
		String userEmail = authenticationRequest.getUserEmail();
		try {
			Users userLogIn = userService.login(authenticationRequest);
			StringBuilder passwordWithSalt = new StringBuilder();
			passwordWithSalt.append(authenticationRequest.getPassword());
			passwordWithSalt.append(userLogIn.getSalt());

			Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
					authenticationRequest.getUserEmail(), passwordWithSalt.toString()));

			SecurityContextHolder.getContext().setAuthentication(authentication);
			Users user = (Users) authentication.getPrincipal();

			if (user.isConfirmed()) {
				String jwt = tokenUtils.generateToken(userEmail);
				logger.info("action=login status=success userEmail={}", userEmail);
				return ResponseEntity
						.ok(new UserTokenState(jwt, tokenUtils.getExpiredIn(), user.getTypeOfUser().name()));
			}
		} catch (Exception e) {
			logger.error("action=login status=failure message={} userEmail={}", e.getMessage(), userEmail);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}

		String message = "Bad request";
		logger.error("action=login status=failure message={} userEmail={}", message, userEmail);
		return new ResponseEntity<>(HttpStatus.BAD_REQUEST);

	}

	@PostMapping(value = "/register", consumes = "application/json")
	public ResponseEntity<Users> addUser(@RequestBody UserDTO userRequest) {
		String userEmail = userRequest.getUserEmail();
		if (userRequest.getTypeOfUser().toUpperCase().equals("ADMIN") || !userValidation.validUser(userRequest)) {
			String message = "User info is not valid";
			logger.error("action=register status=failure message={}", message);
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		} else if (userService.findByUserEmail(userEmail) != null) {
			System.out.println("Username already exists.");
			String message = "User email already exists";
			logger.error("action=register status=failure message={} userEmail={}", message, userEmail);
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}

		try {
			Users userWithPermissions = addPermissionsForUser(userRequest);
			Users userRegistered = userService.save(userWithPermissions);
			logger.info("action=register status=success userEmail={}", userEmail);
			return new ResponseEntity<>(userRegistered, HttpStatus.CREATED);
		} catch (Exception e) {
			e.printStackTrace();
			logger.error("action=register status=failure message={} userEmail={}", e.getMessage(), userEmail);
		}
		String message = "Bad request";
		logger.error("action=register status=failure message={} userEmail={}", message, userEmail);
		return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
	}

	@PutMapping(value = "/confirm_account/{token}", consumes = "application/json")
	public ResponseEntity<Boolean> confirmAccount(@PathVariable String token) {
		try {

			ConfirmationToken confirmationToken = confirmationTokenService.findByConfirmationToken(token);
			if (confirmationToken != null) {
				setConfirmedAccount(confirmationToken);
				logger.info("action=confirmAccount status=success userEmail={}",
						confirmationToken.getUsers().getUserEmail());
				return new ResponseEntity<>(HttpStatus.OK);

			} else {
				String message = "Confirmation token is not valid";
				logger.error("action=confirmAccount status=failure message={} token={}", message, token);
				return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
			}

		} catch (Exception e) {

			logger.error("action=confirmAccount status=failure message={} token={}", e.getMessage(), token);
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
