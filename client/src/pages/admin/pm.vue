<template>
  <div :class='s.wrap'>
    <p-ms></p-ms>
  </div>
</template>

<script type="text/javascript">
import PMs from '@/components/table/PMs'
  export default {
    components: { PMs},
    data () {
      return {
      } 
    },
    methods: {
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  margin: 10px;
  border: solid 1px #4c87b9;
  width: 100%;
  background: #fff;
}
</style>
