package com.rabbit.samples.kubernetes.serverapp.configs;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.config.EnableMongoAuditing;

import javax.annotation.PostConstruct;


/**
 * @author Matteo Baiguini
 */
@Slf4j
@Getter(AccessLevel.PRIVATE)
@Configuration
@EnableMongoAuditing
public class AuditingConfig {

	@Value("${spring.data.mongodb.host}")
	String dbHost;

	@Value("${spring.data.mongodb.port}")
	int dbPort;

	@Value("${spring.data.mongodb.username}")
	String dbUser;

	@Value("${spring.data.mongodb.password}")
	String dbPw;

	@Value("${spring.data.mongodb.database}")
	String dbSchema;

	@PostConstruct
	public void postConstruct() {

		log.warn("MONGO url {}:{}, user {}, pw {}, schema {}", getDbHost(), getDbPort(), getDbUser(), getDbPw(), getDbSchema());
	}
}
