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
          v-for="item in allDataForLikedPosts"
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
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h2>Posts I disliked</h2>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allDataForDislikedPosts"
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
          this.getData(response.data);
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
      this.$http
        .get(
          "https://localhost:8080/api/post/find_post_by_id?post_id=" +
            item.post_id,
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((r) => {
          this.post = r.data;
          console.log(r.data);
          this.addPost(item);
        })
        .catch(console.log);
    },
    addPost(item) {
      this.$http
        .get(
          "https://localhost:8080/api/post/find_selected_post_for_logged_user?id=" +
            item.post_id +
            "&logId=" +
            this.post.user_id,
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log(resp);
          this.post = resp.data;
        })
        .catch(console.log);

      console.log(item.liked_status);
      
      if (item.liked_status == 0) {
        this.text = "I like this post: " + item.post_id;
        this.allDataForLikedPosts.push({
          id: item.id,
          post_id: item.post_id,
          post_user_id: this.post.user_id,
          type: this.post.type,
          path: this.post.path,
          text: this.text,
        });
      } else {
        this.text = "I dislike this post: " + item.post_id;
        this.allDataForDislikedPosts.push({
          id: item.id,
          post_id: item.post_id,
          post_user_id: this.post.user_id,
          type: this.post.type,
          path: this.post.path,
          text: this.text,
        });
      }
    },
    getPost(item) {
      localStorage.setItem("selectedUserId", item.post_user_id);
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
