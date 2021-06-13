<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Stories from highlight</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg4 v-for="item in stories" :key="item.id" class="space-bottom">
          <v-card class="mx-auto" v-on:click="getMyStory(item)">
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
  name: "StoriesOfStoryHighlight",
  data: () => ({
    stories: [],
    logId: null,
    token: null,
    allStoriesIds: [],
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.logId = localStorage.getItem("userId");
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
          console.log(resp.data);
        })
        .catch((er) => {
          window.location.href = "https://localhost:8081/unauthorizedPage";
          console.log(er);
        });

      this.$http
        .get(
          "https://localhost:8080/api/user/auth/check-find-all-stories-for-logged-user-permission/",{
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
          "https://localhost:8080/api/story/find_all_single_story_story_highlights_for_story_highlight?id=" +
            localStorage.getItem("selectedStoryHighlightId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          for (var i = 0; i < response.data.length; i++) {
            this.allStoriesIds.push(response.data[i].single_story_id);
          }
          this.getStories();
        })
        .catch(console.log);
    },
    getStories() {
      this.$http
        .get(
          "https://localhost:8080/api/story/find_all_stories_for_logged_user?id=" +
            localStorage.getItem("userId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          for (var i = 0; i < response.data.length; i++) {
            if (this.allStoriesIds.includes(response.data[i].story_id)) {
              this.stories.push(response.data[i]);
            }
          }
        })
        .catch(console.log);
    },
    getMyStory(item){
      localStorage.setItem("mySelectedUserId", item.user_id);
      localStorage.setItem("mySelectedStoryId", item.story_id);

      window.location.href = "https://localhost:8081/storyByIdWithoutActivity";
    }
  },
};
</script>

<style scoped>
.combo {
  width: 25%;
  margin-left: 42%;
}

.center {
  margin-left: 50%;
  padding: 10px;
}

.spacingOne {
  height: 50px;
}

.title {
  margin-left: 44%;
}

.spacingTwo {
  height: 100px;
}
</style>