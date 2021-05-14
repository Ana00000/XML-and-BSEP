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
  </div>
</template>

<script>
export default {
  name: "SelectedUser",
  data: () => ({
    username: null,
    firstName: null,
    lastName: null,
    token: null
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
        .get("http://localhost:8080/find_selected_user_by_id?id=" + this.selectedUser+"&logId=" + this.logId)
        .then((resp) => {
          this.setUserInfo(resp.data);
          console.log(resp.data);
          console.log(resp.data.profileVisibility);
          console.log(resp.data.followingCheck);

          if(resp.data.profileVisibility == "PUBLIC")
            console.log("PUBLIC JE");
          else if(resp.data.profileVisibility == "PRIVATE")
            console.log("PRIVATE JE");
          else
            console.log("NISTA JE");
          
          if(resp.data.followingCheck == true)
            console.log("TRUE JE");
          else if(resp.data.followingCheck == false)
            console.log("FALSE JE");
          else
            console.log("OPET NISTA JE");
        
        })
        .catch(console.log("Didn't set user info!"));
    },
    setUserInfo(item) {
      this.username = item.username;
      this.firstName = item.firstName;
      this.lastName = item.lastName;
      
    },
    
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