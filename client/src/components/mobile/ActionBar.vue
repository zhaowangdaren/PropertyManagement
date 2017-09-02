<template>
  <div :class='s.wrap'>
    <div :class='s.leftBtns'>
      <div :class='s.btn' v-for='leftBtn in OPTIONS.leftBtns' v-text='leftBtn.text' @click='onLeft(leftBtn)'></div>
    </div>
    <div v-text='OPTIONS.title' :class='s.title'></div>
    <div :class='s.rightBtns'>
      <div :class='s.btn' v-for='rightBtn in OPTIONS.rightBtns' v-text='rightBtn.text' @click='onRight(rightBtn)'></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    OPTIONS: {
      type: Object,
      default: function () {
        return {
          leftBtns:[],
          title: 'Title',
          rightBtns: []
        }
      }
    }
  },
  methods: {
    onLeft (leftBtn) {
      if (leftBtn.event) leftBtn.event()
    },
    onRight (rightBtn) {
      if (rightBtn.event) rightBtn.event()
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  top: 0;
  height: 80px;
  width: 100%;
  position: fixed;
  background-color: #20a0ff;
  display: flex;
  color: #fff;
  justify-content: center;
  align-items: center;
  box-shadow: 1px 0px 2px 1px rgba(0,0,0,0.5);
  z-index: 9999;
  .title{
    flex: 1;
    font-size: 30px;
    overflow: hidden;
    text-overflow:ellipsis;
    white-space: nowrap;
    text-align: center;
  }
  .btn{
    padding: 5px 10px;
    font-size: 20px;
  }
  .leftBtns {
    min-width: 40px;
    margin-left: 5px;
  }
  .rightBtns {
    min-width: 40px;
    margin-right: 5px;
  }
}
</style>