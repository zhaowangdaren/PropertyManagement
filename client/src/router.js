import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  // mode: 'history',
  routes: [
    {
      path: '/test',
      component: require('@/pages/testAPI')
    },
    {
      path: '/',
      component: require('@/pages/index')
    },
    {
      path: '/admin',
      component: require('@/pages/admin'),
      children: [
        {
          path: '/',
          component: require('@/pages/admin/basicInfo')
        }
      ]
    }
  ]
})