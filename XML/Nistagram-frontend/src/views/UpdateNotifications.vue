<template>
  <div class="searchDiv">
    <div class="spacingOne" />
    <div class="title">
      <h1>You can update notifications here!</h1>
    </div>
    <div class="spacingTwo" />

    <v-container>
      <v-layout row wrap>
        <v-card
          class="mx-auto"
          style="width: 90%; height: 300px; overflow-y: scroll"
        >
          <v-toolbar color="light-blue darken-4"> </v-toolbar>
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
                >
                  <template>
                    <v-list-item-content>
                      <v-list-item-subtitle
                        v-text="'USERNAME: ' + user.username"
                        class="containerDiv"
                      ></v-list-item-subtitle>
                      <v-list-item-subtitle
                        v-text="'FIRST NAME: ' + user.firstName"
                        class="containerDiv"
                      ></v-list-item-subtitle>
                      <v-list-item-subtitle
                        v-text="'LAST NAME: ' + user.lastName"
                        class="containerDiv"
                      ></v-list-item-subtitle>
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

    <v-btn color="info mb-5" v-on:click="addPostNotificationForUser">
      Allow post notifications
    </v-btn>

    <v-btn color="info mb-5" v-on:click="addStoryNotificationForUser">
      Allow story notifications
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "UpdateNotifications",
  data: () => ({
    users: [],
    token: null,
    selectedUser: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.id = localStorage.getItem("userId");
      this.token = localStorage.getItem("token");
      this.$http
        .get("https://localhost:8080/api/user/check_if_authentificated/", {
          headers: {
            Authorization: "Bearer " + this.token,
          },
        })
        .then((resp) => {
          console.log(resp.data);
        })
        .catch((er) => {
          window.location.href = "https://localhost:8081/unauthorizedPage";
          console.log(er);
        });
      this.$http
        .get(
          "https://localhost:8080/api/user/find_all_classic_users_but_logged_in?id=" +
            this.id,
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          this.users = resp.data;
        })
        .catch(console.log);
    },
    addPostNotificationForUser() {
      this.$http
        .get(
          "https://localhost:8080/api/settings/find_profile_settings_by_user_id/" +
            localStorage.getItem("userId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.addPostNot(response.data.id);
        })
        .catch(console.log);
    },
    addPostNot(userSettingsId) {
      this.$http
        .post(
          "https://localhost:8080/api/settings/add_post_notifications_for_user/",
          {
            profile_settings_id: userSettingsId,
            post_notifications_profile_id: localStorage.getItem("selectedUserId"),
          },
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          alert("Allowed post notifications of selected user.");
          window.location.href = "https://localhost:8081/updateNotifications";
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },
    addStoryNotificationForUser() {
      this.$http
        .get(
          "https://localhost:8080/api/settings/find_profile_settings_by_user_id/" +
            localStorage.getItem("userId"),
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.addStoryNot(response.data.id);
        })
        .catch(console.log);
    },
    addStoryNot(userSettingsId) {
      this.$http
        .post(
          "https://localhost:8080/api/settings/add_story_notifications_for_user/",
          {
            profile_settings_id: userSettingsId,
            story_notifications_profile_id: localStorage.getItem("selectedUserId"),
          },
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          alert("Allowed story notifications of selected user.");
          window.location.href = "https://localhost:8081/updateNotifications";
        })
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

.containerDiv {
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