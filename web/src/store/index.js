import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    tables: [],
    settings: {
      mainPath: "com.github.crazyjay97",
      pkg: "com.github.crazyjay97.modules",
      author: "crazyjay97",
      email: "zhanweijie111@gmail.com",
      isRemovePrefix: true,
      moduleName: "",
      autoSettingModuleName: true
    }
  },
  mutations: {
    updateTables(state, tables) {
      state.tables = tables
    },
    updateSettings(state, settings) {
      state.settings = settings
    },
  }
})

