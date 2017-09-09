<template>
  <div :class='s.wrap'>
    <user :userType='userType'></user>
  </div>
</template>

<script type="text/javascript">
  import User from '@/components/table/User'
  export default {
    components: {User},
    data () {
      return {
        userType: 3
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
}
</style>