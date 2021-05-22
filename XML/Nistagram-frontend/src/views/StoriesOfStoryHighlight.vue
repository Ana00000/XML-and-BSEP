<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h1>Stories from story highlight</h1>
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
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="item.type == 1">
              <v-list-item-content>
                <video width="320" height="240" controls>
                  <source
                    :src="require(`../../../Media/${item.path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="item.type != 1">
              <v-list-item-content>
                <img
                  :src="require(`../../../Media/${item.path}`)"
                  alt
                  class="icon"
                  width="320"
                  height="240"
                />
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  item.creationDate
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
    allStoriesIds: [],
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get(
          "http://localhost:8086/find_all_single_story_story_highlights_for_story_highlight?id=" +
            localStorage.getItem("selectedStoryHighlightId")
        )
        .then((response) => {
          for (var i = 0; i < response.data.length; i++) {
            this.allStoriesIds.push(response.data[i].single_story_id);
          }

          for (var j = 0; j < this.allStoriesIds.length; j++) {
            this.$http
              .get(
                "http://localhost:8086/find_single_story_for_id?id=" +
                  this.allStoriesIds[j]
              )
              .then((response) => {
                this.setStory(response.data);
              })
              .catch(console.log);
          }
        })
        .catch(console.log);
    },
    setStory(story) {
      for (var j = 0; j < this.allStoriesIds.length; j++) {
        this.$http
          .get(
            "http://localhost:8085/find_single_story_content_for_story_id?id=" +
              this.allStoriesIds[j]
          )
          .then((response) => {
            this.stories.push({
              description: story.description,
              creationDate: story.creationDate,
              path: response.data.path,
              type: response.data.type,
            });
          })
          .catch(console.log);
      }
    },
  },
};
</script>

<style scoped>
.spacingOne {
  height: 50px;
}

.title {
  margin-left: 37%;
}

.spacingTwo {
  height: 100px;
}
</style>