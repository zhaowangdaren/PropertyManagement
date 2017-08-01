<template>
  <div :class='s.wrap'>
    <user :USER_TYPE='userType'></user>
  </div>
</template>

<script type="text/javascript">
  import User from '@/components/table/User'
  export default {
    components: {User},
    data () {
      return {
        userType: 2
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