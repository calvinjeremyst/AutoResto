import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

//Component
import Login from '../views/Login.vue';

const routes = [
    {
        path: '/login',
        name: 'login',
        component: Login,
    },
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    linkActiveClass: 'is-active',
    routes,
    scrollBehavior() {
      window.scrollTo(0, 0);
    },
  });
  
  export default router;