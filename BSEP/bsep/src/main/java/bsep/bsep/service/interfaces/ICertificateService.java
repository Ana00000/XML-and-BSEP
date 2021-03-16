package bsep.bsep.service.interfaces;

import java.util.List;

import bsep.bsep.dto.CertificateDTO;
import bsep.bsep.model.CertificateData;

public interface ICertificateService {
	
	 CertificateData findOne(Long id);
	 List<CertificateDTO> findAll();
	 
	
}
