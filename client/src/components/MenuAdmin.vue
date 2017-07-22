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
            text:'基础信息维护'
          },
          {
            icon: require('@/res/images/refresh.png'),
            text:'数据初始化'
          },
          {
            icon: require('@/res/images/users.png'),
            text:'用户管理'
          },
          {
            icon: require('@/res/images/wechat.png'),
            text:'微信用户绑定管理'
          },
          {
            icon: require('@/res/images/complaints.png'),
            text:'投诉数据维护'
          },
          {
            icon: require('@/res/images/wuye.png'),
            text:'物业信息管理'
          },
          {
            icon: require('@/res/images/build.png'),
            text:'建筑信息管理'
          }
        ]
      }
    }
  }
</script>