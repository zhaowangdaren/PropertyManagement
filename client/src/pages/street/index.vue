<template>
  <div :class='s.wrap'>
    <to-be-done></to-be-done>
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import ToBeDone from '@/components/table/ToBeDone'
  export default {
    components: { ToBeDone },
    filters: {filterEventStatus, filterEventLevel},
    data () {
      return {
        events: []
      }
    },
    mounted () {
    },
    methods: {
    }
  }
</script>
<style lang="less" module='s'>
.wrap{
  margin: 10px; 
}
</style>