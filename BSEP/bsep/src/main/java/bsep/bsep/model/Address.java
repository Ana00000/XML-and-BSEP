package bsep.bsep.model;

public class Address extends Location{
	private String streetName;
	private String streetNumber;

	public Address() {
		// TODO Auto-generated constructor stub
	}
	
	public Address(String streetName, String streetNumber) {
		super();
		this.streetName = streetName;
		this.streetNumber = streetNumber;
	}

	public String getStreetName() {
		return streetName;
	}

	public void setStreetName(String streetName) {
		this.streetName = streetName;
	}

	public String getStreetNumber() {
		return streetNumber;
	}

	public void setStreetNumber(String streetNumber) {
		this.streetNumber = streetNumber;
	}
	
}
