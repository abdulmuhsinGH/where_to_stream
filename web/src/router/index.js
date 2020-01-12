import Vue from 'vue';
import VueRouter from 'vue-router';
import PrimeVideoAndNetflix from '../views/PrimeVideoAndNetflix.vue';
import NetflixInfo from '../views/NetflixInfo.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'primeVideoAndNetflix',
    component: PrimeVideoAndNetflix,
  },
  {
    path: '/netflixinfo',
    name: 'netflixInfo',
    component: NetflixInfo,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
