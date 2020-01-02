import FlagIcon from 'vue-flag-icon';
import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './registerServiceWorker';
import vuetify from './plugins/vuetify';


Vue.config.productionTip = false;
Vue.use(FlagIcon);

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App),
}).$mount('#app');
