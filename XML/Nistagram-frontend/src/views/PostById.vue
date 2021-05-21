<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h1>Selected Post</h1>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg10 class="space-bottom">
          <v-card class="card">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  post.description
                }}</v-list-item-subtitle>
                <v-list-item-title>{{ post.tags }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="post.type == 'VIDEO'">
              <v-list-item-content>
                <video width="320" height="440" controls>
                  <source
                    :src="require(`../../../Media/${post.path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="post.type != 'VIDEO'">
              <v-list-item-content>
                <img
                  :src="require(`../../../Media/${post.path}`)"
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
                    post.country +
                    ' ' +
                    post.city +
                    ' ' +
                    post.street_name +
                    ' ' +
                    post.street_number
                  "
                />
                <v-list-item-subtitle>{{
                  post.creation_date
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>

    <v-btn color="info mb-5" v-on:click="likePost" class="likeButton">
      Like
    </v-btn>

    <v-btn
      color="info mb-5"
      v-on:click="removeLike"
      class="removeLikeButton"
      v-if="!isHiddenRemoveLike"
    >
      Remove like
    </v-btn>

    <v-btn color="info mb-5" v-on:click="dislikePost" class="dislikeButton">
      Dislike
    </v-btn>

    <v-btn
      color="info mb-5"
      v-on:click="removeDislike"
      class="removeDislikeButton"
      v-if="!isHiddenRemoveDislike"
    >
      Remove dislike
    </v-btn>

    <v-btn color="info mb-5" v-on:click="favoritePost" class="favoriteButton">
      Favorite
    </v-btn>

    <v-btn
      color="info mb-5"
      v-on:click="removeFavorite"
      class="removeFavoriteButton"
      v-if="!isHiddenRemoveFavorite"
    >
      Remove Favorite
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "PostById",
  data: () => ({
    post: null,
    activities: [],
    likeActivityId: null,
    dislikeActivityId: null,
    favoriteActivityId: null,
    isHiddenRemoveLike: true,
    isHiddenRemoveDislike: true,
    isHiddenRemoveFavorite: true,
    likeabilityStatus: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get(
          "http://localhost:8084/find_selected_post_for_logged_user?id=" +
            localStorage.getItem("selectedPostId") +
            "&logId=" +
            localStorage.getItem("selectedUserId")
        )
        .then((response) => {
          this.post = response.data;
        })
        .catch(console.log);
    },
    likePost() {
      this.$http
        .get(
          "http://localhost:8084/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId")
        )
        .then((response) => {
          this.activities = response.data;
          for (var i = 0; i < this.activities.length; i++) {
            if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              this.activities[i].liked_status == 0
            ) {
              alert("You have already liked this post!");
              this.isHiddenRemoveLike = false;
              this.likeActivityId = this.activities[i].id;
              return;
            } else if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              (this.activities[i].liked_status == 1 ||
                this.activities[i].liked_status == 2)
            ) {
              this.likeActivityId = this.activities[i].id;
              this.$http
                .post("http://localhost:8084/update_activity/", {
                  id: this.likeActivityId,
                  likedStatus: 0,
                  IsFavorite: false,
                })
                .then((response) => {
                  console.log(response);
                  this.isHiddenRemoveLike = false;
                  this.isHiddenRemoveDislike = true;
                  alert("You have liked this post.");
                })
                .catch((er) => {
                  console.log(er.response.data);
                });
              return;
            }
          }

          this.$http
            .post("http://localhost:8084/activity/", {
              postID: localStorage.getItem("selectedPostId"),
              userID: localStorage.getItem("selectedUserId"),
              likedStatus: 0,
              IsFavorite: false,
            })
            .then((response) => {
              this.likeActivityId = response.data;
              this.isHiddenRemoveLike = false;
              this.isHiddenRemoveDislike = true;
              alert("You have liked this post.");
            })
            .catch((er) => {
              console.log(er.response.data);
            });
        })
        .catch(console.log);
    },
    dislikePost() {
      this.$http
        .get(
          "http://localhost:8084/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId")
        )
        .then((response) => {
          this.activities = response.data;
          for (var i = 0; i < this.activities.length; i++) {
            if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              this.activities[i].liked_status == 1
            ) {
              alert("You have already disliked this post!");
              this.isHiddenRemoveDislike = false;
              this.dislikeActivityId = this.activities[i].id;
              return;
            } else if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              (this.activities[i].liked_status == 0 ||
                this.activities[i].liked_status == 2)
            ) {
              this.dislikeActivityId = this.activities[i].id;
              this.$http
                .post("http://localhost:8084/update_activity/", {
                  id: this.dislikeActivityId,
                  likedStatus: 1,
                  IsFavorite: false,
                })
                .then((response) => {
                  console.log(response);
                  this.isHiddenRemoveDislike = false;
                  this.isHiddenRemoveLike = true;
                  alert("You have disliked this post.");
                })
                .catch((er) => {
                  console.log(er.response.data);
                });
              return;
            }
          }

          this.$http
            .post("http://localhost:8084/activity/", {
              postID: localStorage.getItem("selectedPostId"),
              userID: localStorage.getItem("selectedUserId"),
              likedStatus: 1,
              IsFavorite: false,
            })
            .then((response) => {
              this.dislikeActivityId = response.data;
              this.isHiddenRemoveDislike = false;
              this.isHiddenRemoveLike = true;
              alert("You have disliked this post.");
            })
            .catch((er) => {
              console.log(er.response.data);
            });
        })
        .catch(console.log);
    },
    favoritePost() {
      this.$http
        .get(
          "http://localhost:8084/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId")
        )
        .then((response) => {
          this.activities = response.data;
          for (var i = 0; i < this.activities.length; i++) {
            if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              this.activities[i].is_favorite == true
            ) {
              alert("You have already favorited this post!");
              this.isHiddenRemoveFavorite = false;
              this.favoriteActivityId = this.activities[i].id;
              this.likeabilityStatus = this.activities[i].liked_status;
              return;
            } else if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              this.activities[i].is_favorite == false
            ) {
              this.favoriteActivityId = this.activities[i].id;
              this.likeabilityStatus = this.activities[i].liked_status;
              this.$http
                .post("http://localhost:8084/update_activity/", {
                  id: this.favoriteActivityId,
                  likedStatus: this.likeabilityStatus,
                  IsFavorite: true,
                })
                .then((response) => {
                  console.log(response);
                  this.isHiddenRemoveFavorite = false;
                  alert("You have favorited this post.");
                })
                .catch((er) => {
                  console.log(er.response.data);
                });
              this.isHiddenRemoveFavorite = false;
              return;
            }
          }

          this.$http
            .post("http://localhost:8084/activity/", {
              postID: localStorage.getItem("selectedPostId"),
              userID: localStorage.getItem("selectedUserId"),
              likedStatus: this.likeabilityStatus,
              IsFavorite: true,
            })
            .then((response) => {
              this.favoriteActivityId = response.data;
              this.isHiddenRemoveFavorite = false;
              alert("You have favorited this post.");
            })
            .catch((er) => {
              console.log(er.response.data);
            });
        })
        .catch(console.log);
    },
    removeFavorite() {
      this.$http
        .get(
          "http://localhost:8084/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId")
        )
        .then((response) => {
          this.activities = response.data;
          for (var i = 0; i < this.activities.length; i++) {
            if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              (this.activities[i].is_favorite == true ||
                this.activities[i].is_favorite == false)
            ) {
              this.favoriteActivityId = this.activities[i].id;
              this.likeabilityStatus = this.activities[i].liked_status;
            }
          }
        })
        .catch(console.log);

      if (this.favoriteActivityId == null) {
        alert("You have not favorited this post.");
        return;
      } else if (this.likeabilityStatus == null) {
        this.likeabilityStatus = 2;
      }

      this.$http
        .post("http://localhost:8084/update_activity/", {
          id: this.favoriteActivityId,
          likedStatus: this.likeabilityStatus,
          IsFavorite: false,
        })
        .then((response) => {
          console.log(response);
          this.isHiddenRemoveFavorite = true;
          alert("You have removed favorite for this post.");
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    removeLike() {
      this.$http
        .get(
          "http://localhost:8084/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId")
        )
        .then((response) => {
          this.activities = response.data;
          for (var i = 0; i < this.activities.length; i++) {
            if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              this.activities[i].liked_status == 0
            ) {
              this.likeActivityId = this.activities[i].id;
            } else if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              (this.activities[i].liked_status == 1 ||
                this.activities[i].liked_status == 2)
            ) {
              this.isHiddenRemoveLike = true;
              alert("You have not liked this post.");
              return;
            }
          }
        })
        .catch(console.log);

      if (this.likeActivityId == null) {
        alert("You have not liked this post.");
        return;
      }

      this.$http
        .post("http://localhost:8084/update_activity/", {
          id: this.likeActivityId,
          likedStatus: 2,
          IsFavorite: false,
        })
        .then((response) => {
          console.log(response);
          this.isHiddenRemoveLike = true;
          alert("You have removed like for this post.");
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    removeDislike() {
      this.$http
        .get(
          "http://localhost:8084/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId")
        )
        .then((response) => {
          this.activities = response.data;
          for (var i = 0; i < this.activities.length; i++) {
            if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              this.activities[i].liked_status == 0
            ) {
              this.dislikeActivityId = this.activities[i].id;
            } else if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              (this.activities[i].liked_status == 0 ||
                this.activities[i].liked_status == 2)
            ) {
              this.isHiddenRemoveDislike = true;
              alert("You have not disliked this post.");
              return;
            }
          }
        })
        .catch(console.log);

      if (this.dislikeActivityId == null) {
        alert("You have not disliked this post.");
        return;
      }

      this.$http
        .post("http://localhost:8084/update_activity/", {
          id: this.dislikeActivityId,
          likedStatus: 2,
          IsFavorite: false,
        })
        .then((response) => {
          console.log(response);
          this.isHiddenRemoveDislike = true;
          alert("You have removed dislike for this post.");
        })
        .catch((er) => {
          console.log(er.response.data);
        });
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

.likeButton {
  width: 120px;
  margin-left: 20%;
}

.removeLikeButton {
  width: 150px;
  margin-left: 3%;
}

.dislikeButton {
  width: 120px;
  margin-left: 3%;
}

.removeDislikeButton {
  width: 160px;
  margin-left: 3%;
}

.favoriteButton {
  width: 120px;
  margin-left: 3%;
}

.removeFavoriteButton {
  width: 190px;
  margin-left: 3%;
}
</style>