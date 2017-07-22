<template>
  <div :class='s.wrap'>
    <img src="~@/res/images/logo.png" :class='s.logo'>
    <div :class='s.title'>信息系统</div>
    <div :class='s.noticeWrap'>
      <img src="~@/res/images/notice.png" :class='s.notice'>
      <div :class='s.noticeNum'>0</div>
    </div>
    <live-user/>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    width: 100%;
    display: flex;
    background-color: skyblue;
    align-items: center;
    padding: 5px;
    .logo{
      height: 0.6rem;
    }
    .title{
      flex: 1;
      margin-left: 5px;
      font-size: 40px;
    }
    .noticeWrap{
      position: relative;
      .notice{
        width: 40px;
      }
      .noticeNum{
        position: absolute;
        top: 0;
        right: 0;
        width: 0.15rem;
        height: 0.15rem;
        border-radius: 50%;
        background-color: red;
        color: #fff;
        text-align: center;
        line-height: 0.15rem;
        font-size: 0.12rem;
      }
    }
  }
</style>

<script>
import LiveUser from '@/components/LiveUser'
export default {
  components: {LiveUser}
}
</script>