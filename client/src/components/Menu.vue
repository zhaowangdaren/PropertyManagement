<template>
  <div :class='s.wrap'>
    <div v-for='menu in menus' @click='onClick(menu)' :class='{[s.active]: curPath == menu.path}'>
      <i :class='"iconfont" + " " + menu.icon'></i>
      {{menu.text}}
    </div>
  </div>
</template>

<script type="text/javascript">
  export default {
    props: {
      menus: Array
    },
    data () {
      return {
        curPath: ''
      }
    },
    mounted () {
      console.info('menuPath', this.$route.path)
      this.curPath = this.$route.path
      console.info(this.menus)
    },
    methods: {
      onClick (menu) {
        this.$router.push({path: menu.path})
        this.$nextTick(() => {
          this.curPath = this.$route.path
        })
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
      background: #379ad3;
      color: #fff;
    }
  }
</style>