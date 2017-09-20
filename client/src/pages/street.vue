
<template>
<!--  首页
 居民物业纠纷处理
 物业信息管理
 物业考核查询
 申请法院调解请求
 操作记录管理 -->
  <div :class='s.wrap'>
    <div :class='s.left'>
      <left-menu :menus='menus' :NEXT='nextPath'/>
    </div>
    <div :class='s.right'>
      <action-bar></action-bar>
      <div :class='s.content'>
        <router-view :class='s.body'></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import LeftMenu from '@/components/Menu'
import ActionBar from '@/components/ActionBar'
import fetchpm from '@/fetchpm'
export default {
  components: {LeftMenu, ActionBar},
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
          text:'居民物业纠纷处理' //居民物业纠纷处理
        },
        {
          path:'/street/pms',
          icon: 'icon-wuyeguanli',
          text:'物业信息管理'//物业信息管理
        },
        {
          path:'/street/pmkpis',
          icon: 'icon-kaohechengji',
          text:'物业考核查询'//物业考核查询
        },
        {
          path:'/street/operationRecs',
          icon: 'icon-records',
          text:'操作记录'//操作记录管理
        },
        {
          icon: 'icon-gonggao',
          text:'政府公告',
          path: '/street/announcement'
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
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  .right{
    overflow: auto;
    flex: 1;
    width: 100%;
  }
  .content{
    margin: 10px;
  }
}
</style>
