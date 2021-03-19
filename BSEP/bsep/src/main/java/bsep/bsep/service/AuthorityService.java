package bsep.bsep.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import bsep.bsep.model.Authority;
import bsep.bsep.repository.IAuthorityRepository;
import bsep.bsep.service.interfaces.IAuthorityService;

@Service
public class AuthorityService implements IAuthorityService {

	@Autowired
	private IAuthorityRepository authorityRepository;

	@Override
	public Authority findById(Long id) {
		return authorityRepository.getOne(id);
	}

	@Override
	public Authority findByName(String name) {
		return authorityRepository.findByName(name);
	}
}