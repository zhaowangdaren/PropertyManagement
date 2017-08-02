<template>
  <div :class='s.wrap'>
    <left-menu :menus='menus' />
  </div>
</template>

<style lang='less' module='s'>
  .wrap{
    height: 100%;
  }
</style>

<script>
  import LeftMenu from '@/components/Menu'
  export default {
    components: {LeftMenu},
    data () {
      return {
        menus: [
          {
            icon: 'icon-home',
            text:'Home',//基础信息维护
            path:'/admin'
          },
          {
            icon: 'icon-shuaxin',
            text:'Init',//数据初始化
            path:'/admin/init'
          },
          {
            icon: 'icon-users',
            text:'Users',//用户数据管理
            path:'/admin/users'
          },
          {
            icon: 'icon-wechat',
            text:'wx',//微信用户绑定管理
            path:'/admin/wx'
          },
          {
            icon: 'icon-tousu',
            text:'complaints',//微信用户绑定管理
            path:'/admin/complaints'
          },
          {
            icon: 'icon-wuyeguanli',
            text:'pm',//物业信息管理
            path:'/admin/pm'
          },
          {
            icon: 'icon-build',
            text:'build',//物业信息管理
            path:'/admin/build'
          }
        ]
      }
    }
  }
</script>