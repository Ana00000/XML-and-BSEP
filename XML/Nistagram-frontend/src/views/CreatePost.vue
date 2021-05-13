<template>
  <div>
    <div class="spacing" />
    <v-card width="400" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1">Location creation</h1>
      </v-card-title>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Longitude"
            v-model="longitude"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Latitude"
            v-model="latitude"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Country"
            v-model="country"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="City"
            v-model="city"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Street name"
            v-model="streetName"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Street number"
            v-model="streetNumber"
            prepend-icon="mdi-address-circle"
          />
        </v-form>
      </v-card-text>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Tag name"
            v-model="tagName"
            prepend-icon="mdi-address-circle"
          />
        </v-form>
      </v-card-text>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Description"
            v-model="postDescription"
            prepend-icon="mdi-address-circle"
          />
        </v-form>
      </v-card-text>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Path"
            v-model="path"
            prepend-icon="mdi-address-circle"
          />
          <v-select
            class="typeCombo"
            v-model="selectedType"
            hint="Choose your type."
            :items="types"
            item-text="state"
            :label="label1"
            return-object
            single-line
          />
        </v-form>
      </v-card-text>
      <v-card-actions class="justify-center mb-5">
        <v-btn color="info mb-5" v-on:click="createPost"> Create </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
export default {
  name: "Register",
  data: () => ({
    longitude: "",
    latitude: "",
    country: "",
    city: "",
    streetName: "",
    streetNumber: "",
    tagName: null,
    postDescription: "",
    locationId: null,
    path: "",
    types: ["PICTURE", "VIDEO"],
    selectedType: "PICTURE",
    label1: "Type",
    postId: null,
  }),
  methods: {
    createPost() {
      this.createTag();

      if (
        this.longitude == "" &&
        this.latitude == "" &&
        this.country == "" &&
        this.city == "" &&
        this.streetName == "" &&
        this.streetNumber == ""
      ) {
        this.createPostWithoutLocation();
        return;
      }
      if (
        !this.validLongitude() ||
        !this.validLatitude() ||
        !this.validCountry() ||
        !this.validCity() ||
        !this.validStreetName() ||
        !this.validStreetNumber()
      )
        return;

      this.$http
        .post("http://localhost:8083/", {
          longitude: this.longitude,
          latitude: this.latitude,
          country: this.country,
          city: this.city,
          streetName: this.streetName,
          streetNumber: this.streetNumber,
        })
        .then((response) => {
          this.locationId = response.data;
          this.createPostWithLocation();
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    createTag() {
      if (this.tagName == null) return;
      if (!this.validTag()) return;

      this.$http
        .post("http://localhost:8082/post_tag/", {
          name: this.tagName,
        })
        .then((response) => {
          console.log(response.data);
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    createPostWithLocation() {
      if (!this.validPostDescription()) return;
      this.$http
        .post("http://localhost:8084/single_post/", {
          description: this.postDescription,
          userID: localStorage.getItem("userId"),
          locationId: this.locationId,
        })
        .then((response) => {
          console.log(response.data);
          this.postId = response.data;
          this.createPostContent();
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    createPostWithoutLocation() {
      if (!this.validPostDescription()) return;
      this.$http
        .post("http://localhost:8084/single_post/", {
          description: this.postDescription,
          userID: localStorage.getItem("userId"),
        })
        .then((response) => {
          console.log(response.data);
          this.postId = response.data;
          this.createPostContent();
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    createPostContent() {
      if (!this.validPath()) return;

      setTimeout(5000);

      this.$http
        .post("http://localhost:8085/single_post_content/", {
          path: this.path,
          type: 0,
          single_post_id: this.postId
        })
        .then((response) => {
          console.log(response.data);
          alert("Successful creation.");
          window.location.href = "http://localhost:8081/";
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    validTag() {
      if (this.tagName.length > 30) {
        alert("Your tag name shouldn't contain more than 30 characters!");
        return false;
      } else if (this.tagName.match(/[!#$%^&*:'<>+/\\"]/g)) {
        alert("Your tag name shouldn't contain those special characters.");
        return false;
      } else if (this.tagName.match(/\d/g)) {
        alert("Your tag name shouldn't contain numbers!");
        return false;
      }
      return true;
    },
    validLongitude() {
      if (this.longitude.length < 2) {
        alert("Your longitude should contain at least 2 characters!");
        return false;
      } else if (this.longitude.length > 30) {
        alert("Your longitude shouldn't contain more than 30 characters!");
        return false;
      } else if (this.longitude.match(/[!@#$%^&*:'<>+-/\\"]/g)) {
        alert("Your longitude shouldn't contain those special characters.");
        return false;
      }
      return true;
    },
    validLatitude() {
      if (this.latitude.length < 2) {
        alert("Your latitude should contain at least 2 characters!");
        return false;
      } else if (this.latitude.length > 30) {
        alert("Your latitude shouldn't contain more than 30 characters!");
        return false;
      } else if (this.latitude.match(/[!@#$%^&*:'<>+-/\\"]/g)) {
        alert("Your latitude shouldn't contain those special characters.");
        return false;
      }
      return true;
    },
    validCountry() {
      if (this.country.length < 2) {
        alert("Your country should contain at least 2 characters!");
        return false;
      } else if (this.country.length > 30) {
        alert("Your country shouldn't contain more than 30 characters!");
        return false;
      } else if (this.country.match(/[!@#$%^&*:'<>+-/\\"]/g)) {
        alert("Your country shouldn't contain those special characters.");
        return false;
      }
      return true;
    },
    validCity() {
      if (this.city.length < 2) {
        alert("Your city should contain at least 2 characters!");
        return false;
      } else if (this.city.length > 30) {
        alert("Your city shouldn't contain more than 30 characters!");
        return false;
      } else if (this.city.match(/[!@#$%^&*:'<>+-/\\"]/g)) {
        alert("Your city shouldn't contain those special characters.");
        return false;
      }
      return true;
    },
    validStreetName() {
      if (this.streetName.length < 2) {
        alert("Your street name should contain at least 2 characters!");
        return false;
      } else if (this.streetName.length > 30) {
        alert("Your street name shouldn't contain more than 30 characters!");
        return false;
      } else if (this.streetName.match(/[!@#$%^&*:'<>+-/\\"]/g)) {
        alert("Your street name shouldn't contain those special characters.");
        return false;
      } else if (this.streetName.match(/\d/g)) {
        alert("Your street name shouldn't contain numbers!");
        return false;
      }
      return true;
    },
    validStreetNumber() {
      if (this.streetNumber.length < 1) {
        alert("Your street number should contain at least 1 character!");
        return false;
      } else if (this.streetNumber.length > 10) {
        alert("Your street number shouldn't contain more than 10 characters!");
        return false;
      } else if (this.streetNumber.match(/[!@#$%^&*:'<>+-/\\"]/g)) {
        alert("Your street number shouldn't contain those special characters.");
        return false;
      } else if (this.streetNumber.match(/\d/g) == null) {
        alert("Your street number should contain at least 1 number!");
        return false;
      }
      return true;
    },
    validPostDescription() {
      if (this.postDescription.length > 50) {
        alert(
          "Your post description shouldn't contain more than 50 characters!"
        );
        return false;
      }
      return true;
    },
    validPath() {
      if (this.path.length < 3) {
        alert("Your path should contain at least 3 characters!");
        return false;
      } else if (this.path.length > 50) {
        alert("Your path shouldn't contain more than 50 characters!");
        return false;
      } else if (this.path.match(/[!@#$%^&*'<>+"]/g)) {
        alert("Your path shouldn't contain those special characters.");
        return false;
      }
      return true;
    },
  },
};
</script>

<style scoped>
.spacing {
  height: 100px;
}

.typeCombo {
  width: 90%;
  margin-left: 10%;
}
</style>