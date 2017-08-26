<template>
  <div :class='s.wrap'>
    <div :class='s.up'>
      <div :class='s.title'>事件总数</div>
      <div :class='s.value'>{{sum}}</div>
    </div>
    <div :class='s.down'>
      <div v-for='status in eventStatus' :class='s.item'>
        <div :class='s.title'>{{status.status}}</div>
        <div :class='s.value'>{{status.num}}</div>
      </div>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
import filterEventStatus from '@/filters/filterEventStatus'
export default {
  data () {
    return {
      eventStatus: [],
      sum: 0
    }
  },
  mounted () {
    this.fetchEvenStatus()
  },
  methods: {
    fetchEvenStatus () {
      fetchpm(this, true, '/pm/event/key/nums/status', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEvenStatus', body)
        if (body.error === 0) {
          body.data = body.data || []
          this.eventStatus = this.formatStatus(body.data)
        }
      })
    },
    formatStatus (statusNums) {
      this.sum = 0
      for (var i = 0; i < statusNums.length; i++) {
        statusNums[i].status = filterEventStatus(statusNums[i]._id.status)
        this.sum += statusNums[i].num
      }
      return statusNums
    }
  }
}
</script>
<style lang='less' module='s'>
.wrap{
  background-color: #1f2d3d;
  border: solid 1px #ddd;
  border-radius: 5px;
  padding: 5px;
  box-shadow: 1px 1px 4px 0px #ddedf7,-1px -1px 4px 0px #ddedf7;
}
.title{
  color: #bac9d8;
  font-size: 16px;
  font-weight: bold;
}
.value{
  color: #fff;
  font-size: 30px;
  font-weight: bold;
}
.down{
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
}
.item{
  margin: 5px;
  flex: 1;
}
</style>