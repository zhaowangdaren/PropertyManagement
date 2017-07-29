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
      path: '/login',
      component: require('@/pages/login')
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
          path: 'eventDetail/:index',
          name: 'eventDetail',
          component: require('@/components/DetailEvent')
        },
        {
          path: 'handle',
          component: require('@/components/InfosEventHandles')
        },
        {
          path: 'pms',
          component: require('@/pages/street/pms')
        },
        {
          path: 'pmkpis',
          component: require('@/pages/street/pmsKPI')
        },
        {
          path: 'court',
          component: require('@/components/InfosCourt')
        },
        {
          path: 'operationRecs',
          component: require('@/pages/street/recs')
        }
      ]
    },
    {
      path: '/gov',
      component: require('@/pages/gov'),
      children: [
        {
          path:'',
          component: require('@/pages/gov/index')
        },
        {
          path:'eventHandle',
          component: require('@/pages/gov/eventHandle')
        },{
          path:'pms',
          component: require('@/pages/gov/pm')
        },{
          path:'build',
          component: require('@/pages/gov/build')
        },{
          path:'pmkpis',
          component: require('@/pages/gov/pmkpi')
        },{
          path:'court',
          component: require('@/pages/gov/court')
        },{
          path:'announcement',
          component: require('@/pages/gov/announcement')
        },
      ]
    }
  ]
})