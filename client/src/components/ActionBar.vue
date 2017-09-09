<template>
  <div :class='s.wrap'>
    <span :class='s.userName'>欢迎，{{userName}}</span>
    <span :class='s.item' @click='onChangePassword'>
      <i class='iconfont icon-xiugaimima'></i>修改密码
    </span>
    <span :class='s.item' @click='onLogout'>
      <i class='iconfont icon-kaiguan'></i>退出
    </span>
    <el-dialog
      title='修改密码'
      :visible.sync='showChangePassword'>
      <change-password v-if='showChangePassword' />
    </el-dialog>
  </div>
</template>

<script>
import ChangePassword from '@/components/dialog/ChangePassword'
export default {
  components: { ChangePassword },
  data () {
    return {
      userName: '',
      showChangePassword: false
    }
  },
  mounted () {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.userName = user.UserName
  },
  methods:{
    onChangePassword () {
      this.showChangePassword = true;
    },
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
  display: flex;
}
.userName{
  font-size: 30px;
  font-family: "Times New Roman";
  font-family: "Arial";
  flex: 1;
}
.item{
  display: flex;
  justify-content: flex-end;
  align-items: center;
  font-weight: bold;
  font-size: 15px;
  margin: auto 10px;
  i{
    margin-right: 10px;
    color: #999;
  }
}
</style>