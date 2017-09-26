<template>
  <div :class='s.wrap'>
    <div :class='s.left'>
      <left-menu :menus='menus' :NEXT='nextPath'/>
    </div>
    <div :class='s.right'>
      <action-bar></action-bar>
      <div :class='s.content'>
        <router-view :class='s.body'></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import LeftMenu from '@/components/Menu'
import ActionBar from '@/components/ActionBar'
export default {
  components: {LeftMenu, ActionBar},
  beforeRouteUpdate(to, from, next) {
    console.info('beforeRouteUpdate', to)
    this.nextPath = to.path
    next()
  },
  data () {
    return {
      realName: '',
      nextPath: '',
      menus: [
        {
          path: '/court',
          icon: 'icon-home',
          text:'居民投诉调解'
        },
        {
          path:'/court/event/handle/recs',
          icon: 'icon-records',
          text: '操作记录'//操作记录管理
        }
      ]
    }
  },
  created () {
    var user = JSON.parse(sessionStorage.getItem('user'))
    if (user == null) user = {}
    else {
      this.realName = user.RealName
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  .right{
    overflow: auto;
    flex: 1;
    width: 100%;
  }
  .content{
    margin: 10px;
  }
}
</style>
