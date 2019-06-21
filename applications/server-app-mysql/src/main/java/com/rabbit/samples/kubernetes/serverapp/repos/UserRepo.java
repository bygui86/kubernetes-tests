package com.rabbit.samples.kubernetes.serverapp.repos;

import com.rabbit.samples.kubernetes.serverapp.domain.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;


/**
 * @author Matteo Baiguini
 */
@Repository
public interface UserRepo extends JpaRepository<User, Long> {

	User findByEmail(final String email);

}
