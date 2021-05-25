package bsep.bsep.service.interfaces;

import bsep.bsep.model.Authority;

public interface IAuthorityService {
    Authority findById(Long id);
    Authority findByName(String name);
}