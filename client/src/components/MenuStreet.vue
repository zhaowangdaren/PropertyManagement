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
            text:'物业信息管理'
          },
          {
            icon: require('@/res/images/kpi.png'),
            text:'物业考核查询'
          },
          {
            icon: require('@/res/images/court.png'),
            text:'申请法院调解请求'
          },
          {
            icon: require('@/res/images/tasks.png'),
            text:'操作记录管理'
          }
        ]
      }
    }
  }
</script>