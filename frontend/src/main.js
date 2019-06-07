import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import mavonEditor from 'mavon-editor'
import 'element-ui/lib/theme-chalk/index.css'
import 'mavon-editor/dist/css/index.css'
import 'font-awesome/css/font-awesome.min.css'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(mavonEditor)

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  store
})

var currentUser = localStorage.getItem("currentUser")

if (currentUser) {
  currentUser = JSON.parse(currentUser)
  store.commit("changeHasLogin", true)
  store.commit("changeIsAdmin", currentUser.role === "admin")
  store.commit("changeUserInfo", currentUser)
}
