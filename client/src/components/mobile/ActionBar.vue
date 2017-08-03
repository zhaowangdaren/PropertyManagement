<template>
  <div :class='s.wrap'>
    <div :class='s.leftBtns'>
      <div :class='s.btn' v-for='leftBtn in OPTIONS.leftBtns' v-text='leftBtn'></div>
    </div>
    <div v-text='OPTIONS.title' :class='s.title'></div>
    <div :class='s.rightBtns'>
      <div :class='s.btn' v-for='leftBtn in OPTIONS.rightBtns' v-text='leftBtn'></div>
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
  box-shadow: 0 1px 1px 1px #ddd;
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