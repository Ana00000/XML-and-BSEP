<template>
  <div>
    <br />
    <v-container fluid class="container">
      <v-combobox
        class="combo"
        :items="certificatePurpose"
        :item-text="text"
        v-model="selectedCertificatePurpose"
        :label="label2"
        hint="Choose certificate purpose."
      />

      <v-row aria-rowspan="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Common name"
            v-model="commonName"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Given name"
            v-model="givenName"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Surname"
            v-model="surname"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Organization"
            v-model="organization"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Organizational unit name"
            v-model="organizationalUnitName"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Organization email"
            v-model="organizationEmail"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Country code"
            v-model="countryCode"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Alias "
            v-model="alias"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="End date"
            v-model="endDate"
            :hint="'End date should be in format yyyy-mm-dd and must be before '+issuerEndDate+'!'"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Issuer Alias"
            v-model="issuerAlias"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-row rows="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Issuer Serial Number"
            v-model="issuerSerialNumber"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>

      <v-combobox
        class="combo"
        :item-text="text"
        :items="certificateTypes"
        v-model="selectedCertificateType"
        :label="label1"
        hint="Choose certificate type."
        />

        <v-combobox
        class="combo"
        :item-text="text"
        :items="usersEmailsItems"
        v-model="selectedUserEmail"
        :label="'USER EMAIL'"
        hint="Choose user email."
      />
    </v-container>

    <div class="center">
      <v-btn v-on:click="createCertificate" color="primary" large elevation="20"
        >Create</v-btn
      >
    </div>
  </div>
</template>

<script>
export default {
  name: "CreateOtherCertificates",
  data: () => ({
    commonName: "",
    givenName: "",
    surname: "",
    organization: "",
    organizationalUnitName: "",
    organizationEmail: "",
    countryCode: "",
    alias: "",
    endDate: "",
    issuerEndDate: null,
    issuerSerialNumber: null,
    issuerCertificateType: null,
    issuerAlias: "",
    correctIssuerAlias: null,
    certificates: [],
    dateValid: true,
    selectedCertificateType: "INTERMEDIATE",
    selectedCertificatePurpose: "NONE",
    certificatePurpose: ["SERVICE", "SUBSYSTEM", "USER", "NONE"],
    certificateTypes: ["ENDENTITY", "INTERMEDIATE"],
    label1: "Certificates type",
    label2: "Certificate purpose",
    usersEmailsItems: [],
    selectedUserEmail: null,
    token: null
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
        this.correctIssuerAlias = localStorage.getItem("issuerAlias");
        this.issuerSerialNumber = localStorage.getItem("issuerSerialNumber");
        this.issuerEndDate = localStorage.getItem("endDate");
        this.issuerCertificateType = localStorage.getItem("issuerCertificateType");
        this.token = localStorage.getItem("token");
        this.getEmails();
    },
    getEmails(){
        this.$http
        .get(
          "http://localhost:8080/users/getUsersEmails",{
        headers:{
            'Authorization':"Bearer "+ this.token
        }})
        .then((resp) => {
          this.usersEmailsItems = resp.data;
        })
        .catch(console.log("Didn't set certificate info!"));
    },

    createCertificate() {
      if(!this.validCertificate()) return;

      this.$http
        .post("http://localhost:8080/certificate/createCertificate", {
          commonName: this.commonName,
          givenName: this.givenName,
          surname: this.surname,
          organization: this.organization,
          organizationalUnitName: this.organizationalUnitName,
          organizationEmail: this.organizationEmail,
          countryCode: this.countryCode,
          alias: this.alias,
          endDate: this.endDate,
          certificateType: this.selectedCertificateType,
          certificatePurposeType: this.selectedCertificatePurpose,
          issuerSerialNumber: this.issuerSerialNumber,
          issuerAlias: this.issuerAlias,
          userEmail: this.selectedUserEmail
        },{
        headers:{
            'Authorization':"Bearer "+ this.token
        }})
        .then((resp) => {
          console.log(resp.data);
          alert("Created certificate.");
          window.location.href = "http://localhost:8081/certificates";
        })
        .catch((err) => {
          alert("Certificate wasn't created, sorry.");
          console.log(err.response.data);
        });
    },

    validCertificate() {
      if (this.validCommonName() && this.validGivenName() && this.validSurname() 
      && this.validOrganization() &&this.validOrganizationalUnitName() &&
      this.validOrganizationEmail() && this.validCountryCode() &&
      this.validAlias() && this.validEndDate() && this.validIssuerAlias() && this.validSelectedUserEmail()) return true;
      return false;
    },
    validCommonName() {
      if (this.commonName.length < 2) {
        alert("Your common name should contain at least 2 characters!");
        return false;
      } else if (this.commonName.length > 20) {
        alert("Your common name shouldn't contain more than 20 characters!");
        return false;
      } else if(this.commonName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) { 
          alert("Your common name shouldn't contain special characters.");
          return false;
      } else if (this.commonName.match(/[ ]/g)) {
        alert("Your common name shouldn't contain spaces!");
        return false;
      } else if (this.commonName.match(/\d/g)) {
        alert("Your common name shouldn't contain numbers!");
        return false;
      } else if (!this.commonName.match(/[A-Z][a-z]+/g)) {
        alert("Your common name needs to have upper letter at the start!");
        return false;
      } else if (this.commonName.match(/[A-Z][a-z]+[A-Z]+/g)) {
        alert("Your common name needs to have upper letter only at the start!");
        return false;
      }
      return true;
    },
    validGivenName() {
      if (this.givenName.length < 2) {
        alert("Your given name should contain at least 2 characters!");
        return false;
      } else if (this.givenName.length > 20) {
        alert("Your given name shouldn't contain more than 20 characters!");
        return false;
      } else if(this.givenName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) { 
          alert("Your given name shouldn't contain special characters.");
          return false;
      } else if (this.givenName.match(/[ ]/g)) {
        alert("Your given name shouldn't contain spaces!");
        return false;
      } else if (this.givenName.match(/\d/g)) {
        alert("Your given name shouldn't contain numbers!");
        return false;
      } else if (!this.givenName.match(/[A-Z][a-z]+/g)) {
        alert("Your given name needs to have upper letter at the start!");
        return false;
      } else if (this.givenName.match(/[A-Z][a-z]+[A-Z]+/g)) {
        alert("Your given name needs to have upper letter only at the start!");
        return false;
      }
      return true;
    },
    validSurname() {
      if (this.surname.length < 2) {
        alert("Your surname should contain at least 2 characters!");
        return false;
      } else if (this.surname.length > 35) {
        alert("Your surname shouldn't contain more than 35 characters!");
        return false;
      } else if(this.surname.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) { 
          alert("Your surname shouldn't contain special characters.");
          return false;
      } else if (this.surname.match(/[ ]/g)) {
        alert("Your surname shouldn't contain spaces!");
        return false;
      } else if (this.surname.match(/\d/g)) {
        alert("Your surname shouldn't contain numbers!");
        return false;
      } else if (!this.surname.match(/[A-Z][a-z]+/g)) {
        alert("Your surname needs to have upper letter at the start!");
        return false;
      } else if (this.surname.match(/[A-Z][a-z]+[A-Z]+/g)) {
        alert("Your surname needs to have upper letter only at the start!");
        return false;
      }
      return true;
    },
    validOrganization() {
      if (this.organization.length < 2) {
        alert("Your organization should contain at least 2 characters!");
        return false;
      } else if (this.organization.length > 20) {
        alert("Your organization shouldn't contain more than 20 characters!");
        return false;
      } else if(this.organization.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) { 
          alert("Your organization shouldn't contain special characters.");
          return false;
      }
      return true;
    },
    validOrganizationalUnitName() {
       if (this.organizationalUnitName.length < 2) {
        alert("Your organizational unit name should contain at least 2 characters!");
        return false;
      } else if (this.organizationalUnitName.length > 20) {
        alert("Your organizational unit name shouldn't contain more than 20 characters!");
        return false;
      } else if(this.organizationalUnitName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) { 
          alert("Your organizational unit name shouldn't contain special characters.");
          return false;
      } else if (this.organizationalUnitName.match(/\d/g)) {
        alert("Your organizational unit name shouldn't contain numbers!");
        return false;
      }
      return true;
    },
    validOrganizationEmail() {
      if (!/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(this.organizationEmail)) {
        alert("You have entered an invalid organization email address!");
        return false;
      } else if (this.organizationEmail.length > 35) {
        alert("Organization email address shouldn't contain more than 35 characters!");
        return false;
      }
      return true;
    },
    validCountryCode() {
      if(this.countryCode.match(/[a-zA-Z]/g)) {
          alert("Your country code shouldn't contain letters.");
          return false;
      } else if (this.countryCode.match(/[ ]/g)) {
        alert("Your country code shouldn't contain spaces!");
        return false;
      } else if(this.countryCode.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) {
          alert("Your country code shouldn't contain special characters.");
          return false;
      } else if(this.countryCode.match(/\d/g) == null){
        alert("Your country code needs numbers!");
        return false;
      } else if (this.countryCode.match(/\d/g).length < 2) {
        alert("Your country code should contain at least 2 numbers!");
        return false;
      } else if (this.countryCode.match(/\d/g).length > 7) {
        alert("Your country code shouldn't contain more than 7 numbers!");
        return false;
      }
      return true;
    },
    validAlias() {
      if(this.alias.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) {
          alert("Your alias shouldn't contain special characters.");
          return false;
      } else if (this.alias.match(/[ ]/g)) {
        alert("Your alias shouldn't contain spaces!");
        return false;
      } else if (this.alias.length < 1) {
        alert("Your alias should contain at least 1 character!");
        return false;
      } else if (this.alias.length > 20) {
        alert("Your alias shouldn't contain more than 20 characters!");
        return false;
      } 
      return true;
    },
    validIssuerAlias() {
      if(this.issuerAlias.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) {
          alert("Your issuer alias shouldn't contain special characters.");
          return false;
      } else if (this.issuerAlias.match(/[ ]/g)) {
        alert("Your issuer alias shouldn't contain spaces!");
        return false;
      }  else if (this.issuerAlias.length < 1) {
        alert("Your issuer alias should contain at least 1 character!");
        return false;
      } else if (this.issuerAlias.length > 20) {
        alert("Your issuer alias shouldn't contain more than 20 characters!");
        return false;
      } else if (this.issuerAlias != this.correctIssuerAlias){
        alert("You didn't enter a correct issuer alias!");
        return false;
      }
      return true;
    },
    validEndDate() {
       if(this.endDate.match(/\d/g) == null){
        alert("Your end date needs numbers!");
        return false;
      } else if (this.endDate.match(/\d/g).length < 6) {
        alert("Your end date should contain at least 6 numbers!");
        return false;
      } else if (this.endDate.match(/\d/g).length > 8) {
        alert("Your end date shouldn't contain more than 8 numbers!");
        return false;
      } else if(this.endDate.match(/[a-zA-Z]/g)) {
          alert("Your end date shouldn't contain letters.");
          return false;
      } else if(this.endDate.match(/[!@#$%^&*,:'/.<>+\\"]/g)) {
          alert("Your end date shouldn't contain special character other than [-].");
          return false;
      } else if (this.endDate.match(/[ ]/g)) {
        alert("Your end date shouldn't contain spaces!");
        return false;
      } else if(!this.endDate.match(/[2][0-9]{3}-[0-1][0-9]-[0-3][0-9]/g)) {
          alert("Your end date is not set in right format.");
          return false;
      } else if(this.endDate.match(/[2][0-9]{3}-[0-1][0-9]-[0-3][0-9][-]+/g)) {
          alert("Your end date can't contain - at end of input.");
          return false;
      } else if(!this.validEndDateWithIssuerDate()){
        return false;
      }
      return true;
    },
    validEndDateWithIssuerDate() {
      var issuerEndDateSplit = this.issuerEndDate.split('-');
      var issuerEDSDay = issuerEndDateSplit[2];
      var issuerEDSMonth = issuerEndDateSplit[1];
      var issuerEDSYear = issuerEndDateSplit[0];

      var endDateSplit = this.endDate.split('-');
      var eDSDay = endDateSplit[2];
      var eDSMonth = endDateSplit[1];
      var eDSYear = endDateSplit[0];

      var issuerEndDateO = new Date(issuerEDSYear,issuerEDSMonth,issuerEDSDay);
      var endDateO = new Date(eDSYear,eDSMonth,eDSDay);

      if (issuerEndDateO<endDateO){
        alert("Your expired date must be before issuer expired date!");
        return false;
      }
      console.log(issuerEDSDay+' '+issuerEDSMonth+' '+issuerEDSYear);
      return true;
    },
    validSelectedUserEmail(){
      if(this.selectedUserEmail == null){
        alert("Your user email wasn't selected!");
        return false;
      }
      return true;
    }
  },
};
</script>

<style scoped>
.combo {
  width: 25%;
  margin-left: 42%;
}

.center {
  margin-left: 50%;
  padding: 10px;
}
</style>
