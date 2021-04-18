<template>
  <div>
    <v-card width="400" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1">Registration</h1>
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
          <v-text-field
            :type="showPassword ? 'text' : 'password'"
            label="Repeat password"
            v-model="passwordAgain"
            prepend-icon="mdi-lock"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
          />
          <v-text-field
            label="First name"
            v-model="firstName"
            prepend-icon="mdi-name-circle"
          />
          <v-text-field
            label="Last name"
            v-model="lastName"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Phone number"
            v-model="phoneNumber"
            prepend-icon="mdi-address-circle"
          />
        </v-form>
      </v-card-text>
      <v-card-actions class="justify-center mb-5">
        <v-btn color="info mb-5" v-on:click="register"> Register </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
export default {
  name: "Register",
  data: () => ({
    showPassword: false,
    userEmail: "",
    password: "",
    passwordAgain: "",
    phoneNumber: "",
    firstName: "",
    lastName: "",
    users: []
  }),
  methods: {
    register() {
      if (!this.validEmail() || !this.validPassword() || !this.validFirstName() || !this.validLastName() || !this.validPhoneNumber()) return;
      this.$http
        .post("https://localhost:8080/users/register", {
          userEmail: this.userEmail,
          lastName: this.lastName,
          password: this.password,
          firstName: this.firstName,
          phoneNumber: this.phoneNumber,
          typeOfUser: "USER",
        })
        .then((response) => {
          console.log(response.data.firstName);
          alert("Successfully registered.");
          window.location.href = "https://localhost:8081/login";
        })
        .catch((er) => {
          alert("Email already exists.");
          console.log("Error while registering in");
          console.log(er.response.data);
        });
    },
    validEmail() {
      if (!/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(this.userEmail)) {
        alert("You have entered an invalid email address!");
        return false;
      } else if (this.userEmail.length > 35) {
        alert("Your email shouldn't contain more than 35 characters!");
        return false;
      }
      return true;
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
          this.passwordAgain='';
          return false;
      } else if(!this.password.match(/[A-Z]/g)) {
          alert("Your password should contain at least one big letter.");
          this.passwordAgain='';
          return false;
      } else if(!this.password.match(/[0-9]/g)) {
          alert("Your password should contain at least one number.");
          this.passwordAgain='';
          return false;
      } else if(!this.password.match(/[!@#$%^&*.,:'+-/\\"]/g)) { 
          alert("Your password should contain at least one special character (other than <>).");
          return false;
      } else if(this.password.match(/[<>]/g)) { 
          alert("Your password shouldn't contain special character < or >.");
          return false;
      } else if (this.password.match(/[ ]/g)) {
        alert("Your password shouldn't contain spaces!");
        return false;
      } else if (this.password !== this.passwordAgain){
          alert("Passwords don't match !!!");
          this.passwordAgain='';
          return false;
      }
      return true;
    },
    validFirstName() {
      if (this.firstName.length < 2) {
        alert("Your first name should contain at least 2 character!");
        return false;
      } else if (this.firstName.length > 20) {
        alert("Your first name shouldn't contain more than 20 characters!");
        return false;
      } else if(this.firstName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) { 
          alert("Your first name shouldn't contain special characters.");
          return false;
      } else if (this.firstName.match(/[ ]/g)) {
        alert("Your first name shouldn't contain spaces!");
        return false;
      } else if (this.firstName.match(/\d/g)) {
        alert("Your first name shouldn't contain numbers!");
        return false;
      } else if (!/^[A-Z][a-z]+$/.test(this.firstName)) {
        alert("Your first name needs to have one upper letter at the start!");
        return false;
      }
      return true;
    },
    validLastName() {
      if (this.lastName.length < 2) {
        alert("Your last name should contain at least 2 character!");
        return false;
      } else if (this.lastName.length > 35) {
        alert("Your last name shouldn't contain more than 35 characters!");
        return false;
      } else if(this.lastName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) { 
          alert("Your last name shouldn't contain special characters.");
          return false;
      } else if (this.lastName.match(/[ ]/g)) {
        alert("Your last name shouldn't contain spaces!");
        return false;
      } else if (this.lastName.match(/\d/g)) {
        alert("Your last name shouldn't contain numbers!");
        return false;
      } else if (!/^[A-Z][a-z]+$/.test(this.lastName)) {
        alert("Your last name needs to have one upper letter at the start!");
        return false;
      }
      return true;
    },
    validPhoneNumber() {
      if(this.phoneNumber.match(/[a-zA-Z]/g)) {
          alert("Your phone number shouldn't contain letters.");
          return false;
      } else if (this.phoneNumber.match(/[ ]/g)) {
          alert("Your phone number shouldn't contain spaces!");
          return false;
      } else if (!/^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\\s./0-9]*$/.test(this.phoneNumber)) {
          alert("Your phone number is not in right form!");
          return false;
      }
      return true;
    }
  },
};
</script>