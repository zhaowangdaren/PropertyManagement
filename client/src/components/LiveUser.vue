<template>
  <div :class='s.wrap'>
    <img src="~@/res/images/header.jpg">
    <div :class='s.userWrap'>
      <div :class='s.user' >
        <div v-text='userName' :class='s.name'></div>
        <i :class='"iconfont icon-return " + s.unfold'></i>
      </div>
      <div :class='s.operations'>
        <div :class='s.item' @click='onHome'>
          <i class="iconfont icon-Return"></i>返回控制台
        </div>
        <div :class='s.item' @click='onChangePassword'>
          <i class="iconfont icon-xiugaimima"></i>修改密码
        </div>
        <div :class='s.item' @click='onLogout'>
          <i class="iconfont icon-tuichu"></i>退出登录
        </div>
      </div>
    </div>
    <i :class='"iconfont icon-tuichu " + s.logout'
      @click='onLogout'></i>
    <el-dialog
      title='修改密码'
      :visible.sync='showChangePassword'>
      <change-password v-if='showChangePassword' />
    </el-dialog>
  </div>
</template>

<script>
  import ChangePassword from '@/components/dialog/ChangePassword'
  import BasicDialog from '@/components/dialog/BasicDialog'
  export default {
    components: { ChangePassword, BasicDialog },
    data () {
      return {
        showOperations: false,
        showChangePassword: false,
        userName: ''
      }
    },
    mounted () {
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      this.userName = user.UserName
    },
    methods: {
      onOperations () {
        this.showOperations = true;
      },
      onLogout () {
        sessionStorage.clear()
        this.$router.push({path: '/'})
      },
      onChangePassword () {
        this.showChangePassword = true;
      },
      onHome () {
        console.info('onHome', this.$route)
        this.$router.push({path: this.$route.matched[0].path})
      }
    }
  }
</script>

<style lang="less" module='s'>
  .wrap{
    display: flex;
    align-items: center;
    justify-content: center;
    margin-left: 20px;
    position: relative;
    .userWrap{
      position: relative;
      padding: 5px 20px;
      cursor: pointer;
      &:hover .operations{
        display: block;
      }
      .user{
        display: flex;
        width: 100%;
        align-items: center;
        justify-content: center;
        font-size: 22px;
        color: #fff;
        .name{
          // margin: auto 10px;
        }
        .unfold{
          transform: rotate(-90deg);
        }
      }
      .operations{
        display: none;
        position: absolute;
        width: 120%;
        left: 0;
        background-color: #fff;
        box-shadow: 0 2px 1px 1px #ddd;
        z-index: 2;
        .item {
          padding: 5px 5px;
          &:hover {
            background-color: #f0f0f0;
          }
        }
        i{
          width: 20px;
          margin-right: 10px;
        }
      }
    }

    .logout{
      font-weight: bold;
      color: #fff;
      height: 30px;
      margin-left: 20px;
      margin-right: 20px;
    }
  }
</style>
