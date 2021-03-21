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
            hint="Common name should contain at least 2 characters!"
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
            hint="Given name should contain at least 2 characters!"
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
            hint="Surname should contain at least 2 characters!"
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
            hint="Organization should contain at least 2 characters!"
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
            hint="Organizational unit name should contain at least 1 character!"
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
            hint="Organization email should contain at least 10 characters!"
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
            hint="Country code should contain at least 2 characters!"
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
            hint="Alias should contain at least 2 characters!"
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
            hint="End date should contain at least 6 characters!"
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
    commonName: null,
    givenName: null,
    surname: null,
    organization: null,
    organizationalUnitName: null,
    organizationEmail: null,
    countryCode: null,
    alias: null,
    endDate: null,
    issuerSerialNumber: null,
    issuerAlias: null,
    certificates: [],
    selectedCertificateType: "ROOT",
    selectedCertificatePurpose: "NONE",
    certificatePurpose: ["SERVICE", "SUBSYSTEM", "USER", "NONE"],
    label1: "Certificates type",
    label2: "Certificate purpose",
  }),
  methods: {
    createCertificate() {
      //this.validation();
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
          issuerSerialNumber: null,
          issuerAlias: this.alias,
        })
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
    validation() {
      this.validationOfCommonName();
      this.validationOfGivenName();
      this.validationOfSurname();
      this.validationOfOrganization();
      this.validationOfOrganizationalUnitName();
      this.validationOfOrganizationEmailLength();
      this.validationOfCountryCodeLength();
      this.validationOfAlias();
      this.validationOfEndDate();
    },
    validationOfCommonName() {
      if (this.commonName.length < 2) {
        alert("Your common name should contain at least 2 characters!");
        return;
      } else if (this.name.length > 20) {
        alert("Your common name shouldn't contain more than 20 characters!");
        return;
      }
    },
    validationOfGivenName() {
      if (this.givenName.length < 2) {
        alert("Your given name should contain at least 2 characters!");
        return;
      } else if (this.givenName.length > 20) {
        alert("Your given name shouldn't contain more than 20 characters!");
        return;
      }
    },
    validationOfSurname() {
      if (this.surname.length < 2) {
        alert("Your surname should contain at least 2 characters!");
        return;
      } else if (this.surname.length > 35) {
        alert("Your surname shouldn't contain more than 35 characters!");
        return;
      }
    },
    validationOfOrganization() {
      if (this.organization.length < 2) {
        alert("Your organization should contain at least 23 characters!");
        return;
      } else if (this.organization.length > 20) {
        alert("Your organization shouldn't contain more than 20 characters!");
        return;
      }
    },
    validationOfOrganizationalUnitName() {
      if (this.organizationalUnitName.length < 1) {
        alert(
          "Your organizational unit name should contain at least 1 character!"
        );
        return;
      } else if (this.organizationalUnitName.length > 20) {
        alert(
          "Your organizational unit name shouldn't contain more than 20 characters!"
        );
        return;
      }
    },
    validationOfOrganizationEmailLength() {
      if (
        /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(
          this.organizationEmail
        )
      ) {
        return true;
      }
      alert("You have entered an invalid organization email address!");
      return false;
    },
    validationOfCountryCodeLength() {
      if (this.countryCode.length < 2) {
        alert("Your country code should contain at least 2 characters!");
        return;
      } else if (this.countryCode.length > 20) {
        alert("Your country code shouldn't contain more than 20 characters!");
        return;
      }
    },
    validationOfAlias() {
      if (this.alias.length < 2) {
        alert("Your alias should contain at least 2 character!");
        return;
      } else if (this.alias.length > 20) {
        alert("Your alias shouldn't contain more than 20 characters!");
        return;
      }
    },
    validationOfEndDate() {
      if (this.endDate.length < 6) {
        alert("Your end date should contain at least 6 character!");
        return;
      } else if (this.endDate.length > 20) {
        alert("Your end date shouldn't contain more than 20 characters!");
        return;
      }
    },
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
