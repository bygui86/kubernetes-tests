package com.rabbit.samples.kubernetes.serverapp.repos;

import com.rabbit.samples.kubernetes.serverapp.domain.User;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;


/**
 * @author Matteo Baiguini
 */
@Repository
public interface UserRepo extends MongoRepository<User, String> {

	User findByEmail(final String email);

}
