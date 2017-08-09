<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
      <div :class='s.warn' v-text='warn'></div>
      <div :class='s.item'>
        <div :class='s.key'>进度：</div>
        <div :class='s.value'>{{event.Status | filterEventStatus }}</div>
        <el-button 
          v-if='event.Status == 0'
          type='danger' 
          :class='s.revoke'
          @click='onRevoke(event.Index)'>撤销事件</el-button>
      </div>
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
        title: '查看进度详情'
      },
      eventIndex: '',
      eventStatus: 0,
      event: {
        Status: 0
      },
      eventHandles: []
    }
  },
  mounted () {
    this.eventIndex = this.$route.query.index
    this.eventStatus = this.$route.query.status
    this.fetchEvent()
    this.fetchEventHandle()
  },
  methods: {
    onRevoke (index) {
      fetch(this.host + '/open/event/del/' + this.eventIndex, {
        method: 'POST',
        body: '{}'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onCancelEvent', body)
        if (body.error === 0 ) this.warn = '撤销成功'
        else this.warn = body.data
      })
    },
    fetchEvent () {
      fetch(this.host + '/open/event/id/' + this.eventIndex, {
        method: 'POST',
        body:'{}'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEvent', body)
        if (body.error === 0 ) {
          this.event = body.data[0] || {}
        }
      })
    },
    fetchEventHandle () {
      fetch(this.host + '/open/eventHandle/kvs', {
        method: 'POST',
        body: JSON.stringify({index: this.eventIndex})
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventHandle', body)
        if (body.error === 0){
          this.eventHandles = body.data || []
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
  .item{
    padding: 10px;
    display: flex;
    align-items: center;
    .key{
      color: #555;
    }
    .value{
      flex: 1;
    }
    .revoke {

    }
  }
}
</style>