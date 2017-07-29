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
            path:'/street',
            text:'Home'//首页
          },
          {
            path:'/street/handle',
            icon: require('@/res/images/complaints.png'),
            text:'complaintsHandler' //居民物业纠纷处理
          },
          {
            path:'/street/pms',
            icon: require('@/res/images/wuye.png'),
            text:'PM infos'//物业信息管理
          },
          {
            path:'/street/pmkpis',
            icon: require('@/res/images/kpi.png'),
            text:'PM KPI'//物业考核查询
          },
          {
            path:'/street/court',
            icon: require('@/res/images/court.png'),
            text:'Court'//申请法院调解请求
          },
          {
            path:'/street/operationRecs',
            icon: require('@/res/images/tasks.png'),
            text:'Operation recs'//操作记录管理
          }
        ]
      }
    }
  }
</script>