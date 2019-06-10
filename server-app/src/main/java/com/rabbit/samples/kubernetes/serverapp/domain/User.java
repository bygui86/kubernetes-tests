package com.rabbit.samples.kubernetes.serverapp.domain;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;
import lombok.ToString;
import lombok.experimental.FieldDefaults;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.SequenceGenerator;
import javax.persistence.Table;


/**
 * @author Matteo Baiguini
 */
@FieldDefaults(level = AccessLevel.PRIVATE)
@AllArgsConstructor
@NoArgsConstructor
@Data
@Builder
@EqualsAndHashCode(callSuper = true)
@ToString(callSuper = true)
@Entity
@Table(name = "users")
public class User extends Auditable {

	@Id
	@GeneratedValue(generator = "user_generator")
	@SequenceGenerator(
			name = "user_generator",
			sequenceName = "user_sequence",
			initialValue = 1000
	)
	Long id;

	@Column(unique = true)
	String email;

	@Column
	String name;

	@Column
	int age;

	String errorMessage;

}
