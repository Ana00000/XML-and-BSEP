package bsep.bsep.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import bsep.bsep.model.RecoverPasswordToken;

public interface IRecoverPasswordTokenRepository extends JpaRepository<RecoverPasswordToken, Long> {

}
