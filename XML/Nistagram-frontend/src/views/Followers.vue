<template>
  <v-container>
    <v-layout row wrap>
      <v-card loading class="mx-auto"
        id="followerCard"
      >
      
        <v-toolbar color="light-blue darken-4" dark><v-toolbar-title class="flex text-center">FOLLOWERS</v-toolbar-title></v-toolbar>
        <v-list two-line>
          <v-list-item-group
            active-class="indigo--text"
            v-model="selectedFollower"
            single
          >
            <template v-for="(follower, id) in followers">
              <v-list-item 
              :key="follower.id" 
              :value="follower"
                v-on:click="redirectToFollower"
              >
                <template>
                  <v-list-item-content class = "center text-wrap">
                  

                    <v-list-item-subtitle v-text="'Follower id: ' + follower.id" />
                   
                    <v-list-item-subtitle
                      v-text="
                        'Username: ' + follower.username
                      "
                    />
                    <v-list-item-subtitle
                      v-text="'First Name: ' + follower.firstName"
                    />
                    <v-list-item-subtitle
                      v-text="' Last name: ' + follower.lastName"
                    />
            
                  </v-list-item-content>
                </template>
              </v-list-item>
              <v-divider
                v-if="`A-${id}` < followers.length - 1"
                :key="`A-${id}`"
              />
            </template>
          </v-list-item-group>
        </v-list>
      </v-card>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "followers",
  data: () => ({
    followers: [],
    user: null,
    userId: null,
    selectedFollower: null,
    token: null
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token =localStorage.getItem("token");
      this.getFollowers();
    },
    getFollowers() {
      var userId = localStorage.getItem("userId");
      console.log("USER ID: " + userId);
      this.$http
        .get("https://localhost:8080/find_all_followers_for_user?id=" + userId)
        .then((res) => {
          this.followers = res.data;
        })
        .catch((err) => console.log(err));
    },
    redirectToFollower() {
      localStorage.setItem(
        "selectedFollowerId",
        this.selectedFollower.id
      );
      window.location.href = "https://localhost:8081/selectedFollower";
    },
  },
};
</script>

<style scoped>
.helloMessage {
  font-weight: bolder;
  font-size: 20px;
  height: 50px;
}
.center {
  
  padding: 10px;
  text-align: center;
}
#followerCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>