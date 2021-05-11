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
    userId: "",
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
      this.userId = hrefPaths[5];

      localStorage.setItem("userId", this.userId);
      alert(this.confirmationToken+"/"+this.userId)
      
      this.$http
        .post("http://localhost:8080/confirm_registration/", {
          confirmation_token: this.confirmationToken,
          user_id: this.userId,
        })
        .then((res) => {
          console.log(res);
          this.message = "You have successfully verified your account! You can log in on system!"
        })
        .catch((err) => {
          console.log(err);
          this.message = "Your token is invalid or expiried! Please, contact system admin!"
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