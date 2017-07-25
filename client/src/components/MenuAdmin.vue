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
            icon: require('@/res/images/home.png'),
            text:'Home',//基础信息维护
            path:'/admin/'
          },
          {
            icon: require('@/res/images/refresh.png'),
            text:'Init',//数据初始化
            path:'/admin/init'
          },
          {
            icon: require('@/res/images/users.png'),
            text:'Users',//用户数据管理
            path:'/admin/users'
          },
          {
            icon: require('@/res/images/wechat.png'),
            text:'wx',//微信用户绑定管理
            path:'/admin/wx'
          },
          {
            icon: require('@/res/images/complaints.png'),
            text:'complaints',//微信用户绑定管理
            path:'/admin/complaints'
          },
          {
            icon: require('@/res/images/wuye.png'),
            text:'pm',//物业信息管理
            path:'/admin/pm'
          },
          {
            icon: require('@/res/images/build.png'),
            text:'build',//物业信息管理
            path:'/admin/build'
          }
        ]
      }
    }
  }
</script>