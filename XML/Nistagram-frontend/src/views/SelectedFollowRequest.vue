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
  name: "SelectedRequest",
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
        .get("http://localhost:80807/find_request_by_id?id=" + this.selectedRequestId)
        .then((resp) => {
          this.setRequestInfo(resp.data);
          this.requestFollower = resp.data.followerUserId
          console.log(resp.data);
          
        })
        .catch(console.log("Didn't set user info!"));
    },
    setRequestInfo(item) {
      this.id = item.id;
      this.followerUserId = item.followerUserId;
      
    },
    acceptRequest() {
      
     this.$http
        .post("http://localhost:8080/accept_follow_request/", {
          classic_user_id: this.logId,
          follower_user_id: this.requestFollower,
      
        })
        .then((resp) => {
          console.log(resp.data);
          alert("Accepted follow!");
          window.location.reload();
        })
        .catch((err) => console.log(err));
      
    },
    rejectRequest(){
       this.$http
        .post("http://localhost:8087/reject_follow_request?id=?"+this.selectedRequestId, {})
        .then((resp) => {
          console.log(resp.data);
          alert("Rejected  follow!");
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