import { Commit, createStore } from 'vuex'
import { number } from 'yup'

export default createStore({
  state: {
    auth: false,
    user_id: 0,
    user: [{
      user_id: "",
      nick_name: "",
      email: "",
    }]
  },
  mutations: {
    setAuth: (state: { auth: boolean }, auth: boolean) => state.auth = auth,
    setUserId: (state: { user_id: number }, user_id: number) => state.user_id = user_id,
  },
  actions: {
    // ログイン済みかどうかの状態を管理する
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('setAuth', auth),
    //ユーザーIDを取得する
    setUserId: ({ commit }: { commit: Commit }, user_id: number) => commit('setUserId', user_id)
  },
  modules: {
  }
})
