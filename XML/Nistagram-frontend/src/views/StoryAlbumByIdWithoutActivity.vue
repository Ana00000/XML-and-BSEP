<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h1>Selected Story Album</h1>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg4 class="space-bottom">
          <v-card class="mx-auto">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  storyAlbum.description
                }}</v-list-item-subtitle>
                <v-list-item-title>{{ storyAlbum.tags }}</v-list-item-title>
                <v-list-item-subtitle
                  v-text="
                    storyAlbum.country +
                    ' ' +
                    storyAlbum.city +
                    ' ' +
                    storyAlbum.street_name +
                    ' ' +
                    storyAlbum.street_number
                  "
                />
                <v-list-item-subtitle>{{
                  storyAlbum.creation_date
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>

    <v-container grid-list-lg>
      <v-layout row>
        <v-flex
          lg4
          v-for="item in storyAlbumContents"
          :key="item.id"
          class="space-bottom"
        >
          <v-card class="mx-auto">
            <v-list-item three-line v-if="item.type == 'VIDEO'">
              <v-list-item-content>
                <video width="320" height="240" controls>
                  <source
                    :src="require(`../../../Media/${item.path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="item.type != 'VIDEO'">
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
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "StoryAlbumByIdWithoutActivity",
  data: () => ({
    storyAlbum: null,
    storyAlbumContents: [],
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get(
          "http://localhost:8084/find_selected_story_album_for_logged_user?id=" +
            localStorage.getItem("mySelectedStoryAlbumId") +
            "&logId=" +
            localStorage.getItem("mySelectedUserId")
        )
        .then((response) => {
          this.storyAlbum = response.data;
          for (var i = 0; i < response.data.types.length; i++) {
            this.storyAlbumContents.push({
              type: response.data.types[i],
              path: response.data.paths[i],
            });
          }
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
</style>