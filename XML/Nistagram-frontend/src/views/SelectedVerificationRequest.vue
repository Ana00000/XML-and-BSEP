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
            label="First Name"
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
            label="Last Name"
            v-model="lastName"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>
      <v-row rows="1">
        <v-col cols="5" />
        <v-col cols="3">
          <v-text-field
            label="Status"
            v-model="status"
            color="light-blue darken-4"
            readonly
          />
        </v-col>
      </v-row>
    </v-container>

    <!-- IMG -->
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h1>Official Document Path</h1>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg10 class="space-bottom">
          <v-card class="card">
            
        
            <v-list-item three-line>
              <v-list-item-content>
                <img
                  :src="require(`/app/public/uploads/${path}`)"
                  alt
                  class="icon"
                  width="320"
                  height="440"
                />
              </v-list-item-content>
            </v-list-item>
          </v-card>
        </v-flex>
      </v-layout>
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
  name: "SelectedVerificationRequest",
  data: () => ({
    id: null,
    token: null,
    firstName: null,
    lastName: null,
    status: null,
    path: null,
   
  }),
  mounted() {
    this.selectedRequestId = localStorage.getItem("selectedVerificationRequestId");
    this.token = localStorage.getItem("token");
    console.log(this.selectedRequestId);
    this.init();
  },
  methods: {
    init() {

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
          "https://localhost:8080/api/user/auth/check-find-verification-request-by-id-permission/",{
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

      this.getRequest();
    },
    getRequest() {
      console.log(this.selectedVerificationRequestId);
      this.$http
        .get("https://localhost:8080/api/requests/find_verification_request_by_id?id=" + this.selectedVerificationRequestId,{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        })
        .then((resp) => {
          this.setRequestInfo(resp.data);
          console.log(resp.data);
          console.log(resp.data.classic_user_id);
          
        })
        .catch(console.log("Didn't set user info!"));
    },
    setRequestInfo(item) {
      this.id = item.id;
      this.firstName = item.first_name;
      this.lastName =  item.last_name;
      this.path = item.official_document_path;
      if(item.registered_user_category == 0){
          this.status = "INFLUENCER";
      }else if(item.registered_user_category == 1){
          this.status = "SPORTS";
      }else if(item.registered_user_category == 2){
          this.status = "NEW_MEDIA";
      }else if(item.registered_user_category == 3){
          this.status = "BUSINESS";
      }else if(item.registered_user_category == 4){
          this.status = "BRAND";
      }else if(item.registered_user_category == 5){
          this.status = "ORGANIZATION";
      }else if(item.registered_user_category == 6){
          this.status = "NONE";
      }
      
    },
    acceptRequest() {
    
     this.$http
        .post("https://localhost:8080/api/user/accept_verification_request/", {
          id: this.selectedVerificationRequest,
        },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          })
        .then((resp) => {
          console.log(resp.data);
          alert("Accepted verification request!");
           window.location.href = "https://localhost:8081/verificationRequests";
        })
        .catch((err) => console.log(err));
      
    },
    rejectRequest(){
       this.$http
        .post("https://localhost:8080/api/requests/reject_verification_request?id="+this.selectedVerificationRequest, {},{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        })
        .then((resp) => {
          console.log(resp.data);
          alert("Rejected  verification request!");
           window.location.href = "https://localhost:8081/verificationRequests";
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