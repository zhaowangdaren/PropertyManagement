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
          path:'',
          component: require('@/pages/admin/basicInfo')
        },
        {
          path: 'users',
          component: require('@/pages/admin/users')
        },{
          path: 'init',
          component: require('@/pages/admin/init')
        },{
          path: 'wx',
          component: require('@/pages/admin/wx')
        },{
          path: 'complaints',
          component: require('@/pages/admin/complaints')
        },{
          path: 'pm',
          component: require('@/pages/admin/pm')
        },{
          path: 'build',
          component: require('@/pages/admin/build')
        },
      ]
    },
    {
      path: '/street',
      component: require('@/pages/street'),
      children: [
        {
          path: '',
          component: require('@/pages/street/index')
        },
        {
          path: 'detail/:index',
          component: require('@/components/DetailEvent')
        },
        {
          path: 'handle',
          component: require('@/components/InfosEventHandles')
        },
        {
          path: 'pms',
          component: require('@/components/InfosStreetPM')
        },
        {
          path: 'pms/kpi',
          component: require('@/components/InfosStreetPMsKPI')
        },
        {
          path: 'court',
          component: require('@/components/InfosCourt')
        },
        {
          path: 'operationRecs',
          component: require('@/components/InfosOperationRecs')
        }
      ]
    },
    {
      path: '/gov',
      component: require('@/pages/gov')
    }
  ]
})