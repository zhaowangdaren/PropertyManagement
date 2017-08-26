<template>
  <div :class='s.wrap'>
    <p-ms :EDITABLE='editable'></p-ms>
  </div>
</template>

<script type="text/javascript">
import PMs from '@/components/table/PMs'
  export default {
    components: { PMs},
    data () {
      return {
        editable: true
      } 
    },
    methods: {
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
}
</style>
