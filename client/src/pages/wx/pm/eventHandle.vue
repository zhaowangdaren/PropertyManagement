<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
      <div :class='s.item'>小区名称：<span>{{ xq.Name }}</span></div>
      <div :class='s.item'>投诉时间：<span>{{ event.Time | filterTime}}</span></div>
      <div :class='s.item'>处理说明：</div>
      <div :class='s.textarea'>
        <textarea
        rows='10'
        placeholder="请输入处理内容"
        v-model="eventHandle.HandleInfo" 
        @focus='onFocus'></textarea>
      </div>
      <div :class='s.btn' @click='onHandle'>确认已处理</div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import filterTime from '@/filters/filterTime'
import fetchpm from '@/fetchpm'
export default {
  components: { ActionBar },
  filters: { filterTime },
  data () {
    return {
      headerOptions: {
        title: '确认处理'
      },
      event: {
        Index: '',
        Time: 0,
        XQID:''
      },
      eventHandle: {
        HandleInfo: '',
        Index: '',
        AuthorCategory: 4,
        AuthorName: '物业公司'
      },
      xq: {
        Name: ''
      }
    }
  },
  mounted () {
    this.event.Index = this.$route.query.index
    this.fetchEvent(this.event.Index)
  },
  methods: {
    onHandle () {
      fetchpm(this, false, '/open/eventHandle/pm/deal', {
        method: 'POST',
        body: this.eventHandle
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0) {
          this.$router.push({path:'/wx/complaint/finish'})
        } else {
          Message({type:'error', message: body.data})
        }
      })
    },
    fetchEvent (index) {
      fetchpm(this, false, '/open/event/id/' + index, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEvent', body)
        if (body.error === 0 ) {
          this.event = body.data.events[0] || {}
          this.fetchXQName(this.event.XQID)
          this.eventHandle = Object.assign(this.eventHandle, this.event)
        }
      })
    },
    fetchXQName (id) {
      fetchpm(this, false, '/open/xq/ids', {
        method: 'POST',
        body: {Values: [id]}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0 && body.data ) {
          this.xq = body.data[0] || {Name: ''}
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  height: 100%;
  // background-color: #fff;
}
.content{
  padding-top: 100px;
}
.item{
  padding: 10px;
  font-size: 25px;
  margin: 10px auto;
  display: flex;
  align-items: center;
  span{
    flex: 1;
    border: solid 1px #ddd;
    padding: 10px;
    background-color: #fff;
    border-radius: 5px;
  }
}
.textarea{
  margin: 10px;
  text-align: center;
  textarea{
    width: 100%;
    font-size: 20px;
  }
}
.btn{
  width: 60%;
  color: #fff;
  background-color: #1eb504;
  font-size: 22px;
  padding: 10px;
  text-align: center;
  border-radius: 5px;
  margin: 10px auto;
}
</style>