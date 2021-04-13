package bsep.bsep.controller;

import java.io.Console;
import java.io.IOException;
import java.security.cert.CertificateEncodingException;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
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

@Controller
@CrossOrigin(origins = "http://localhost:8081")
@RequestMapping(value = "/certificate", produces = MediaType.APPLICATION_JSON_VALUE)
public class CertificateController {

	private CertificateService certificateService;
	private UserService userService;

	@Autowired
	public CertificateController(CertificateService certificateService, UserService userService) {
		this.certificateService = certificateService;
		this.userService = userService;
	}

	
	@CrossOrigin(origins = "/*")	
	@PutMapping(value = "/revokeCertificate/{serialNumber}")
	public ResponseEntity<Boolean> revokeCertificate(@PathVariable String serialNumber) {
		try {
			certificateService.revokeCertificate(serialNumber);
			return new ResponseEntity<>(HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}

	}
	
	@GetMapping(value = "/getCertificate/{serialNumber}")
	public ResponseEntity<CertificateDTO> getCertificateDTOBySerialNumber(@PathVariable String serialNumber) {
		return new ResponseEntity<>(certificateService.findCertificateDTOBySerialNumber(serialNumber), HttpStatus.OK);
	}

	@GetMapping(value = "/allValid/{userEmail}")
	public ResponseEntity<List<CertificateDTO>> allValidCertificates(@PathVariable String userEmail) {
		return new ResponseEntity<>(getListOfCertificate(userEmail), HttpStatus.OK);
	}
	
	private List<CertificateDTO> getListOfCertificate(String userEmail){
		String typeOfUser = findTypeByEmail(userEmail);
		return typeOfUser.equals("ADMIN")? certificateService.findAllValid() : certificateService.findAllForUser(userEmail); 
	}

	private String findTypeByEmail(String userEmail) {
		return userService.findByUserEmail(userEmail).getTypeOfUser().name();
	}

	@GetMapping(value = "/allRevokedOrExpired")
	public ResponseEntity<List<CertificateDTO>> allRevokedOrExpiredCertificates() {

		return new ResponseEntity<>(certificateService.findAllRevokedOrExpired(), HttpStatus.OK);
	}

	@GetMapping(value = "/loadToFile/{serialNumber}")
	public ResponseEntity<Boolean> getAllValidCertificatesDTO(@PathVariable String serialNumber) {
		try {
			certificateService.loadCertificateToFile(serialNumber);
			return new ResponseEntity<>(HttpStatus.OK);
		} catch (CertificateEncodingException | IOException e) {
			e.printStackTrace();
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		} catch (Exception e) {
			e.printStackTrace();
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}
	}

	@PostMapping(value = "/createCertificate", consumes = "application/json")
	public ResponseEntity<CertificateData> createCertificate(@RequestBody CertificateInfoDTO certificateInfoDTO) {
		try {
			CertificateData certificateData = certificateService.createCertificate(certificateInfoDTO);
			return new ResponseEntity<>(certificateData, checkStatusForCreatingCertificate(certificateData));
		} catch (Exception e) {
			e.printStackTrace();
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
	}

	private HttpStatus checkStatusForCreatingCertificate(CertificateData certificateData) {
		return certificateData != null ? HttpStatus.CREATED : HttpStatus.BAD_REQUEST;
	}
}
