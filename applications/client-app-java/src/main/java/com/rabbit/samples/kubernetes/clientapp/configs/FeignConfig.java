package com.rabbit.samples.kubernetes.clientapp.configs;

import org.springframework.cloud.openfeign.EnableFeignClients;
import org.springframework.context.annotation.Configuration;


/**
 * @author Matteo Baiguini
 */
@Configuration
@EnableFeignClients(basePackages = "com.rabbit.samples.kubernetes.clientapp.feign.clients")
public class FeignConfig {

	// no-op
}
