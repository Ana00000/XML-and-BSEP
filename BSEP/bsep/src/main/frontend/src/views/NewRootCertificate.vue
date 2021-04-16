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
            hint="End date should be in format yyyy-mm-dd"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      <v-combobox
        class="combo"
        :item-text="text"
        v-model="selectedCertificateType"
        :label="label1"
        hint="Choose certificate type."
        disable
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
  name: "NewRootCertificate",
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
    issuerSerialNumber: null,
    issuerAlias: null,
    certificates: [],
    selectedCertificateType: "ROOT",
    selectedCertificatePurpose: "NONE",
    certificatePurpose: ["SERVICE", "SUBSYSTEM", "USER", "NONE"],
    label1: "Certificates type",
    label2: "Certificate purpose",
    userEmail: null,
    token: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.userEmail = localStorage.getItem("userEmail");
      this.token = localStorage.getItem("token");
    },
    createCertificate() {
      if (!this.validCertificate()) return;
      this.$http
        .post(
          "http://localhost:8080/certificate/createCertificate",
          {
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
            issuerSerialNumber: null,
            issuerAlias: this.alias,
            userEmail: this.userEmail,
          },
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log(resp.data);
          alert("Created Root certificate.");
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
      this.validAlias() && this.validEndDate()) return true;
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
      }
      var endDateSplit = this.endDate.split('-');
      var eDSYear = endDateSplit[0];
      var eDSMonth = endDateSplit[1];
      var eDSDay = endDateSplit[2];

      if (eDSYear > 3000 || eDSYear < 2021){
        alert("Year of end date isn't valid");
        return false;
      }else if(eDSYear < 2025){
        alert("End date must me valid 5 years from current date.");
        return false;
      }else if (eDSMonth > 12 || eDSMonth < 0){
        alert("Month of end date isn't valid");
        return false;
      }else if (eDSDay > 31 || eDSDay < 1){
        alert("Day of end date isn't valid");
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
