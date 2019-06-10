package com.rabbit.samples.kubernetes.serverapp.configs;

import com.google.common.base.Predicate;
import lombok.AccessLevel;
import lombok.Getter;
import lombok.experimental.FieldDefaults;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.RequestHandler;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.service.Contact;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

import java.util.Collections;


/**
 * @author Matteo Baiguini
 */
@FieldDefaults(level = AccessLevel.PRIVATE)
@Getter(AccessLevel.PRIVATE)
@Configuration
@EnableSwagger2
public class SwaggerConfig {

	@Bean
	public Docket apiDocket() {

		return new Docket(DocumentationType.SWAGGER_2)
				.select()
				.apis(
						getApiSelector()
				)
				.paths(
						getPathSelector()
				)
				.build()
				.apiInfo(
						getApiInfo()
				);
	}

	private Predicate<RequestHandler> getApiSelector() {

		return RequestHandlerSelectors
				.basePackage("com.rabbit.samples.kubernetes.serverapp.controllers");
	}

	private Predicate<String> getPathSelector() {

		return PathSelectors
				.any();
	}

	private ApiInfo getApiInfo() {

		return new ApiInfo(
				"Server application",
				"Kubernetes server application",
				"0.0.1",
				"https://raw.githubusercontent.com/bygui86/kubernetes-tests/master/terms-of-service.txt",
				new Contact("Matteo Baiguini","https://mbsolidconsulting.com",""),
				"Apache 2.0",
				"http://www.apache.org/licenses/LICENSE-2.0",
				Collections.emptyList()
		);
	}

}
