import FlagIcon from 'vue-flag-icon';
import Vue from 'vue';
import * as Sentry from '@sentry/browser';
import * as Integrations from '@sentry/integrations';
import App from './App.vue';
import router from './router';
import store from './store';
import './registerServiceWorker';
import vuetify from './plugins/vuetify';


Vue.config.productionTip = false;
Vue.use(FlagIcon);
Sentry.init({
  dsn: 'https://26d47886ac9347ad86538e2ef5cbaa87@sentry.io/1872654',
  integrations: [new Integrations.Vue({Vue, attachProps: true, logErrors: true})],
});

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App),
}).$mount('#app');
