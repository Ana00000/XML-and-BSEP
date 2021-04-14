<template>
  <v-container fluid class="container">
    <h1 class="text-center display-1 mt-10 mb-10">
      Check the validity of the certificate based on the serial number
    </h1>
    <v-row aria-rowspan="1">
      <v-col cols="4" />
      <v-col cols="4">
        <v-text-field
          label="Input certificate serial number"
          v-model="serialNumber"
          color="light-blue darken-4"
        />
      </v-col>
    </v-row>
    <v-row rows="2">
      <v-flex text-center align-center>
        <v-btn
          class="mb-10 mt-10"
          color="primary"
          elevation="9"
          large
          v-on:click="checkValidity"
        >
          Check validity
        </v-btn>
      </v-flex>
    </v-row>
    <v-row rows="3" class="mt-20">
      <v-col cols="4" />
      <v-col cols="4">
        <v-text-field
          v-show="isShowValidity"
          class="mt-50"
          label="Certificate status"
          v-model="certificateStatus"
          color="light-blue darken-4"
          readonly
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Check certificate validity",
  data: () => ({
    certificate: null,
    serialNumber: "",
    certificateStatus: null,
    isShowValidity: false,
    token: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");
    },
    checkValidity() {
      this.$http
        .get(
          "http://localhost:8080/certificate/" +
            this.serialNumber,
          {
            headers: {
              'Authorization': "Bearer " + this.token
            }
          }
        )
        .then((resp) => {
          console.log(resp.data.certificateStatus);
          if (resp.data.certificateStatus == "VALID") {
            this.certificateStatus = "Certificate is valid";
          }
          else if (resp.data.certificateStatus == "REVOKED") {
            this.certificateStatus = "Certificate is revoked";
          }
          else if (resp.data.certificateStatus == "EXPIRED"){
            this.certificateStatus = "Certificate is expired";
          }
          else
          {
            this.certificateStatus = "Certificate not found with this serial number";
          }
          this.isShowValidity = true;
          
        })
        .catch((err) => {
          if (this.serialNumber == "") {
            alert("Please enter the certificate serial number");
          }
          this.isShowValidity = false;
          console.log(err);
        });
    },
  },
};
</script>
