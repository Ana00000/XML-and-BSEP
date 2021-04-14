package bsep.bsep.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import bsep.bsep.model.ConfirmationToken;

public interface IConfirmationTokenRepository extends JpaRepository<ConfirmationToken, Long> {
}
