<template>
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <menu-admin :class='s.menu'/>
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
import MenuAdmin from '@/components/MenuAdmin'
import ActionBar from '@/components/ActionBar'
export default {
  components: {MenuAdmin, ActionBar}
}
</script>