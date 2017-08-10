<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div>Content</div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: { ActionBar }
}
</script>

<style lang="less" module='s'>
.wrap{
  margin-top: 80px;
}
</style>