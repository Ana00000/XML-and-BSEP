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
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
