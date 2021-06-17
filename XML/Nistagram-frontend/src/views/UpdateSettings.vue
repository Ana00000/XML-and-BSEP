<template>
  <div>
    <v-container fluid class="container mt-1">
      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-select
            v-model="selectedUserVisibility"
            hint="Choose your visibility."
            :items="userVisibilities"
            item-text="state"
            return-object
            single-line
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="5" />
        <v-col cols="3">
          <v-select
            v-model="selectedMessageApprovalType"
            hint="Choose your visibility."
            :items="messageApprovalTypes"
            item-text="state"
            return-object
            single-line
          />
        </v-col>
      </v-row>

      <input type="checkbox" id="checkbox" v-model="isPostTaggable" />
      <label for="checkbox">{{ isPostTaggable }}</label>

      <input type="checkbox" id="checkbox" v-model="isStoryTaggable" />
      <label for="checkbox">{{ isStoryTaggable }}</label>

      <input type="checkbox" id="checkbox" v-model="isCommentTaggable" />
      <label for="checkbox">{{ isCommentTaggable }}</label>

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
          "https://localhost:8080/api/settings/find_profile_settings_by_user_id/",
          {
            userId: this.id,
          },
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
      if (item.user_visibility == 0) {
        this.selectedUserVisibility = "PRIVATE_VISIBILITY";
      } else if (item.user_visibility == 1) {
        this.selectedUserVisibility = "PUBLIC_VISIBILITY";
      }
      if (item.message_approval_type == 0) {
        this.selectedMessageApprovalType = "PUBLIC";
      } else if (item.message_approval_type == 1) {
        this.selectedMessageApprovalType = "FRIENDS_ONLY";
      }
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