<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h1>Selected Story</h1>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg10 class="space-bottom">
          <v-card class="card">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  story.description
                }}</v-list-item-subtitle>
                <v-list-item-title>{{ story.tags }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="story.type == 'VIDEO'">
              <v-list-item-content>
                <video width="320" height="440" controls>
                  <source
                    :src="require(`/app/public/uploads/${story.path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="story.type != 'VIDEO'">
              <v-list-item-content>
                <img
                  :src="require(`/app/public/uploads/${story.path}`)"
                  alt
                  class="icon"
                  width="320"
                  height="440"
                />
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle
                  v-text="
                    story.country +
                    ' ' +
                    story.city +
                    ' ' +
                    story.street_name +
                    ' ' +
                    story.street_number
                  "
                />
                <v-list-item-subtitle>{{
                  story.creation_date
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>

        <v-btn
          color="info mb-5"
          v-on:click="addToStoryHighlight"
          class="addButton"
        >
          Add to story highlight
        </v-btn>
      </v-layout>

      <v-container grid-list-lg>
        <div class="spacingOne" />
        <v-card-title class="justify-center">
          <h1 class="display-1">My Highlighted Stories</h1>
        </v-card-title>
        <div class="spacingTwo" />
        <v-layout row>
          <v-flex
            lg4
            v-for="item in highlightedStories"
            :key="item.id"
            class="space-bottom"
          >
            <v-card class="mx-auto" v-on:click="getStoryHighlight(item)">
              <v-list-item three-line>
                <v-list-item-content>
                  <v-list-item-subtitle>{{ item.title }}</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "StoryByIdWithoutActivity",
  data: () => ({
    story: null,
    token: null,
    highlightedStoryId: null,
    highlightedStories: [],
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");
      alert(localStorage.getItem("mySelectedStoryId"));
      alert(localStorage.getItem("mySelectedUserId"));
      this.$http
        .get(
          "https://localhost:8080/api/story/find_selected_story_reg?id=" +
            localStorage.getItem("mySelectedStoryId") +
            "&logId=" +
            localStorage.getItem("mySelectedUserId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        })
        .then((response) => {
          this.story = response.data;
        })
        .catch(console.log);

      this.$http
        .get(
          "https://localhost:8080/api/story/find_all_story_highlights_for_user?id=" +
            localStorage.getItem("userId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
            }
        )
        .then((response) => {
          this.highlightedStories = response.data;
        })
        .catch(console.log);
    },
    getStoryHighlight(item) {
      this.highlightedStoryId = item.id;
    },
    addToStoryHighlight() {
      if (this.highlightedStoryId == null) {
        alert("You have not selected story highlight.");
        return;
      }

      this.$http
        .get(
          "https://localhost:8080/api/story/find_all_single_story_story_highlights_for_story?id=" +
            localStorage.getItem("mySelectedStoryId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
            }
        )
        .then((response) => {
          for (var i = 0; i < response.data.length; i++) {
            if (
              this.highlightedStoryId == response.data[i].story_highlight_id
            ) {
              alert(
                "You have already added this story to selected story highlight."
              );
              return;
            }
          }
          this.$http
            .post("https://localhost:8080/api/story/single_story_story_highlights/", {
              story_highlight_id: this.highlightedStoryId,
              single_story_id: localStorage.getItem("mySelectedStoryId"),
            },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
            })
            .then((response) => {
              console.log(response.data);
              alert("You have added this story to selected story highlight.");
            })
            .catch((er) => {
              console.log(er.response.data);
            });
        })
        .catch(console.log);
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
  height: 50px;
}

.card {
  margin-left: 20%;
}

.addButton {
  margin-top: 2%;
  margin-left: 43%;
}
</style>