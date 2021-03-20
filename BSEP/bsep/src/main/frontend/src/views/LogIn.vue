<template>
  <v-card style="margin-top: 10%;" width="20%" class="mx-auto">
    <v-card-title class="justify-center">
      <h1 class="display-1 mt-5">Login</h1>
    </v-card-title>
    <v-card-text>
      <v-form class="mx-auto ml-5 mr-5">
        <v-text-field
          label="Email"
          v-model="userEmail"
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
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" class="mx-auto mb-5" large v-on:click="logIn">
        Log in
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import axios from "axios";
function redirectLoggedUser() {
  var tokenString = "";
  tokenString = localStorage.getItem("token");
  const config = {
    headers: { Authorization: `Bearer ${tokenString}` },
  };
  axios.get("http://localhost:8080/users/redirectMeToMyHomePage", config).then(
    (response) => {
      console.log(response);
      window.location.href = response.data;
    },
    (error) => {
      console.log(error);
    }
  );
}

export default {
  name: "Login",
  data: () => ({
    showPassword: false,
    userEmail: "",
    password: "",
    users: [],
  }),
  computed: {
    user() {
      return { userEmail: this.userEmail, password: this.password };
    },
  },
  methods: {
    logIn() {
      this.$http
        .post("http://localhost:8080/users/login", this.user)
        .then((resp) => {
          console.log(resp.data);
          localStorage.setItem("token", resp.data.accessToken);
          localStorage.setItem("userEmail", this.user.userEmail);
          redirectLoggedUser();
        })
        .catch((er) => {
          console.log("Error while logging in");
          console.log(er.response.data);
        });
    }
  },
};
</script>