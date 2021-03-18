import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/certificates",
    name: "Certificates",
    component: () =>
      import("../views/Certificates.vue"),
  },
  {
    path: "/newCertificate",
    name: "NewCertificate",
    component: () =>
      import("../views/NewCertificate.vue"),
  },
  {
    path: "/logIn",
    name: "LogIn",
    component: () =>
      import("../views/LogIn.vue"),
  },
  {
    path: "/registration",
    name: "Registration",
    component: () =>
      import("../views/Registration.vue"),
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
