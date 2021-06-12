<template>
  <v-card style="margin-top: 2%" width="25%" class="mx-auto">
    <v-img height="600" width="550" :src="require('../assets/403.png')">
    </v-img>
    <v-card-title class="justify-center">
      <h1 class="display-1 mt-5">403 - You shall not pass!</h1>
      <h3>You are forbidden to see this page</h3>
    </v-card-title>
    <div></div>
    <v-card-actions>
      <v-btn color="primary" class="mx-auto mb-5" large v-on:click="back">
        Back Home
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "ForbiddenPage",
  data: () => ({}),
  methods: {
    back() {
      window.location.href = "https://localhost:8081/";
    },
  },
};
</script>