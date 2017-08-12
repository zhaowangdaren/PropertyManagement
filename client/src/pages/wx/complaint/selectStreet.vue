<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div :class='s.item' v-for='street in streets' @click='onStreet(street)'>
        <div v-text='street.Name' :class='s.name'></div>
        <i class='iconfont icon-next'></i>
      </div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: { ActionBar },
  data () {
    return {
      headerOptions: {
        title: 'Select Street',
        rightBtns: [{text:'Cancel', event: null}]
      },
      streets: [],
      host:'https://127.0.0.1:3000'
    }
  },
  mounted () {
    this.headerOptions.rightBtns[0].event = this.onCancel
    this.fetchStreets()
  },
  methods: {
    onCancel () {
      this.$router.go(-1)
    },
    fetchStreets () {
      fetch(this.host + '/open/street',{
        method: 'POST',
        body: '{}'
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.streets = data.data
        }
      })
    },
    onStreet (street) {
      this.$router.push({path: '/wx/complaint/selectCommunity', query: {streetID: street.ID}})
    }
  }
}
</script>

<style lang="less" module='s'>
.body{
  margin-top: 80px;
  font-size: 25px;
  .item {
    color: #20a0ff;
    background-color: #fff;
    box-shadow: 0 1px 1px 1px #ddd;
    padding: 15px 10px;
    margin-top: 2px;
    display: flex;
    justify-content: center;
    .name{
      flex: 1;
    }
  }
}
</style>