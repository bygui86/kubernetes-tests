package com.rabbit.samples.kubernetes.clientapp.controllers;

import com.rabbit.samples.kubernetes.clientapp.domain.User;
import com.rabbit.samples.kubernetes.clientapp.feign.clients.UserServiceClient;
import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.experimental.FieldDefaults;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;


@Slf4j
@FieldDefaults(level = AccessLevel.PRIVATE)
@AllArgsConstructor
@Getter(value = AccessLevel.PRIVATE)
@RestController
@RequestMapping("/apis/users")
class UserController {

	UserServiceClient userServiceClient;

	@GetMapping
	public List<User> getAll() {

		log.info("Get all users");
		List<User> users = getUserServiceClient().getAll();
		log.info("All users found: {}", users.size());
		return users;
	}

	@GetMapping("/{email}")
	public User getByEmail(@PathVariable final String email) {

		log.info("Get by email: {}", email);
		User user = getUserServiceClient().getByEmail(email);
		log.info("User by email: {}", user);
		return user;
	}

	@PostMapping
	public User insert(@RequestBody final User user) {

		log.info("Insert new: {}", user);
		User newUser = getUserServiceClient().insert(user);
		log.info("New user: {}", newUser);
		return newUser;
	}

	@PutMapping
	public User update(@RequestBody final User user) {

		log.info("Update: {}", user);
		User updUser = getUserServiceClient().update(user);
		log.info("User updated: {}", updUser);
		return updUser;
	}

	@DeleteMapping("/{email}")
	public void deleteByEmail(@PathVariable final String email) {

		log.info("Delete by email: {}", email);
		getUserServiceClient().deleteByEmail(email);
	}

	@DeleteMapping
	public void deleteAll() {

		log.info("Delete all");
		getUserServiceClient().deleteAll();
	}

}
