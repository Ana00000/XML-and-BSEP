package bsep.bsep.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

import bsep.bsep.security.RestAuthenticationEntryPoint;
import bsep.bsep.security.TokenAuthenticationFilter;
import bsep.bsep.security.TokenUtils;
import bsep.bsep.service.CustomUserDetailsService;

@Configuration
@EnableGlobalMethodSecurity(prePostEnabled = true)
@EnableWebSecurity
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {
	@Bean
	public PasswordEncoder passwordEncoder() {
		return new BCryptPasswordEncoder();
	}

	// Service used for reading application users data
	@Autowired
	private CustomUserDetailsService jwtUserDetailsService;

	// Handler for returning 401 when a client with invalid username and password
	// tries to access the resource
	@Autowired
	private RestAuthenticationEntryPoint restAuthenticationEntryPoint;

	// Registering the authentication manager which will do the user authentication
	// for us
	@Bean
	@Override
	public AuthenticationManager authenticationManagerBean() throws Exception {
		return super.authenticationManagerBean();
	}

	// Definition of instructions for the authentication manager
	// Defining which service should the authentication manager use to extract user
	// for authentication data
	// Defining which encoder should be used to encode the password from the user's
	// request so the resulting hash can be compared with the hash from the database
	// using the bcrypt algorithm
	// (the password in the database is not in plain text)
	@Autowired
	public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {
		auth.userDetailsService(jwtUserDetailsService).passwordEncoder(passwordEncoder());
	}

	// Injection of TokenUtils class implementation so we can use it's methods for
	// JWT in TokenAuthenticationFilteru
	@Autowired
	private TokenUtils tokenUtils;

	// Defining access rules for specific URL
	@Override
	public void configure(HttpSecurity http) throws Exception {
		http

				// Communication between the client and server is statelss because it is a REST
				// application
				.sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS).and()

				// All unauthenticated request should be uniform processed and return 401 error
				.exceptionHandling().authenticationEntryPoint(restAuthenticationEntryPoint).and()

				// Allow access to paths /auth/**, (/h2-console/** if H2 database is used) and
				// /api/foo to all users
				.authorizeRequests().antMatchers("/auth/**").permitAll().antMatchers("/auth/**/*").permitAll()
				.antMatchers("/h2-console/**").permitAll().antMatchers("/api/foo").permitAll().antMatchers("/api/**")
				.permitAll().antMatchers("/users/login").permitAll().antMatchers("/users/register").permitAll()
				.antMatchers("/users/changePassword").permitAll().antMatchers("/users/findUserWithToken").permitAll()
				.antMatchers("/users/recoverPasswordWithToken").permitAll().antMatchers("/users/confirm_account/*")
				.permitAll()

				// For any other request the user must be authenticated
				.anyRequest().authenticated().and()

				// For development purposes turn on CORS configuration from WebConfig class
				// .cors().disable()
				.cors().and()

				// Add custom filter TokenAuthenticationFilter for JWT token check instead of
				// using the username and password (which is done with the
				// BasicAuthenticationFilter)
				.addFilterBefore(new TokenAuthenticationFilter(tokenUtils, jwtUserDetailsService),
						BasicAuthenticationFilter.class)

				// Cross-Site Scripting (XSS) Attack
				.headers().xssProtection().and().contentSecurityPolicy("script-src 'self'");

		// Because of the simplicity of the example
		http.csrf().disable();
	}

	// General application security
	@Override
	public void configure(WebSecurity web) throws Exception {
		// TokenAuthenticationFilter will ignore everything under the listed path
		web.ignoring().antMatchers(HttpMethod.PUT, "/users/confirm_account/*", "/users/changePassword");
		web.ignoring().antMatchers(HttpMethod.POST, "/users/login", "/users/register",
				"/users/recoverPasswordWithToken", "/users/findUserWithToken");
		web.ignoring().antMatchers(HttpMethod.GET, "/", "/webjars/**", "/*.html", "/favicon.ico", "/**/*.html",
				"/**/*.css", "/**/*.js", "/auth/getRole");
	}
}