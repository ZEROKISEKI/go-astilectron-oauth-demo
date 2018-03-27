import Vue from 'vue'
import Vuex from 'vuex'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuex)
Vue.use(Vuetify)

import VuexStore from './store'

const store = new Vuex.Store(VuexStore)

import App from './App.vue'
import router from './router'

/*new Vue({
  components: { App },
  router,
  template: '<App/>'
}).$mount('#app')*/
new Vue({
  el: '#app',
  render: h => h(App),
  router,
  store,
});