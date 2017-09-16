<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
      <div :class='s.up'>
        <div :class='s.item'>
          <span :class='s.key'>投诉状态</span>
          <select v-model='selectedStatus'>
            <option v-for='status in eventStatus' 
              :value='status.value'
              v-text='status.label'></option>
          </select>
        </div>
        <div :class='s.item'>
          <span :class='s.key'>投诉时间</span>
          <select v-model='selectedTime'>
            <option v-for='time in times' 
              :value='time.value'
              v-text='time.label'></option>
          </select>
        </div>
      </div>
      <div v-for='event, index in showEvents' :class='s.event'>
        <div :class='s.left'>
          <div><span :class='s.key'>编号：</span><span v-text='event.Index'></span></div>
          <div><span :class='s.key'>小区名：</span><span v-if='xqs[index]' v-text='xqs[index].Name'></span></div>
          <div><span :class='s.key'>进度：</span><span>{{event.Status | filterEventStatus}}</span></div>
          <div><span :class='s.key'>提交时间:</span><span>{{event.Time | filterTime}}</span></div>
        </div>
        <div :class='s.btn' @click='onDetails(event)'>
          查看<br/>详情
        </div>
      </div>
      <div v-if='events.length == 0' :class='s.event + " "+ s.tip'>无记录</div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import filterEventStatus from '@/filters/filterEventStatus'
import filterTime from '@/filters/filterTime'
import fetchpm from '@/fetchpm'
export default {
  components: { ActionBar },
  filters: { filterEventStatus, filterTime },
  data () {
    return {
      headerOptions: {
        title: '查看进度'
      },
      openid: '',
      events: [],
      xqs:[],
      selectedStatus: 0,
      eventStatus:[{value: 0, label:"全部", status: null}, {value: 4, label:'未解决', status: 4}, {value: 3, label: '已解决', status: 3}],
      aMonth: 2592000, // 30 * 24 * 60 * 60 * 1000
      threeMonth: 7776000, // 3 * 30 * 24 * 60 * 60 * 1000
      aYear: 31104000, // 12 * 30 * 24 * 60 * 60 * 1000
      selectedTime: 0,
      times:[{value: 0, label:"近一个月内", startTime: 0, endTime: 0}, {value: 1, label: '近三个月内', startTime: 0, endTime: 0}, {value: 2, label:'近一年内', startTime: 0, endTime: 0}, {value: 3, label: '所有'}]
    }
  },
  computed: {
    showEvents: function () {
      return this.events.filter((item) => {
        switch(this.selectedStatus){
          case 0:
            return true
          case 3:
            return item.Status === 3
          case 4:
            return item.Status === 4
          default:
            return true
        }
      })
    }
  },
  watch: {
    selectedTime: function (val) {
      switch(this.selectedTime){
        case 0:
        case 1:
        case 2:
          this.fetchEventsByTime(this.openid, this.times[this.selectedTime].startTime, this.times[this.selectedTime].endTime)
          break
        case 3:
          this.fetchEvents({OpenID: this.openid})
          break
      }
    },
    showEvents: function (val) {
      let xqIDs = val.map((item) => {
        return item.XQID
      })
      this.fetchXQs(xqIDs)
    }
  },
  mounted () {
    var WXUser = JSON.parse(sessionStorage.getItem('WXUser')) || {openid: ''}
    this.openid = WXUser.openid
    this.initData()
    this.fetchEventsByTime(this.openid, this.times[this.selectedTime].startTime, this.times[this.selectedTime].endTime)
  },
  methods: {
    initData () {
      var today = new Date().getTime() / 1000
      this.times[0].startTime = parseInt(today - this.aMonth)
      this.times[0].endTime = parseInt(today)
      this.times[1].startTime = parseInt(today - this.threeMonth)
      this.times[1].endTime = this.times[0].endTime
      this.times[2].startTime = parseInt(today - this.aYear)
      this.times[2].endTime = this.times[0].endTime
      console.info('times', this.times)
    },
    onDetails (event) {
      this.$router.push({name:'detailsProgress', query: {index: event.Index, status: event.Status, identity: "user"}})
    },
    fetchEvents (query){
      fetchpm(this, false,'/open/event/kvs', {
        method: 'POST',
        body: query
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEvents', body)
        if (body.error === 0) {
          this.events = body.data || []
        }
        this.xqs = []
      })
    },
    fetchEventsByTime (openID, startTime, endTime) {
      fetchpm(this, false, '/open/event/check/progress', {
        method: 'POST',
        body: {OpenID: openID, StartTime: startTime, EndTime: endTime}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventsByTime', body)
        if (body.error === 0) this.events = body.data || []
        else this.events = []
        this.events.sort((a, b) => {
          return b.Time - a.Time
        })
        this.xqs = []
      })
    },
    fetchXQs(ids){
      fetchpm(this, false, '/open/xq/ids', {
        method: 'POST',
        body:{values: ids}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0){
          this.xqs = body.data || []
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.content{
  margin-top: 80px;
  font-size: 25px;
  .up{
    background-color: #fff;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.3);
    .item{
      padding: 10px;
      display: flex;
      align-items: center;
      .key{
        color: #555;
      }
      select{
        flex: 1;
        padding: 10px;
        border-radius: 5px;
        margin: 10px;
      }
    }
  }
  
  .event{
    margin: 10px;
    background-color: #fff;
    border-radius: 2px;
    padding: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    .left{
      flex: 1;
      .key{
        color: #555;
      }
    }
    .btn{
      background-color: #20a0ff;
      border-radius: 5px;
      padding: 10px;
      color: #fff;
      font-size: 12px;
    }
    
  }
  .tip{
    text-align: center;
    color: #999;
    padding: 10px;
  }
}

</style>