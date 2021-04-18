package bsep.bsep.security;

import java.io.IOException;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.web.filter.OncePerRequestFilter;

public class TokenAuthenticationFilter extends OncePerRequestFilter {

	private TokenUtils tokenUtils;

	private UserDetailsService userDetailsService;

	public TokenAuthenticationFilter(TokenUtils tokenHelper, UserDetailsService userDetailsService) {
		this.tokenUtils = tokenHelper;
		this.userDetailsService = userDetailsService;
	}

	@Override
	public void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain)
			throws IOException, ServletException {

		String userEmail;
		String authToken = tokenUtils.getToken(request);

		if (authToken != null) {
			// Get username from the token
			userEmail = tokenUtils.getUsernameFromToken(authToken);

			if (userEmail != null) {
				// Get user with the username
				UserDetails userDetails = userDetailsService.loadUserByUsername(userEmail);

				// Check if the token is valid
				if (tokenUtils.validateToken(authToken, userDetails)) {
					// Create authentication
					TokenBasedAuthentication authentication = new TokenBasedAuthentication(userDetails);
					authentication.setToken(authToken);
					SecurityContextHolder.getContext().setAuthentication(authentication);
				}
			}
		}

		// Forward the request to the next filter
		chain.doFilter(request, response);
	}

}