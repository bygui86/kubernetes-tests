package com.rabbit.samples.kubernetes.serverapp.configs;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Profile;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

import javax.annotation.PostConstruct;


/**
 * @author Matteo Baiguini
 */
@Slf4j
@Getter(AccessLevel.PRIVATE)
@Configuration
@EnableJpaAuditing
@Profile("in-memory")
public class InMemoryAuditingConfig {

	// @Value("${postgres.cluster.host}")
	// String dbHost;

	// @Value("${postgres.cluster.port}")
	// int dbPort;

	// @Value("${postgres.cluster.schema}")
	// String dbSchema;

	// @Value("${postgres.cluster.user}")
	// String dbUser;

	// @Value("${postgres.cluster.pw}")
	// String dbPw;

	// @PostConstruct
	// public void postConstruct() {

	// 	log.warn("POSTGRES host {}, port {}, schema {}, user {}, pw {}", getDbHost(), getDbPort(), getDbSchema(), getDbUser(), getDbPw());
	// }
}
