<template>
  <div>
    <br />
    <v-container fluid class="container">
      <v-row aria-rowspan="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Certificate purpose type"
            v-model="certificatePurposeType"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>

      <v-row aria-rowspan="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Common name"
            v-model="commonName"
            color="light-blue darken-4"
            readonly
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
            readonly
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
            readonly
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
            readonly
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
            readonly
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
            readonly
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
            readonly
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
            readonly
          />
        </v-col>
      </v-row>

      <v-row rows="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="End date"
            v-model="endDate"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>

      <v-row rows="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Certificate type"
            v-model="certificateType"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>

      <div class="center">
        <v-btn
          v-on:click="createCertificate"
          color="primary"
          large
          elevation="20"
          v-show="!isHidden"
          >Create</v-btn
        >
      </div>

      <div class="center">
        <v-btn
          v-on:click="revokeCertificate"
          v-show="!isHiddenRole"
          color="primary"
          large
          elevation="20"
          >Revoke</v-btn
        >
      </div>

      <div class="center">
        <v-btn
          v-on:click="downloadCertificate"
          color="primary"
          large
          elevation="20"
          >Download</v-btn
        >
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "SelectedCertificate",
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
    certificateType: null,
    certificatePurposeType: "",
    serialNbr: null,
    isHiddenRole: false,
    isHidden: false,
    token: null
  }),
  mounted() {
    this.serialNbr = localStorage.getItem("serialNumber");
    this.token = localStorage.getItem("token");
    console.log(this.serialNbr);
    this.init();
  },
  methods: {
    init() {
      this.getCertificate();
    },
    getCertificate() {
      console.log(this.serialNbr);
      this.$http
        .get(
          "http://localhost:8080/certificate/getCertificate/" + this.serialNbr
        ,{
        headers:{
            'Authorization':"Bearer "+ this.token
        }})
        .then((resp) => {
          this.setCertificateInfo(resp.data);
          console.log(resp.data);
          if (resp.data.certificateType == "ENDENTITY" || (resp.data.certificateType == "INTERMEDIATE"  && localStorage.getItem('role') != 'ADMIN')){
            this.isHidden = true;
          }
          if (localStorage.getItem('role') != 'ADMIN'){
            this.isHiddenRole = true;
          }
        })
        .catch(console.log("Didn't set certificate info!"));
    },
    setCertificateInfo(item) {
      this.commonName = item.commonName;
      this.givenName = item.givenName;
      this.surname = item.surname;
      this.organization = item.organization;
      this.organizationalUnitName = item.organizationalUnitName;
      this.organizationEmail = item.organizationEmail;
      this.countryCode = item.countryCode;
      this.alias = item.alias;
      this.endDate = item.endDate;
      this.certificateType = item.certificateType;
      this.certificatePurposeType = item.certificatePurposeType;
      this.issuerSerialNumber = item.issuerSerialNumber;
      this.issuerAlias = item.issuerSerialNumber;
    },
    createCertificate() {
       localStorage.setItem("endDate", this.endDate);
       localStorage.setItem("issuerAlias", this.alias);
       localStorage.setItem("issuerSerialNumber", this.serialNbr);
       window.location.href = "http://localhost:8081/createOtherCertificates";
    },
    revokeCertificate() {
      this.$http
        .put(
          "http://localhost:8080/certificate/revokeCertificate/" + this.serialNbr
          , {}, {
            headers:{
              'Authorization':"Bearer "+ this.token
            }
        })
        .then((resp) => {
          console.log(resp.data);
          alert("You have successfully revoked the certificate.");
          window.location.href = "http://localhost:8081/invalidCertificates";
        })
        .catch((err) => console.log(err));
    },
    downloadCertificate() {
      this.$http
        .get("http://localhost:8080/certificate/loadToFile/" + this.serialNbr, {
        headers:{
            'Authorization':"Bearer "+ this.token
        }})
        .then((resp) => {
          console.log(resp.data);
          alert("Successfully downloaded the certificate.");
        });
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
