import {createStore } from 'vuex'

export default createStore({
  state: {
    auth: false
  },
  mutations: {
    setAuth: (state, auth) => state.auth = auth
  },
  actions: {
     // ログイン済みかどうかの状態を管理する
    setAuth: (commit, auth) => commit('setAuth', auth)
  },
  modules: {
  }
})
