import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    tables: []
  },
  mutations: {
    updateTables(state, tables) {
      state.tables = tables
    }
  }
})

