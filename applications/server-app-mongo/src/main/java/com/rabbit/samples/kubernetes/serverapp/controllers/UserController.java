package com.rabbit.samples.kubernetes.serverapp.controllers;


import com.rabbit.samples.kubernetes.serverapp.repos.UserRepo;
import com.rabbit.samples.kubernetes.serverapp.domain.User;
import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.experimental.FieldDefaults;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;


/**
 * @author Matteo Baiguini
 */
@Slf4j
@FieldDefaults(level = AccessLevel.PRIVATE)
@AllArgsConstructor
@Getter(value = AccessLevel.PRIVATE)
@RestController
@RequestMapping("/users")
public class UserController {

	static final String USER_NOT_FOUND_ERR_MSG = "User not found";

	UserRepo userRepo;

	@GetMapping
	public List<User> getAll() {

		log.info("Get all");
		List<User> users = getUserRepo().findAll();
		log.info("All users found: {}", users.size());
		return users;
	}

	@GetMapping("/{email}")
	public User getByEmail(@PathVariable final String email) {

		log.info("Get by email: {}", email);
		User user = getUserRepo().findByEmail(email);
		if(user == null) {
			user = buildErrorUserByEmail(email);
		}
		log.info("User by email: {}", user);
		return user;
	}

	@PostMapping
	public User insert(@RequestBody final User user) {

		log.info("Insert new: {}", user);
		User newUser = getUserRepo().save(user);
		log.info("New user: {}", newUser);
		return newUser;
	}

	@PutMapping
	public User update(@RequestBody final User user) {

		log.info("Update: {}", user);
		User updUser = getUserRepo().save(user);
		log.info("User updated: {}", updUser);
		return updUser;
	}

	@DeleteMapping("/{email}")
	public void deleteByEmail(@PathVariable final String email) {

		log.info("Delete by email: {}", email);
		final User user = getByEmail(email);
		if(user.getId() != null) {
			getUserRepo().deleteById(user.getId());
		}
	}

	@DeleteMapping
	public void deleteAll() {

		log.info("Delete all");
		getUserRepo().deleteAll();
	}

	/* ADDITIONAL */

	@GetMapping("/additional/{id}")
	public User getById(@PathVariable final String id) {

		log.info("Get by id: {}", id);
		User user = getUserRepo()
				.findById(id)
				.orElse(buildErrorUserById(id));
		log.info("User by id: {}", user);
		return user;
	}

	@DeleteMapping("/additional/{id}")
	public void deleteById(@PathVariable final String id) {

		log.info("Delete by id: {}", id);
		getUserRepo().deleteById(id);
	}

	private User buildErrorUserById(final String id){

		return User.builder().id(id).errorMessage(USER_NOT_FOUND_ERR_MSG).build();
	}

	private User buildErrorUserByEmail(final String email){

		return User.builder().email(email).errorMessage(USER_NOT_FOUND_ERR_MSG).build();
	}

}
