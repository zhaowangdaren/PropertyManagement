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
            text:'首页'
          },
          {
            icon: require('@/res/images/complaints.png'),
            text:'居民物业纠纷处理'
          },
          {
            icon: require('@/res/images/wuye.png'),
            text:'查看物业信息'
          },
          {
            icon: require('@/res/images/build.png'),
            text:'查看建筑信息'
          },
          {
            icon: require('@/res/images/kpi.png'),
            text:'物业考核查询'
          },
          {
            icon: require('@/res/images/court.png'),
            text:'推送法院调解请求'
          },
          {
            icon: require('@/res/images/gonggao.png'),
            text:'推送政府公告'
          }
        ]
      }
    }
  }
</script>