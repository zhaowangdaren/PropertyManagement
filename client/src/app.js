import Vue from 'vue'
import router from './router'
import App from './App'
import 'whatwg-fetch'

new Vue ({
  el: '#app',
  router,
  render: h => h(App)
})