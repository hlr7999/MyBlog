import Vue from 'vue'
import Vuex from 'vuex'

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
    }
  },

  mutations: {
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
    }
  }
})
