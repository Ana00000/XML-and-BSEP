package bsep.bsep.model;

public class IntermediateCA {

	private Long id;
	private String organization;
	private String serialNumber;
	private String commonName;
	private CategoryType category;
	private Address companyAddress;
	private Location incLocation;
	private String organizationEmail;
	private String organizationalUnitName;
	private boolean isValid;

	public IntermediateCA() {

	}

	public IntermediateCA(Long id, String organization, String serialNumber, String commonName, CategoryType category,
			Address companyAddress, Location incLocation, String organizationEmail, String organizationalUnitName,
			boolean isValid) {
		super();
		this.id = id;
		this.organization = organization;
		this.serialNumber = serialNumber;
		this.commonName = commonName;
		this.category = category;
		this.companyAddress = companyAddress;
		this.incLocation = incLocation;
		this.organizationEmail = organizationEmail;
		this.organizationalUnitName = organizationalUnitName;
		this.isValid = isValid;
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

	public String getOrganizationEmail() {
		return organizationEmail;
	}

	public void setOrganizationEmail(String organizationEmail) {
		this.organizationEmail = organizationEmail;
	}

	public String getOrganizationalUnitName() {
		return organizationalUnitName;
	}

	public void setOrganizationalUnitName(String organizationalUnitName) {
		this.organizationalUnitName = organizationalUnitName;
	}

	public void setCategory(CategoryType category) {
		this.category = category;
	}

	public CategoryType getCategory() {
		return category;
	}

	public boolean isValid() {
		return isValid;
	}

	public void setValid(boolean isValid) {
		this.isValid = isValid;
	}

}
