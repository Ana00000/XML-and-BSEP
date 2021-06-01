package bsep.bsep.service.interfaces;

import bsep.bsep.model.RecoverPasswordToken;

public interface IRecoverPasswordTokenService {
	RecoverPasswordToken findRecoverPasswordTokenByToken(String recoverPasswordToken);
	RecoverPasswordToken saveTokenAndSendEmailToUser(RecoverPasswordToken token);
}
