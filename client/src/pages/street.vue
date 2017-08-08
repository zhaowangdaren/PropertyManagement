
<template>
<!--  首页
 居民物业纠纷处理
 物业信息管理
 物业考核查询
 申请法院调解请求
 操作记录管理 -->
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <menu-street :class='s.menu' :menus='menus' :NEXT='nextPath'/>
      <div :class='s.body'>
        <div :class='s.street'>Street Name: {{streetName}}</div>
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import MenuStreet from '@/components/Menu'
import ActionBar from '@/components/ActionBar'
import fetchpm from '@/fetchpm'
export default {
  components: {MenuStreet, ActionBar},
  beforeRouteUpdate(to, from, next) {
    console.info('beforeRouteUpdate', to)
    this.nextPath = to.path
    next()
  },
  data () {
    return {
      nextPath: '',
      streetName: '',
      menus: [
        {
          icon: 'icon-home',
          path:'/street',
          text:'首页'//首页
        },
        {
          path:'/street/handle',
          icon: 'icon-chulizhong',
          text:'complaintsHandler' //居民物业纠纷处理
        },
        {
          path:'/street/pms',
          icon: 'icon-wuyeguanli',
          text:'PM infos'//物业信息管理
        },
        {
          path:'/street/pmkpis',
          icon: 'icon-kaohe',
          text:'PM KPI'//物业考核查询
        },
        {
          path:'/street/court',
          icon: 'icon-court',
          text:'Court'//申请法院调解请求
        },
        {
          path:'/street/operationRecs',
          icon: 'icon-records',
          text:'Operation recs'//操作记录管理
        }
      ]
    }
  },
  mounted () {
    this.fetchUserStreet()
  },
  methods: {
    fetchUserStreet() {
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      fetchpm(this, true, '/pm/street/ids', {
        method: 'POST',
        body:{values: [user.StreetID]}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchUserStreet', body)
        if (body.error === 0 && body.data != null && body.data.length > 0 ) {
          this.streetName = body.data[0].Name
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
  .wrap{
    display: flex;
    flex-direction: column;
    .content{
      display: flex;
      // height: 100%;
      position: absolute;
      top: 0;
      bottom: 0;
      left: 0;
      width: 100%;
      // z-index: 2;
      margin-top: 100px;
      .menu{
        height: 100%;
      }
      .body{
        overflow-y: scroll;
        flex: 1;
        .street{
          background-color: #ddd;
          font-size: 25px;
          padding: 5px;
          margin-bottom: 20px;
        }
      }
    }
  }
</style>
