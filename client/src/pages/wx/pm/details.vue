<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div>Content</div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: { ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: [{text:'上一步', event: null}],
        title: 'PM Details',
        rightBtns: []
      },
      xqs: [],
      host:'http://10.176.118.61:3000'
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  margin-top: 80px;
}
</style>