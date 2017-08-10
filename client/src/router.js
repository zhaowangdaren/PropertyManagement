import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  // mode: 'history',
  routes: [
    {
      path: '/testapi',
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
          component: require('@/components/table/EventHandles')
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
          component: require('@/components/table/Courts')
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
    },
    {
      path: '/wx/startComplaint',
      component: require('@/pages/wx/startComplaint')
    },
    {
      path: '/wx/complaint/selectStreet',
      component: require('@/pages/wx/complaint/selectStreet')
    },
    {
      path: '/wx/complaint/selectCommunity',
      component: require('@/pages/wx/complaint/selectCommunity')
    },
    {
      path: '/wx/complaint/selectXQ',
      component: require('@/pages/wx/complaint/selectXQ')
    },
    {
      path: '/wx/complaint/addEvent',
      component: require('@/pages/wx/complaint/addEvent')
    },
    {
      path: '/wx/complaint/uploadImage',
      component: require('@/pages/wx/complaint/uploadImage')
    },
    {
      path: '/wx/complaint/finish',
      component: require('@/pages/wx/complaint/finish')
    },
    {
      path: '/wx/viewProgress',
      component: require('@/pages/wx/viewProgress')
    },
    {
      path: '/wx/detailsProgress',
      name: 'detailsProgress',
      component: require('@/pages/wx/progress/detailsProgress')
    }

  ]
})