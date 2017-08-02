<template>
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <menu-admin :class='s.menu'/>
      <router-view :class='s.body'></router-view>
    </div>
  </div>
</template>

<script>
import MenuAdmin from '@/components/MenuAdmin'
import ActionBar from '@/components/ActionBar'
export default {
  components: {MenuAdmin, ActionBar}
}
</script>

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
        // position: fixed;
      }
      .body{
        flex: 1;
      }
    }
  }
</style>