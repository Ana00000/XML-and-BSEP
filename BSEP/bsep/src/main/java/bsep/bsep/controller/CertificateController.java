package bsep.bsep.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
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

	@GetMapping(value = "/all")
	public ResponseEntity<List<CertificateDTO>> getAllValidCertificatesDTO() {

		return new ResponseEntity<>(certificateService.findAll(), HttpStatus.OK);
	}

	@PostMapping(value = "/createRootCertificate", consumes = "application/json")
	public ResponseEntity<CertificateData> createRootCertificate(@RequestBody CertificateInfoDTO certificateInfoDTO) {
		try {
			return new ResponseEntity<>(certificateService.createRootCertificate(certificateInfoDTO), HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
	}

	
	
}
