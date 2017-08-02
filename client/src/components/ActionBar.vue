<template>
  <div :class='s.wrap' :style='{backgroundImage: "url(" + bgImg +  ")"}'>
    <img src="~@/res/images/logo_jh.png" :class='s.logo'>
    <div :class='s.title'>Title</div>
    <div :class='s.noticeWrap'>
      <img src="~@/res/images/notice.png" :class='s.notice'>
      <div :class='s.noticeNum'>0</div>
    </div>
    <live-user/>
  </div>
</template>

<script>
import LiveUser from '@/components/LiveUser'
export default {
  components: {LiveUser},
  data () {
    return {
      bgImg: require('@/res/images/bluesky_top_bg1.jpg')
    }
  }
}
</script>

<style lang="less" module='s'>
  .wrap{
    width: 100%;
    display: flex;
    background-color: #379ad3;
    background-position: bottom;
    background-repeat: repeat-x;
    background-size: 100% 100%;
    align-items: center;
    padding: 5px;
    height: 90px;
    box-shadow: 2px 0 2px 2px #ddd;
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