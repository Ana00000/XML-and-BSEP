package bsep.bsep.service.interfaces;

import bsep.bsep.model.RecoverPasswordToken;
import bsep.bsep.model.Users;

public interface IRecoverPasswordTokenService {
	RecoverPasswordToken findRecoverPasswordTokenByToken(String recoverPasswordToken);
	RecoverPasswordToken saveTokenAndSendEmailToUser(RecoverPasswordToken token);
}
