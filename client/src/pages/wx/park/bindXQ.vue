<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <streets @onSelect='onSelectStreet'></streets>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import Streets from '@/components/mobile/Streets'
export default {
  components: { Streets, ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: [{text:'上一步', event: this.onBack}],
        title: '选择街道',
        rightBtns: []
      }
    }
  },
  methods: {
    onBack () {
      this.$router.go(-1)
    },
    onSelectStreet (street) {
      this.$router.push({path: '/wx/park/manager/select/community', query: {streetID: street.ID}})
    }
  }
}
</script>

<style lang="less" module='s'>
.body{
  margin-top: 80px;
  font-size: 25px;
}
</style>