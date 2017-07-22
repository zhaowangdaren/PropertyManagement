<template>
  <div :class='s.wrap'>
    <img src="~@/res/images/header.jpg">
    <div :class='s.userWrap'>
      <div :class='s.user' >
        <div>Admin</div>
        <img src="~@/res/images/unfold.png" :class='s.unfold'>
      </div>
      <div :class='s.operations'>
        <div :class='s.item'>
          <img src="~@/res/images/return.png">返回控制台
        </div>
        <div :class='s.item'>
          <img src="~@/res/images/password.png">
          修改密码
        </div>
        <div :class='s.item'>
          <img src="~@/res/images/logout.png">
          退出登录
        </div>
      </div>
    </div>
    <img src="~@/res/images/logout.png" :class='s.logout'>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    display: flex;
    align-items: center;
    margin-left: 20px;
    .userWrap{
      position: relative;
      padding: 5px 20px;
      &:hover .operations{
        display: block;
      }
      .user{
        display: flex;
        width: 100%;
        align-items: center;
        justify-content: center;
        .unfold{
          height: 20px;
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
        img{
          width: 20px;
          margin-right: 10px;
        }
      }
    }
    
    .logout{
      height: 30px;
      margin-left: 20px;
      margin-right: 20px;
    }
  }
</style>

<script>
  export default {
    data () {
      return {
        showOperations: false
      }
    },
    methods: {
      onOperations () {
        this.showOperations = true;
      }
    }
  }
</script>