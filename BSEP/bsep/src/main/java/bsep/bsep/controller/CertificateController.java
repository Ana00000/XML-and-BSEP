package bsep.bsep.controller;

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
	public ResponseEntity<List<CertificateDTO>> getAllValidCertificatesDTO() {

		return new ResponseEntity<>(certificateService.findAll(), HttpStatus.OK);
	}

	@PostMapping(value = "/createCertificate", consumes = "application/json")
	public ResponseEntity<CertificateData> createCertificate(@RequestBody CertificateInfoDTO certificateInfoDTO) {
		try {
			return new ResponseEntity<>(certificateService.createCertificate(certificateInfoDTO), HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
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
