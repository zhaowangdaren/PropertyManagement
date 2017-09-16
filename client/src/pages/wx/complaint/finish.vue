<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <i class='iconfont icon-succ'></i>
      <p>提交成功</p>
      <!-- <div :class='s.btn' @click='onFinish'>好的</div> -->
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: {ActionBar},
  data () {
    return {
      headerOptions: {
        title: '完成'
      }
    }
  }
}
</script>

<style lang="less" module='s'>
.body{
  padding-top: 100px;
  margin: auto 5px;
  font-size: 30px;
  text-align: center;
  i{
    color: #1eb504;
    font-size: 60px;
  }
  .btn{
    background-color: #1eb504;
    font-size: 30px;
    color: #fff;
    text-align: center;
    padding: 15px 0;
    margin: 20px;
  }
}
</style>