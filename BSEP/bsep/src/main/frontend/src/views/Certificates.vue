<template>
  <div class="page">
    <br/><br/><br/>
    <v-container >
      <v-layout row wrap>
      <v-card
          class="mx-auto" style="width: 70%; height: 760px; overflow-y: scroll">
          <v-toolbar
          color="light-blue darken-4" dark>
          </v-toolbar>
          <v-list two-line>
          <v-list-item-group
              active-class="indigo--text"
              multiple>
              <template v-for="(certificate, id) in certificates">
              <v-list-item :key="certificate.id">
                  <template>  
                  <v-list-item-content>
                      <v-list-item-subtitle v-text="certificate.commonName+' '+certificate.givenName"/>
                      <v-list-item-subtitle v-text="certificate.surname+' '+certificate.organization"/>
                      <v-list-item-subtitle v-text="certificate.organizationalUnitName+' '+certificate.organizationEmail"/>
                      <v-list-item-subtitle v-text="certificate.countryCode+' '+certificate.alias"/>
                      <v-list-item-subtitle v-text="certificate.endDate"/>
                  </v-list-item-content>
                  </template>
              </v-list-item>
              <v-divider v-if="`A-${id}` <certificates.length-1" :key="`A-${id}`"/>
              </template>
          </v-list-item-group>
          </v-list>
      </v-card>
      </v-layout>
    </v-container>
    <br/><br/><br/>
  </div>  
</template>

<script>
export default {
  name: 'Certificates',
  components: {
    
  }
}
</script>

<style scoped>
.helloMessage {
    font-weight: bolder;
    font-size: 20px;
    height: 50px;
}
.page {
  backgRound-color: #b5dafd;
}
</style>