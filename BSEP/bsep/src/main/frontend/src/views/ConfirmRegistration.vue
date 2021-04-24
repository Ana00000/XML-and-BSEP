<template>
  <div >
    <h1 class="center">{{ message }}</h1>
  </div>
</template>

<script>
export default {
  name: "ConfirmRegistration",
  data: () => ({
    confirmationToken:"",
    message: ""
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      var hrefPath = window.location.href;
      var hrefPaths = [];
      hrefPaths=  hrefPath.split('/');
      this.confirmationToken = hrefPaths[4];
      this.$http
        .put("https://localhost:8080/users/confirm_account/" + this.confirmationToken, {})
        .then((res) => {
          console.log(res);
          this.message = "You are successfully verified your account! You can log in on system!"
        })
        .catch((err) => {
          console.log(err);
          this.message = "Your token are invalid or expiried! Please, contact system admin!"
        });
    },
    redirectToLogIn() {
      window.location.href = "https://localhost:8081/logIn";
    },
  },
};
</script>

<style scoped>
.helloMessage {
  font-weight: bolder;
  font-size: 20px;
  height: 50px;
}

.center {
  margin-top: 10%;
  padding: 10px;
  text-align: center;
}

#certificateCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>