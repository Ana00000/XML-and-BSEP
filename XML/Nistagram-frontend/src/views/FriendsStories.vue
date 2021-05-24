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
    stories: [],
    publicPath: process.env.VUE_APP_BASE_URL,
    albumStories: [],
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get("http://localhost:8080/api/story/find_all_following_stories?id=" + localStorage.getItem("userId"))
        .then((response) => {
          this.stories = response.data;
          console.log(response.data);
        })
        .catch(console.log);

      this.getStoryAlbums();
    },
    getStoryAlbums() {
      this.$http
        .get(
          "http://localhost:8086/find_all_following_story_albums?id=" +
            localStorage.getItem("userId")
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
        "http://localhost:8081/storyAlbumByIdWithoutActivity";
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
</style>