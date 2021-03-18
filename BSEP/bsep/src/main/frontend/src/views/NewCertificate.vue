<template>
  <div class="page">
      <br/><br/>
      <v-container fluid class="container">
          <v-row>
          <v-col cols="5"/>
          <v-col cols="1">
              <v-text-field
                  label="Common name"
                  v-model="commonName"
                  hint="Common name should contain at least 2 characters!"
                  color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-row>
          <v-col cols="5"/>
          <v-col cols="2">
              <v-text-field
              label="Given name"
              v-model="givenName"
              hint="Given name should contain at least 2 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-row>
          <v-col cols="5"/>
          <v-col cols="3">
              <v-text-field
              label="Surname"
              v-model="surname"
              hint="Surname should contain at least 2 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-row>
          <v-col cols="5"/>
          <v-col cols="4">
              <v-text-field
              label="Organization"
              v-model="organization"
              hint="Organization should contain at least 2 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-row>
          <v-col cols="5"/>
          <v-col cols="5">
              <v-text-field
              label="Organizational unit name"
              v-model="organizationalUnitName"
              hint="Organizational unit name should contain at least 1 character!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>
          
          <v-row>
          <v-col cols="5"/>
          <v-col cols="6">
              <v-text-field
              label="Organization email"
              v-model="organizationEmail"
              hint="Organization email should contain at least 10 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>
          
          <v-row>
          <v-col cols="5"/>
          <v-col cols="7">
              <v-text-field
              label="Country code"
              v-model="countryCode"
              hint="Country code should contain at least 2 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-row>
          <v-col cols="5"/>
          <v-col cols="8">
              <v-text-field
              label="Alias "
              v-model="alias"
              hint="Alias should contain at least 2 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-row>
          <v-col cols="5"/>
          <v-col cols="9">
              <v-text-field
              label="End date"
              v-model="endDate"
              hint="End date should contain at least 6 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-combobox :items="certificates" :item-text="text" v-model="selectedCertificateType" 
          :label="label" hint="Choose certificate type."/>

          <v-combobox :items="certificates" :item-text="text" v-model="selectedCertificatePurpose"
          :label="label" hint="Choose certificate purpose."/>

          <v-row v-bind:hidden="selectedCertificateType=='ROOT'">
          <v-col cols="5"/>
          <v-col cols="10">
              <v-text-field
              label="Issuer serial number"
              v-model="issuerSerialNumber"
              hint="Issuer serial number should contain at least 3 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>

          <v-row v-bind:hidden="selectedCertificateType=='ROOT'">
          <v-col cols="5"/>
          <v-col cols="11">
              <v-text-field
              label="Issuer alias"
              v-model="issuerAlias"
              hint="Issuer alias should contain at least 2 characters!"
              color="light-blue darken-4"/>
          </v-col>
          </v-row>
      </v-container>

      <div class="createButton">
          <v-btn
              v-on:click="createCertificate" 
              color="light-blue darken-4"
              elevation="24"
              x-large
              raised
              rounded
          >Create</v-btn>
      </div>
  </div>
</template>

<script>
export default {
  name: 'NewCertificate',
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
        selectedCertificatePurpose: null,
        label: 'Certificates',
    }),
}
</script>

<style scoped>
.page {
  backgRound-color: #b5dafd;
}
</style>
