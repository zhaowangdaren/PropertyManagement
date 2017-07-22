<template>
  <div :class='s.wrap'>
    <menu-admin :class='s.menu'/>
    <router-view></router-view>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    display: flex;
    .menu{
      height: 100%;
    }
  }
</style>

<script>
import MenuAdmin from '@/components/MenuAdmin'
  export default {
    components: {MenuAdmin}
  }
</script>