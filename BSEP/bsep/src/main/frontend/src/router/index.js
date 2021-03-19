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
    path: '/certificates',
    name: 'Certificates',
    component: () => import('../views/Certificates.vue')
  },
  {
    path: '/newCertificate',
    name: 'NewCertificate',
    component: () => import('../views/NewCertificate.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
