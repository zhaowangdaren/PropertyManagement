<template>
  <div :class='s.wrap'>
    <infospm 
      :identityLevel='identityLevel'
      :curStreet='curStreet'
      >
    </infospm>
  </div>
</template>

<style lang='less' module='s'>
  .wrap{
    width: 100%;
  }
</style>

<script type="text/javascript">
  import InfosPM from '@/components/InfosPM'
  export default {
    components: { 'infospm': InfosPM },
    data () {
      return {
        identityLevel: 1,
        curStreet:'Street'
      }
    }
  }
</script>

