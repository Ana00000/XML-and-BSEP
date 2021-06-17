<template>
  <div>
    <h1>Posts I liked</h1>
    <v-container grid-list-lg>
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allDataForLikedPosts"
          :key="item.id"
          class="space-bottom"
        >
          <div class="spacingOne" />
          <v-card color="info" dark max-width="500">
            <v-card-title>
              <span class="title text-xs-center font-weight-light">
                Liked post
              </span>
            </v-card-title>

            <v-card-text class="headline font-weight-bold">
              {{ item.text }}
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <h2>Posts I disliked</h2>
    <v-container grid-list-lg>
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allDataForDislikedPosts"
          :key="item.id"
          class="space-bottom"
        >
          <div class="spacingOne" />
          <v-card color="info" dark max-width="500">
            <v-card-title>
              <span class="title text-xs-center font-weight-light">
                Disliked post
              </span>
            </v-card-title>

            <v-card-text class="headline font-weight-bold">
              {{ item.text }}
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
</div>
</template>

<script>
export default {
  name: "ReactedPosts",
  data: () => ({
    publicPath: process.env.VUE_APP_BASE_URL,
    allLikedPosts: [],
    allDislikedPosts: [],
    token: null,
    userId: null,
    allDataForLikedPosts: [],
    allDataForDislikedPosts: [],
    text: "",
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");
      this.userId = localStorage.getItem("userId");

      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_liked_posts_by_user_id?user_id=" +
            this.userId,
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          this.allLikedPosts = response.data;
          this.getData(response.data);
        })
        .catch(console.log);

      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_disliked_posts_by_user_id?user_id=" +
            this.userId,
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          this.allDislikedPosts = response.data;
          this.getData(response.data);
        })
        .catch(console.log);
    },
    getData(items) {
      for (var i = 0; i < items.length; i++) {
        this.getItem(items[i]);
      }
    },
    getItem(item) {
      console.log(item);
      if (item.liked_status == 0) {
        this.text = "I like this post: " + item.post_id;
        this.allDataForLikedPosts.push({
          id: item.id,
          post_id: item.post_id,
          text: this.text,
        });
      } else {
        this.text = "I dislike this post: " + item.post_id;
        this.allDataForDislikedPosts.push({
          id: item.id,
          post_id: item.post_id,
          text: this.text,
        });
      }
    },
  },
};
</script>
