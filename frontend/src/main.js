import Vue from 'vue';
import App from './App.vue';
import router from './router/router';
import Buefy from 'buefy';
import Axios from 'axios';
import VueAxios from 'vue-axios';

import { library } from '@fortawesome/fontawesome-svg-core';
import { fab } from '@fortawesome/free-brands-svg-icons';
import { fas } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';


library.add(fas, fab);
Vue.component('vue-fontawesome', FontAwesomeIcon);
Vue.config.productionTip = false;

Vue.use(Buefy, {
  defaultIconComponent: 'vue-fontawesome',
  defaultIconPack: 'fas',
  customIconPacks: {
    fas: {
      sizes: {
        default: 'lg',
        'is-small': '1x',
        'is-medium': '2x',
        'is-large': '3x',
      },
      iconPrefix: '',
    },
  },
});
Vue.use(VueAxios, Axios);

Vue.prototype.$axios = Axios;
Vue.axios.defaults.baseURL = `http://localhost:8080`;

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');

