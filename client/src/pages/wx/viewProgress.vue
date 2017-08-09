<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
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
      <div v-for='event in events' :class='s.event'>
        <div><span :class='s.key'>编号：</span><span v-text='event.Index'></span></div>
      </div>
      <div v-if='events.length == 0' :class='s.event + " "+ s.tip'>无记录</div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: { ActionBar },
  data () {
    return {
      host:'http://10.176.118.61:3000',
      headerOptions: {
        title: '查看进度'
      },
      events: [],
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
  // color: #555;
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
  .event{
    border-top: solid 1px #000;
    width: 100%;
    .key{
      color: #555;
    }
  }
  .tip{
    text-align: center;
    color: #999;
    padding: 10px;
  }
}

</style>