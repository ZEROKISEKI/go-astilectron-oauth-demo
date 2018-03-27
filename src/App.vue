<template>
    <div id="app" class="text-xs-center d-flex align-center">
        <v-app>
            <v-content>
                <v-btn v-if="!access_token" color="success" @click="login">OAuth(With Github Demo)</v-btn>
                <v-avatar class="grey lighten-4" v-if="avatar">
                    <img :src="avatar" alt="avatar">
                </v-avatar>
                <p v-if="name">{{ name }}</p>
                <v-btn v-if="access_token" color="error" @click="logout">Logout</v-btn>
            </v-content>
        </v-app>
    </div>
</template>

<script>
import { mapState, mapActions } from "vuex"

export default {
  name: 'app',
  data() {
    return {
      avatar: '',
      name: ''
    }
  },
  computed: {
    ...mapState({
      access_token: ({ common }) => common.access_token
    })
  },
  methods: {
    ...mapActions(['setAccessToken', 'getUser', 'removeAccessToken']),
    login() {
      const that = this
      astilectron.onMessage(function (message) {
        if (message.name === "window.event.will.navigate" ||
          message.name === "window.event.did.get.redirect.request") {
          that.setAccessToken(message.payload.access_token)
          console.log(that.access_token)
          that.getUser(that.access_token).then(data => {
            that.avatar = data.avatar_url
            that.name = data.login
          })
        }
        return { "name": "world" }
      })

      astilectron.sendMessage({ "name": "login" }, function(message) {})
    },
    logout() {
      this.removeAccessToken()
      this.avatar = ''
      this.name = ''
    }
  },
  created() {
    const token = localStorage.getItem('ls-access-token') || ""
    if(token !== "") {
      this.getUser(token).then(data => {
        this.avatar = data.avatar_url
        this.name = data.login
      })
    }
  }
}
</script>

<style scoped>

</style>