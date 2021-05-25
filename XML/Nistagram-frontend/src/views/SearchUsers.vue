<template>
        <div class="searchDiv">

        <div class="spacingOne" />
        <div class="title">
          <h1>You can find users here!</h1>
          <div class="welcoming"> Search users by username</div>
        </div>
        <div class="spacingTwo" />

            <v-container>
                <v-layout row wrap>
                <v-card
                    class="mx-auto" style="width: 90%; height: 300px; overflow-y: scroll">
                    <v-toolbar
                    color="light-blue darken-4">
                    <v-text-field
                        hide-details
                        prepend-icon="mdi-magnify"
                        single-line
                        v-model="searchInput"
                        v-on:keyup="searchQuery()"
                        />
                    </v-toolbar>
                    <v-list two-line>
                    <v-list-item-group active-class="indigo--text" v-model="selectedUser" single>
                        <template v-for="(user,id) in users" >
                        <v-list-item :key="user.id" :value="user" v-on:click="redirectToSelectedUser" >
                            <template >  
                            <v-list-item-content>
                              <v-list-item-subtitle v-text="'USERNAME: ' + user.username" class="containerDiv"></v-list-item-subtitle>
                              <v-list-item-subtitle v-text="'FIRST NAME: ' + user.firstName" class="containerDiv"></v-list-item-subtitle>
                              <v-list-item-subtitle v-text="'LAST NAME: ' + user.lastName" class="containerDiv"></v-list-item-subtitle>
                              <v-list-item-subtitle v-text="' '" class="emptyContentClass"></v-list-item-subtitle>
                            </v-list-item-content>
                            </template>
                        </v-list-item>
                       <v-divider
                v-if="`A-${id}` < users.length - 1"
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
    name: 'SearchUsers',   
    data: () => ({
      searchInput: "",
      users: [],
      usersCopy : [],
      selectedUser: null,
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
        .get("http://localhost:8080/api/user/find_all_classic_users_but_logged_in?id=" + this.id)
        .then((resp) => {
          console.log("USAO")
          this.users = resp.data
          this.usersCopy = resp.data
          console.log("duzina: " + this.users.length)
        })
        .catch(console.log);
    },
          
  searchQuery() {
        var resultOfSearch = [];
        for(var i = 0; i < this.usersCopy.length; i++) {
        
          if(this.usersCopy[i].username.toLowerCase().includes(this.searchInput.toLowerCase()))
                resultOfSearch.push(this.usersCopy[i])
        }
        this.users = resultOfSearch;
    },
    
    redirectToSelectedUser() {
    console.log("OK");
     console.log("usao ovde i id je: "+ this.selectedUser.id);
        localStorage.setItem("selectedUserUsername",this.selectedUser.username);
        localStorage.setItem("selectedUserId",this.selectedUser.id);
        window.location.href = "http://localhost:8081/selectedUser";
      
    }
    }
  }
</script>

<style scoped>
.sort {
  padding-top: 15px;
  padding-bottom: 15px;
}

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
    margin-left: 7%;
}
.searchDiv {
     height: 840px; 
    
}
.containerDiv{
  font-weight: bolder;
  font-size: 20px;
}

.spacingOne {
  height: 50px;
}

.title {
  margin-left: 40%;
}

.spacingTwo {
  height: 100px;
}
</style>