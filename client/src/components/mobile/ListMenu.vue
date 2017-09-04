<!-- item 告诉父组件选择
{
  lable: '名称',
  value: ''
}
 -->
<template>
  <div>
    <div :class='s.item' v-for='item in list' @click='onItem(item)'>
      <div v-text='item.label' :class='s.name'></div>
      <i class='iconfont icon-next'></i>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    list: Array
  },
  methods: {
    onItem (item) {
      this.$emit('item', item)
    }
  }
}
</script>

<style lang="less" module='s'>
.item {
  color: #555;
  background-color: #fff;
  box-shadow: 0 1px 1px 1px #ddd;
  padding: 20px 10px;
  margin-top: 2px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
  .name{
    flex: 1;
  }
  i{
    color: #999;
  }
}
</style>