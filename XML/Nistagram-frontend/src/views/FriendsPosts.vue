<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Friends Posts</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg4 v-for="item in posts" :key="item.id" class="space-bottom">
          <v-card class="mx-auto" v-on:click="getPost(item)">
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
        <h2>Friends Album Posts</h2>
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
  </div>
</template>

<script>
export default {
  name: "FriendsPosts",
  data: () => ({
    posts: [],
    albumPosts: [],
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get(
          "http://localhost:8084/find_all_following_posts?id=" +
            localStorage.getItem("userId")
        )
        .then((response) => {
          this.posts = response.data;
          console.log(response.data);
        })
        .catch(console.log);

      this.getPostAlbums();
    },
    getPost(item) {
      localStorage.setItem("selectedUserId", item.user_id);
      localStorage.setItem("selectedPostId", item.post_id);

      window.location.href = "http://localhost:8081/postById";
    },
    getPostAlbums() {
      this.$http
        .get(
          "http://localhost:8084/find_all_following_post_albums?id=" +
            localStorage.getItem("userId")
        )
        .then((response) => {
          this.albumPosts = response.data;
        })
        .catch(console.log);
    },
    getMyAlbumPosts(item) {
      localStorage.setItem("mySelectedUserId", item.user_id);
      localStorage.setItem("mySelectedPostAlbumId", item.post_album_id);

      window.location.href =
        "http://localhost:8081/postAlbumByIdWithoutActivity";
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