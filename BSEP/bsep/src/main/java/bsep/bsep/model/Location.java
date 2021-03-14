package bsep.bsep.model;

public class Location  {
	private int postalCode;
	private String city;
	private String state;
	private String country;
	    
	public Location() {
		// TODO Auto-generated constructor stub
	}
	
	public Location(int postalCode, String city, String state, String country) {
		super();
		this.postalCode = postalCode;
		this.city = city;
		this.state = state;
		this.country = country;
	}

	public int getPostalCode() {
		return postalCode;
	}

	public void setPostalCode(int postalCode) {
		this.postalCode = postalCode;
	}

	public String getCity() {
		return city;
	}

	public void setCity(String city) {
		this.city = city;
	}

	public String getState() {
		return state;
	}

	public void setState(String state) {
		this.state = state;
	}

	public String getCountry() {
		return country;
	}

	public void setCountry(String country) {
		this.country = country;
	}
}
