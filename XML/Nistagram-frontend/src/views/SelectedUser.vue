<template>
  <div>
    <br />
    <v-container fluid class="container">
      <v-row aria-rowspan="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Username"
            v-model="username"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="First name"
            v-model="firstName"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Last name"
            v-model="lastName"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>
    </v-container>

    <div class="FollowButton">
      <v-btn
        v-if="!isHiddenFollow"
        v-on:click="followProfile"
        color="info mb-5"
        elevation="24"
        x-large
        raised
        >Follow</v-btn
      >
    </div>
    <div class="FollowingButton">
      <v-btn
        v-if="!isHiddenFollowing"
        color="info mb-5"
        elevation="24"
        x-large
        raised
        >Following</v-btn
      >
    </div>
    <div class="SendFollowRequestButton">
      <v-btn
        v-if="!isHiddenSendFollowRequest"
        v-on:click="sendFollowRequest"
        color="info mb-5"
        elevation="24"
        x-large
        raised
        >Send Follow Request</v-btn
      >
    </div>
    <div class="FollowRequestSentButton">
      <v-btn
        v-if="!isHiddenFollowRequestSent"
        color="info mb-5"
        elevation="24"
        x-large
        raised
        >Follow Request Sent</v-btn
      >
    </div>

    <v-container grid-list-lg >
      <div class="spacingOne" />
      <div class="title">
        <h1>Stories</h1>
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
                  <source :src="require(`/app/public/uploads/${ item.path }`)" type="video/mp4">
                </video>
              </v-list-item-content>
            </v-list-item>

             <v-list-item three-line v-if="item.type != 'VIDEO'">
              <v-list-item-content> 
                <img :src="require(`/app/public/uploads/${ item.path }`)" alt class="icon" width="320" height="240"/>
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

    <v-container grid-list-lg >
      <div class="spacingOne" />
      <div class="title">
        <h1>Posts</h1>
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
                  <source :src="require(`/app/public/uploads/${ item.path }`)" type="video/mp4">
                </video>
              </v-list-item-content>
            </v-list-item>

             <v-list-item three-line v-if="item.type != 'VIDEO'">
              <v-list-item-content> 
                <img :src="require(`/app/public/uploads/${ item.path }`)" alt class="icon" width="320" height="240"/>
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

  </div>
</template>

<script>
export default {
  name: "SelectedUser",
  data: () => ({
    publicPath: process.env.VUE_APP_BASE_URL,
    username: null,
    firstName: null,
    lastName: null,
    token: null,
    isHiddenFollow: true,
    isHiddenFollowing: true,
    isHiddenSendFollowRequest: true,
    isHiddenFollowRequestSent: true,
    posts: [],
    stories: [],
  }),
  mounted() {
    this.selectedUser = localStorage.getItem("selectedUserId");
    this.token = localStorage.getItem("token");
    this.logId = localStorage.getItem("userId");
    console.log(this.selectedUser);
    this.init();
  },
  methods: {
    init() {
      this.getUser();
    },
    getUser() {
      console.log(this.selectedUser);
      this.$http
        .get("https://localhost:8080/api/user/find_selected_user_by_id?id=" + this.selectedUser+"&logId=" + this.logId)
        .then((resp) => {
          this.setUserInfo(resp.data);
          console.log(resp.data);
          console.log(resp.data.profileVisibility);
          console.log(resp.data.followingStatus);

          if(resp.data.profileVisibility == "PUBLIC_VISIBILITY")
            console.log("PUBLIC JE");
          else if(resp.data.profileVisibility == "PRIVATE")
            console.log("PRIVATE JE");
          else console.log("NISTA JE");

          if (resp.data.followingStatus == "FOLLOWING") {
            this.isHiddenFollowing = false;
            this.isHiddenFollow = true;
            this.isHiddenSendFollowRequest = true;
            this.isHiddenFollowRequestSent = true;
            console.log("TRUE JE");
          } else if (
            resp.data.followingStatus == "NOT FOLLOWING" &&
            resp.data.profileVisibility == "PRIVATE"
          ) {
            this.isHiddenFollowing = true;
            this.isHiddenFollow = true;
            this.isHiddenSendFollowRequest = false;
            this.isHiddenFollowRequestSent = true;
            console.log("FALSE JE I PRIVATE JE");
          }else if(resp.data.followingStatus == "NOT FOLLOWING" && resp.data.profileVisibility == "PUBLIC_VISIBILITY"){
            this.isHiddenFollowing  = true
            this.isHiddenFollow = false
            this.isHiddenSendFollowRequest = true
            this.isHiddenFollowRequestSent = true;
            console.log("FALSE JE I PUBLIC");
          } else if (
            resp.data.followingStatus == "PENDING" &&
            resp.data.profileVisibility == "PRIVATE"
          ) {
            this.isHiddenFollowing = true;
            this.isHiddenFollow = true;
            this.isHiddenSendFollowRequest = true;
            this.isHiddenFollowRequestSent = false;
            console.log("PENDING JE I PRIVATE");
          } else console.log("OPET NISTA JE");
        })
        .catch(console.log("Didn't set user info!"));

        this.$http
        .get("https://localhost:8080/api/post/find_all_posts_for_reg?id=" + this.selectedUser + "&logId=" + localStorage.getItem("userId"),{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        }
        )
        .then((response) => {
          this.posts = response.data;
        })
        .catch(console.log);

        this.$http
        .get("https://localhost:8080/api/story/find_all_stories_for_reg?id=" + this.selectedUser + "&logId=" + localStorage.getItem("userId"),{
          headers: {
            Authorization: "Bearer " + this.token,
          },
        })
        .then((response) => {
          this.stories = response.data;
        })
        .catch(console.log);
    },
    setUserInfo(item) {
      this.username = item.username;
      this.firstName = item.firstName;
      this.lastName = item.lastName;
    },
    followProfile() {
      this.isHiddenFollow = true;
      this.isHiddenFollowing = false;
      this.isHiddenSendFollowRequest = true;
      this.isHiddenFollowRequestSent = true;
      
     this.$http
        .post("https://localhost:8080/api/user/create_following/", {
          classic_user_id: this.logId,
          following_user_id: this.selectedUser,
        },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
        .then((resp) => {
          console.log(resp.data);
          alert("Successfully followed profile!");
          window.location.reload();
        })
        .catch((err) => console.log(err));
    },
    sendFollowRequest() {
      this.isHiddenFollow = true;
      this.isHiddenFollowing = true;
      this.isHiddenSendFollowRequest = true;
      this.isHiddenFollowRequestSent = false;

       this.$http
        .post("https://localhost:8080/api/requests/create_follow_request/", {
          classic_user_id: this.logId,
          follower_user_id: this.selectedUser,
        },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        })
        .then((resp) => {
          console.log(resp.data);
          alert("Successfully sent follow request!");
          window.location.reload();
        })
        .catch((err) => console.log(err));
    },
    getPost(item){
      localStorage.setItem("selectedUserId", item.user_id);
      localStorage.setItem("selectedPostId", item.post_id);

      window.location.href = "https://localhost:8081/postById";
    }
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