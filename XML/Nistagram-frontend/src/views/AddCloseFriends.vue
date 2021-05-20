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
          style="width: 90%; height: 300px; overflow-y: scroll"
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
                  v-on:click="addCloseFriend(user)"
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
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.id = localStorage.getItem("userId");
      this.$http
        .get(
          "http://localhost:8080/find_all_mutual_followers_for_user?id=" +
            this.id
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
    addCloseFriend(item) {
      this.$http
        .post("http://localhost:8080/create_close_friend/", {
          classic_user_id: item.classic_user_id,
          close_friend_user_id: item.follower_user_id,
        })
        .then(alert("You have added this friend as close friend."))
        .catch((er) => {
          console.log(er.response.data);
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
</style>