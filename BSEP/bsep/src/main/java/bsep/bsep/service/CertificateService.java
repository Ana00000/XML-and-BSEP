package bsep.bsep.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import bsep.bsep.model.Certificate;
import bsep.bsep.repository.ICertificateRepository;
import bsep.bsep.service.interfaces.ICertificateService;

@Service
public class CertificateService implements ICertificateService{
	
	private ICertificateRepository certificateRepository;

    @Autowired
    public CertificateService(ICertificateRepository certificateRepository) {
        this.certificateRepository = certificateRepository;
    }
    
    @Override
    public Certificate findOne(Long id) {
        return certificateRepository.findById(id).orElseGet(null);
    }

    @Override
    public List<Certificate> findAll() {
        return certificateRepository.findAll();
    }

    @Override
    public Certificate save(Certificate certificate) {
        return certificateRepository.save(certificate);
    }

       
}
