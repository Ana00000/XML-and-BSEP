package bsep.bsep.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import bsep.bsep.model.Certificate;

public interface ICertificateRepository extends JpaRepository<Certificate, Long>{

}
