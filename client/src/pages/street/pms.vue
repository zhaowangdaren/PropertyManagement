<template>
  <div :class='s.wrap'>
    <p-ms :EDITABLE='editable'></p-ms>
  </div>
</template>

<style lang='less' module='s'>
  .wrap{
    width: 100%;
  }
</style>

<script type="text/javascript">
  import PMs from '@/components/table/PMs'
  export default {
    components: { PMs },
    data () {
      return {
        editable: false
      }
    }
  }
</script>

