package bsep.bsep.service;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import bsep.bsep.model.RecoverPasswordToken;
import bsep.bsep.model.Users;
import bsep.bsep.repository.IRecoverPasswordTokenRepository;
import bsep.bsep.service.interfaces.IRecoverPasswordTokenService;

@Service
public class RecoverPasswordTokenService implements IRecoverPasswordTokenService {
	
	private IRecoverPasswordTokenRepository iRecoverPasswordTokenRepository;
	private EmailService emailService;
	
	private Logger logger = LoggerFactory.getLogger(UserService.class);

	@Autowired
	public RecoverPasswordTokenService(IRecoverPasswordTokenRepository iRecoverPasswordTokenRepository, EmailService emailService) {
		this.iRecoverPasswordTokenRepository = iRecoverPasswordTokenRepository;
		this.emailService = emailService;
	}

	@Override
	public RecoverPasswordToken findRecoverPasswordTokenByToken(String recoverPasswordToken) {
		for (RecoverPasswordToken tokenIt : iRecoverPasswordTokenRepository.findAll()) {
			if(tokenIt.getRecoveryPasswordToken().equals(recoverPasswordToken)) return tokenIt;
		}
		return null;
	}

	@Override
	public RecoverPasswordToken saveTokenAndSendEmailToUser(RecoverPasswordToken token) {
		RecoverPasswordToken recoverPasswordToken = iRecoverPasswordTokenRepository.save(token);
		sendRecoverPasswordEmail(recoverPasswordToken.getUsers(), recoverPasswordToken);
		return recoverPasswordToken;
	}
	
	private void sendRecoverPasswordEmail(Users user, RecoverPasswordToken recoverPasswordToken) {
		try {
			String supplierEmail = user.getUserEmail();
			String subject = "Recover password";
			String text = "Dear "+user.getFirstName()+",\n   You can recover your password by clicking on the link below: \n\n" + "http://localhost:8081/changePasswordByToken/"+ recoverPasswordToken.getRecoveryPasswordToken()+"\n Best regards,\nTeam BSEP";
			emailService.sendNotificaitionAsync(supplierEmail, subject, text);
			
			System.out.println("Email sent");
	
		}catch(Exception e) {
			logger.info("Error sending email: "+ e.getMessage());
		}
	}
}
