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
      console.log(this.username);
      console.log(this.password);
      this.$http
        .post("http://localhost:8080/login/", {
           username: this.username,
           password: this.password
        })
        .then((resp) => {
          console.log(resp.data);
          localStorage.setItem("username", this.user.username);
          localStorage.setItem("password", this.user.password);
          localStorage.setItem("token", resp.data)
          window.location.href = "http://localhost:8081/";
        })
        .catch((er) => {
          alert("Invalid username and/or password! Please, try again!");
          this.username = '';
          this.password = '';
          console.log(er.response.data);
        });
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