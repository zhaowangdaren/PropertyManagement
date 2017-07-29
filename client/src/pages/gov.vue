<template>
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <menu-gov :class='s.menu'></menu-gov>
      <div :class='s.body'>
        <div :class='s.userName'>UserName</div>
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import MenuGov from '@/components/MenuGov'
import ActionBar from '@/components/ActionBar'
export default {
  components: {MenuGov, ActionBar}
}
</script>

<style lang="less" module='s'>
  .wrap{
    display: flex;
    flex-direction: column;
    .content{
      display: flex;
      .menu{
        height: 100%;
      }
      .body{
        flex: 1;
        margin: 10px;
        .userName{
          background-color: #ddd;
          font-size: 20px;
          padding: 5px;
        }
      }

    }
  }
</style>