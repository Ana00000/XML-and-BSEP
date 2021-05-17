<template>
  <div>
    <br /><br />

    <v-container fluid class="container">
      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field label="First name" v-model="firstName" color="indigo" v-bind:readonly="isReadOnly"/>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field label="Last name" v-model="lastName" color="indigo" v-bind:readonly="isReadOnly"/>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field label="Email" v-model="email" color="indigo" v-bind:readonly="isReadOnly"/>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Phone number"
            v-model="phoneNumber"
            color="indigo"
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-select
            class="genderCombo"
            v-model="selectedGender"
            hint="Choose your gender."
            :items="genders"
            item-text="state"
            :label="label1"
            return-object
            single-line
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Date of birth"
            v-model="dateOfBirth"
            color="indigo"
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field label="Username" v-model="username" color="indigo" v-bind:readonly="isReadOnly"/>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field label="Website" v-model="website" color="indigo" v-bind:readonly="isReadOnly"/>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field label="Biography" v-model="biography" color="indigo" v-bind:readonly="isReadOnly"/>
        </v-col>
      </v-row>

      <br /><br />
    </v-container>

    <div class="changeButton">
      <v-btn
        v-if="!isHiddenChange"
        v-on:click="isHiddenUpdate = false, isHiddenChange = true, isHiddenCancel = false, isReadOnly = false"
        color="info mb-5"
        elevation="24"
        x-large
        raised
        rounded
        >Change</v-btn
      >
    </div>

    <div class="updateButton">
      <v-btn
        v-if="!isHiddenUpdate"
        v-on:click="updateProfile"
        color="info mb-5"
        elevation="24"
        x-large
        raised
        rounded
        >Save</v-btn
      >
    </div>

    <div class="cancelButton">
      <v-btn
        v-if="!isHiddenCancel"
        v-on:click="cancelChanges"
        color="info mb-5"
        elevation="24"
        x-large
        raised
        rounded
        >Cancel</v-btn
      >
    </div>
  </div>
</template>

<script>
export default {
  name: "UpdateProfile",
  data: () => ({
    id: "",
    firstName: "",
    lastName: "",
    email: "",
    phoneNumber: "",
    genders: ["FEMALE", "MALE", "OTHER"],
    selectedGender: "FEMALE",
    dateOfBirth: "",
    username: "",
    website: "",
    biography: "",
    userType: "",
    isHiddenChange: false,
    isHiddenUpdate: true,
    isHiddenCancel: true,
    isReadOnly: true,
    user: []
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.id = localStorage.getItem("userId");
      this.token = localStorage.getItem("token");

      this.$http
        .get("http://localhost:8080/find_user_by_id?id=" + this.id)
        .then((resp) => {
          this.user = resp.data;
          console.log(resp.data)
          this.setUserInfo(this.user);
        })
        .catch(console.log);
    },
    setUserInfo(item) {
      this.firstName = item.firstName
      this.lastName = item.lastName
      this.email = item.email
      this.phoneNumber = item.phoneNumber
      this.gender = item.gender
      this.dateOfBirth = item.dateOfBirth
      this.username = item.username
      this.website = item.website
      this.biography = item.biography
      this.userType = item.userType
    },
    cancelChanges() {
      window.location.href = "http://localhost:8081/updateProfile"
    },
    updateProfile() {
      if (
        !this.validFirstName() ||
        !this.validLastName() ||
        !this.validEmail() ||
        !this.validPhoneNumber() ||
        !this.validDateOfBirth() ||
        !this.validUsername() ||
        !this.validWebsite() ||
        !this.validBiography()
      )
        return;

      this.$http
        .post("http://localhost:8080/update_user_profile_info/", {
          id: this.id,
          firstName: this.firstName,
          lastName: this.lastName,
          email: this.email,
          phoneNumber: this.phoneNumber,
          gender: this.selectedGender,
          dateOfBirth: this.dateOfBirth,
          username: this.username,
          website: this.website,
          biography: this.biography,
          userType: this.userType
        },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        
        )
        .then((resp) => {
          console.log(resp.data);
          alert("Successfully updated profile informations!");
        })
        .catch((err) => console.log(err));
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
    validEmail() {
      if (!/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(this.email)) {
        alert("You have entered an invalid email address!");
        return false;
      } else if (this.email.length > 35) {
        alert("Your email shouldn't contain more than 35 characters!");
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
      } else if (this.dateOfBirth.match(/\d/g).length > 30) {
        alert("Your date of birth shouldn't contain more than 30 numbers!");
        return false;
      } else if (this.dateOfBirth.match(/[!@#$%^&*,'/<>\\"]/g)) {
        alert(
          "Your date of birth shouldn't contain special character other than [-,:,+ or .]."
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
      if (dOBSYear > 3000 || dOBSYear < 1900) {
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
  },
};
</script>

<style scoped>
.changeButton {
  position: absolute;
  right: 275px;
  top: 100px;
}

.updateButton {
  position: absolute;
  right: 250px;
  top: 300px;
}

.cancelButton {
  position: absolute;
  right: 240px;
  top: 500px;
}
</style>