<template>
  <div>
    <v-card width="400" class="mx-auto mt-5" color="white">
      <v-card-title class="justify-center">
        <h1 class="display-1 ">Registration</h1>
      </v-card-title>
      <v-card-text>
        <v-form class="mx-auto ml-5 mr-5">
          <v-text-field
              label="Username/Email"
              v-model="userEmail"
              prepend-icon="mdi-account-circle"/>
          <v-text-field
              :type="showPassword ? 'text' : 'password'"
              label="Password"
              v-model="password"
              prepend-icon="mdi-lock"
              :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              @click:append="showPassword = !showPassword"/>
          <v-text-field
              label="First name"
              v-model="firstName"
              prepend-icon="mdi-name-circle"/>
          <v-text-field
              label="Last name"
              v-model="lastName"
              prepend-icon="mdi-address-circle"/>
          <v-text-field
              label="Phone number"
              v-model="phoneNumber"
              prepend-icon="mdi-address-circle"/> 

          <v-combobox :items="typesOfUsers" :item-text="text" v-model="selectedTypeOfUser" 
          :label="label1" hint="Choose type of user."/>
        </v-form>
      </v-card-text>
      <v-card-actions class="justify-center mb-5">
        <v-btn color="info mb-5" v-on:click="register">
          Register
        </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data: () => ({
    showPassword: false,
    userEmail: '',
    password: '',
    phoneNumber:'',
    firstName: '',
    lastName: '',
    users: [],
    selectedTypeOfUser: "ADMIN",
    label1: 'Type of user'
  }),
  methods: {
    register() {
      if(!this.ValidateEmail()){
        return;
      }/*
      this.$http.get('https://localhost:8080/users/findAll')
      .then(response => {
        console.log(response.data);

        //window.location.href = 'http://localhost:8081/login';
      })
      .catch(er => {
        console.log('Error while registering in');
        console.log(er.response.data);
      })*/
      
      this.$http.post('/api/users/register',
        {         
          userEmail : this.userEmail,
          lastName : this.lastName,
          password : this.password,
          firstName : this.firstName,
          phoneNumber : this.phoneNumber,
          typeOfUser : this.selectedTypeOfUser
      })
      .then(response => {
        
        console.log(response.data.firstName);
        
        console.log(this.phoneNumber);
        //window.location.href = 'http://localhost:8081/login';
      })
      .catch(er => {
        console.log('Error while registering in');
        console.log(er.response.data);
      })
    },
    ValidateEmail() {
      if (/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(this.userEmail))
      {
        return (true)
      }
        alert("You have entered an invalid email address!")
        return (false)
      }
  }
};
</script>