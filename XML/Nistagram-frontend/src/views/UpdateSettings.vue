<template>
  <div>
    <v-container fluid class="container mt-1">
      <v-row>
        <v-col cols="5">
           <label>Selected User Visibility</label>
        </v-col>
        <v-col cols="3">
          <v-select
            v-model="selectedUserVisibility"
            hint="Choose your visibility."
            :items="userVisibilities"
            item-text="state"
            return-object
            single-line
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5">
           <label>Selected Message Approval Type</label>
        </v-col>
        <v-col cols="3">
          <v-select
            v-model="selectedMessageApprovalType"
            hint="Choose your visibility."
            :items="messageApprovalTypes"
            item-text="state"
            return-object
            single-line
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

       <v-row>
        <v-col cols="5">
           <label>Selected Likes Notifications</label>
        </v-col>
        <v-col cols="3">
          <v-select
            v-model="selectedLikesNotifications"
            hint="Choose likes notifications option."
            :items="notificationTypes"
            item-text="state"
            return-object
            single-line
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

       <v-row>
        <v-col cols="5">
           <label>Selected Comments Notifications</label>
        </v-col>
        <v-col cols="3">
          <v-select
            v-model="selectedCommentsNotifications"
            hint="Choose comments notifications option."
            :items="notificationTypes"
            item-text="state"
            return-object
            single-line
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5">
           <label>Selected Messages Notifications</label>
        </v-col>
        <v-col cols="3">
          <v-select
            v-model="selectedMessagesNotifications"
            hint="Choose messages notifications option."
            :items="notificationTypes"
            item-text="state"
            return-object
            single-line
            v-bind:readonly="isReadOnly"
          />
        </v-col>
      </v-row>

      <input type="checkbox" id="checkboxPostTaggable" v-model="isPostTaggable" v-bind:readonly="isReadOnly"/>
      <label for="checkboxPostTaggable" v-if="isPostTaggable == false"> Post can't be tagged</label>
      <label for="checkboxPostTaggable" v-if="isPostTaggable == true"> Post can be tagged</label>

      <input type="checkbox" id="checkboxStoryTaggable" v-model="isStoryTaggable" v-bind:readonly="isReadOnly"/>
      <label for="checkboxStoryTaggable" v-if="isStoryTaggable == false" > Story can't be tagged</label>
      <label for="checkboxStoryTaggable" v-if="isStoryTaggable == true"> Story can be tagged</label>

      <input type="checkbox" id="checkboxCommentTaggable" v-model="isCommentTaggable" v-bind:readonly="isReadOnly"/>
      <label for="checkboxCommentTaggable" v-if="isCommentTaggable == false"> Comment can't be tagged</label>
      <label for="checkboxCommentTaggable" v-if="isCommentTaggable == true"> Comment can be tagged</label>

      <div class="changeButton">
        <v-btn
          v-if="!isHiddenChange"
          v-on:click="
            (isHiddenUpdate = false),
              (isHiddenChange = true),
              (isHiddenCancel = false),
              (isReadOnly = false)
          "
          color="info mb-5"
          x-large
          >Change</v-btn
        >
      </div>

      <div class="updateButton">
        <v-btn
          v-if="!isHiddenUpdate"
          v-on:click="changeSettings"
          color="info mb-5"
          x-large
          >Save</v-btn
        >
      </div>

      <div class="cancelButton">
        <v-btn
          v-if="!isHiddenCancel"
          v-on:click="cancelChanges"
          color="info mb-5"
          x-large
          >Cancel</v-btn
        >
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "UpdateSettings",
  data: () => ({
    isReadOnly: true,
    id: "",
    token: null,
    settings: [],
    isHiddenChange: false,
    isHiddenUpdate: true,
    isHiddenCancel: true,
    userVisibilities: ["PRIVATE_VISIBILITY", "PUBLIC_VISIBILITY"],
    selectedUserVisibility: "",
    messageApprovalTypes: ["PUBLIC", "FRIENDS_ONLY"],
    selectedMessageApprovalType: "",
    notificationTypes: ["ALL_NOTIFICATIONS", "FRIENDS_NOTIFICATIONS", "NONE"],
    selectedLikesNotifications: "",
    selectedCommentsNotifications: "",
    selectedMessagesNotifications: "",
    isPostTaggable: null,
    isStoryTaggable: null,
    isCommentTaggable: null,
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
          console.log("User is authentificated!");
          console.log(resp.data);
        })
        .catch((er) => {
          window.location.href = "https://localhost:8081/unauthorizedPage";
          console.log(er);
        });

      this.$http
        .get(
          "https://localhost:8080/api/settings/find_profile_settings_by_user_id/"+this.id,
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          this.settings = response.data;
          this.setSettings(this.settings);
        })
        .catch(console.log);
    },
    setSettings(item) {
      console.log(item);
      this.selectedUserVisibility = item.user_visibility;
      if (this.selectedUserVisibility=="PUBLIC_VISIBILITY"){
        localStorage.setItem("userPrivacy", "PUBLIC");
      } else {
        localStorage.setItem("userPrivacy", "PRIVATE");
      }
      this.selectedMessageApprovalType = item.message_approval_type;
      this.selectedLikesNotifications = item.likes_notifications;
      this.selectedCommentsNotifications = item.comments_notifications;
      this.selectedMessagesNotifications = item.messages_notifications;
      this.isPostTaggable = item.is_post_taggable;
      this.isStoryTaggable = item.is_story_taggable;
      this.isCommentTaggable = item.is_comment_taggable;
    },
    cancelChanges() {
      window.location.href = "https://localhost:8081/updateSettings";
    },
    changeSettings() {
      this.$http
        .post(
          "https://localhost:8080/api/settings/update_profile_settings/",
          {
            user_id: this.id,
            user_visibility: this.selectedUserVisibility,
            message_approval_type: this.selectedMessageApprovalType,
            likes_notifications: this.selectedLikesNotifications,
            comments_notifications: this.selectedCommentsNotifications,
            messages_notifications: this.selectedMessagesNotifications,
            is_post_taggable: this.isPostTaggable,
            is_story_taggable: this.isStoryTaggable,
            is_comment_taggable: this.isCommentTaggable,
          },
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          alert("Successfully updated settings!");
          window.location.href = "https://localhost:8081/updateSettings";
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>

<style scoped>
.changeButton {
  position: absolute;
  right: 275px;
  top: 100px;
}

.updateButton {
  position: absolute;
  right: 250px;
  top: 300px;
}

.cancelButton {
  position: absolute;
  right: 240px;
  top: 500px;
}
</style>