import Vue from 'vue'
import Vuex from 'vuex'

const state = {
  loading: false
}

Vue.use(Vuex)

export default new Vuex.Store({
  state
})
