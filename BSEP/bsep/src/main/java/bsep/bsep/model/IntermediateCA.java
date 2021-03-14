package bsep.bsep.model;

public class IntermediateCA {
	
	private Long id;	
	private String organization;
	private String serialNumber;
	private String commonName;
	private String category;
	private Address companyAddress;
	private Location incLocation;

	public IntermediateCA() {

	}

	public IntermediateCA(Long id, String organization, String serialNumber, String commonName, String category,
			Address companyAddress, Location incLocation) {
		super();
		this.id = id;
		this.organization = organization;
		this.serialNumber = serialNumber;
		this.commonName = commonName;
		this.category = category;
		this.companyAddress = companyAddress;
		this.incLocation = incLocation;
	}

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public String getOrganization() {
		return organization;
	}

	public void setOrganization(String organization) {
		this.organization = organization;
	}

	public String getSerialNumber() {
		return serialNumber;
	}

	public void setSerialNumber(String serialNumber) {
		this.serialNumber = serialNumber;
	}

	public String getCommonName() {
		return commonName;
	}

	public void setCommonName(String commonName) {
		this.commonName = commonName;
	}

	public String getCategory() {
		return category;
	}

	public void setCategory(String category) {
		this.category = category;
	}

	public Address getCompanyAddress() {
		return companyAddress;
	}

	public void setCompanyAddress(Address companyAddress) {
		this.companyAddress = companyAddress;
	}

	public Location getIncLocation() {
		return incLocation;
	}

	public void setIncLocation(Location incLocation) {
		this.incLocation = incLocation;
	}
}
