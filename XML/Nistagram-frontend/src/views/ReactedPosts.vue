<template>
  <div>
    <div class="spacingOne" />
    <div class="title">
      <h1>Reacted Posts</h1>
    </div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Posts I liked</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allLikedPosts"
          :key="item.id"
          class="space-bottom"
        >
          <v-card
            class="mx-auto"
            color="info"
            dark
            max-width="500"
            v-on:click="getPost(item)"
          >
            <v-card-title>
              <span class="title text-xs-center font-weight-light">
                Liked post
              </span>
            </v-card-title>

            <v-card-text class="headline font-weight-bold">
              {{ item.text }}
            </v-card-text>

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
                  width="320"
                  height="240"
                />
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Posts I disliked</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allDislikedPosts"
          :key="item.id"
          class="space-bottom"
        >
          <div class="spacingOne" />
          <v-card
            class="mx-auto"
            color="info"
            dark
            max-width="500"
            v-on:click="getPost(item)"
          >
            <v-card-title>
              <span class="title text-xs-center font-weight-light">
                Disliked post
              </span>
            </v-card-title>

            <v-card-text class="headline font-weight-bold">
              {{ item.text }}
            </v-card-text>

            <v-list-item three-line v-if="item.type == 'VIDEO'">
              <v-list-item-content>
                <video width="320" height="240" controls>
                  <source
                    :src="require(`/app/public/uploads/${item.path}`)"
                    alt="post_video"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="item.type != 'VIDEO'">
              <v-list-item-content>
                <img
                  :src="require(`/app/public/uploads/${item.path}`)"
                  alt="post_image"
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
  name: "ReactedPosts",
  data: () => ({
    allLikedPosts: [],
    allDislikedPosts: [],
    token: null,
    userId: null,
    allDataForLikedPosts: [],
    allDataForDislikedPosts: [],
    text: "",
    post: null,
  }),
  mounted() {
    this.token = localStorage.getItem("token");
    this.userId = localStorage.getItem("userId");
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get(
          "https://localhost:8080/api/user/auth/check-find-all-liked-post-by-user-id-permission/",
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
          "https://localhost:8080/api/user/auth/check-find-all-disliked-post-by-user-id-permission/",
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

      this.getLikedPostsByUser();
      this.getDislikedPostsByUser();
    },
    getLikedPostsByUser() {
      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_liked_posts_by_user_id?id=" + localStorage.getItem("userId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          this.allLikedPosts = response.data;
        })
        .catch(console.log);
    },
    getDislikedPostsByUser() {
      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_disliked_posts_by_user_id?id=" + localStorage.getItem("userId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          this.allDislikedPosts = response.data;
        })
        .catch(console.log);
    },
    getPost(item) {
      localStorage.setItem("selectedPostId", item.post_id);

      window.location.href = "https://localhost:8081/postById";
    },
  },
};
</script>

<style scoped>
.spacingOne {
  height: 2%;
}
.spacingTwo {
  height: 5%;
}
.title {
  margin-left: 44%;
}
</style>
