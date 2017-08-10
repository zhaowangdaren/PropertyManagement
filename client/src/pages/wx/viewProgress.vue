<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
      <div :class='s.up'>
        <div :class='s.item'>
          <span :class='s.key'>Evnet Status</span>
          <select v-model='selectedStatus'>
            <option v-for='status in eventStatus' 
              :value='status.value'
              v-text='status.label'></option>
          </select>
        </div>
        <div :class='s.item'>
          <span :class='s.key'>Time</span>
          <select v-model='selectedTime'>
            <option v-for='time in times' 
              :value='time.value'
              v-text='time.label'></option>
          </select>
        </div>
      </div>
      <div v-for='event, index in events' :class='s.event'>
        <div :class='s.left'>
          <div><span :class='s.key'>编号：</span><span v-text='event.Index'></span></div>
          <div><span :class='s.key'>XQ Name：</span><span v-if='xqs[index]' v-text='xqs[index].Name'></span></div>
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
export default {
  components: { ActionBar },
  filters: { filterEventStatus, filterTime },
  data () {
    return {
      host:'http://10.176.118.61:3000',
      headerOptions: {
        title: '查看进度'
      },
      events: [],
      xqs:[],
      host:'http://localhost:3000',
      selectedStatus: 0,
      eventStatus:[{value: 0, label:"全部"}, {value: 4, label: '已解决'}, {value: 1, label:'未解决'}],
      selectedTime: 0,
      times:[{value: 0, label:"近一个月内"}, {value: 1, label: '近三个月内'}, {value: 2, label:'近一年内'}, {value: 3, label: '所有'}]
    }
  },
  mounted () {
    this.fetchEvents({})
  },
  methods: {
    onDetails (event) {
      this.$router.push({name:'detailsProgress', query: {index: event.Index, status: event.Status}})
    },
    fetchEvents (query){
      fetch(this.host + '/open/event/kvs', {
        method: 'POST',
        body: JSON.stringify(query)
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEvents', body)
        if (body.error === 0) {
          this.events = body.data || []
          let xqIDs = this.events.map((item) => {
            return item.XQID
          })
          this.fetchXQs(xqIDs)
        }
      })
    },
    fetchXQs(ids){
      fetch(this.host + '/open/xq/ids', {
        method: 'POST',
        body:JSON.stringify({values: ids})
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