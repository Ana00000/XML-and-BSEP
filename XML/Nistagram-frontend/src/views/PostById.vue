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
                    :src="require(`/app/public/uploads/${post.path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="post.type != 'VIDEO'">
              <v-list-item-content>
                <img
                  :src="require(`/app/public/uploads/${post.path}`)"
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

    <v-btn
      color="info mb-5"
      v-on:click="setVisibleCommentTextArea"
      class="commentButton"
    >
      Add comment
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
    
  <v-container grid-list-lg v-if="!isHiddenTagComment">
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allUserTags"
          :key="item.id"
          class="space-bottom"
        >
          <v-card class="mx-auto" v-on:click="getUserTag(item)">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  item.name
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>

    <v-container grid-list-lg v-if="!isHiddenTagComment">
      <v-layout row>
        <v-flex
          lg4
          v-for="item in allHashTags"
          :key="item.id"
          class="space-bottom"
        >
          <v-card class="mx-auto" v-on:click="getHashTag(item)">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  item.name
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>

    <v-textarea
      class="textArea"
      v-if="!isHiddenComment"
      v-model="text"
      solo
      name="input-5-4"
      label="Add description"
    ></v-textarea>


    <v-btn
      color="info mb-10"
      v-if="!isHiddenTagComment"
      v-on:click="isHiddenTagComment = true, isHiddenComment = false"
      class="cancelCommentButton"
    >
      Skip
    </v-btn>

    <v-btn
      color="info mb-10"
      v-if="!isHiddenTagComment"
      v-on:click="isHiddenTagComment = true, isHiddenComment = false"
    >
      Add tag
    </v-btn>

    <v-btn
      color="info mb-10"
      v-if="!isHiddenComment"
      v-on:click="cancleComment"
      class="cancelCommentButton"
    >
      Cancel
    </v-btn>

     <v-btn
      color="info mb-10"
      v-if="!isHiddenComment"
      v-on:click="createComment"
    >
      Finish
    </v-btn>

    <v-container grid-list-lg>
      <v-layout row>
        <v-flex lg4 v-for="item in allData" :key="item.id" class="space-bottom">
          <div class="spacingOne" />
          <v-card color="info" dark max-width="500">
            <v-card-title>
              <span class="title text-xs-center font-weight-light"
                >User comment</span
              >
            </v-card-title>

            <v-card-text class="headline font-weight-bold">
              {{ item.text }}
            </v-card-text>

            <v-card-text class="headline font-weight-bold">
              {{ item.tags }}
            </v-card-text>

            <v-card-actions>
              <v-list-item class="grow" v-if="!isHiddenUserName">
                <v-list-item-avatar
                  v-model="gender"
                  color="grey darken-3"
                  id="avatar"
                >
                  <v-img
                    v-if="item.gender == 'MALE'"
                    class="elevation-6"
                    alt=""
                    src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                  ></v-img>
                  <v-img
                    v-if="item.gender == 'FEMALE'"
                    class="elevation-6"
                    alt=""
                    src="https://avataaars.io/"
                  ></v-img>
                </v-list-item-avatar>
                <v-list-item-content>
                  <v-list-item-title>{{ item.userName }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-card-actions>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>

    <v-container grid-list-lg v-if="!isHiddenCollections">
      <div class="spacingOne" />
      <v-card-title class="justify-center">
        <h1 class="display-1">My Collections</h1>
      </v-card-title>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex
          lg4
          v-for="item in postCollections"
          :key="item.id"
          class="space-bottom"
        >
          <v-card class="mx-auto" v-on:click="getCollection(item)">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{ item.title }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
      <v-btn color="info mb-5" v-on:click="addPostToCollection"> Add </v-btn>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "PostById",
  data: () => ({
    publicPath: process.env.VUE_APP_BASE_URL,
    post: null,
    activities: [],
    likeActivityId: null,
    dislikeActivityId: null,
    favoriteActivityId: null,
    isHiddenRemoveLike: true,
    isHiddenRemoveDislike: true,
    isHiddenRemoveFavorite: true,
    likeabilityStatus: null,
    isHiddenCollections: true,
    isHiddenComment: true,
    isHiddenUserName: true,
    postCollections: [],
    postCollectionId: null,
    allPostComments: [],
    creationDate: "",
    text: "",
    userName: "",
    gender: "",
    allData: [],
    itemsHashtag: [],
    items: [],
    itemsUsertag:[],
    allTags: [],
    allUserTags: [],
    allHashTags: [],
    selectedTags: [],
    selectedUserTags: [],
    selectedHashTags: [],
    token: null,
    isHiddenTagComment: true,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");

      this.$http
        .get(
          "https://localhost:8080/api/user/check_if_authentificated/",{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
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
          "https://localhost:8080/api/user/auth/check-create-activity-permission/",{
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
          "https://localhost:8080/api/post/find_all_post_collections_for_reg?id=" +
            localStorage.getItem("userId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.postCollections = response.data;
        })
        .catch(console.log);

      this.$http
        .get(
          "https://localhost:8080/api/post/find_selected_post_for_logged_user?id=" +
            localStorage.getItem("selectedPostId") +
            "&logId=" +
            localStorage.getItem("selectedUserId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.post = response.data;
        })
        .catch(console.log);

      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_comments_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response)
          this.allPostComments = response.data;
          this.getData(response.data);
        })
        .catch(console.log);
      
      this.$http
        .get("https://localhost:8080/api/tag/find_all_taggable_users_comment/",{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
        .then((response) => {
          console.log(response.data);
          for (var i = 0; i < response.data.length; i++) {
            if (response.data[i].tag_type == 0 && response.data[i].user_id != this.userId) {
              this.itemsUsertag.push(response.data[i].name);
              this.allUserTags.push(response.data[i]);
            }
          }
        })
        .catch(console.log);
        ///find_all_hashtags/
      this.$http
        .get("https://localhost:8080/api/tag/find_all_hashtags/",{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
        .then((response) => {
          console.log(response.data);
          for (var i = 0; i < response.data.length; i++) {
            if (response.data[i].tag_type == 1) {
              this.allHashTags.push(response.data[i]);
              this.itemsHashtag.push(response.data[i].name);
            }
          }
        })
        .catch(console.log);
    },
    onOpen (key) {
      this.items = key === '@' ? this.itemsUsertag : this.itemsHashtag
    },
    getData(items) {
      for (var i = 0; i < items.length; i++) {
        this.getItem(items[i]);
      }
    },
    getUserTag(item) {
        this.selectedUserTags.push(item);
    },
    getHashTag(item) {
        this.selectedHashTags.push(item);
    },
    getItem(item){
      console.log(item);
       this.$http
          .get("https://localhost:8080/api/user/find_user_by_id?id=" + item.user_id,{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
          .then((r) => {
            this.userName = r.data.username;
            if (r.data.gender == 0) {
              this.gender = "MALE";
            } else if (r.data.gender == 1) {
              this.gender = "FEMALE";
            } else {
              this.gender = "OTHER";
            }
            this.$http
              .get("https://localhost:8080/api/tag/find_comment_tag_comments_for_comment/"+item.id,{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
              .then((resp) => {
                console.log(resp.data)
                this.allData.push({
                  id: r.data.id,
                  text: item.text,
                  gender: this.gender,
                  userName: this.userName,
                  tags: resp.data,
                });
                }).catch(console.log);
            })
          .catch(console.log);

        this.isHiddenUserName = false;
    },
    likePost() {
      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
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
                .post("https://localhost:8080/api/post/update_activity/", {
                  id: this.likeActivityId,
                  likedStatus: 0,
                  IsFavorite: false,
                },{
                  headers: {
                    Authorization: "Bearer " + this.token,
                  },
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
            .post("https://localhost:8080/api/post/activity/", {
              postID: localStorage.getItem("selectedPostId"),
              userID: localStorage.getItem("selectedUserId"),
              likedStatus: 0,
              IsFavorite: false,
            },{
              headers: {
                Authorization: "Bearer " + this.token,
              },
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
          "https://localhost:8080/api/post/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
              headers: {
                Authorization: "Bearer " + this.token,
              },
            }
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
                .post("https://localhost:8080/api/post/update_activity/", {
                  id: this.dislikeActivityId,
                  likedStatus: 1,
                  IsFavorite: false,
                },{
                  headers: {
                    Authorization: "Bearer " + this.token,
                  },
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
            .post("https://localhost:8080/api/post/activity/", {
              postID: localStorage.getItem("selectedPostId"),
              userID: localStorage.getItem("selectedUserId"),
              likedStatus: 1,
              IsFavorite: false,
            },{
                  headers: {
                    Authorization: "Bearer " + this.token,
                  },
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
          "https://localhost:8080/api/post/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
                  headers: {
                    Authorization: "Bearer " + this.token,
                  },
                }
        )
        .then((response) => {
          this.activities = response.data;
          for (var i = 0; i < this.activities.length; i++) {
            if (
              this.activities[i].user_id ==
                localStorage.getItem("selectedUserId") &&
              this.activities[i].is_favorite == true
            ) {
              alert(
                "You have already favorited this post!You can add this post to your collections now."
              );
              this.isHiddenCollections = false;
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
                .post("https://localhost:8080/api/post/update_activity/", {
                  id: this.favoriteActivityId,
                  likedStatus: this.likeabilityStatus,
                  IsFavorite: true,
                },{
                  headers: {
                    Authorization: "Bearer " + this.token,
                  },
                })
                .then((response) => {
                  console.log(response);
                  this.isHiddenRemoveFavorite = false;
                  alert(
                    "You have favorited this post.You can add this post to your collections now."
                  );
                  this.isHiddenCollections = false;
                })
                .catch((er) => {
                  console.log(er.response.data);
                });
              this.isHiddenRemoveFavorite = false;
              return;
            }
          }

          this.$http
            .post("https://localhost:8080/api/post/activity/", {
              postID: localStorage.getItem("selectedPostId"),
              userID: localStorage.getItem("selectedUserId"),
              likedStatus: this.likeabilityStatus,
              IsFavorite: true,
            },{
                  headers: {
                    Authorization: "Bearer " + this.token,
                  },
            })
            .then((response) => {
              this.favoriteActivityId = response.data;
              this.isHiddenRemoveFavorite = false;
              alert(
                "You have favorited this post.You can add this post to your collections now."
              );
              this.isHiddenCollections = false;
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
          "https://localhost:8080/api/post/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
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
        .post("https://localhost:8080/api/post/update_activity/", {
          id: this.favoriteActivityId,
          likedStatus: this.likeabilityStatus,
          IsFavorite: false,
        },{
          headers: {
            Authorization: "Bearer " + this.token,
          },
        })
        .then((response) => {
          console.log(response);
          this.isHiddenRemoveFavorite = true;
          this.isHiddenCollections = true;
          alert("You have removed favorite for this post.");
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    removeLike() {
      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
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
        .post("https://localhost:8080/api/post/update_activity/", {
          id: this.likeActivityId,
          likedStatus: 2,
          IsFavorite: false,
        },{
          headers: {
            Authorization: "Bearer " + this.token,
          },
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
          "https://localhost:8080/api/post/find_all_activities_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }  
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
        .post("https://localhost:8080/api/post/update_activity/", {
          id: this.dislikeActivityId,
          likedStatus: 2,
          IsFavorite: false,
        },{
          headers: {
            Authorization: "Bearer " + this.token,
          },
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
    getCollection(item) {
      this.postCollectionId = item.id;
    },
    addPostToCollection() {
      if (this.postCollectionId == null) {
        alert("You have not selected collection.");
        return;
      }

      this.$http
        .get(
          "https://localhost:8080/api/post/find_all_post_collection_posts_for_post?id=" +
            localStorage.getItem("selectedPostId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          for (var i = 0; i < response.data.length; i++) {
            if (this.postCollectionId == response.data[i].post_collection_id) {
              alert("You have already added this post to selected collection.");
              return;
            }
          }
          this.$http
            .post("https://localhost:8080/api/post/post_collection_posts/", {
              post_collection_id: this.postCollectionId,
              single_post_id: localStorage.getItem("selectedPostId"),
            },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
            .then((response) => {
              console.log(response.data);
              alert("You have added this post to your collection.");
            })
            .catch((er) => {
              console.log(er.response.data);
            });
        })
        .catch(console.log);
    },
    setVisibleCommentTextArea() {
      this.isHiddenTagComment = false;
    },
    createComment() {
      if (!this.validComment()) return;

      var currentDate = new Date();
      var date = currentDate.toISOString();
      console.log(date);

      this.$http.post("https://localhost:8080/api/post/comment/", {
        creation_date: date,
        user_id: localStorage.getItem("userId"),
        post_id: localStorage.getItem("selectedPostId"),
        text: this.text,
      },{
        headers: {
          Authorization: "Bearer " + this.token,
        },
      }).then((response) => {
        if(this.selectedUserTags.length != 0) {
          for(var i = 0; i < this.selectedUserTags.length; i++) {
            this.$http
              .post("https://localhost:8080/api/tag/comment_tag_comments/", {
                tag_id: this.selectedUserTags[i].id,
                comment_id: response.data,
              },{
                  headers: {
                    Authorization: "Bearer " + this.token,
                  },
               })
              .then((r) => {
                console.log(r.data);
              })
              .catch((er) => {
                console.log(er.response.data);
              });
          }
        }

        if(this.selectedHashTags.length != 0) {
          for(var j = 0; j < this.selectedHashTags.length; j++) {
            this.$http
              .post("https://localhost:8080/api/tag/comment_tag_comments/", {
                tag_id: this.selectedHashTags[j].id,
                comment_id: response.data,
              },{
                headers: {
                  Authorization: "Bearer " + this.token,
                },
              })
              .then((response) => {
                console.log(response.data);
              })
              .catch((er) => {
                console.log(er.response.data);
              });
          }
        }
        console.log(response.data);
        alert("Successfully created comment.");
        window.location.href = "https://localhost:8081/postById"
      })
      .catch((er) => {
        console.log(er.response.data);
      });
      this.text = "";
      this.isHiddenComment = true;
    },
    cancleComment() {
      console.log("Cancel comment");
      this.text = "";
      this.isHiddenComment = true;
    },
    validComment() {
      if (this.text.length < 1) {
        alert("Your text should contain at least 1 character!");
        return false;
      } else if (this.text.length > 200) {
        alert("Your text shouldn't contain more than 200 characters!");
        return false;
      }
      return true;
    },
  },
};
</script>

<style scoped>

.mention-item {
  padding: 4px 10px;
  border-radius: 4px;
}

.mention-selected {
  background: rgb(192, 250, 153);
}

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
  margin-left: 30%;
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

.commentButton {
  width: 180px;
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

.textArea {
  margin-left: 20%;
  width: 60%;
}

.tags {
  margin-left: 20%;
}

.addCommentButton {
  width: 120px;
  margin-left: 30%;
}

.cancelCommentButton {
  width: 120px;
  margin-left: 25%;
}

</style>