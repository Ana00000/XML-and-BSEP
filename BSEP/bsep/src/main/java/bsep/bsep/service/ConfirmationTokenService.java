package bsep.bsep.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import bsep.bsep.model.ConfirmationToken;
import bsep.bsep.model.Users;
import bsep.bsep.repository.IConfirmationTokenRepository;
import bsep.bsep.service.interfaces.IConfirmationTokenService;

@Service
public class ConfirmationTokenService implements IConfirmationTokenService {

	private final IConfirmationTokenRepository iConfirmationTokenRepository;

	@Autowired
	public ConfirmationTokenService(IConfirmationTokenRepository iConfirmationTokenRepository) {
		this.iConfirmationTokenRepository = iConfirmationTokenRepository;
	}

	public List<ConfirmationToken> findAll() {
		return iConfirmationTokenRepository.findAll();
	}

	@Override
	public ConfirmationToken findByConfirmationToken(String confirmationToken) {

		for (ConfirmationToken confirmationTokenIt : findAll()) {

			if (confirmationToken.equals(confirmationTokenIt.getConfirmationToken())) {

				return confirmationTokenIt;
			}
		}

		return null;
	}

	@Override
	public ConfirmationToken save(Users user) {
		return iConfirmationTokenRepository.save(new ConfirmationToken(user));
	}

}
