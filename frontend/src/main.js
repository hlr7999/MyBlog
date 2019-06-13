import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import ElementUI, { Message } from 'element-ui'
import mavonEditor from 'mavon-editor'
import 'element-ui/lib/theme-chalk/index.css'
import 'mavon-editor/dist/css/index.css'
import 'font-awesome/css/font-awesome.min.css'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(mavonEditor)

var currentUser = localStorage.getItem("currentUser")

if (currentUser) {
  currentUser = JSON.parse(currentUser)
  store.commit("login", currentUser)
}

import { GetFirstClass } from "./api/api"

GetFirstClass()
.then(res => {
  store.commit("changeClassList", res.data)
})
.catch(() => {
  Message.error("获取分类失败")
})

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  store
})
