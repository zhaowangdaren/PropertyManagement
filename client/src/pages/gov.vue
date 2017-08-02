<template>
  <div :class='s.wrap'>
    <action-bar></action-bar>
    <div :class='s.content'>
      <menu-gov :class='s.menu'></menu-gov>
      <div :class='s.body'>
        <div :class='s.userName' v-text='userRealName'></div>
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import MenuGov from '@/components/MenuGov'
import ActionBar from '@/components/ActionBar'
export default {
  components: {MenuGov, ActionBar},
  data () {
    return {
      userRealName: ''
    }
  },
  mounted () {
    this.userRealName = sessionStorage.getItem('RealName')
  }
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
      }
      .body{
        flex: 1;
        margin: 10px;
        .userName{
          background-color: #ddd;
          font-size: 20px;
          padding: 5px;
          margin-bottom: 20px;
        }
      }

    }
  }
</style>