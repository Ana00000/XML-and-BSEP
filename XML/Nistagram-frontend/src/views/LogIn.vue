<template>
  <v-card style="margin-top: 10%" width="20%" class="mx-auto">
    <v-card-title class="justify-center">
      <h1 class="display-1 mt-5">Login</h1>
    </v-card-title>
    <v-card-text>
      <v-form class="mx-auto ml-5 mr-5">
        <v-text-field
          label="Username"
          v-model="username"
          prepend-icon="mdi-account-circle"
        /> 
        <v-text-field
          :type="showPassword ? 'text' : 'password'"
          label="Password"
          v-model="password"
          prepend-icon="mdi-lock"
          :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
          @click:append="showPassword = !showPassword"
        />
        <a class="center" href="/recoverPasswordEmail">You forget your password?</a>
      </v-form>
      
    </v-card-text>
    <div >
        
    </div>
    <v-card-actions>
      <v-btn color="primary" class="mx-auto mb-5" large v-on:click="logIn">
        Log in
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "Login",
  data: () => ({
    showPassword: false,
    username: "",
    password: "",
    users: [],
  }),
  computed: {
    user() {
      return { username: this.username, password: this.password };
    },
  },
  methods: {
    logIn() {
      this.$http
        .post("https://localhost:8080/api/user/login/", {
           username: this.username,
           password: this.password
        })
        .then((resp) => {
          console.log(resp.data);
          localStorage.setItem("username", this.user.username);
          localStorage.setItem("token", resp.data.token)
          localStorage.setItem("userId", resp.data.id)
          localStorage.setItem("userType", resp.data.userType)
         
          this.$http
          .get("https://localhost:8080/api/settings/find_profile_settings_by_user_id/"+resp.data.id)
          .then((resp) => {
          console.log("FOUND PROFILE SETTINGS")
          console.log("USER PRIVACY")
          console.log(resp.data.user_visibility)

          if (resp.data.user_visibility == "PRIVATE_VISIBILITY"){
            console.log("u 0")
             localStorage.setItem("userPrivacy", "PRIVATE")
          }else if (resp.data.user_visibility == "PUBLIC_VISIBILITY"){
            console.log("u 1")
             localStorage.setItem("userPrivacy", "PUBLIC")
          }

          window.location.href = "https://localhost:8081/";

         }).catch(console.log);
         
        })
        .catch((er) => {
          alert("Invalid username and/or password! Please, try again!");
          this.username = '';
          this.password = '';
          console.log(er.response.data);
        });
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