<template>
  <div>
    <div class="spacing" />
    <v-card width="600" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1">Verification request creation</h1>
      </v-card-title>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
            label="First Name"
            v-model="firstName"
            prepend-icon="mdi-address-circle"
          />
          <v-text-field
            label="Last name"
            v-model="lastName"
            prepend-icon="mdi-address-circle"
          />
          <v-select
            class="typeCombo"
            v-model="selectedType"
            hint="Choose category"
            :items="types"
            item-text="state"
            :label="label1"
            return-object
            single-line
          />
          <iframe
            name="dummyframe"
            id="dummyframe"
            style="display: none"
          ></iframe>
          <form
            action="https://localhost:8080/api/requests/uploadOfficialDocument/"
            enctype="multipart/form-data"
            method="post"
            target="dummyframe"
            class="uploadButton"
          >
            <template>
              <input id="pic" type="file" accept="image/*" name="myPostFile" />
            </template>
            <input
              type="submit"
              value=" <- Upload file"
              v-on:click="ValidteType"
            />
          </form>
        </v-form>
      </v-card-text>
        <v-btn color="info mb-5" v-on:click="createRequest" v-if="isVisibleFinishButton">
          Create
        </v-btn>
      </v-card-actions>
    </v-card>
    <div class="spacing" />
  </div>
</template>

<script>
export default {
  name: "CreateVerificationRequest",
  data: () => ({

    types: ["INFLUENCER", "SPORTS", "NEW_MEDIA", "BUSINESS", "BRAND", "ORGANIZATION"],
    selectedType: "INFLUENCER",
    label1: "Type",
    postId: null,
    extension: "",
    allUserTags: [],
    userId: null,
    token: null,
    path: null,
    isVisibleFinishButton: false
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.userId = localStorage.getItem("userId");
      this.token = localStorage.getItem("token");
      this.$http
        .get(
          "https://localhost:8080/api/user/check_if_authentificated/",{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log("User is authentificated!");
           console.log(resp.data);
        })
        .catch((er) => {
           console.log(er);
          window.location.href = "https://localhost:8081/unauthorizedPage";
        });

      this.$http
        .get(
          "https://localhost:8080/api/user/auth/check-create-verification-request-permission/",{
            headers: {
              Authorization: "Bearer " + this.token,
            },
          }
        )
        .then((resp) => {
          console.log("User is authorized!");
          console.log(resp.data);
        })
        .catch((er) => {
          window.location.href = "https://localhost:8081/forbiddenPage";
          console.log(er);
        });

    },
    GetExtension(pathFile) {
      console.log(pathFile);
      let out = pathFile.split("\\");
      let fileName = out[out.length - 1];
      let dotSplit = fileName.split(".");
      this.extension = dotSplit[dotSplit.length - 1];
      console.log(this.extension);
    },
    createRequest() {
      this.$http
        .post("https://localhost:8080/api/requests/verificationRequest", {
          first_name: this.firstName,
          last_name: this.lastName,  
          official_document_path: this.path,
          registered_user_category: this.selectedType,
          user_id: this.userId,
        },{
            headers: {
              Authorization: "Bearer " + this.token,
            },
        })
        .then((response) => {
          console.log("Successful creation!");
          console.log(response.data);
          window.location.href = "https://localhost:8081/";
        })
        .catch((er) => {
          console.log(er.response.data);
        });
    },,
    ValidteType() {
      let pathFile = "";
      if (this.selectedType === "PICTURE") {
        pathFile = document.getElementById("pic").value;
        this.GetExtension(pathFile);
        console.log(this.extension);
        if (
          this.extension === "PNG" ||
          this.extension === "png" ||
          this.extension === "JPG" ||
          this.extension === "jpg" ||
          this.extension === "jpeg" ||
          this.extension === "JPEG"
        ) {
          this.isVisibleFinishButton = true;
        } else {
          this.isVisibleFinishButton = false;
          alert(
            "Please, choose a picture in a correct format e.g. png, jpg or jpeg."
          );
        }
      } 
    }
};
</script>

<style scoped>
.spacing {
  height: 100px;
}

.uploadButton {
  margin-left: 6%;
}

.typeCombo {
  width: 94%;
  margin-left: 6%;
}
</style>