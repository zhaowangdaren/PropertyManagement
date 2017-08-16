<template>
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <left-menu :menus='menus' :NEXT='nextPath'/>
      <router-view :class='s.body'></router-view>
    </div>
  </div>
</template>

<script>
import LeftMenu from '@/components/Menu'
import ActionBar from '@/components/ActionBar'
export default {
  components: {LeftMenu, ActionBar},
  data() {
    return {
      nextPath: '',
      menus:[
        {
          icon: 'icon-home',
          text:'基础信息维护',//基础信息维护
          path:'/admin'
        },
        {
          icon: 'icon-shuaxin',
          text:'数据初始化',//数据初始化
          path:'/admin/init'
        },
        {
          icon: 'icon-users',
          text:'用户数据管理',//用户数据管理
          path:'/admin/users'
        },
        {
          icon: 'icon-wechat',
          text:'微信用户绑定管理',//微信用户绑定管理
          path:'/admin/wx'
        },
        {
          icon: 'icon-tousu',
          text:'投诉数据维护',//微信用户绑定管理
          path:'/admin/complaints'
        },
        {
          icon: 'icon-wuyeguanli',
          text:'物业信息管理',//物业信息管理
          path:'/admin/pm'
        },
        {
          icon: 'icon-build',
          text:'建筑信息管理',//物业信息管理
          path:'/admin/build'
        }
      ]
    }
  },
  beforeRouteUpdate(to, from, next) {
    console.info('beforeRouteUpdate', to)
    this.nextPath = to.path
    next()
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
    flex-direction: column;
    .content{
      flex: 1;
      margin-top: 100px;
      display: flex;
      width: 100%;

      .menu{
        height: 100%;
        // position: fixed;
      }
      .body{
        flex: 1;
      }
    }
  }
</style>
