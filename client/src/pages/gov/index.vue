<template>
  <div :class='s.wrap'>
    <to-be-done from='gov'></to-be-done>
  </div>
</template>

<style lang="less" module='s'>
.wrap{
  width: 100%;
}
</style>
<script>
  import ToBeDone from '@/components/table/ToBeDone'
  export default {
    components: { ToBeDone }
  }
</script>