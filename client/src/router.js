import Vue from 'vue'
import Router from 'vue-router'


const Login = require('@/pages/login')
const Regist = require('@/pages/regist')

const Admin = resolve => require(['@/pages/admin'], resolve)

const AdminHome = resolve => require(['@/pages/admin/home'], resolve)
const AdminBasicInfo = resolve => require(['@/pages/admin/basicInfo'], resolve)
const AdminUsers = resolve => require(['@/pages/admin/users'], resolve)
const AdminInit = resolve => require(['@/pages/admin/init'], resolve)
const AdminWX = resolve => require(['@/pages/admin/wx'], resolve)
const AdminComplaint = resolve => require(['@/pages/admin/complaints'], resolve)
const AdminPM = resolve => require(['@/pages/admin/pm'], resolve)
const AdminBuild = resolve => require(['@/pages/admin/build'], resolve)
const Street = resolve => require(['@/pages/street'], resolve)
const StreetIndex = resolve => require(['@/pages/street/index'], resolve)
const StreetDetailEvent = resolve => require(['@/components/DetailEvent'], resolve)
const StreetEventHandle = resolve => require(['@/components/table/EventHandles'], resolve)
const StreetPM= resolve => require(['@/pages/street/pms'], resolve)
const StreetPMKPI= resolve => require(['@/pages/street/pmsKPI'], resolve)
const StreetCourts= resolve => require(['@/components/table/Courts'], resolve)
const StreetRecs= resolve => require(['@/pages/street/recs'], resolve)
const Gov = resolve => require(['@/pages/gov'], resolve)
const GovIndex = resolve => require(['@/pages/gov/index'], resolve)
const GovEventHandle = resolve => require(['@/pages/gov/eventHandle'], resolve)
const GovPM = resolve => require(['@/pages/gov/pm'], resolve)
const GovBuild = resolve => require(['@/pages/gov/build'], resolve)
const GovPMKPI = resolve => require(['@/pages/gov/pmkpi'], resolve)
const GovCourt = resolve => require(['@/pages/gov/court'], resolve)
const GovAnnouncement = resolve => require(['@/pages/gov/announcement'], resolve)
const GovDetailEvent = resolve => require(['@/components/DetailEvent'], resolve)

const WXStartComplaint = require('@/pages/wx/startComplaint')
const WXComplaintSelectStreet = resolve => require(['@/pages/wx/complaint/selectStreet'], resolve)
const WXComplaintSelectCommunity = resolve => require(['@/pages/wx/complaint/selectCommunity'], resolve)
const WXComplaintSelectXQ = resolve => require(['@/pages/wx/complaint/selectXQ'], resolve)
const WXComplaintAddEvent = resolve => require(['@/pages/wx/complaint/addEvent'], resolve)
const WXComplaintUploadImage = resolve => require(['@/pages/wx/complaint/uploadImage'], resolve)
const WXComplaintFinish = resolve => require(['@/pages/wx/complaint/finish'], resolve)
const WXViewProgress = require('@/pages/wx/complaint/viewProgress')
const WXComplaintMethod = resolve => require(['@/pages/wx/complaint/method'], resolve)
const WXDetailsProgress = resolve => require(['@/pages/wx/progress/detailsProgress'], resolve)
const WXCheckPM = require('@/pages/wx/checkPM')
const WXPMSelectCommunity = resolve => require(['@/pages/wx/pm/selectCommunity'], resolve)
const WXPMSelectXQ = resolve => require(['@/pages/wx/pm/selectXQ'], resolve)
const WXPMDetails = resolve => require(['@/pages/wx/pm/details'], resolve)
const WXPMBind = resolve => require(['@/pages/wx/pm/bind'], resolve)

const PMRule = resolve => require(['@/pages/wx/pmRule'], resolve)
const PMRules = resolve => require(['@/pages/wx/pm/rules'], resolve)
const PMMPRules = resolve => require(['@/pages/wx/pm/managementPractices'], resolve)
const PMKPIRules = resolve => require(['@/pages/wx/pm/kpiRules'], resolve)
const PMQA = resolve => require(['@/pages/wx/pm/qa'], resolve)
const PMEventHandle = resolve => require(['@/pages/wx/pm/eventHandle'], resolve)

const PMCharge = resolve => require(['@/pages/wx/pmCharge'], resolve)
const ParkMain = require('@/pages/wx/parkMain')
const ParkManager = resolve => require(['@/pages/wx/park/ParkManager'], resolve)
const ParkManagerBindXQ = resolve => require(['@/pages/wx/park/bindXQ'], resolve)
const ParkManagerSelectCommunity = resolve => require(['@/pages/wx/park/selectCommunity'], resolve)
const ParkManagerSelectXQ = resolve => require(['@/pages/wx/park/selectXQ'], resolve)
const ParkManagerBindWX = resolve => require(['@/pages/wx/park/bindWX'], resolve)

const CommunityService = require('@/pages/wx/communityService')
const CSConstructs = resolve => require(['@/pages/wx/community/constructs'], resolve)
const CSCFurniture = resolve => require(['@/pages/wx/community/construct/furniture'], resolve)
const CSCFix = resolve => require(['@/pages/wx/community/construct/fix'], resolve)
const CSCompany = resolve => require(['@/pages/wx/community/company'], resolve) // 公司详情

const ContactUs = resolve => require(['@/pages/wx/contactUs'], resolve)
Vue.use(Router)

export default new Router({
  // mode: 'history',
  routes: [
    {
      path: '/test',
      component: require('@/pages/test')
    },
    {
      path: '/',
      component: require('@/pages/index')
    },
    {
      path: '/login',
      component: Login
    },
    {
      path: '/regist',
      name: 'Regist',
      component: Regist
    },
    {
      path: '/admin',
      component: Admin,
      children: [
        {
          path: '',
          component: AdminHome
        },
        {
          path:'basic',
          component: AdminBasicInfo
        },
        {
          path: 'users',
          name:'adminUsers',
          component: AdminUsers
        },{
          path: 'init',
          component: AdminInit
        },{
          path: 'wx',
          component: AdminWX
        },{
          path: 'complaints',
          component: AdminComplaint
        },{
          path: 'pm',
          component: AdminPM
        },{
          path: 'build',
          component: AdminBuild
        },
      ]
    },
    {
      path: '/street',
      component: Street,
      children: [
        {
          path: '',
          component: StreetIndex
        },
        {
          path: 'eventDetail/:index',
          name: 'eventDetail',
          component: StreetDetailEvent
        },
        {
          path: 'handle',
          component: StreetEventHandle
        },
        {
          path: 'pms',
          component: StreetPM
        },
        {
          path: 'pmkpis',
          component: StreetPMKPI
        },
        {
          path: 'court',
          component: StreetCourts
        },
        {
          path: 'operationRecs',
          component: StreetRecs
        }
      ]
    },
    {
      path: '/gov',
      component: Gov,
      children: [
        {
          path:'',
          component: GovIndex
        },
        {
          path:'eventHandle',
          name: 'GovEventHandle',
          component: GovEventHandle
        },
        {
          path: 'eventDetail/:index',
          name: 'GovDetailEvent',
          component: GovDetailEvent
        },
        {
          path:'pms',
          component: GovPM
        },
        {
          path: 'complaints',
          component: AdminComplaint
        },{
          path:'build',
          component: GovBuild
        },{
          path:'pmkpis',
          component: GovPMKPI
        },{
          path:'court',
          component: GovCourt
        },{
          path:'announcement',
          component: GovAnnouncement
        },
      ]
    },
    {
      path: '/wx/startComplaint',
      component: WXStartComplaint
    },
    {
      path: '/wx/complaint/selectStreet',
      component: WXComplaintSelectStreet
    },
    {
      path: '/wx/complaint/selectCommunity',
      component: WXComplaintSelectCommunity
    },
    {
      path: '/wx/complaint/selectXQ',
      component: WXComplaintSelectXQ
    },
    {
      path: '/wx/complaint/addEvent',
      component: WXComplaintAddEvent
    },
    {
      path: '/wx/complaint/uploadImage',
      component: WXComplaintUploadImage
    },
    {
      path: '/wx/complaint/finish',
      component: WXComplaintFinish
    },
    {
      path: '/wx/viewProgress',
      component: WXViewProgress
    },
    {
      path: '/wx/complaint/method',
      component: WXComplaintMethod
    },
    {
      path: '/wx/detailsProgress',
      name: 'detailsProgress',
      component: WXDetailsProgress
    },
    {
      path: '/wx/checkPM',
      component: WXCheckPM
    },
    {
      path: '/wx/pm/selectCommunity',
      component: WXPMSelectCommunity
    },
    {
      path: '/wx/pm/selectXQ',
      component: WXPMSelectXQ
    },
    {
      path: '/wx/pm/details',
      component: WXPMDetails
    },
    {
      path: '/wx/pm/bind',
      component: WXPMBind
    },
    {
      path: '/wx/pm/rule',
      component: PMRule
    },
    {
      path: '/wx/pm/rules',
      component: PMRules
    },
    {
      path: '/wx/pm/rules/mp',
      component: PMMPRules
    },
    {
      path: '/wx/pm/rules/kpi',
      component: PMKPIRules
    },
    {
      path: '/wx/pm/rules/qa',
      component: PMQA
    },
    {
      path: '/wx/pm/event/handle',
      name: 'PMEvneHandle',
      component: PMEventHandle
    },
    {
      path: '/wx/pm/charge',
      component: PMCharge
    },
    {
      path: '/wx/park/main',
      component: ParkMain
    },
    {
      path: '/wx/park/manager',
      component: ParkManager
    },
    {
      path: '/wx/park/manager/bind/xq',
      component: ParkManagerBindXQ
    },
    {
      path: '/wx/park/manager/select/community',
      component: ParkManagerSelectCommunity
    },
    {
      path: '/wx/park/manager/select/xq',
      component: ParkManagerSelectXQ
    },
    {
      path: '/wx/park/manager/bind/wx',
      component: ParkManagerBindWX
    },
    {
      path: '/wx/community/service',
      component: CommunityService
    },
    {
      path: '/wx/community/constructs',
      component: CSConstructs
    },
    {
      path: '/wx/community/construct/furniture',
      name: 'CSCFurniture',
      component: CSCFurniture
    },
    {
      path: '/wx/community/construct/fix',
      name: 'CSCFix',
      component: CSCFix
    },
    {
      path: '/wx/community/company',
      name: 'CSCompany',
      component: CSCompany
    },
    {
      path: '/wx/contact',
      component: ContactUs
    }
  ]
})