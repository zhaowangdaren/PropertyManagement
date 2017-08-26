<template>
  <div :class='s.wrap'>
    <div :class='s.logo'>
      <div><img src="~@/res/images/logo.png"></div>
      <div>花山物业管理信息平台</div>
    </div>
    <div v-for='menu in menus' @click='onClick(menu)' :class='[s.menu, {[s.active]: next == menu.path}]'>
      <i :class='"iconfont" + " " + menu.icon'></i>
      {{menu.text}}
    </div>
  </div>
</template>

<script>
  export default {
    props: {
      menus: Array,
      NEXT: {
        type: String,
        default: ''
      }
    },
    data () {
      return {
        next: ''
      }
    },
    watch: {
    },
    mounted () {
      console.info('menuPath', this.$route.path)
      this.next = this.NEXT
      if (this.next === '') this.next = this.$route.path
      console.warn('NEXT', this.NEXT)
      console.info('menus', this.menus)
    },
    methods: {
      onClick (menu) {
        this.next = menu.path
        this.$router.push({path: menu.path})
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  height: 100%;
  background-color: #1f2d3d;
  min-width: 150px;
  color: #f1f3f6;
  .menu{
    padding: 20px 20px;
    font-size: 18px;
    display: flex;
    align-items: center;
    cursor: pointer;
    i{
      width: 30px;
      margin-right: 10px;
    }
  }
  .active{
    background: #4c87b9;
  }
}
.logo{
  border-bottom: solid rgba(255, 255, 255, 0.5) 1px;
  font-size: 20px;
  text-align: center;
  padding: 10px 0;
  img{
    width: 30%;
  }
}
</style>
