import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/logIn',
    name: 'LogIn',
    component: () => import('../views/LogIn.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue')
  },
  {
    path: '/confirmRegistration/:confirmationToken/:userId',
    name: 'ConfirmRegistration',
    component: () => import('../views/ConfirmRegistration.vue')
  },
  {
    path: '/recoverPasswordEmail',
    name: 'RecoverPasswordEmail',
    component: () => import('../views/RecoverPasswordEmail.vue')
  },
  {
    path: '/changePasswordByToken/:token/:id',
    name: 'ChangePasswordByToken',
    component: () => import('../views/ChangePasswordByToken.vue')
  },
  {
    path: '/updateProfile',
    name: 'UpdateProfile',
    component: () => import('../views/UpdateProfile.vue')
  },
  {
    path: '/searchUsers',
    name: 'SearchUsers',
    component: () => import('../views/SearchUsers.vue')
  },
  {
    path: '/selectedUser',
    name: 'SelectedUser',
    component: () => import('../views/SelectedUser.vue')
  },
  {
    path: '/createPost',
    name: 'CreatePost',
    component: () => import('../views/CreatePost.vue')
  },
  {
    path: '/createStory',
    name: 'CreateStory',
    component: () => import('../views/CreateStory.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
