<template>
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
export default {
  components: {LeftMenu, ActionBar},
  beforeRouteUpdate(to, from, next) {
    console.info('beforeRouteUpdate', to)
    this.nextPath = to.path
    next()
  },
  data () {
    return {
      realName: '',
      nextPath: '',
      menus: [
        {
          path: '/gov',
          icon: 'icon-home',
          text:'主页'
        },
        {
          icon: 'icon-tousu',
          text:'居民物业纠纷处理',
          path:'/gov/eventHandle'
        },
        {
          icon: 'icon-tousu',
          text:'投诉数据分析',//微信用户绑定管理
          path:'/gov/complaints'
        },
        {
          icon: 'icon-wuyeguanli',
          text:'查看物业信息',
          path:'/gov/pms'
        },
        {
          icon: 'icon-kaohechengji',
          text:'物业考核查询',
          path:'/gov/pmkpis'
        },
        {
          icon: 'icon-build',
          text:'查看建筑信息',
          path:'/gov/build'
        },
        {
          icon: 'icon-gonggao',
          text:'政府公告',
          path: '/gov/announcement'
        },
        {
          icon: 'icon-gonggao',
          text:'投诉数据导出',
          path: '/gov/event/overview'
        },
        {
          path:'/gov/event/handle/recs',
          icon: 'icon-records',
          text: '操作记录'//操作记录管理
        }
      ]
    }
  },
  created () {
    var user = JSON.parse(sessionStorage.getItem('user'))
    if (user == null) user = {}
    else {
      this.realName = user.RealName
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
