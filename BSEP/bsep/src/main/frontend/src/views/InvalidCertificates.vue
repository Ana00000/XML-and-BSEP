<template>
  <v-container>
    <v-layout row wrap>
      <v-card loading class="mx-auto"
        id="certificateCard"
      >
      
        <v-toolbar color="light-blue darken-4" dark><v-toolbar-title class="flex text-center">INVALID CERTIFICATES</v-toolbar-title></v-toolbar>
        <v-list two-line>
          <v-list-item-group active-class="indigo--text" single>
            <template v-for="(certificate, id) in certificates">
              <v-list-item :key="certificate.id">
                <template>
                  <v-list-item-content class = "center text-wrap">
                  

                    <v-list-item-subtitle v-text="'Serial number: ' + certificate.serialNumber" />
                    <v-list-item-subtitle
                      v-text="
                        'Common name: '+ certificate.commonName
                      "
                    />
                    <v-list-item-subtitle
                      v-text="
                        'Given name: '+ certificate.givenName + ' --- Surname: ' + certificate.surname
                      "
                    />
                    <v-list-item-subtitle
                      v-text="
                        'Organization: ' + certificate.organization + ' --- Organization unit name: ' + certificate.organizationalUnitName
                      "
                    />
                    <v-list-item-subtitle
                      v-text="
                        'Organization email: ' + certificate.organizationEmail
                      "
                    />
                    <v-list-item-subtitle
                      v-text="'Alias: ' + certificate.alias"
                    />
                    <v-list-item-subtitle
                      v-text="'Country code: ' + certificate.countryCode"
                    />
                    <v-list-item-subtitle v-text="'Expired date: '+certificate.endDate" />
                  </v-list-item-content>
                </template>
              </v-list-item>
              <v-divider
                v-if="`A-${id}` < certificates.length - 1"
                :key="`A-${id}`"
              />
            </template>
          </v-list-item-group>
        </v-list>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "InvalidCertificates",
  data: () => ({
    certificates: [],
    user: null,
    token: null
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");
      this.getInvalidCertificates();
    },
    getInvalidCertificates() {
      this.$http
        .get("https://localhost:8080/certificate/allRevokedOrExpired",{
        headers:{
            'Authorization':"Bearer "+ this.token
        }})
        .then((res) => {
          this.certificates = res.data;
          console.log(res.data);
        })
        .catch((err) => console.log(err));
    },
  },
};
</script>

<style scoped>
.helloMessage {
  font-weight: bolder;
  font-size: 20px;
  height: 50px;
}

.center {
  
  padding: 10px;
  text-align: center;
}

#certificateCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>