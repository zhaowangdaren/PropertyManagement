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
  data() {
    return {
      nextPath: '',
      menus:[
        {
          icon: 'icon-home',
          text: '主页',
          path: '/admin'
        },
        {
          icon: 'icon-jichu',
          text:'基础信息维护',//基础信息维护
          path:'/admin/basic'
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
  .right{
    flex: 1;
    width: 100%;
  }
  .content{
    margin: 10px;
  }
}
</style>
