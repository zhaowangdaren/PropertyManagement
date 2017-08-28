<template>
  <div :class='s.wrap'>
    <div :class='s.start' @click='onCheck'>查看停车位</div>
    <div :class='s.start' @click='onManager'>我是管理员</div>
  </div>
</template>

<script>
export default {
  methods: {
    onCheck () {

    },
    onManager () {
      
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  background-color: #20a0ff;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .start{
    text-align: center;
    width: 60%;
    border:solid 2px #fff;
    border-radius: 5px;
    padding: 20px;
    margin: 5px;
    color: #fff;
    font-size: 30px;
    font-weight: bold;
    &:active {
      background-color: rgba(254,254,254, 0.1);
    }
  }
}
</style>