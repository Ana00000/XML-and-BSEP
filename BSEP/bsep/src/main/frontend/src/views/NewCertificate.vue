<template>
  <div>
    <br/>
    <v-container fluid class="container">

    <v-combobox
        class="combo"
        :items="certificates"
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
        :items="certificates"
        :item-text="text"
        v-model="selectedCertificateType"
        :label="label1"
        hint="Choose certificate type."
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
  name: "NewCertificate",
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
    label1: "Certificates type",
    label2: "Certificate purpose",
  }),
  methods: {
     createCertificate() {
      this.$http.post ('http://localhost:8080/certificate/createCertificate',
          {
              "commonName":this.commonName,
              "givenName":this.givenName,
              "surname":this.surname,
              "organization":this.organization,
              "organizationalUnitName":this.organizationalUnitName,
              "organizationEmail":this.organizationEmail,
              "countryCode":this.countryCode,
              "alias":this.alias,
              "endDate":this.endDate,
              "certificateType":this.selectedCertificateType,
              "certificatePurposeType":this.selectedCertificatePurpose,
              "issuerSerialNumber":null,
              "issuerAlias":this.alias
            }).then(resp => {
               console.log(resp.data);
                alert("Created Root certificate.");
            }).catch(err => {
                alert("Doctor or patient is busy at this time.");
                console.log(err.response.data);
            })
      }
  }
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
