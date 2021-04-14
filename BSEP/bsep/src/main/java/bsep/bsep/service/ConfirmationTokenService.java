package bsep.bsep.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import bsep.bsep.model.ConfirmationToken;
import bsep.bsep.model.Users;
import bsep.bsep.repository.IConfirmationTokenRepository;
import bsep.bsep.service.interfaces.IConfirmationTokenService;

@Service
public class ConfirmationTokenService implements IConfirmationTokenService {

	private IConfirmationTokenRepository iConfirmationTokenRepository;
	
	@Autowired
	public ConfirmationTokenService(IConfirmationTokenRepository iConfirmationTokenRepository) {
		this.iConfirmationTokenRepository = iConfirmationTokenRepository;
	}
	
	@Override
	public ConfirmationToken findByConfirmationToken(String confirmationToken) {
		for (ConfirmationToken confirmationTokenIt : iConfirmationTokenRepository.findAll()) {
			if (confirmationToken.equals(confirmationTokenIt.getConfirmationToken())) return confirmationTokenIt;
		}
		return null;
	}

	@Override
	public ConfirmationToken save(Users user) {
		return iConfirmationTokenRepository.save(new ConfirmationToken(user));
	}

}
