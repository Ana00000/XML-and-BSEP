package bsep.bsep.model;

import javax.persistence.Embeddable;

@Embeddable
public class Location {
	private int postalCode;
	private String city;
	private String state;
	private String country;
	private String countryCode;

	public Location() {
	}

	public Location(int postalCode, String city, String state, String country, String countryCode) {
		super();
		this.postalCode = postalCode;
		this.city = city;
		this.state = state;
		this.country = country;
		this.countryCode = countryCode;
	}

	public String getCountryCode() {
		return countryCode;
	}

	public void setCountryCode(String countryCode) {
		this.countryCode = countryCode;
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
