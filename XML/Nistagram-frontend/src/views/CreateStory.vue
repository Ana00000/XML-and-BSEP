<template>
  <div>
    <div class="spacing" />
    <v-card width="600" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1">Story creation</h1>
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
            v-model="storyDescription"
            prepend-icon="mdi-address-circle"
            v-if="!isHiddenDescription"
          />
          <v-select
            class="typeCombo"
            v-model="selectedStoryType"
            v-if="!isHiddenDescription"
            hint="Choose your publicity story type."
            :items="storyTypes"
            item-text="state"
            :label="label2"
            return-object
            single-line
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
            action="http://localhost:8080/api/content/uploadStoryMedia/"
            enctype="multipart/form-data"
            method="post"
            v-if="!isHiddenContent"
            target="dummyframe"
            class="uploadButton"
          >
            <template v-if="selectedType === 'PICTURE'">
              <input id="pic" type="file" accept="image/*" name="myStoryFile" />
            </template>
            <template v-else>
              <input
                id="vid"
                type="file"
                accept="video/*, .mp4"
                name="myStoryFile"
              />
            </template>
            <input
              type="submit"
              value=" <- Upload file"
              v-on:click="ValidteType"
            />
          </form>
        </v-form>
        <v-text-field
          label="Tag name"
          v-model="tagName"
          prepend-icon="mdi-address-circle"
          v-if="!isHiddenTag && selectedTagType === 'HASH_TAG'"
        />
        <v-select
          class="typeCombo"
          v-model="selectedTagType"
          v-on:click="checkTag"
          v-if="!isHiddenTag"
          hint="Choose your tag type."
          :items="tagTypes"
          item-text="state"
          :label="label3"
          return-object
          single-line
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
              (isVisibleContentButton = false),
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
          v-if="isVisibleContentButton"
          v-on:click="addContent"
        >
          Add content
        </v-btn>
        <v-btn color="info mb-5" v-if="!isHiddenTagButton" v-on:click="addTag">
          Add tag
        </v-btn>
        <v-btn color="info mb-5" v-if="!isHiddenTagButton" v-on:click="finish">
          Finish
        </v-btn>
      </v-card-actions>
    </v-card>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allUserTags"
          :key="item.id"
          class="space-bottom"
        >
          <v-card
            class="mx-auto"
            v-on:click="getTag(item)"
            v-if="isVisibleTags"
          >
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{ item.name }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <div class="spacing" />
  </div>
</template>

<script>
export default {
  name: "CreateStory",
  data: () => ({
    longitude: "",
    latitude: "",
    country: "",
    city: "",
    streetName: "",
    streetNumber: "",
    tagName: null,
    storyDescription: "",
    locationId: null,
    storyTypes: ["CLOSE_FRIENDS", "ALL_FRIENDS", "PUBLIC"],
    selectedStoryType: "CLOSE_FRIENDS",
    label2: "Story publicity type",
    path: "",
    types: ["PICTURE", "VIDEO"],
    selectedType: "PICTURE",
    label1: "Type",
    label3: "Tag type",
    tagTypes: ["USER_TAG", "HASH_TAG"],
    selectedTagType: "HASH_TAG",
    storyId: null,
    isHiddenLocationButton: false,
    isHiddenLocation: false,
    isHiddenDescriptionButton: true,
    isHiddenDescription: true,
    isVisibleContentButton: false,
    isHiddenContent: true,
    isHiddenTagButton: true,
    isHiddenTag: true,
    isValidLocation: false,
    isValidStoryDescription: false,
    storyTagId: null,
    extension: "",
    isVisibleTags: false,
    allUserTags: [],
    userTag: null,
    userId: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
        this.userId = localStorage.getItem("userId");
        this.$http
        .get("http://localhost:8080/api/tag/find_all_taggable_users_story/")
        .then((response) => {
          for (var i = 0; i < response.data.length; i++) {
            if (response.data[i].tag_type == 0 && response.data[i].user_id != this.userId) {
              this.allUserTags.push(response.data[i]);
            }
          }
        })
        .catch(console.log);
    },
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
    GetExtension(pathFile) {
      console.log(pathFile);
      let out = pathFile.split("\\");
      let fileName = out[out.length - 1];
      let dotSplit = fileName.split(".");
      this.extension = dotSplit[dotSplit.length - 1];
      console.log(this.extension);
    },
    ValidteType() {
      let pathFile = "";
      if (this.selectedType === "PICTURE") {
        pathFile = document.getElementById("pic").value;
        this.GetExtension(pathFile);
        console.log(this.extension);
        if (
          this.extension === "PNG" ||
          this.extension === "png" ||
          this.extension === "JPG" ||
          this.extension === "jpg" ||
          this.extension === "jpeg" ||
          this.extension === "JPEG"
        ) {
          this.isVisibleContentButton = true;
        } else {
          this.isVisibleContentButton = false;
          alert(
            "Please, choose a picture in a correct format e.g. png, jpg or jpeg."
          );
        }
      } else {
        pathFile = document.getElementById("vid").value;
        this.GetExtension(pathFile);
        console.log(this.extension);
        if (this.extension === "mp4" || this.extension === "MP4") {
          this.isVisibleContentButton = true;
        } else {
          this.isVisibleContentButton = false;
          alert("Please, choose a video in a correct format mp4.");
        }
      }
    },
    addDescription() {
      if (!this.validStoryDescription()) return;

      this.isValidStoryDescription = true;
      this.isHiddenDescriptionButton = true;
      this.isHiddenDescription = true;
      this.isVisibleContentButton = false;
      this.isHiddenContent = false;
    },
    addContent() {
      this.isVisibleContentButton = false;
      this.isHiddenContent = true;
      this.isHiddenTagButton = false;
      this.isHiddenTag = false;

      if (this.isValidLocation) {
        this.$http
          .post("http://localhost:8080/api/location/", {
            longitude: this.longitude,
            latitude: this.latitude,
            country: this.country,
            city: this.city,
            streetName: this.streetName,
            streetNumber: this.streetNumber,
          })
          .then((response) => {
            this.locationId = response.data;
            this.createStoryDescription();
          })
          .catch((er) => {
            console.log(er.response.data);
          });
      } else {
        this.createStoryDescription();
      }
    },
    createStoryDescription() {
      if (this.isValidStoryDescription) {
        this.$http
          .post("http://localhost:8080/api/story/single_story/", {
            description: this.storyDescription,
            userID: localStorage.getItem("userId"),
            locationId: this.locationId,
            storyType: this.selectedStoryType,
          })
          .then((response) => {
            this.storyId = response.data;
            this.createContent();
          })
          .catch((er) => {
            console.log(er.response.data);
          });
      } else {
        this.$http
          .post("http://localhost:8080/api/story/single_story/", {
            description: "",
            userID: localStorage.getItem("userId"),
            locationId: this.locationId,
            storyType: this.selectedStoryType,
          })
          .then((response) => {
            this.storyId = response.data;
            this.createContent();
          })
          .catch((er) => {
            console.log(er.response.data);
          });
      }
    },
    createContent() {
      this.$http
        .post("http://localhost:8080/api/content/single_story_content/", {
          path: this.path,
          type: this.selectedType,
          single_story_id: this.storyId,
        })
        .then((response) => {
          console.log(response.data);
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    finish() {
      alert("Successful creation.");
      window.location.href = "http://localhost:8081/";
    },
    addTag() {
      
      if (this.selectedTagType == "USER_TAG") {
        if (this.userTag==null){
          alert("Tag is not selected");
          return;
        }
        this.createStoryUserTagStories();
      } else {
         if (!this.validTag()) return;
         this.$http
          .post("http://localhost:8080/api/tag/tag/", {
            name: this.tagName,
            tag_type: this.selectedTagType,
          })
          .then((response) => {
            this.storyTagId = response.data;
            this.createStoryTagStories();
          })
          .catch((er) => {
            console.log(er.response.data);
          });
      }
    },
    createStoryTagStories() {
      this.$http
        .post("http://localhost:8080/api/tag/story_tag_stories/", {
          tag_id: this.storyTagId,
          story_id: this.storyId,
        })
        .then((response) => {
          console.log(response.data);
          alert("Tag is created! Add more tags or finish creation.");
          this.tagName = null;
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    createStoryUserTagStories() {
      this.$http
        .post("http://localhost:8080/api/tag/story_tag_stories/", {
          tag_id: this.userTag.id,
          story_id: this.storyId,
        })
        .then((response) => {
          console.log(response.data);
          alert("Tag is created! Add more tags or finish creation.");
          this.userTag = null;
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    checkTag() {
      if (this.selectedTagType == "USER_TAG") {
        this.isVisibleTags = true;
      }
    },
    getTag(item) {
      this.userTag = item;
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
    validStoryDescription() {
      if (this.storyDescription.length < 1) {
        alert("Your story description should contain at least 1 character!");
        return;
      } else if (this.storyDescription.length > 50) {
        alert(
          "Your story description shouldn't contain more than 50 characters!"
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

.uploadButton {
  margin-left: 6%;
}

.typeCombo {
  width: 94%;
  margin-left: 6%;
}
</style>