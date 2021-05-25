package bsep.bsep.security;

public class UserTokenState {
	private String accessToken;
	private Long expiresIn;
	private String role;

	public UserTokenState() {
		this.accessToken = null;
		this.expiresIn = null;
		this.role = "NONE";
	}

	public UserTokenState(String accessToken, long expiresIn, String role) {
		this.accessToken = accessToken;
		this.expiresIn = expiresIn;
		this.role = role.equals("ADMIN") || role.equals("USER") ? role : "NONE";
	}

	public String getAccessToken() {
		return accessToken;
	}

	public void setAccessToken(String accessToken) {
		this.accessToken = accessToken;
	}

	public Long getExpiresIn() {
		return expiresIn;
	}

	public void setExpiresIn(Long expiresIn) {
		this.expiresIn = expiresIn;
	}

	public String getRole() {
		return role;
	}

	public void setRole(String role) {
		this.role = role;
	}
}