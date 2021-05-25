package bsep.bsep.security;

import java.util.Date;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Component;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

@Component
public class TokenUtils {
	// Token issuer
	@Value("spring-security-example")
	private String APP_NAME;

	// Secret that only the application backend should know so it can generate and
	// check JWT https://jwt.io/
	@Value("somesecret")
	public String SECRET;

	// Validity period
	@Value("300000000")
	private int EXPIRES_IN;

	private static final String AUDIENCE_WEB = "web";

	// Header title which will be used to transfer JWT in the client-server
	// communication
	@Value("Authorization")
	private String AUTH_HEADER;

	// Algorith used for signing JWT
	private SignatureAlgorithm SIGNATURE_ALGORITHM = SignatureAlgorithm.HS512;

	public String generateToken(String userEmail) {
		return Jwts.builder().setIssuer(APP_NAME).setSubject(userEmail).setAudience(generateAudience())
				.setIssuedAt(new Date()).setExpiration(generateExpirationDate())
				// .claim("key", value) // it is possible to set custom data in the body of the
				// JWT token
				.signWith(SIGNATURE_ALGORITHM, SECRET).compact();
	}

	public String getToken(HttpServletRequest request) {
		String authHeader = getAuthHeaderFromHeader(request);

		// JWT is transfered trough the Authorization header in the format:
		// Bearer
		// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
		if (authHeader != null && authHeader.startsWith("Bearer ")) {
			return authHeader.substring(7);
		}

		return null;
	}

	public String getUsernameFromToken(String token) {
		String userEmail;
		try {
			final Claims claims = this.getAllClaimsFromToken(token);
			userEmail = claims.getSubject();
		} catch (Exception e) {
			userEmail = null;
		}
		return userEmail;
	}

	public int getExpiredIn() {
		return EXPIRES_IN;
	}

	private Claims getAllClaimsFromToken(String token) {
		Claims claims;
		try {
			claims = Jwts.parser().setSigningKey(SECRET).parseClaimsJws(token).getBody();
		} catch (Exception e) {
			claims = null;
		}
		return claims;
	}

	private String generateAudience() {
		return AUDIENCE_WEB;
	}

	private Date generateExpirationDate() {
		return new Date(new Date().getTime() + EXPIRES_IN);
	}

	public String getAuthHeaderFromHeader(HttpServletRequest request) {
		return request.getHeader(AUTH_HEADER);
	}

	// JWT token validation function
	public Boolean validateToken(String token, UserDetails userDetails) {
		final String userEmail = getUsernameFromToken(token);
		return (userEmail != null && userEmail.equals(userDetails.getUsername()));
	}

	public Date getIssuedAtDateFromToken(String token) {
		Date issueAt;
		try {
			final Claims claims = this.getAllClaimsFromToken(token);
			issueAt = claims.getIssuedAt();
		} catch (Exception e) {
			issueAt = null;
		}
		return issueAt;
	}
}