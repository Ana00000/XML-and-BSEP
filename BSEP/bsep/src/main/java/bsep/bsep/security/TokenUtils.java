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
	// Izdavac tokena
	@Value("spring-security-example")
	private String APP_NAME;

	// Tajna koju samo backend aplikacija treba da zna kako bi mogla da generise i
	// proveri JWT https://jwt.io/
	@Value("somesecret")
	public String SECRET;

	// Period vazenja
	@Value("300000000")
	private int EXPIRES_IN;

	private static final String AUDIENCE_WEB = "web";

	// Naziv headera kroz koji ce se prosledjivati JWT u komunikaciji server-klijent
	@Value("Authorization")
	private String AUTH_HEADER;

	// Algoritam za potpisivanje JWT
	private SignatureAlgorithm SIGNATURE_ALGORITHM = SignatureAlgorithm.HS512;

	public String generateToken(String userEmail) {
		return Jwts.builder().setIssuer(APP_NAME).setSubject(userEmail).setAudience(generateAudience())
				.setIssuedAt(new Date()).setExpiration(generateExpirationDate())
				// .claim("key", value) //moguce je postavljanje proizvoljnih podataka u telo
				// JWT tokena
				.signWith(SIGNATURE_ALGORITHM, SECRET).compact();
	}

	public String getToken(HttpServletRequest request) {
		String authHeader = getAuthHeaderFromHeader(request);

		// JWT se prosledjuje kroz header Authorization u formatu:
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

	// Funkcija za validaciju JWT tokena
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