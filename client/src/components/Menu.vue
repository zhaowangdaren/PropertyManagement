<template>
  <div :class='s.wrap'>
    <div v-for='menu in menus'>
      <img :src="menu.icon">
      {{menu.text}}
    </div>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    border-right: solid 1px #aaa;
    height: 100%;
    div{
      border-bottom: solid 1px #ddd;
      padding: 20px 10px;
      font-size: 20px;
      display: flex;
      align-items: center;
      img{
        width: 30px;
        margin-right: 10px;
      }
    }
  }
</style>
<script type="text/javascript">
  export default {
    props: {
      menus: Array
    }
  }
</script>