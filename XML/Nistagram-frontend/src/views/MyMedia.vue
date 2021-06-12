<template>
  <div>
    <div class="spacingOne" />
    <div class="title">
      <h1>My Media</h1>
    </div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Stories</h2>
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

    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Posts</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg4 v-for="item in posts" :key="item.id" class="space-bottom">
          <v-card class="mx-auto" v-on:click="getMyPost(item)">
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
        <h2>Album Posts</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in albumPosts"
          :key="item.id"
          class="space-bottom"
        >
          <v-card class="mx-auto" v-on:click="getMyAlbumPosts(item)">
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

    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Album Stories</h2>
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
  name: "MyMedia",
  data: () => ({
    posts: [],
    stories: [],
    logId: null,
    albumPosts: [],
    token: null,
    albumStories: [],
  }),
  mounted() {
    this.logId = localStorage.getItem("userId");
    this.token = localStorage.getItem("token");
    this.init();
  },
  methods: {
    init() {
      this.getPosts();
      this.getStories();
      this.getPostAlbums();
      this.getStoryAlbums();
    },
    getPosts() {
      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_posts_for_logged_user?id=" +
            localStorage.getItem("userId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        }
        )
        .then((response) => {
          this.posts = response.data;
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
          this.stories = response.data;
        })
        .catch(console.log);
    },
    getPostAlbums() {
      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_album_posts_for_logged_user?id=" +
            localStorage.getItem("userId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        }
        )
        .then((response) => {
          this.albumPosts = response.data;
        })
        .catch(console.log);
    },
    getStoryAlbums() {
      this.$http
        .get(
          "https://localhost:8080/api/story/find_all_album_stories_for_logged_user?id=" +
            localStorage.getItem("userId"),{
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
    getMyPost(item) {
      localStorage.setItem("mySelectedUserId", item.user_id);
      localStorage.setItem("mySelectedPostId", item.post_id);

      window.location.href = "https://localhost:8081/postByIdWithoutActivity";
    },
    getMyStory(item) {
      localStorage.setItem("mySelectedUserId", item.user_id);
      localStorage.setItem("mySelectedStoryId", item.story_id);

      window.location.href = "https://localhost:8081/storyByIdWithoutActivity";
    },
    getMyAlbumPosts(item) {
      localStorage.setItem("mySelectedUserId", item.user_id);
      localStorage.setItem("mySelectedPostAlbumId", item.post_album_id);

      window.location.href =
        "https://localhost:8081/postAlbumByIdWithoutActivity";
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