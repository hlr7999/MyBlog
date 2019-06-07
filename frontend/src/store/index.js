import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    baseUrl: "http://localhost:2333/",
    loading: false,
    hasLogin: false,
    isAdmin: false,
    userInfo: {
      userToken: "",
      userId: "",
      userAvatar: "/static/img/head.jpg"
    }
  },

  mutations: {
    changeLoading(state, l) {
      state.loading = l
    },
    changeHasLogin(state, h) {
      state.hasLogin = h
    },
    changeIsAdmin(state, i) {
      state.isAdmin = i
    },
    changeUserInfo(state, info) {
      state.userInfo.userAvatar = info.avatar
      state.userInfo.userId = info.id
      state.userInfo.userToken = info.token
    }
  }
})
