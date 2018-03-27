import {
  SET_ACCESS_TOKEN,
  REMOVE_ACCESS_TOKEN,
} from "../mutations"

export default {
  state: {
    access_token: localStorage.getItem('ls-access-token') || null,
  },
  mutations: {
    [SET_ACCESS_TOKEN](state, token) {
      state.access_token = token
      localStorage.setItem('ls-access-token', token)
    },
    [REMOVE_ACCESS_TOKEN](state) {
      state.access_token = null
    }
  }
}