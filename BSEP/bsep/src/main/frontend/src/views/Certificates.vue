<template>
  <v-container>
    <v-layout row wrap>

      <v-card style="width: 40%; height: 700px; overflow-y: scroll" id="certificateCard">
        <v-toolbar color="light-blue darken-4" dark> </v-toolbar>
        <v-list two-line>
          <v-list-item-group active-class="indigo--text" single>
            <template v-for="(certificate, id) in certificates">
              <v-list-item :key="certificate.id">
                <template>
                  <v-list-item-content>
                    <v-list-item-subtitle
                      v-text="
                        certificate.commonName + ' ' + certificate.givenName
                      "
                    />
                    <v-list-item-subtitle
                      v-text="
                        certificate.surname + ' ' + certificate.organization
                      "
                    />
                    <v-list-item-subtitle
                      v-text="
                        certificate.organizationalUnitName +
                        ' ' +
                        certificate.organizationEmail
                      "
                    />
                    <v-list-item-subtitle
                      v-text="certificate.countryCode + ' ' + certificate.alias"
                    />
                    <v-list-item-subtitle v-text="certificate.endDate" />
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
      
      <v-card class="mx-auto" style="width: 40%; height: 700px; overflow-y: scroll">
        <v-toolbar color="light-blue darken-4" dark> </v-toolbar>
        <v-list two-line>
          <v-list-item-group active-class="indigo--text" disabled>
            <template v-for="(certificate, id) in certificates">
              <v-list-item :key="certificate.id">
                <template>
                  <v-list-item-content>
                    <v-list-item-subtitle
                      v-text="
                        certificate.alias + ' ' + certificate.organization
                      "
                    />
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
  name: "Certificates",
  data: () => ({
    certificates: [],
    user: null,
    userEmail: null
   }),
   mounted() {
        this.init();
        //this.getValidCertificates();
    },
  methods: {
      init() {
    this.getValidCertificates()},
     getValidCertificates() {
          var userEmail = localStorage.getItem("userEmail");
          this.$http.get('http://localhost:8080/certificate/allValid/'+ userEmail)
            .then(res => {
              this.certificates = res.data;
              console.log(res.data);
              console.log("then usao");
              console.log(userEmail);
                          })
            .catch(err => console.log(err));
      },
  }
};
</script>

<style scoped>

.helloMessage {
  font-weight: bolder;
  font-size: 20px;
  height: 50px;
}

#certificateCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>