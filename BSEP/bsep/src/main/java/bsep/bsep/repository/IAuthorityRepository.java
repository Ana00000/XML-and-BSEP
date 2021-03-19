package bsep.bsep.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import bsep.bsep.model.Authority;

public interface IAuthorityRepository extends JpaRepository<Authority, Long> {
	Authority findByName(String name);
}