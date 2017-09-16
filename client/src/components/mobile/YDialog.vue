<template>
  <div :class="s.wrap" v-if='visible' @click.self='onClose'>
    <div :class="s.close"><i class="iconfont icon-close" @click='onClose'></i></div>
    <div :class="s.content">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    visible: Boolean
  },
  methods: {
    onClose () {
      this.$emit('update:visible', false)
    }
  }
}
</script>
<style lang="less" module="s">
.wrap{
  background-color: rgba(31, 45, 61, .61);
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  // height: 100%;
  // width: 100%;
  z-index: 101;
  padding-top: 120px;
  overflow: auto;
}
.close{
  display: flex;
  justify-content: flex-end;
  align-items: center;
  i{
    color: #fff;
    font-size: 40px;
    margin-right: 50px;
  }
}
.content{
  background-color: #fff;
  margin: 20px 30px;
  overflow: auto;
  border-radius: 5px;
}
</style>