<template>
  <div :class='s.wrap'>
    <div :class='s.item' @click='onLogout'><i class='iconfont icon-kaiguan'></i>退出</div>
  </div>
</template>

<script>
export default {
  data () {},
  methods:{
    onLogout () {
      sessionStorage.clear()
      this.$router.push({path: '/'})
    },
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  background-color: rgb(235, 232, 241);
  padding: 10px;
}
.item{
  display: flex;
  justify-content: flex-end;
  align-items: center;
  font-weight: bold;
  i{
    margin-right: 5px;
    color: #999;
  }
}
</style>