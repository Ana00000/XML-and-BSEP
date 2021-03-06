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

    <v-btn
      color="info mb-5"
      v-on:click="
        (isHiddenReportStoryFinal = false), (isHiddenReportStory = true)
      "
      v-if="!isHiddenReportStory"
    >
      Report Story
    </v-btn>

    <v-text-field
      label="Note"
      v-model="note"
      prepend-icon="mdi-address-circle"
      v-if="!isHiddenReportStoryFinal"
    />

    <v-btn
      color="info mb-5"
      v-on:click="reportStory"
      v-if="!isHiddenReportStoryFinal"
    >
      Report
    </v-btn>
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
    isHiddenReportStory: true,
    isHiddenReportStoryFinal: true,
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
          window.location.href = "https://localhost:8081/unauthorizedPage";
          console.log(er);
        });

      this.$http
        .get(
          "https://localhost:8080/api/user/auth/check-find-selected-story-by-id-for-registered-users-permission/",
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

      alert(localStorage.getItem("mySelectedStoryId"));
      alert(localStorage.getItem("mySelectedUserId"));
      this.$http
        .get(
          "https://localhost:8080/api/story/find_selected_story_reg?id=" +
            localStorage.getItem("mySelectedStoryId") +
            "&logId=" +
            localStorage.getItem("mySelectedUserId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.story = response.data;
        })
        .catch(console.log);

      this.$http
        .get(
          "https://localhost:8080/api/story/find_all_story_highlights_for_user?id=" +
            localStorage.getItem("userId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.highlightedStories = response.data;
        })
        .catch(console.log);

      if (localStorage.getItem("userPrivacy") != null) {
        this.isHiddenReportStory = false;
      }
    },
    reportStory() {
      if (!this.validReportNote()) return;

      this.$http
        .post(
          "https://localhost:8080/api/requests/storyICR/",
          {
            note: this.note,
            userId: localStorage.getItem("userId"),
            storyId: localStorage.getItem("mySelectedStoryId"),
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
          window.location.href = "https://localhost:8081/";
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
            localStorage.getItem("mySelectedStoryId"),
          {
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
            .post(
              "https://localhost:8080/api/story/single_story_story_highlights/",
              {
                story_highlight_id: this.highlightedStoryId,
                single_story_id: localStorage.getItem("mySelectedStoryId"),
              },
              {
                headers: {
                  Authorization: "Bearer " + this.token,
                },
              }
            )
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