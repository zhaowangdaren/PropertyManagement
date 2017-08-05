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
            path:'/street',
            text:'Home'//首页
          },
          {
            path:'/street/handle',
            icon: 'icon-chulizhong',
            text:'complaintsHandler' //居民物业纠纷处理
          },
          {
            path:'/street/pms',
            icon: 'icon-wuyeguanli',
            text:'PM infos'//物业信息管理
          },
          {
            path:'/street/pmkpis',
            icon: 'icon-kaohe',
            text:'PM KPI'//物业考核查询
          },
          {
            path:'/street/court',
            icon: 'icon-court',
            text:'Court'//申请法院调解请求
          },
          {
            path:'/street/operationRecs',
            icon: 'icon-records',
            text:'Operation recs'//操作记录管理
          }
        ]
      }
    }
  }
</script>
