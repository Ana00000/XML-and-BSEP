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
        color="#aba7ff"
        elevation="24"
        x-large
        raised
        rounded
        >Follow</v-btn
      >
    </div>
    <div class="FollowingButton">
      <v-btn
        v-if="!isHiddenFollowing"
        color="#aba7ff"
        elevation="24"
        x-large
        raised
        rounded
        >Following</v-btn
      >
    </div>
    <div class="SendFollowRequestButton">
      <v-btn
        v-if="!isHiddenSendFollowRequest"
        v-on:click="sendFollowRequest"
        color="#aba7ff"
        elevation="24"
        x-large
        raised
        rounded
        >Send Follow Request</v-btn
      >
    </div>
    <div class="FollowRequestSentButton">
      <v-btn
        v-if="!isHiddenFollowRequestSent"
        color="#aba7ff"
        elevation="24"
        x-large
        raised
        rounded
        >Follow Request Sent</v-btn
      >
    </div>
     </div>
</template>

<script>
export default {
  name: "SelectedUser",
  data: () => ({
    username: null,
    firstName: null,
    lastName: null,
    token: null,
    isHiddenFollow: true,
    isHiddenFollowing: true,
    isHiddenSendFollowRequest: true,
    isHiddenFollowRequestSent: true,
  }),
  mounted() {
    this.selectedUser = localStorage.getItem("selectedUserId");
    this.token = localStorage.getItem("token");
    this.logId= localStorage.getItem("userId");
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
        .get("http://localhost:8080/api/user/find_selected_user_by_id?id=" + this.selectedUser+"&logId=" + this.logId)
        .then((resp) => {
          this.setUserInfo(resp.data);
          console.log(resp.data);
          console.log(resp.data.profileVisibility);
          console.log(resp.data.followingStatus);

          if(resp.data.profileVisibility == "PUBLIC_VISIBILITY")
            console.log("PUBLIC JE");
          else if(resp.data.profileVisibility == "PRIVATE")
            console.log("PRIVATE JE");
          else
            console.log("NISTA JE");
          
          if(resp.data.followingStatus == "FOLLOWING"){
            this.isHiddenFollowing  = false
            this.isHiddenFollow = true
            this.isHiddenSendFollowRequest = true
            this.isHiddenFollowRequestSent = true;
            console.log("TRUE JE");
          }else if(resp.data.followingStatus == "NOT FOLLOWING" && resp.data.profileVisibility == "PRIVATE"){
             this.isHiddenFollowing  = true
            this.isHiddenFollow = true
            this.isHiddenSendFollowRequest = false
            this.isHiddenFollowRequestSent = true;
            console.log("FALSE JE I PRIVATE JE");
          }else if(resp.data.followingStatus == "NOT FOLLOWING" && resp.data.profileVisibility == "PUBLIC_VISIBILITY"){
            this.isHiddenFollowing  = true
            this.isHiddenFollow = false
            this.isHiddenSendFollowRequest = true
            this.isHiddenFollowRequestSent = true;
            console.log("FALSE JE I PUBLIC");
          }else if(resp.data.followingStatus == "PENDING" && resp.data.profileVisibility == "PRIVATE"){
            this.isHiddenFollowing  = true
            this.isHiddenFollow = true
            this.isHiddenSendFollowRequest = true
            this.isHiddenFollowRequestSent = false;
             console.log("PENDING JE I PRIVATE");
          }else
            console.log("OPET NISTA JE");
        
        })
        .catch(console.log("Didn't set user info!"));
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
        .post("http://localhost:8080/api/user/create_following/", {
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
    sendFollowRequest(){
       this.isHiddenFollow = true; 
       this.isHiddenFollowing = true; 
       this.isHiddenSendFollowRequest = true;
       this.isHiddenFollowRequestSent = false;

       this.$http
        .post("http://localhost:8080/api/requests/create_follow_request/", {
          classic_user_id: this.logId,
          follower_user_id: this.selectedUser,
      
        })
        .then((resp) => {
          console.log(resp.data);
          alert("Successfully sent follow request!");
          window.location.reload();
        })
        .catch((err) => console.log(err));
    }
    
  }
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
</style>