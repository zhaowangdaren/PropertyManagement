import Vue from 'vue'
import router from './router'
// import axios from 'axios'
// import VueAxios from 'vue-axios'

import Promise from 'promise-polyfill'
if (!window.Promise) window.Promise = Promise

import App from './App'
import 'whatwg-fetch'
import { Upload, Dialog, Button, Input, Select,Option, DatePicker } from 'element-ui'
import '@/res/common.css'
import fetchpm from '@/fetchpm'
// Vue.use(VueAxios, axios)

Vue.use(Upload)
Vue.use(Dialog)
Vue.use(Button)
Vue.use(Input)
Vue.use(Select)
Vue.use(Option)
Vue.use(DatePicker)

new Vue ({
  el: '#app',
  router,
  render: h => h(App)
})
