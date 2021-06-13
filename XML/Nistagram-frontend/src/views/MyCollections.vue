<template>
  <div>
    <div class="spacing" />
    <v-card width="600" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1">Collection creation</h1>
      </v-card-title>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="Title"
            v-model="title"
            prepend-icon="mdi-address-circle"
          />
        </v-form>
      </v-card-text>
      <v-card-actions class="justify-center mb-5">
        <v-btn color="info mb-5" v-on:click="createCollection"> Create </v-btn>
      </v-card-actions>
    </v-card>
    <div class="spacing" />

    <v-container grid-list-lg>
      <div class="spacingOne" />
      <v-card-title class="justify-center">
        <h1 class="display-1">My Collections</h1>
      </v-card-title>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in postCollections"
          :key="item.id"
          class="space-bottom"
        >
          <v-card class="mx-auto">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{ item.title }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "MyCollections",
  data: () => ({
    title: "",
    postCollectionId: null,
    token: null,
    postCollections: [],
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");

      this.$http
        .get(
          "https://localhost:8080/api/user/check_if_authentificated/",{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log("User is authentificated!");
        })
        .catch((er) => {
          window.location.href = "https://localhost:8081/unauthorizedPage";
        });

      this.$http
        .get(
          "https://localhost:8080/api/user/auth/check-find-all-post-collections-for-user-registered-user-permission/",{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log("User is authorized!");
        })
        .catch((er) => {
          window.location.href = "https://localhost:8081/forbiddenPage";
        });

      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_post_collections_for_reg?id=" +
            localStorage.getItem("userId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.postCollections = response.data;
        })
        .catch(console.log);
    },
    createCollection() {
      if (!this.validTitle()) return;

      for (var i = 0; i < this.postCollections.length; i++) {
        if (this.title == this.postCollections[i].title) {
          alert("You have already created collection with this name!");
          return;
        }
      }

      this.$http
        .post("https://localhost:8080/api/post/post_collection/", {
          title: this.title,
          userID: localStorage.getItem("userId"),
        },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
        .then((response) => {
          this.postCollectionId = response.data;
          alert("Successful creation of collection.");
          window.location.href = "https://localhost:8081/myCollections";
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    validTitle() {
      if (this.title.length < 2) {
        alert("Your title should contain at least 2 characters!");
        return false;
      } else if (this.title.length > 30) {
        alert("Your title shouldn't contain more than 30 characters!");
        return false;
      } else if (this.title.match(/[!@#$%^&*:'<>+-/\\"]/g)) {
        alert("Your title shouldn't contain those special characters.");
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