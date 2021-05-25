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

    <v-btn
      color="info mb-5"
      v-on:click="setVisibleCommentTextArea"
      class="commentButton"
    >
      Add comment
    </v-btn>

    <template>
      <v-container>
        <v-textarea
          class="textArea"
          v-if="!isHiddenComment"
          v-model="text"
          solo
          name="input-5-4"
          label="Add Comment"
        ></v-textarea>
      </v-container>
    </template>

    <v-btn
      color="info mb-10"
      v-if="!isHiddenComment"
      v-on:click="createComment"
      class="addCommentButton"
    >
      Add
    </v-btn>

    <v-btn
      color="info mb-10"
      v-if="!isHiddenComment"
      v-on:click="cancleComment"
      class="cancelCommentButton"
    >
      Cancel
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
  </div>
</template>

<script>
export default {
  name: "PostByIdWithoutActivity",
  data: () => ({
    post: null,
    allPostComments: [],
    allData: [],
    text: "",
    userName: "",
    gender: "",
    isHiddenComment: true,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get(
          "http://localhost:8080/api/post/find_selected_post_for_logged_user?id=" +
            localStorage.getItem("mySelectedPostId") +
            "&logId=" +
            localStorage.getItem("mySelectedUserId")
        )
        .then((response) => {
          this.post = response.data;
        })
        .catch(console.log);

      this.$http
        .get(
          "http://localhost:8080/api/post/find_all_comments_for_post?id=" +
            localStorage.getItem("mySelectedPostId")
        )
        .then((response) => {
          alert(response.data);
          this.allPostComments = response.data;
          this.getData(response.data);
        })
        .catch(console.log);
    },
    setVisibleCommentTextArea() {
      this.isHiddenComment = false;
    },
    getData(items) {
      for (var i = 0; i < items.length; i++) {
        this.getItem(items[i]);
      }
    },
    getItem(item) {
      console.log(item);
      this.$http
        .get("http://localhost:8080/api/user/find_user_by_id?id=" + item.user_id)
        .then((r) => {
          this.userName = r.data.username;
          if (r.data.gender == 0) {
            this.gender = "MALE";
          } else if (r.data.gender == 1) {
            this.gender = "FEMALE";
          } else {
            this.gender = "OTHER";
          }
          this.allData.push({
            id: r.data.id,
            text: item.text,
            gender: this.gender,
            userName: this.userName,
          });
        })
        .catch(console.log);

      this.isHiddenUserName = false;
    },
    createComment() {
      if (!this.validComment()) return;

      var currentDate = new Date();
      var date = currentDate.toISOString();
      console.log(date);

      this.$http.post("http://localhost:8080/api/post/comment/", {
        creation_date: date,
        user_id: localStorage.getItem("userId"),
        post_id: this.post.post_id,
        text: this.text,
      }).then((r) => {
          console.log(r);
          alert("Successfully created comment");
          window.location.href = "http://localhost:8081/postByIdWithoutActivity"
        })
        .catch(console.log);
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

.commentButton {
  width: 180px;
  margin-left: 46%;
}

.textArea {
  margin-left: 20%;
  width: 60%;
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