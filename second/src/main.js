import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueCookies from 'vue-cookies'
import iView from 'view-design';//引入iview
import 'view-design/dist/styles/iview.css';    // 使用 CSS'


// Vue.use(iView);//vue使用iview
Vue.use(iView);
Vue.config.productionTip = false

Vue.use(VueCookies)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
