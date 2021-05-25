package bsep.bsep.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import bsep.bsep.model.CertificateData;

public interface ICertificateRepository extends JpaRepository<CertificateData, Long> {

}
