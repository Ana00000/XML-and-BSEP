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
  </div>
</template>

<script>
export default {
  name: "CreateCollection",
  data: () => ({
    title: "",
    postCollectionId: null,
  }),
  methods: {
    createCollection() {
      if (!this.validTitle()) return;

      this.$http
        .get(
          "http://localhost:8084/find_all_post_collections_for_reg?id=" +
            localStorage.getItem("userId")
        )
        .then((response) => {
          for (var i = 0; i < response.data.length; i++) {
            if (this.title == response.data[i].title) {
              alert("You have already created collection with this name!");
              return;
            }
          }
          this.$http
            .post("http://localhost:8084/post_collection/", {
              title: this.title,
              userID: localStorage.getItem("userId"),
            })
            .then((response) => {
              this.postCollectionId = response.data;
              alert("Successful creation of collection.");
              this.title = "";
            })
            .catch((er) => {
              console.log(er.response.data);
            });
        })
        .catch(console.log);
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