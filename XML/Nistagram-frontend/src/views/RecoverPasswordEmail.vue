<template>
  <div>
    <br />
    <v-container fluid class="container">

      <v-row rows="2" >
        <v-col cols="5" />
        <v-col cols="2">
          <v-text-field 
            label="Email for recover password"
            v-model="emailRecover"
            hint="Email for recover password should be in format, for example,julias@uns.ac.rs"
            color="light-blue darken-4"
          />
        </v-col>
      </v-row>

      
    </v-container>

    <div class="center">
        <v-btn class="center" v-on:click="recoverPassword" color="primary" large elevation="20">Send recover email</v-btn>
        <br />
        <v-btn class="center" v-on:click="redirectToLogIn" color="primary" large elevation="20">Cancel</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  name: "RecoverPasswordEmail",
  data: () => ({
    emailRecover:""
  }),
  methods: {
    recoverPassword() {
      if (!this.validCertificate()) return;
      this.$http
        .post(
          "https://localhost:8080/api/user/recovery_password/",
          {
              email: this.emailRecover
          }
          
        )
        .then((resp) => {
          console.log(resp.data);
          alert("Send email on "+this.emailRecover+"!");
          window.location.href = "https://localhost:8081/";
        })
        .catch((err) => {
          alert("Your email is incorect or don't exist account");
          console.log(err.response.data);
        });
    },
    redirectToLogIn() {
      window.location.href = "https://localhost:8081/logIn";
    },
    validCertificate() {
      if (this.validRecoverEmail()) return true;
      return false;
    },
    validRecoverEmail() {
      if (!/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(this.emailRecover)) {
        alert("You have entered an invalid recover email address!");
        return false;
      } else if (this.emailRecover.length > 35) {
        alert("Recover email address shouldn't contain more than 35 characters!");
        return false;
      }
      return true;
    }
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
  margin:5px;
}

#certificateCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>