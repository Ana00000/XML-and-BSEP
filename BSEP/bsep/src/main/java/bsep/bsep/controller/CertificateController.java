package bsep.bsep.controller;

import java.io.IOException;
import java.security.cert.CertificateEncodingException;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
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
@CrossOrigin(origins = "", allowedHeaders = "")
@RequestMapping(value = "/certificate")
public class CertificateController {

	private CertificateService certificateService;

	
	@Autowired
	public CertificateController(CertificateService certificateService) {
		this.certificateService = certificateService;
	}

	@GetMapping(value = "/allValid")
	public ResponseEntity<List<CertificateDTO>> allValidCertificates() {
	
		return new ResponseEntity<>(certificateService.findAllValid(), HttpStatus.OK);
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
			// TODO Auto-generated catch block
			e.printStackTrace();
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		} catch (Exception e) {
			// TODO Auto-generated catch block
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
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
	}

	private HttpStatus checkStatusForCreatingCertificate(CertificateData certificateData) {
		return certificateData != null ? HttpStatus.CREATED : HttpStatus.BAD_REQUEST;
	}

	@PutMapping(value = "/revokeCertificate/{serialNumber}", consumes = "application/json")
	public ResponseEntity<Boolean> revokeCertificate(@PathVariable String serialNumber) {
		try {
			certificateService.revokeCertificate(serialNumber);
			return new ResponseEntity<>(HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
		}

	}
}
