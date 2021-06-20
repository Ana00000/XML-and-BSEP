<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Friends Stories</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg4 v-for="item in stories" :key="item.id" class="space-bottom">
          <v-card class="mx-auto">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  item.description
                }}</v-list-item-subtitle>
                <v-list-item-title>{{ item.tags }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="item.type == 'VIDEO'">
              <v-list-item-content>
                <video width="320" height="240" controls>
                  <source
                    :src="require(`/app/public/uploads/${item.path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="item.type != 'VIDEO'">
              <v-list-item-content>
                <img
                  :src="require(`/app/public/uploads/${item.path}`)"
                  alt
                  class="icon"
                  width="320"
                  height="240"
                />
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle
                  v-text="
                    item.country +
                    ' ' +
                    item.city +
                    ' ' +
                    item.street_name +
                    ' ' +
                    item.street_number
                  "
                />
                <v-list-item-subtitle>{{
                  item.creation_date
                }}</v-list-item-subtitle>

                <v-list-item-content>
                  <v-text-field
                    label="Note"
                    v-model="note"
                    prepend-icon="mdi-address-circle"
                    class="note"
                  />
                  <v-btn color="info mb-5" v-on:click="reportStory(item.id)">
                    Report Story
                  </v-btn>
                </v-list-item-content>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>

    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Friends Album Stories</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in albumStories"
          :key="item.id"
          class="space-bottom"
        >
          <v-card class="mx-auto" v-on:click="getMyAlbumStories(item)">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  item.description
                }}</v-list-item-subtitle>
                <v-list-item-title>{{ item.tags }}</v-list-item-title>
                <v-list-item-subtitle
                  v-text="
                    item.country +
                    ' ' +
                    item.city +
                    ' ' +
                    item.street_name +
                    ' ' +
                    item.street_number
                  "
                />
                <v-list-item-subtitle>{{
                  item.creation_date
                }}</v-list-item-subtitle>
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
  name: "FriendsStories",
  data: () => ({
    token: null,
    stories: [],
    publicPath: process.env.VUE_APP_BASE_URL,
    albumStories: [],
    note: "",
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");

      this.$http
        .get("https://localhost:8080/api/user/check_if_authentificated/", {
          headers: {
            Authorization: "Bearer " + this.token,
          },
        })
        .then((resp) => {
          console.log("User is authentificated!");
          console.log(resp.data);
        })
        .catch((er) => {
          console.log(er);
          window.location.href = "https://localhost:8081/unauthorizedPage";
        });

      this.$http
        .get(
          "https://localhost:8080/api/user/auth/check-find-all-following-stories-permission/",
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log("User is authorized!");
          console.log(resp.data);
        })
        .catch((er) => {
          window.location.href = "https://localhost:8081/forbiddenPage";
          console.log(er);
        });

      this.$http
        .get(
          "https://localhost:8080/api/story/find_all_following_stories?id=" +
            localStorage.getItem("userId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.stories = response.data;
          console.log(response.data);
        })
        .catch(console.log);

      this.getStoryAlbums();
    },
    reportStory(selectedStoryId) {
      if (!this.validReportNote()) return;

      this.$http
        .post(
          "https://localhost:8080/api/requests/storyICR/",
          {
            note: this.note,
            userId: localStorage.getItem("userId"),
            storyId: selectedStoryId,
          },
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          alert("Story was reported.");
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    validReportNote() {
      if (this.note.length < 2) {
        alert("Your story report note should contain at least 2 characters!");
        return false;
      } else if (this.note.length > 30) {
        alert(
          "Your story report note shouldn't contain more than 30 characters!"
        );
        return false;
      } else if (this.note.match(/[&<>/\\"]/g)) {
        alert(
          "Your story report note shouldn't contain those special characters."
        );
        return false;
      }
      return true;
    },
    getStoryAlbums() {
      this.$http
        .get(
          "https://localhost:8080/api/story/find_all_following_story_albums?id=" +
            localStorage.getItem("userId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.albumStories = response.data;
        })
        .catch(console.log);
    },
    getMyAlbumStories(item) {
      localStorage.setItem("mySelectedUserId", item.user_id);
      localStorage.setItem("mySelectedStoryAlbumId", item.story_album_id);

      window.location.href =
        "https://localhost:8081/storyAlbumByIdWithoutActivity";
    },
  },
};
</script>

<style scoped>
.spacingOne {
  height: 50px;
}

.title {
  margin-left: 44%;
}

.spacingTwo {
  height: 100px;
}

.note {
  width: 120px;
  margin-left: 7%;
}
</style>