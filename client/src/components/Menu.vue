<template>
  <div :class='s.wrap'>
    <div v-for='menu in menus' @click='onClick(menu)' :class='{[s.active]: NEXT == menu.path}'>
      <i :class='"iconfont" + " " + menu.icon'></i>
      {{menu.text}}
    </div>
  </div>
</template>

<script>
  export default {
    props: {
      menus: Array,
      NEXT: String
    },
    data () {
      return {
      }
    },
    mounted () {
      console.info('menuPath', this.$route.path)
      if (this.NEXT == '') this.NEXT = this.$route.path
      console.warn('NEXT', this.NEXT)
      console.info('menus', this.menus)
    },
    methods: {
      onClick (menu) {
        this.$router.push({path: menu.path})
      }
    }
  }
</script>

<style lang="less" module='s'>
  .wrap{
    border-right: solid 1px #aaa;
    height: 100%;
    background-color: #f0f0f0;
    div{
      border-bottom: solid 1px #ddd;
      padding: 20px 10px;
      font-size: 20px;
      display: flex;
      align-items: center;
      i{
        width: 30px;
        margin-right: 10px;
      }
    }
    .active{
      background: #20a0ff;
      color: #fff;
    }
  }
</style>
