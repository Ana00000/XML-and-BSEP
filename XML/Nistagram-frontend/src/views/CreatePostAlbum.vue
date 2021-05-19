<template>
  <div>
    <div class="spacing" />
    <v-card width="600" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1">Post album creation</h1>
      </v-card-title>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Longitude"
            v-model="longitude"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenLocation"
          />
          <v-text-field
            label="Latitude"
            v-model="latitude"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenLocation"
          />
          <v-text-field
            label="Country"
            v-model="country"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenLocation"
          />
          <v-text-field
            label="City"
            v-model="city"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenLocation"
          />
          <v-text-field
            label="Street name"
            v-model="streetName"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenLocation"
          />
          <v-text-field
            label="Street number"
            v-model="streetNumber"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenLocation"
          />
          <v-text-field
            label="Description"
            v-model="postAlbumDescription"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenDescription"
          />
          <v-select
            class="typeCombo"
            v-model="selectedType"
            v-if="!isHiddenContent"
            hint="Choose your type."
            :items="types"
            item-text="state"
            :label="label1"
            return-object
            single-line
          />
          <iframe
            name="dummyframe"
            id="dummyframe"
            style="display: none"
          ></iframe>
          <form
            action="http://localhost:8085/uploadPostAlbumMedia/"
            enctype="multipart/form-data"
            method="post"
            v-if="!isHiddenContent"
            target="dummyframe"
          >
            <input
              type="file"
              accept="image/*,video/*,.mkv"
              name="myPostAlbumFile"
            />
            <input type="submit" value="upload file .." />
          </form>
          <v-select
            class="typeCombo"
            v-model="selectedType"
            v-if="!isHiddenAdditionalContent"
            hint="Choose your type."
            :items="types"
            item-text="state"
            :label="label1"
            return-object
            single-line
          />

          <iframe
            name="dummyframe"
            id="dummyframe"
            style="display: none"
          ></iframe>
          <form
            action="http://localhost:8085/uploadPostAlbumMedia/"
            enctype="multipart/form-data"
            method="post"
            v-if="!isHiddenAdditionalContent"
            target="dummyframe"
          >
            <input
              type="file"
              accept="image/*,video/*,.mkv"
              name="myPostAlbumFile"
            />
            <input type="submit" value="upload file .." />
          </form>
        </v-form>
        <v-text-field
          label="Tag name"
          v-model="tagName"
          prepend-icon="mdi-address-circle"
          v-if="!isHiddenTag"
        />
      </v-card-text>
      <v-card-actions class="justify-center mb-5">
        <v-btn
          color="info mb-5"
          v-on:click="
            (isHiddenLocation = true),
              (isHiddenDescription = false),
              (isHiddenDescriptionButton = false),
              (isHiddenLocationButton = true)
          "
          v-if="!isHiddenLocationButton"
        >
          Skip location
        </v-btn>
        <v-btn
          color="info mb-5"
          v-on:click="addLocation"
          v-if="!isHiddenLocationButton"
        >
          Add location
        </v-btn>
        <v-btn
          color="info mb-5"
          v-on:click="
            (isHiddenDescriptionButton = true),
              (isHiddenDescription = true),
              (isHiddenContentButton = false),
              (isHiddenContent = false)
          "
          v-if="!isHiddenDescriptionButton"
        >
          Skip description
        </v-btn>
        <v-btn
          color="info mb-5"
          v-on:click="addDescription"
          v-if="!isHiddenDescriptionButton"
        >
          Add description
        </v-btn>
        <v-btn
          color="info mb-5"
          v-if="!isHiddenContentButton"
          v-on:click="addContent"
        >
          Add content
        </v-btn>
        <v-btn
          color="info mb-5"
          v-if="!isHiddenAdditionalContentButton"
          v-on:click="createAdditionalContent"
        >
          Add more content
        </v-btn>
        <v-btn
          color="info mb-5"
          v-if="!isHiddenAdditionalContentButton"
          v-on:click="continueCreation"
        >
          Continue
        </v-btn>
        <v-btn color="info mb-5" v-if="!isHiddenTagButton" v-on:click="addTag">
          Add tag
        </v-btn>
        <v-btn color="info mb-5" v-if="!isHiddenTagButton" v-on:click="finish">
          Finish
        </v-btn>
      </v-card-actions>
    </v-card>
    <div class="spacing" />
  </div>
</template>

<script>
export default {
  name: "CreatePostAlbum",
  data: () => ({
    longitude: "",
    latitude: "",
    country: "",
    city: "",
    streetName: "",
    streetNumber: "",
    tagName: null,
    postAlbumDescription: "",
    locationId: null,
    path: "",
    types: ["PICTURE", "VIDEO"],
    selectedType: "PICTURE",
    label1: "Type",
    postAlbumId: null,
    isHiddenLocationButton: false,
    isHiddenLocation: false,
    isHiddenDescriptionButton: true,
    isHiddenDescription: true,
    isHiddenContentButton: true,
    isHiddenContent: true,
    isHiddenTagButton: true,
    isHiddenTag: true,
    isValidLocation: false,
    isValidPostAlbumDescription: false,
    postAlbumTagId: null,
    isHiddenAdditionalContentButton: true,
    isHiddenAdditionalContent: true,
  }),
  methods: {
    addLocation() {
      if (
        !this.validLongitude() ||
        !this.validLatitude() ||
        !this.validCountry() ||
        !this.validCity() ||
        !this.validStreetName() ||
        !this.validStreetNumber()
      )
        return;

      this.isValidLocation = true;
      this.isHiddenLocationButton = true;
      this.isHiddenLocation = true;
      this.isHiddenDescriptionButton = false;
      this.isHiddenDescription = false;
    },
    addDescription() {
      if (!this.validPostAlbumDescription()) return;

      this.isValidPostAlbumDescription = true;
      this.isHiddenDescriptionButton = true;
      this.isHiddenDescription = true;
      this.isHiddenContentButton = false;
      this.isHiddenContent = false;
    },
    addContent() {
      this.isHiddenContentButton = true;
      this.isHiddenContent = true;
      this.isHiddenAdditionalContentButton = false;
      this.isHiddenAdditionalContent = false;

      if (this.isValidLocation) {
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
            this.createPostAlbumDescription();
          })
          .catch((er) => {
            console.log(er.response.data);
          });
      } else {
        this.createPostAlbumDescription();
      }
    },
    createPostAlbumDescription() {
      if (this.isValidPostAlbumDescription) {
        this.$http
          .post("http://localhost:8084/post_album/", {
            description: this.postAlbumDescription,
            userID: localStorage.getItem("userId"),
            locationId: this.locationId,
          })
          .then((response) => {
            this.postAlbumId = response.data;
            this.createContent();
          })
          .catch((er) => {
            console.log(er.response.data);
          });
      } else {
        this.$http
          .post("http://localhost:8084/post_album/", {
            description: "",
            userID: localStorage.getItem("userId"),
            locationId: this.locationId,
          })
          .then((response) => {
            this.postAlbumId = response.data;
            this.createContent();
          })
          .catch((er) => {
            console.log(er.response.data);
          });
      }
    },
    createContent() {
      this.$http
        .post("http://localhost:8085/post_album_content/", {
          path: this.path,
          type: this.selectedType,
          post_album_id: this.postAlbumId,
        })
        .then((response) => {
          console.log(response.data);
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    continueCreation() {
      this.isHiddenAdditionalContentButton = true;
      this.isHiddenAdditionalContent = true;
      this.isHiddenTagButton = false;
      this.isHiddenTag = false;
    },
    createAdditionalContent() {
      if (!this.validPath()) return;

      this.$http
        .post("http://localhost:8085/post_album_content/", {
          path: this.path,
          type: this.selectedType,
          post_album_id: this.postAlbumId,
        })
        .then((response) => {
          console.log(response.data);
        })
        .catch((er) => {
          console.log(er.response.data);
        });

      alert("Content is added! Add more content or continue creation.");
    },
    finish() {
      alert("Successful creation.");
      window.location.href = "http://localhost:8081/";
    },
    addTag() {
      if (!this.validTag()) return;

      this.$http
        .post("http://localhost:8082/post_album_tag/", {
          name: this.tagName,
        })
        .then((response) => {
          this.postAlbumTagId = response.data;
          this.CreatePostAlbumTagPostAlbums();
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    CreatePostAlbumTagPostAlbums() {
      this.$http
        .post("http://localhost:8082/post_album_tag_post_albums/", {
          postAlbumTagId: this.postAlbumTagId,
          postAlbumId: this.postAlbumId,
        })
        .then((response) => {
          console.log(response.data);
          alert("Tag is created! Add more tags or finish creation.");
        })
        .catch((er) => {
          console.log(er.response.data);
        });
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
    validPostAlbumDescription() {
      if (this.postAlbumDescription.length < 1) {
        alert(
          "Your post album description should contain at least 1 character!"
        );
        return;
      } else if (this.postAlbumDescription.length > 50) {
        alert(
          "Your post album description shouldn't contain more than 50 characters!"
        );
        return false;
      }
      return true;
    },
    validTag() {
      if (this.tagName == null) {
        alert("Your tag name should contain at least 1 character!");
        return;
      } else if (this.tagName.length > 30) {
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
  },
};
</script>

<style scoped>
.spacing {
  height: 100px;
}

.typeCombo {
  width: 94%;
  margin-left: 6%;
}
</style>