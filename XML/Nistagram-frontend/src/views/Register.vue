<template>
  <div>
    <v-card width="400" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1">Registration</h1>
      </v-card-title>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Username"
            v-model="username"
            prepend-icon="mdi-account-circle"
          />
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
          <v-select
            class="genderCombo"
            v-model="selectedGender"
            hint="Choose your gender."
            :items="genders"
            item-text="state"
            :label="label1"
            return-object
            single-line
          />
          <v-text-field
            label="Date of birth"
            v-model="dateOfBirth"
            hint="Date of birth should be in format yyyy-mm-dd"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Website"
            v-model="website"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Biography"
            v-model="biography"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Question"
            v-model="question"
            prepend-icon="mdi-name-circle"
          />
          <v-text-field
            label="Answer"
            v-model="answer"
            prepend-icon="mdi-name-circle"
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
    username: "",
    userEmail: "",
    password: "",
    passwordAgain: "",
    phoneNumber: "",
    firstName: "",
    lastName: "",
    genders: ["FEMALE", "MALE", "OTHER"],
    selectedGender: "FEMALE",
    label1: "Gender",
    dateOfBirth: "",
    website: "",
    biography: "",
    question: "",
    answer: "",
    users: [],
  }),
  methods: {
    register() {
      if (
        !this.validUsername() ||
        !this.validEmail() ||
        !this.validPassword() ||
        !this.validFirstName() ||
        !this.validLastName() ||
        !this.validPhoneNumber() ||
        !this.validDateOfBirth() ||
        !this.validWebsite() ||
        !this.validBiography() ||
        !this.validQuestion() ||
        !this.validAnswer()
      )
        return;

      this.$http
        .post("https://localhost:8080/api/user/registered_user/", {
          username: this.username,
          password: this.password,
          email: this.userEmail,
          phoneNumber: this.phoneNumber,
          firstName: this.firstName,
          lastName: this.lastName,
          gender: this.selectedGender,
          dateOfBirth: this.dateOfBirth + "T11:45:26.371Z",
          website: this.website,
          biography: this.biography,
          question: this.question,
          answer: this.answer,
          is_blocked: false,
        })
        .then((response) => {
          console.log(response.data.firstName);
          alert("Successfully registered.");
          window.location.href = "https://localhost:8081/logIn";
        })
        .catch((er) => {
          if (er.response.status == 409) alert("Username already exists.");
          else if (er.response.status == 417) alert("Email already exists.");
          else alert("Registration failed!");
          console.log("Error while registering in");
          console.log(er.response.data);
        });
    },
    validUsername() {
      if (this.username.length < 1) {
        alert("Your username should contain at least 1 character!");
        return false;
      } else if (this.username.length > 20) {
        alert("Your username shouldn't contain more than 20 characters!");
        return false;
      } else if (this.username.match(/[!@#$%^&*'<>+/\\"]/g)) {
        alert("Your username shouldn't contain special characters.");
        return false;
      }
      return true;
    },
    validEmail() {
      if (
        !/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(
          this.userEmail
        )
      ) {
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
      } else if (!this.password.match(/[a-z]/g)) {
        alert("Your password should contain at least one small letter.");
        this.passwordAgain = "";
        return false;
      } else if (!this.password.match(/[A-Z]/g)) {
        alert("Your password should contain at least one big letter.");
        this.passwordAgain = "";
        return false;
      } else if (!this.password.match(/[0-9]/g)) {
        alert("Your password should contain at least one number.");
        this.passwordAgain = "";
        return false;
      } else if (!this.password.match(/[!@#$%^&*.,:'+-/\\"]/g)) {
        alert(
          "Your password should contain at least one special character (other than <>)."
        );
        return false;
      } else if (this.password.match(/[<>]/g)) {
        alert("Your password shouldn't contain special character < or >.");
        return false;
      } else if (this.password.match(/[ ]/g)) {
        alert("Your password shouldn't contain spaces!");
        return false;
      } else if (this.password !== this.passwordAgain) {
        alert("Passwords don't match !!!");
        this.passwordAgain = "";
        return false;
      }
      return true;
    },
    validFirstName() {
      if (this.firstName.length < 2) {
        alert("Your first name should contain at least 2 characters!");
        return false;
      } else if (this.firstName.length > 20) {
        alert("Your first name shouldn't contain more than 20 characters!");
        return false;
      } else if (this.firstName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) {
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
        alert("Your last name should contain at least 2 characters!");
        return false;
      } else if (this.lastName.length > 35) {
        alert("Your last name shouldn't contain more than 35 characters!");
        return false;
      } else if (this.lastName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) {
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
      if (this.phoneNumber.match(/[a-zA-Z]/g)) {
        alert("Your phone number shouldn't contain letters.");
        return false;
      } else if (this.phoneNumber.match(/[ ]/g)) {
        alert("Your phone number shouldn't contain spaces!");
        return false;
      } else if (
        !/^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\\s./0-9]*$/.test(this.phoneNumber)
      ) {
        alert("Your phone number is not in right form!");
        return false;
      }
      return true;
    },
    validDateOfBirth() {
      if (this.dateOfBirth.match(/\d/g) == null) {
        alert("Your date of birth needs numbers!");
        return false;
      } else if (this.dateOfBirth.match(/\d/g).length < 6) {
        alert("Your date of birth should contain at least 6 numbers!");
        return false;
      } else if (this.dateOfBirth.match(/\d/g).length > 8) {
        alert("Your date of birth shouldn't contain more than 8 numbers!");
        return false;
      } else if (this.dateOfBirth.match(/[a-zA-Z]/g)) {
        alert("Your date of birth shouldn't contain letters.");
        return false;
      } else if (this.dateOfBirth.match(/[!@#$%^&*,:'/.<>+\\"]/g)) {
        alert(
          "Your date of birth shouldn't contain special character other than [-]."
        );
        return false;
      } else if (this.dateOfBirth.match(/[ ]/g)) {
        alert("Your date of birth shouldn't contain spaces!");
        return false;
      } else if (
        !this.dateOfBirth.match(/[1-2][0-9]{3}-[0-1][0-9]-[0-3][0-9]/g)
      ) {
        alert("Your date of birth is not set in right format.");
        return false;
      } else if (
        this.dateOfBirth.match(/[1-2][0-9]{3}-[0-1][0-9]-[0-3][0-9][-]+/g)
      ) {
        alert("Your date of birth can't contain - at end of input.");
        return false;
      }
      var dateOfBirthSplit = this.dateOfBirth.split("-");
      var dOBSYear = dateOfBirthSplit[0];
      var dOBSMonth = dateOfBirthSplit[1];
      var dOBSDay = dateOfBirthSplit[2];
      if (dOBSYear > 2021 || dOBSYear < 1900) {
        alert("Year of date of birth isn't valid");
        return false;
      } else if (dOBSMonth > 12 || dOBSMonth < 0) {
        alert("Month of date of birth isn't valid");
        return false;
      } else if (dOBSDay > 31 || dOBSDay < 1) {
        alert("Day of date of birth isn't valid");
        return false;
      }
      return true;
    },
    validWebsite() {
      if (this.website.length < 2) {
        alert("Your website should contain at least 2 characters!");
        return false;
      } else if (this.website.length > 25) {
        alert("Your website shouldn't contain more than 25 characters!");
        return false;
      } else if (this.website.match(/[!@#$%^&*,'<>+"]/g)) {
        alert("Your website shouldn't contain special characters.");
        return false;
      } else if (this.website.match(/[ ]/g)) {
        alert("Your website name shouldn't contain spaces!");
        return false;
      }
      return true;
    },
    validBiography() {
      if (this.biography.length < 1) {
        alert("Your biography should contain at least 1 character!");
        return false;
      } else if (this.biography.length > 100) {
        alert("Your biography shouldn't contain more than 100 characters!");
        return false;
      }
      return true;
    },
    validQuestion() {
      if (this.question.length < 1) {
        alert("Your question should contain at least 1 character!");
        return false;
      } else if (this.question.length > 70) {
        alert("Your question shouldn't contain more than 70 characters!");
        return false;
      }
      return true;
    },
    validAnswer() {
      if (this.answer.length < 1) {
        alert("Your answer should contain at least 1 character!");
        return false;
      } else if (this.answer.length > 70) {
        alert("Your answer shouldn't contain more than 70 characters!");
        return false;
      }
      return true;
    },
  },
};
</script>

<style scoped>
.genderCombo {
  width: 45%;
  margin-left: 25%;
}

.userCategoryCombo {
  width: 45%;
  margin-left: 25%;
}
</style>