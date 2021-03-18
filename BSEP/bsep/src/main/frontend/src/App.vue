<template>
  <div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/certificates">Certificates</router-link> |
    <router-link to="/newCertificate">New certificate</router-link> |
    <router-link to="/logIn">Log in</router-link> |
    <router-link to="/registration">Registration</router-link>
  </div>
  <router-view />
</template>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
