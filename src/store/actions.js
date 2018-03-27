import * as type from './mutations'
import * as axios from 'axios'


export default {
  setAccessToken({ commit }, token) {
    commit(type.SET_ACCESS_TOKEN, token)
  },
  removeAccessToken({ commit }) {
    localStorage.removeItem('ls-access-token')
    commit(type.REMOVE_ACCESS_TOKEN)
  },
  getUser({ commit }, token) {
    return new Promise((resolve, reject) => {
      axios.create({
        baseURL: "https://api.github.com",
        headers: {'Authorization': 'token ' + token},
        responseType: 'json'
      }).get('/user').then(response => {
        resolve(response.data);
      }, err => {
        console.log(err);
        reject(err);
      })
    })
  }
}