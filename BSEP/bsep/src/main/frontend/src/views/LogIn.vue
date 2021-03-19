<template>
 <div class="page">
   <div class="pageUp"/>
    <v-card width="400" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1 mt-5">Login</h1>
      </v-card-title>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Username/email"
            v-model="username"
            prepend-icon="mdi-account-circle"/>
          <v-text-field
            :type="showPassword ? 'text' : 'password'"
            label="Password"
            v-model="password"
            prepend-icon="mdi-lock"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"/>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-btn color="success" class="mx-auto mb-5" v-on:click="logIn">
          Log in
        </v-btn>
      </v-card-actions>
    </v-card>
    <div class="pageDown"/>
  </div>
</template>

<script>
import axios from 'axios';
function redirectLogedUser(){
    var tokenString= '';
    tokenString = localStorage.getItem("token");
    const config = {
        headers: { Authorization: `Bearer ${tokenString}` }
    };
    axios.get( 
        'api/users/redirectMeToMyHomePage',
        config
    ).then((response) => {
      console.log(response);
      window.location.href = response.data;
    }, (error) => {
      console.log(error);
    });
}

export default {
  name: 'Login',
  data: () => ({
    showPassword: false,
    username: '',
    password: '',
    users: []
  }),
  computed: {
    user() {
      return {'email': this.username, 'password': this.password}
    }
  },
  methods: {
    logIn() {
      this.$http.post('api/users/login', this.user)
      .then(resp => {
        console.log(resp.data);
        localStorage.setItem("token", resp.data.accessToken);
        redirectLogedUser();
      })
      .catch(er => {
        console.log('Error while logging in');
        console.log(er.response.data);
      })
    }
  }
};
</script>

<style>
.pageUp {
height: 200px;
}

.pageDown {
height: 364px;
}

.page {
  backgRound-color: #b5dafd;
}
</style>
