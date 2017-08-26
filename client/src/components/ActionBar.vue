<template>
  <div :class='s.wrap'>
    <span :class='s.userName'>{{userName}} ,您好</span>
    <span :class='s.item' @click='onLogout'>
      <i class='iconfont icon-kaiguan'></i>退出
    </span>
  </div>
</template>

<script>
export default {
  data () {
    return {
      userName: ''
    }
  },
  mounted () {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.userName = user.UserName
  },
  methods:{
    onLogout () {
      sessionStorage.clear()
      debugger
      this.$router.push({path: '/'})
    },
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  background-color: rgb(235, 232, 241);
  padding: 10px;
  display: flex;
}
.userName{
  font-size: 20px;
  flex: 1;
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