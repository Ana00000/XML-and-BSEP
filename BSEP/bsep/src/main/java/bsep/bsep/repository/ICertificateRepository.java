package bsep.bsep.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import bsep.bsep.model.CertificateData;

@Repository
public interface ICertificateRepository extends JpaRepository<CertificateData, Long>{

}
