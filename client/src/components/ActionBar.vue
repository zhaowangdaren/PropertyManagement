<template>
  <div :class='s.wrap' :style='{backgroundImage: "url(" + bgImg +  ")"}'>
    <img src="~@/res/images/logo.png" :class='s.logo'>
    <div :class='s.title'>
      <img src="~@/res/images/title.png">
    </div>
    <div :class='s.noticeWrap'>
      <i :class='"iconfont icon-notice " + s.notice'></i>
      <span :class='s.noticeNum'>0</span>
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
    position: absolute;
    left: 0;
    top: 0;
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
      color: #fff;
      img{
        width: 90%;
      }
    }
    .noticeWrap{
      position: relative;
      .notice{
        font-size: 30px;
        vertical-align: middle;
        color: #fff;
      }
      .noticeNum{
        min-width: 14px;
        text-align: center;
        line-height: 14px;
        display: inline-block;
        position: absolute;
        right: -5px;
        top: -8px;
        background: red;
        color: #fff;
        border-radius: 17px;
        padding: 6px;
        font-size: 16px;
        transform: scale(.7);
        font-family: Tahoma!important;
      }
    }
  }
</style>
