<template>
  <div>
    <v-container grid-list-lg>
      <div class="spacingOne" />
      <div class="title">
        <h1>Public Posts</h1>
      </div>
      <div class="spacingTwo" />
      <v-layout row>
        <v-flex lg4 v-for="item in posts" :key="item.id" class="space-bottom">
          <v-card class="mx-auto" max-width="475">
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle>{{
                  item.description
                }}</v-list-item-subtitle>
                <v-list-item-title>{{ item.tags }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-hover>
              <template v-slot:default="{ hover }">
                <v-list-item-content>
                  <img :src="require(`../../../${item.path}`)" alt class="icon" />

                  <v-fade-transition>
                    <v-overlay v-if="hover" absolute color="#036358">
                      <v-btn>Open</v-btn>
                    </v-overlay>
                  </v-fade-transition>
                </v-list-item-content>
              </template>
            </v-hover>
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-subtitle
                  v-text="
                    item.country +
                    ' ' +
                    item.city +
                    ' ' +
                    item.street_name +
                    ' ' +
                    item.street_number
                  "
                />
                <v-list-item-subtitle>{{
                  item.creation_date
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
  name: "PublicPostsAll",
  data: () => ({
    posts: [],
    allTags: "",
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get("http://localhost:8084/find_all_public_posts_not_reg/")
        .then((response) => {
          this.posts = response.data;
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
  height: 100px;
}
</style>