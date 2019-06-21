package com.rabbit.samples.kubernetes.serverapp.configs;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

import javax.annotation.PostConstruct;


/**
 * @author Matteo Baiguini
 */
@Slf4j
@Getter(AccessLevel.PRIVATE)
@Configuration
@EnableJpaAuditing
public class AuditingConfig {

	@Value("${spring.datasource.url}")
	String dbUrl;

	@Value("${spring.datasource.username}")
	String dbUser;

	@Value("${spring.datasource.password}")
	String dbPw;

	@PostConstruct
	public void postConstruct() {

		log.warn("MYSQL url {}, user {}, pw {}", getDbUrl(), getDbUser(), getDbPw());
	}
}
