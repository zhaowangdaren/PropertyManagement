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
            text:'首页',
            path:'/gov'
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
    }
  }
</script>
