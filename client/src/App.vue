<template>
  <div>
  <action-bar/>
    <router-view></router-view>
  </div>
</template>

<script>
  import ActionBar from '@/components/ActionBar'
  export default {
    components: {ActionBar}
  }
</script>