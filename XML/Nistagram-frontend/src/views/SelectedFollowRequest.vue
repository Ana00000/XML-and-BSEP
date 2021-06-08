<template>
  <div>
    <br />
    <v-container fluid class="container">
      <v-row aria-rowspan="2">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Id"
            v-model="id"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>

      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Follower"
            v-model="followerUserId"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>
    
    </v-container>
 
  <div class="AcceptButton">
      <v-btn
        v-on:click="acceptRequest"
        color="#aba7ff"
        elevation="24"
        x-large
        raised
        rounded
        >Accept</v-btn
      >
    </div>

    <div class="RejectButton">
      <v-btn
        v-on:click="rejectRequest"
        color="#aba7ff"
        elevation="24"
        x-large
        raised
        rounded
        >Reject</v-btn
      >
    </div>
     </div>
</template>

<script>
export default {
  name: "SelectedFollowRequest",
  data: () => ({
    id: null,
    followerId: null,
    token: null,
    requestFollower: null
   
  }),
  mounted() {
    this.selectedRequestId = localStorage.getItem("selectedRequestId");
    this.token = localStorage.getItem("token");
    this.logId= localStorage.getItem("userId");
    console.log(this.selectedRequestId);
    this.init();
  },
  methods: {
    init() {
      this.getRequest();
    },
    getRequest() {
      console.log(this.selectedRequestId);
      this.$http
        .get("https://localhost:8080/api/requests/find_request_by_id?id=" + this.selectedRequestId)
        .then((resp) => {
          this.setRequestInfo(resp.data);
          this.requestFollower = resp.data.classic_user_id
          console.log(resp.data);
          console.log(resp.data.classic_user_id);
          
        })
        .catch(console.log("Didn't set user info!"));
    },
    setRequestInfo(item) {
      this.id = item.id;
      this.followerUserId = item.follower_user_id;
      
    },
    acceptRequest() {
      console.log(this.logId)
      console.log(this.requestFollower)
      
     this.$http
        .post("https://localhost:8080/api/user/accept_follow_request/", {
          classic_user_id: this.requestFollower,
          follower_user_id: this.logId,
      
        })
        .then((resp) => {
          console.log(resp.data);
          alert("Accepted follow!");
           window.location.href = "https://localhost:8081/followRequests";
        })
        .catch((err) => console.log(err));
      
    },
    rejectRequest(){
       this.$http
        .post("https://localhost:8080/api/requests/reject_follow_request?id="+this.selectedRequestId, {})
        .then((resp) => {
          console.log(resp.data);
          alert("Rejected  follow!");
           window.location.href = "https://localhost:8081/followRequests";
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