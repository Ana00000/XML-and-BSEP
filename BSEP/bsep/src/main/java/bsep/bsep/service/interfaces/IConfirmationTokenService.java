package bsep.bsep.service.interfaces;

import bsep.bsep.model.ConfirmationToken;
import bsep.bsep.model.Users;

public interface IConfirmationTokenService {
	ConfirmationToken findByConfirmationToken(String confirmationToken);
	ConfirmationToken save(Users user);
}
