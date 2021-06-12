<template>
  <v-card style="margin-top: 10%" width="20%" class="mx-auto">
    
    <v-img height = "150" width = "250" :src = "require('../assets/403.png')">
    </v-img>
    <v-card-title class="justify-center">
      <h1 class="display-1 mt-5">403 - You shall not pass!</h1>
      <h3>You are unauthorized to see this page</h3>
    </v-card-title>
    <div >
        
    </div>
    <v-card-actions>
      <v-btn color="primary" class="mx-auto mb-5" large v-on:click="back">
        Back Home
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "UnauthorizedPage",
  data: () => ({
    showPassword: false,

  }),
  computed: {
    user() {
      return { username: this.username, password: this.password };
    },
  },
  methods: {
      back(){
           window.location.href =
        "https://localhost:8081/postAlbumByIdWithoutActivity";
      }
    
    
  }
};
</script>

<style scoped>
.helloMessage {
  font-weight: bolder;
  font-size: 20px;
  height: 50px;
}

.center {
  padding: 31px;
  text-align: center;
}

#certificateCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>