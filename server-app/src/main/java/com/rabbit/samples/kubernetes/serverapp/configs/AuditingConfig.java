package com.rabbit.samples.kubernetes.serverapp.configs;

import lombok.extern.slf4j.Slf4j;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;


/**
 * @author Matteo Baiguini
 */
@Configuration
@EnableJpaAuditing
public class AuditingConfig {

	// no-op
}
