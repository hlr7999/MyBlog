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
    },
    host: "http://localhost:2333/",
    classList: []
  },

  mutations: {
    changeClassList(state, cl) {
      state.classList = cl
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
