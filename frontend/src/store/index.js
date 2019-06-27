import Vue from 'vue'
import Vuex from 'vuex'
import { GetFirstClass } from "../api/api"
import {Message} from 'element-ui'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    loading: false,
    hasLogin: false,
    isAdmin: false,
    userInfo: {
      userToken: "",
      userId: "",
      userAvatar: ""
    },
    host: "http://47.100.225.59:2333/",
    classList: []
  },

  mutations: {
    getClassList(state) {
      GetFirstClass()
      .then(res => {
        state.classList = res.data
      })
      .catch(() => {
        Message.error("获取分类失败")
      })
    },
    changeAvatar(state, data) {
      state.userInfo.userAvatar = data
    },
    login(state, data) {
      state.hasLogin = true
      state.isAdmin = (data.role === "admin")
      state.userInfo = {
        userToken: data.token,
        userId: data.id,
        userAvatar: data.avatar
      }
    },
    logout(state) {
      state.hasLogin = false
      state.isAdmin = false
      state.userInfo = {
        userToken: "",
        userId: "",
        userAvatar: ""
      }
    },
    changeLoading(state, l) {
      state.loading = l
    }
  }
})
