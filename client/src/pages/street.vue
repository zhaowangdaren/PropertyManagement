
<template>
<!--  首页
 居民物业纠纷处理
 物业信息管理
 物业考核查询
 申请法院调解请求
 操作记录管理 -->
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <menu-street :class='s.menu'/>
      <router-view></router-view>
    </div>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    display: flex;
    flex-direction: column;
    .content{
      .menu{
        height: 100%;
      }
      display: flex;
    }
  }
</style>

<script>
import MenuStreet from '@/components/MenuStreet'
import ActionBar from '@/components/ActionBar'
export default {
  components: {MenuStreet, ActionBar}
}
</script>