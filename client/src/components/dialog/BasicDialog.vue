<template lang="html">
  <div :class="s.wrap" @click.self='onClose' v-if='visible'>
    <div :class="s.body">
      <div :class="s.header">
        <div :class="s.title" v-text='title'></div>
        <i class='iconfont icon-close' @click='onClose'></i>
      </div>
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    visible: Boolean,
    title: String
  },
  methods: {
    onClose () {
      this.$emit('update:visible', false)
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  position: fixed;
  z-index: 2000;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0,0.3);
  border-radius: 2px;
  .body{
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.3);
    top: 15%;
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    background: #fff;
    border-radius: 0.013333rem;
    box-sizing: border-box;
    margin-bottom: 0.333333rem;
    padding: 10px;
    .header{
      display: flex;
      align-items: center;
      margin-bottom: 20px;
      .title{
        color: #1f2d3d;
        font-size: 20px;
        font-weight: bold;
        flex: 1;
        // text-align: left;
      }
      i{
        color: #bfcbd9;
      }
    }
  }

}
</style>
