
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
      <div :class='s.body'>
        <div :class='s.street'>Street</div>
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    display: flex;
    flex-direction: column;
    .content{
      display: flex;
      height: 100%;
      position: absolute;
      top: 100px;
      bottom: 0;
      left: 0;
      width: 100%;
      .menu{
        height: 100%;
      }
      .body{
        flex: 1;
        .street{
          background-color: #ddd;
          font-size: 25px;
          padding: 5px;
          margin-bottom: 20px;
        }
      }
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