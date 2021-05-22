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
                <v-list-item-subtitle>{{ description }}</v-list-item-subtitle>
                <v-list-item-title>{{ tags }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="type == 'VIDEO'">
              <v-list-item-content>
                <video width="320" height="440" controls>
                  <source
                    :src="require(`/app/public/uploads/${path}`)"
                    type="video/mp4"
                  />
                </video>
              </v-list-item-content>
            </v-list-item>

            <v-list-item three-line v-if="type != 'VIDEO'">
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

    <v-btn color="info mb-5" v-on:click="likePost" class="likeButton">
      Like
    </v-btn>
    <v-btn color="info mb-5" v-on:click="dislikePost" class="dislikeButton">
      Dislike
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "PostById",
  data: () => ({
    publicPath: process.env.VUE_APP_BASE_URL,
    description: "",
    tags: [],
    path: "",
    type: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.description = localStorage.getItem("selectedPostDescription");
      this.tags = localStorage.getItem("selectedPostTags");
      this.path = localStorage.getItem("selectedPostPath");
      this.type = localStorage.getItem("selectedPostType");
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

.likeButton {
  width: 120px;
  margin-left: 30%;
}

.dislikeButton {
  width: 120px;
  margin-left: 30%;
}

.card {
  margin-left: 20%;
}
</style>