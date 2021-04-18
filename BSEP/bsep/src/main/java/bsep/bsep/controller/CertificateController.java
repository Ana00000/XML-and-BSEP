package bsep.bsep.controller;

import java.io.IOException;
import java.security.cert.CertificateEncodingException;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PostAuthorize;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

import bsep.bsep.dto.CertificateDTO;
import bsep.bsep.dto.CertificateInfoDTO;
import bsep.bsep.model.CertificateData;
import bsep.bsep.service.CertificateService;
import bsep.bsep.service.UserService;
import bsep.bsep.validation.CertificateValidation;

@Controller
@CrossOrigin(origins = "https://localhost:8081")
@RequestMapping(value = "/certificate", produces = MediaType.APPLICATION_JSON_VALUE)
public class CertificateController {

	private CertificateService certificateService;
	private UserService userService;
	private CertificateValidation certificateValidation;
	private Logger logger = LoggerFactory.getLogger(UserService.class);

	@Autowired
	public CertificateController(CertificateService certificateService, UserService userService) {
		this.certificateService = certificateService;
		this.userService = userService;
		this.certificateValidation = new CertificateValidation();
	}

	@PostAuthorize("hasAuthority('USER_REVOKE_CERTIFICATE_PRIVILEGE')")
	@PutMapping(value = "/revokeCertificate/{serialNumber}")
	public ResponseEntity<Boolean> revokeCertificate(@PathVariable String serialNumber) {
		try {
			certificateService.revokeCertificate(serialNumber);
			logger.info("action=revokeCertificate status=success certificateSerialNumber={}", serialNumber);
			return new ResponseEntity<>(HttpStatus.OK);
		} catch (Exception e) {
			logger.error("action=revokeCertificate status=failure message={} certificateSerialNumber={}",
					e.getMessage(), serialNumber);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}
	}

	@PostAuthorize("hasAuthority('USER_GET_CERTIFICATE_DTO_BY_SERIAL_NUMBER_PRIVILEGE')")
	@GetMapping(value = "/{serialNumber}")
	public ResponseEntity<CertificateData> checkCertificateValidity(@PathVariable String serialNumber) {
		logger.info("action=checkCertificateValidity status=success certificateSerialNumber={}", serialNumber);
		return new ResponseEntity<>(certificateService.findCertificateDataBySerialNumber(serialNumber), HttpStatus.OK);
	}

	@PostAuthorize("hasAuthority('USER_GET_CERTIFICATE_DTO_BY_SERIAL_NUMBER_PRIVILEGE')")
	@GetMapping(value = "/getCertificate/{serialNumber}")
	public ResponseEntity<CertificateDTO> getCertificateDTOBySerialNumber(@PathVariable String serialNumber) {
		logger.info("action=getCertificate status=success certificateSerialNumber={}", serialNumber);
		return new ResponseEntity<>(certificateService.findCertificateDTOBySerialNumber(serialNumber), HttpStatus.OK);
	}

	@PostAuthorize("hasAuthority('USER_ALL_VALID_CERTIFICATES_PRIVILEGE')")
	@GetMapping(value = "/allValid/{userEmail}")
	public ResponseEntity<List<CertificateDTO>> allValidCertificates(@PathVariable String userEmail) {
		logger.info("action=allValidCertificates status=success");
		return new ResponseEntity<>(getListOfCertificate(userEmail), HttpStatus.OK);
	}

	@PostAuthorize("hasAuthority('USER_ALL_REVOKED_OR_EXPIRED_CERTIFICATES_PRIVILEGE')")
	@GetMapping(value = "/allRevokedOrExpired")
	public ResponseEntity<List<CertificateDTO>> allRevokedOrExpiredCertificates() {
		logger.info("action=allRevokedOrExpired status=success");
		return new ResponseEntity<>(certificateService.findAllRevokedOrExpired(), HttpStatus.OK);
	}

	@PostAuthorize("hasAuthority('USER_GET_ALL_VALID_CERTIFICATES_DTO_PRIVILEGE')")
	@GetMapping(value = "/loadToFile/{serialNumber}")
	public ResponseEntity<Boolean> loadToFile(@PathVariable String serialNumber) {
		try {
			certificateService.loadCertificateToFile(serialNumber);
			logger.info("action=loadToFile status=success certificateSerialNumber={}", serialNumber);
			return new ResponseEntity<>(HttpStatus.OK);
		} catch (CertificateEncodingException | IOException e) {
			e.printStackTrace();
			logger.error("action=loadToFile status=failure message={} certificateSerialNumber={}", e.getMessage(),
					serialNumber);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		} catch (Exception e) {
			e.printStackTrace();
			logger.error("action=loadToFile status=failure message={} certificateSerialNumber={}", e.getMessage(),
					serialNumber);
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}
	}

	@PostAuthorize("hasAuthority('USER_CREATE_CERTIFICATE_PRIVILEGE')")
	@PostMapping(value = "/createCertificate", consumes = "application/json")
	public ResponseEntity<CertificateData> createCertificate(@RequestBody CertificateInfoDTO certificateInfoDTO) {
		String certificateAlias = certificateInfoDTO.getAlias();
		if (!certificateValidation.validCertificate(certificateInfoDTO)) {
			String message = "Certificate is not valid";
			logger.error("action=createCertificate status=failure message={} certificateAlias={}", message,
					certificateAlias);
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
		try {
			CertificateData certificateData = certificateService.createCertificate(certificateInfoDTO);
			logger.info("action=createCertificate status=success certificateAlias={}", certificateAlias);
			return new ResponseEntity<>(certificateData, checkStatusForCreatingCertificate(certificateData));
		} catch (Exception e) {
			e.printStackTrace();
			logger.error("action=createCertificate status=failure message={} certificateAlias={}", e.getMessage(),
					certificateAlias);
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
	}

	private List<CertificateDTO> getListOfCertificate(String userEmail) {
		String typeOfUser = findTypeByEmail(userEmail);
		return typeOfUser.equals("ADMIN") ? certificateService.findAllValid()
				: certificateService.findAllForUser(userEmail);
	}

	private String findTypeByEmail(String userEmail) {
		return userService.findByUserEmail(userEmail).getTypeOfUser().name();
	}

	private HttpStatus checkStatusForCreatingCertificate(CertificateData certificateData) {
		return certificateData != null ? HttpStatus.CREATED : HttpStatus.BAD_REQUEST;
	}
}
