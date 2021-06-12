<template>
  <div>
    <br />
    <v-container fluid class="container">
      <h1 v-show="showHeader" class="center">BAD TOKEN</h1>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
           <v-text-field
            :type="showPassword ? 'text' : 'password'"
            label="Password"
            v-model="password" 
            v-show="showFields"
            prepend-icon="mdi-lock"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
            />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
            <v-text-field
            :type="showConfirmPassword ? 'text' : 'password'"
            label="Confirmation Password"
            v-show="showFields"
            v-model="confirmPassword" 
            prepend-icon="mdi-lock"
            :append-icon="showConfirmPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showConfirmPassword = !showConfirmPassword"
            />
        </v-col>
      </v-row>
    </v-container>

    <div class="center">
      <v-btn  v-show="showFields" v-on:click="saveNewPassword" color="primary" large elevation="20"
        >Save password</v-btn
      >
    </div>
  </div>
</template>

<script>
export default {
  name: "ChangePasswordByToken",
  data: () => ({
    confirmationToken:"",
    userId:"",
    showPassword:false,
    showConfirmPassword:false,
    password:"",
    confirmPassword:"",
    userEmail:"",
    showFields:false,
    showHeader:false,
    recoveryPasswordTokenId:"",
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      var hrefPath = window.location.href;
      var hrefPaths = [];
      hrefPaths=  hrefPath.split('/');
      this.confirmationToken = hrefPaths[4];
      this.userId = hrefPaths[5];
      this.$http
        .post("https://localhost:8080/api/user/verify_recovery_password_token/", {
            recovery_password_token: this.confirmationToken,
            user_id: this.userId
        })
        .then((res) => {
          alert(res.data)
          this.userEmail = res.data.user_email;
          this.recoveryPasswordTokenId = res.data.recovery_password_token_id;
          this.showFields = true;
          this.showHeader = false;
          alert("Successfully founded user!")
        })
        .catch((err) => {
          console.log(err);
          this.showHeader = true;
          this.showFields = false;
        });
    },
    saveNewPassword() {
        if (!this.validPassword() || !this.validConfirmPassword() || !this.validPasswords()) {
            return;
        }
        this.$http
        .post("https://localhost:8080/api/user/change_user_password/", {
            email: this.userEmail,
            password: this.password,
            confirmed_password: this.confirmPassword,
            recovery_password_token_id: this.recoveryPasswordTokenId,
        })
        .then((res) => {
          console.log(res);
          alert("You are successfully change password on your account! You can log in on system!");
          window.location.href = "https://localhost:8081/logIn";
        })
        .catch((err) => {
          console.log(err);
          alert("Password doesn't changed! Please, contact system admin!");
        });
    },
    redirectToLogIn() {
      window.location.href = "https://localhost:8081/logIn";
    },
    validPassword() {
      if (this.password.length < 10) {
            alert("Your password should contain at least 10 character!");
            return false;
      } else if (this.password.length > 30) {
            alert("Your password shouldn't contain more than 30 characters!");
            return false;
      } else if(!this.password.match(/[a-z]/g)) {
            alert("Your password should contain at least one small letter.");
            this.confirmPassword='';
            return false;
      } else if(!this.password.match(/[A-Z]/g)) {
            alert("Your password should contain at least one big letter.");
            this.confirmPassword='';
            return false;
      } else if(!this.password.match(/[0-9]/g)) {
            alert("Your password should contain at least one number.");
            this.confirmPassword='';
            return false;
      } else if(!this.password.match(/[!@#$%^&*.,:'"]/g)) { 
            alert("Your password should contain at least one special character.");
            return false;
      } 
      return true;
    },
    validConfirmPassword() {
      if (this.confirmPassword.length < 10) {
        alert("Your password should contain at least 10 character!");
        return false;
      } else if (this.confirmPassword.length > 30) {
        alert("Your password shouldn't contain more than 30 characters!");
        return false;
      } else if(!this.confirmPassword.match(/[a-z]/g)) {
          alert("Your confirm password should contain at least one small letter.");
          this.confirmPassword='';
          return false;
      } else if(!this.confirmPassword.match(/[A-Z]/g)) {
          alert("Your confirm password should contain at least one big letter.");
          this.confirmPassword='';
          return false;
      } else if(!this.confirmPassword.match(/[0-9]/g)) {
          alert("Your confirm password should contain at least one number.");
          this.confirmPassword='';
          return false;
      } else if(!this.confirmPassword.match(/[!@#$%^&*.,:'"]/g)) { 
          alert("Your confirm password should contain at least one special character.");
          return false;
      }
      return true;
    },
    validPasswords() {
      if (this.confirmPassword !== this.password) {
        alert("Passwords don't match !!!");
        this.confirmPassword ='';
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
}

#certificateCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>