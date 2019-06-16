package com.rabbit.samples.kubernetes.clientapp.feign.clients;

import com.rabbit.samples.kubernetes.clientapp.domain.User;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;

import java.util.List;


/**
 * @author Matteo Baiguini
 */
@FeignClient(
		name = "serverapp",
		path = "/users"
)
public interface UserServiceClient {

	@GetMapping
	List<User> getAll();

	@GetMapping("/{email}")
	User getByEmail(@PathVariable final String email);

	@PostMapping
	User insert(final User user);

	@PutMapping
	User update(final User user);

	@DeleteMapping("/{email}")
	void deleteByEmail(@PathVariable final String email);

	@DeleteMapping
	void deleteAll();
}
