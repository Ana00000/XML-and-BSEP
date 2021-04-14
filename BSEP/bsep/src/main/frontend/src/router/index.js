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
    path: '/confirmRegistration/:id',
    name: 'ConfirmRegistration',
    component: () => import('../views/ConfirmRegistration.vue')
  },
  {
    path: '/certificates',
    name: 'Certificates',
    component: () => import('../views/Certificates.vue')
  },
  {
    path: '/newRootCertificate',
    name: 'NewRootCertificate',
    component: () => import('../views/NewRootCertificate.vue')
  },
  {
    path: '/invalidCertificates',
    name: 'InvalidCertificates',
    component: () => import('../views/InvalidCertificates.vue')
  },
  {
    path: '/selectedCertificate',
    name: 'SelectedCertificate',
    component: () => import('../views/SelectedCertificate.vue')
  },
  {
    path: '/createOtherCertificates',
    name: 'CreateOtherCertificates',
    component: () => import('../views/CreateOtherCertificates.vue')
  },
  {
    path: '/checkCertificateValidity',
    name: '/CheckCertificateValidity',
    component: () => import('../views/CheckCertificateValidity.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
