<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h1>Selected Post</h1>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg10 class="space-bottom">
          <v-card class="card">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{ post.description }}</v-list-item-subtitle>
                <v-list-item-title>{{ post.tags }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="post.type == 'VIDEO'">
              <v-list-item-content>
                <video width="320" height="440" controls>
                  <source
                    :src="require(`../../../Media/${post.path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="post.type != 'VIDEO'">
              <v-list-item-content>
                <img
                  :src="require(`../../../Media/${post.path}`)"
                  alt
                  class="icon"
                  width="320"
                  height="440"
                />
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle
                  v-text="
                    post.country +
                    ' ' +
                    post.city +
                    ' ' +
                    post.street_name +
                    ' ' +
                    post.street_number
                  "
                />
                <v-list-item-subtitle>{{
                  post.creation_date
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>

          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "PostByIdWithoutActivity",
  data: () => ({
    post: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get(
          "http://localhost:8084/find_selected_post_for_logged_user?id=" +
            localStorage.getItem("mySelectedPostId") +
            "&logId=" +
            localStorage.getItem("mySelectedUserId")
        )
        .then((response) => {
          this.post = response.data;
        })
        .catch(console.log);
    },
  },
};
</script>

<style scoped>
.spacingOne {
  height: 50px;
}

.title {
  margin-left: 44%;
}

.spacingTwo {
  height: 50px;
}

.card {
  margin-left: 20%;
}
</style>