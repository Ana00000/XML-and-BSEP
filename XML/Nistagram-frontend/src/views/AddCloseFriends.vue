<template>
  <div class="searchDiv">
    <div class="spacingOne" />
    <div class="title">
      <h1>You can add close friends here!</h1>
      <div class="welcoming">Search by username</div>
    </div>
    <div class="spacingTwo" />

    <v-container>
      <v-layout row wrap>
        <v-card
          class="mx-auto"
          style="width: 50%; height: 300px; overflow-y: scroll"
        >
          <v-toolbar color="light-blue darken-4">
            <v-text-field
              hide-details
              prepend-icon="mdi-magnify"
              single-line
              v-model="searchInput"
              v-on:keyup="searchQuery()"
            />
          </v-toolbar>
          <v-list two-line>
            <v-list-item-group
              active-class="indigo--text"
              v-model="selectedUser"
              single
            >
              <template v-for="(user, id) in users">
                <v-list-item
                  :key="user.id"
                  :value="user"
                  v-on:click="setCloseFriend(user)"
                >
                  <template>
                    <v-list-item-content>
                      <v-list-item-subtitle
                        v-text="'Classic user id: ' + user.classic_user_id"
                        class="containerDiv"
                      />
                      <v-list-item-subtitle
                        v-text="'Follower user id: ' + user.follower_user_id"
                        class="containerDiv"
                      />
                      <v-list-item-subtitle
                        v-text="' '"
                        class="emptyContentClass"
                      ></v-list-item-subtitle>
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
    <div class="spacingTwo" />
    <v-btn
      color="info mb-5"
      v-on:click="addCloseFriend"
      class="addButton"
      x-large
    >
      Add
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "AddCloseFriends",
  data: () => ({
    searchInput: "",
    users: [],
    usersCopy: [],
    selectedUser: null,
    classicUserId: "",
    followerUserId: "",
    token: null,
    id: "",
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.token = localStorage.getItem("token");
      this.id = localStorage.getItem("userId");
      this.$http
        .get(
          "https://localhost:8080/api/user/find_all_mutual_followers_for_user?id=" +
            this.id,
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log(resp.data);

          this.users = resp.data;
          this.usersCopy = resp.data;
        })
        .catch(console.log);
    },
    searchQuery() {
      var resultOfSearch = [];
      for (var i = 0; i < this.usersCopy.length; i++) {
        if (
          this.usersCopy[i].username
            .toLowerCase()
            .includes(this.searchInput.toLowerCase())
        )
          resultOfSearch.push(this.usersCopy[i]);
      }
      this.users = resultOfSearch;
    },
    setCloseFriend(item) {
      this.classicUserId = item.classic_user_id;
      this.followerUserId = item.follower_user_id;
    },
    addCloseFriend() {
      this.$http
        .post("https://localhost:8080/api/user/create_close_friend/", {
          classic_user_id: this.classicUserId,
          close_friend_user_id: this.followerUserId,
        },
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
          )
        .then((response) => {
          console.log(response)
          alert("You have added this friend as close friend.");
        })
        .catch((er) => {
          console.log(er.response.data);
          if (er.response.status == 409)
            alert("You have already added this friend as close friend.");
        });
    },
  },
};
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
  margin-left: 15%;
}
.searchDiv {
  height: 840px;
}
.containerDiv {
  font-weight: bolder;
  font-size: 20px;
}

.spacingOne {
  height: 50px;
}

.title {
  margin-left: 35%;
}

.spacingTwo {
  height: 100px;
}

.addButton {
  margin-left: 45%;
  width: 150px;
}
</style>