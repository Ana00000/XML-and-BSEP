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
      <v-card-actions class="justify-center mb-5">
        <v-btn color="info mb-5" v-on:click="create"> Create </v-btn>
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
    tagName: "",
    postDescription: ""
  }),
  methods: {
    create() {
      //this.createLocation();
      //this.createTag();
      //this.createPost();
    },
    createLocation() {/*
      if (
        !this.validLongitude() ||
        !this.validLatitude() ||
        !this.validCountry() ||
        !this.validCity() ||
        !this.validStreetName() ||
        !this.validStreetNumber()
      )
        return;*/

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
          console.log(response.data);
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    createTag() {
      //if (!this.validTag()) return;
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
    createPost() {
      this.$http
        .post("http://localhost:8084/single_post/", {
          description: this.postDescription,
          userID: localStorage.getItem("userId"),
          locationId: "8ac6f1b4-4b59-4798-a420-a154e6ce40d6"
        })
        .then((response) => {
          console.log(response.data);
          //alert("Successful creation.");
          //window.location.href = "http://localhost:8081/";
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    }
  },
};
</script>

<style scoped>
.spacing {
  height: 100px;
}
</style>