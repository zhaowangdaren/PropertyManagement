import Vue from 'vue'
import router from './router'
import App from './App'
import 'whatwg-fetch'
import { Upload, Dialog, Button } from 'element-ui'
 
Vue.use(Upload)
Vue.use(Dialog)
Vue.use(Button)


new Vue ({
  el: '#app',
  router,
  render: h => h(App)
})