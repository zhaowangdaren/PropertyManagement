<template>
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <menu-gov :class='s.menu' :menus='menus' :NEXT='nextPath'/>
      <div :class='s.body'>
        <div :class='s.userName'>{{realName}}</div>
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import MenuGov from '@/components/Menu'
import ActionBar from '@/components/ActionBar'
export default {
  components: {MenuGov, ActionBar},
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
          text:'Home'
        },
        {
          icon: 'icon-tousu',
          text:'居民物业纠纷处理',
          path:'/gov/eventHandle'
        },
        {
          icon: 'icon-wuyeguanli',
          text:'查看物业信息',
          path:'/gov/pms'
        },
        {
          icon: 'icon-kaohe',
          text:'物业考核查询',
          path:'/gov/pmkpis'
        },
        {
          icon: 'icon-build',
          text:'查看建筑信息',
          path:'/gov/build'
        },
        {
          icon: 'icon-court',
          text:'推送法院调解请求',
          path: '/gov/court'
        },
        {
          icon: 'icon-fabugonggao',
          text:'推送政府公告',
          path: '/gov/announcement'
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
    display: flex;
    flex-direction: column;
    .content{
      display: flex;
      height: 100%;
      position: absolute;
      top: 100px;
      bottom: 0;
      left: 0;
      width: 100%;
      .menu{
        height: 100%;
      }
      .body{
        flex: 1;
        margin: 10px;
        .userName{
          background-color: #ddd;
          font-size: 20px;
          padding: 5px;
          margin-bottom: 20px;
        }
      }

    }
  }
</style>
