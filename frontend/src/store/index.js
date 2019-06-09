import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    currentPage: 1,
    lastUsePage: "",
    loading: false,
    hasLogin: false,
    isAdmin: false,
    userInfo: {
      userToken: "",
      userId: "",
      userAvatar: ""
    }
  },

  mutations: {
    changPage(state, p) {
      state.currentPage = p
    },
    changLastUsePage(state, lup) {
      state.currentPage = lup
    },
    login(state, data) {
      state.hasLogin = true
      state.isAdmin = (data.role === "admin")
      state.userInfo = {
        userToken: data.roken,
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
