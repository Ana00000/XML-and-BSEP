package bsep.bsep.service.interfaces;

import java.util.List;

import bsep.bsep.model.Certificate;

public interface ICertificateService {
	
	Certificate findOne(Long id);
	List<Certificate> findAll();
	Certificate save(Certificate certificate);
	 
	
}
