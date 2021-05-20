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
                  <source :src="require(`../../../Media/${ item.path }`)" type="video/mp4">
                </video>
              </v-list-item-content>
            </v-list-item>

             <v-list-item three-line v-if="item.type != 'VIDEO'">
              <v-list-item-content> 
                <img :src="require(`../../../Media/${ item.path }`)" alt class="icon" width="320" height="240"/>
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
                  <source :src="require(`../../../Media/${ item.path }`)" type="video/mp4">
                </video>
              </v-list-item-content>
            </v-list-item>

             <v-list-item three-line v-if="item.type != 'VIDEO'">
              <v-list-item-content> 
                <img :src="require(`../../../Media/${ item.path }`)" alt class="icon" width="320" height="240"/>
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
  name: "SelectedUserForNotRegistered",
  data: () => ({
    username: null,
    firstName: null,
    lastName: null,
    token: null,
    posts: [],
    stories: [],
  }),
  mounted() {
    this.selectedUser = localStorage.getItem("selectedUserId");
    this.token = localStorage.getItem("token");
    this.logId = localStorage.getItem("userId");
    this.init();
  },
  methods: {
    init() {
      this.getUser();
    },
    getUser() {
      this.$http
        .get(
          "http://localhost:8080/find_selected_user_by_id?id=" +
            this.selectedUser +
            "&logId=" +
            this.logId
        )
        .then((resp) => {
          this.setUserInfo(resp.data);
        })
        .catch(console.log("Didn't set user info!"));

        this.$http
        .get("http://localhost:8084/find_all_posts_for_not_reg?id=" + this.selectedUser)
        .then((response) => {
          this.posts = response.data;
        })
        .catch(console.log);

        this.$http
        .get("http://localhost:8086/find_all_stories_for_not_reg?id=" + this.selectedUser)
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
      this.$http
        .post("http://localhost:8080/create_following/", {
          classic_user_id: this.logId,
          following_user_id: this.selectedUser,
        })
        .then((resp) => {
          console.log(resp.data);
          alert("Successfully followed profile!");
          window.location.reload();
        })
        .catch((err) => console.log(err));
    },
    sendFollowRequest() {
      this.$http
        .post("http://localhost:8087/create_follow_request/", {
          classic_user_id: this.logId,
          follower_user_id: this.selectedUser,
        })
        .then((resp) => {
          console.log(resp.data);
          alert("Successfully sent follow request!");
          window.location.reload();
        })
        .catch((err) => console.log(err));
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