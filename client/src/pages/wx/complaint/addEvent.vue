<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div :class='s.header'>
        <span :class='s.name' v-text='xq.Name'></span>
        XQ
      </div>
      <div :class='s.warn' v-text='warn'></div>
      <div :class='s.item'>
        <div :class='s.title'>Event Type:</div>
        <select v-model='selectedEventType' :class='s.value' @focus='onFocus'>
          <option disabled value="">请选择</option>
          <option v-for='type in eventTypes' :value='type.Type'>{{type.Type}}</option>
          </option>
        </select>
      </div>
      <div :class='s.item'>
        <div :class='s.title'>联系电话:</div>
        <input 
          :class='s.value'
          type="number"
          v-model='event.Tel'
          @focus='onFocus'>
      </div>
      <div :class='s.contentWrap'>
        <div :class='s.title'>Content</div>
        <div :class='s.value'>
          <textarea 
            rows='10'
            placeholder="请输入内容"
            v-model="event.Content" 
            @focus='onFocus'></textarea>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import fetchpm from '@/fetchpm'
export default{
  components: { ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: [{text:'上一步', event: null}],
        title: 'Fill in Event',
        rightBtns: [{text:'下一步', event: null}]
      },
      xq: {
        ID: '',
        Name: ''
      },
      event: {
        Content: '',
        StreetID:'',
        CommunityID: '',
        XQID: '',
        Complainant:'',//投诉人
        Type: '',//事件类型
        Tel:''
      },
      eventTypes: [],
      selectedEventType:"",
      warn:''
    }
  },
  mounted () {
    this.xq.ID = this.$route.query.XQID
    this.xq.Name = this.$route.query.XQName
    this.headerOptions.leftBtns[0].event = this.onReturn
    this.headerOptions.rightBtns[0].event = this.onNext
    this.fetchEventTypes()
  },
  methods: {
    onReturn () {
      this.$router.go(-1)
    },
    onNext () {
      this.event.StreetID = sessionStorage.getItem('cpStreetID')
      this.event.CommunityID = sessionStorage.getItem('cpCommunityID')
      this.event.XQID = this.xq.ID
      this.event.Type = this.selectedEventType
      if(!this.checkEvent()) return
      this.addEvent()
    },
    addEvent () {
      fetchpm(this, false, '/open/event/add', {
        method: 'POST',
        body: this.event
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('addEvent', body)
        if (body.error === 0) {
          this.$router.push({path: '/wx/complaint/uploadImage', query: {id: body.data}})
        }
      })
    },
    checkEvent () {
      if (this.event.Type === undefined || this.event.Type === '') {
        this.warn = '请选择事件类型'
        return false
      }
      if (this.event.Content === '') {
        this.warn = '请填写内容'
        return false
      }
      return true
    },
    onFocus () {
      this.warn = ''
    },
    fetchEventTypes () {
      fetchpm(this, false, '/open/wx/event/types',{
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventTypes', body)
        if (body.error !== 1){
          this.eventTypes = body.data || []
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  .body{
    margin-top: 80px;
    .header{
      font-size: 30px;
      padding: 20px 10px;
      font-weight: bold;
      color: #999;
      background-color: #fff;
      .name{
        color: #000;
      }
    }
    .warn{
      color: red;
      text-align: center;
      font-size: 20px;
    }
    .item{
      margin: 10px auto;
      font-size: 25px;
      display: flex;
      padding: 0px 10px;
      background-color: #fff;
      .title{
        padding: 15px 0;
        color: #999;
      }
      .value{
        padding: 15px 10px;
        border-bottom: solid 2px #20a0ff;
        border-top: solid 0px;
        border-left: solid 0;
        border-right: solid 0;
        flex: 1;
        font-size: 30px;
        background-color: transparent;
        
      }
    }
    .contentWrap{
      font-size:25px;
      margin-top: 20px;
      background-color: #fff;
      .title{
        margin: 10px;
        font-size: 30px;
        color: #999;
      }
      .value{
        width: 90%;
        margin: 0 auto;
        textarea{
          width: 100%;
          min-width: 50px;
          background-color: #f0f0f0;
        }
      }
    }
    .btn{
      background-color: #20a0ff;
      font-size: 30px;
      color: #fff;
      text-align: center;
      padding: 15px 0;
      margin: 20px;
    }
  }
}
</style>