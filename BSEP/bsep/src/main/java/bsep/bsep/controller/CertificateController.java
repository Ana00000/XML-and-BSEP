package bsep.bsep.controller;

import java.util.ArrayList;
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
import bsep.bsep.model.Certificate;
import bsep.bsep.service.CertificateService;

@Controller
@CrossOrigin(origins = "", allowedHeaders = "")
@RequestMapping(value = "/certificate")
public class CertificateController {

	private final CertificateService certificateService;

	@Autowired
	public CertificateController(CertificateService certificateService) {
		this.certificateService = certificateService;
	}

	@GetMapping(value = "/all")
	public ResponseEntity<List<CertificateDTO>> getAllCertificates() {

		return new ResponseEntity<>(getAllCertificatesDTO(), HttpStatus.OK);
	}

	@PostMapping(value = "/create", consumes = "application/json")
	public ResponseEntity<CertificateDTO> createCertificate(@RequestBody CertificateDTO certificateDTO) {
		try {
			certificateService.save(new Certificate(certificateDTO));
			return new ResponseEntity<>(certificateDTO, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
		}
	}

	private List<CertificateDTO> getAllCertificatesDTO() {

		List<CertificateDTO> certificatesDTO = new ArrayList<>();
		for (Certificate certificate : certificateService.findAll()) {
			certificatesDTO.add(new CertificateDTO(certificate));
		}

		return certificatesDTO;
	}
}
