package com.rabbit.samples.kubernetes.serverapp.configs;

import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Profile;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;


/**
 * @author Matteo Baiguini
 */
@Configuration
@EnableJpaAuditing
@Profile({"local", "docker", "kube", "kube-pg-oper"})
public class AuditingConfig {

	// no-op
}
