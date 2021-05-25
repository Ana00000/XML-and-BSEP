<template>
  <v-container fluid fill-height class="backgroundDiv" style="max-height: 100vh;">
    <v-layout justify-center align-center column pa-5 style="background-color: grey; opacity:0.30">
    </v-layout>
  </v-container>
</template>

<script>
export default {
    name: 'HomeBackground'
}
</script>

<style scoped>
.backgroundDiv {
  background: url('../assets/background.jpg');
  background-size: cover;
  width: 100%;
  height: 100%;
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-position: center;
}
</style>