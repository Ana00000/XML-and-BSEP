<template>
  <v-app>
    <NavigationBar/>
    <v-content>
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>
import NavigationBar from './components/NavigationBar.vue'
export default {
  name: 'App',
  el: '#app',
  components: {
    NavigationBar
  }
};
</script>