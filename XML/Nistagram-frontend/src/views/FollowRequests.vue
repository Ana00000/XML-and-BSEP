<template>
        <div class="searchDiv">
            <br/>
            <div>
                <h1 class="display-2">You can find users here!</h1>
            </div>
            <div class="welcoming"> Search users by username</div>
            <br/>

            <v-container>
                <v-layout row wrap>
                <v-card
                    class="mx-auto" style="width: 90%; height: 300px; overflow-y: scroll">
                    <v-toolbar
                    color="#13077d" dark>
                    </v-toolbar>
                    <v-list two-line>
                    <v-list-item-group active-class="indigo--text" v-model="selectedRequest" single>
                        <template v-for="(request,id) in requests" >
                        <v-list-item :key="request.id" :value="request" v-on:click="redirectToSelectedRequest" >
                            <template >  
                            <v-list-item-content>
                              <v-list-item-subtitle v-text="'FOLLOWER: ' + request.followerUserId" class="containerDiv"></v-list-item-subtitle>
                              <v-list-item-subtitle v-text="'REQUEST ID: ' + request.id" class="containerDiv"></v-list-item-subtitle>
                              <v-list-item-subtitle v-text="' '" class="emptyContentClass"></v-list-item-subtitle>
                            </v-list-item-content>
                            </template>
                        </v-list-item>
                       <v-divider
                v-if="`A-${id}` < requests.length - 1"
                :key="`A-${id}`"
              />
                        
                        </template>
                    </v-list-item-group>
                    </v-list>
                </v-card>
                </v-layout>
                </v-container>

            
                
        </div>
       
     
</template>

<script>

export default {
    name: 'FollowRequests',   
    data: () => ({
      requests: [],
      selectedRequest: null,
    }),
    mounted() {
        this.init();
    },
    methods: {
     
      init() {
      this.id = localStorage.getItem("userId");
      this.token = localStorage.getItem("token");

      console.log(this.id)
      console.log(this.token)
      this.$http
        .get("http://localhost:8087/find_all_pending_requests_for_user?id=" + this.id)
        .then((resp) => {
          console.log("USAO")
          this.requests = resp.data
          console.log("duzina: " + this.requests.length)
        })
        .catch(console.log);
    },
    
    redirectToSelectedRequest() {
    console.log("OK");
     console.log("usao ovde i id je: "+ this.selectedRequest.id);
        localStorage.setItem("selectedRequestId",this.selectedRequest.id);
        window.location.href = "http://localhost:8081/selectedRequest";
      
    }
    }
  }
</script>

<style scoped>

.cardClass {
  display: none; 
}
.template {
    min-height: 1000px;
}
.allUsers {
    position: absolute;
    right: 500px;
    top: 490px;
}
.welcoming {
    font-weight: bolder;
    font-size: 25px;
}
.containerDiv{
  font-weight: bolder;
  font-size: 20px;
  
}


</style>