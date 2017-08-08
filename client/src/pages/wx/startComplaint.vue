<template>
  <div :class='s.wrap'>
    <div :class='s.start' @click='onStart'>Start Complaining</div>
  </div>
</template>

<script>
export default {
  methods: {
    onStart () {
      this.$router.push({path: '/wx/complaint/selectStreet'})
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
  justify-content: center;
  align-items: center;
  .start{
    border:solid 2px #fff;
    border-radius: 5px;
    padding: 20px;
    color: #fff;
    font-size: 30px;
    font-weight: bold;
    &:active {
      background-color: rgba(254,254,254, 0.1);
    }
  }
}
</style>