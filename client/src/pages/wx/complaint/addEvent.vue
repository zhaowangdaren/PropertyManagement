<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div :class='s.header'>
        <span :class='s.name' v-text='xq.Name'></span>
        <span :class='s.tip'>小区</span>
      </div>
      <div :class='s.warn' v-text='warn'></div>
      <div :class='s.item'>
        <div :class='s.title'>投诉类型:</div>
        <!-- <select v-model='selectedEventType' :class='s.value' @focus='onFocus'>
          <option disabled value="">请选择</option>
          <option
            :class='s.option'
            v-for='type in eventTypes'
            :value='type.Type'>{{type.Type}}</option>
          </option>
        </select> -->
        <y-select :class='s.select' :selected.sync='selectedEventType' :options='eventTypes'></y-select>
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
        <div :class='s.title'>投诉内容</div>
        <div :class='s.value'>
          <textarea 
            rows='10'
            placeholder="请输入内容"
            v-model="event.Content" 
            @focus='onFocus'></textarea>
        </div>
      </div>
      <div :class='s.btn' @click='onNext'>下一步</div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import YSelect from '@/components/mobile/YSelect'
import fetchpm from '@/fetchpm'
export default{
  components: { ActionBar, YSelect },
  data () {
    return {
      headerOptions: {
        leftBtns: [{text:'上一步', event: null}],
        title: '填写投诉内容',
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
        Complainant:'',//投诉人,
        OpenID: '', // 投诉人open ID
        Type: '',//事件类型
        Tel:''
      },
      eventTypes: [],
      selectedEventType:{value:'', label: ''},
      warn:''
    }
  },
  mounted () {
    this.xq.ID = this.$route.query.XQID
    this.xq.Name = this.$route.query.XQName
    this.headerOptions.leftBtns[0].event = this.onReturn
    this.headerOptions.rightBtns[0].event = this.onNext
    this.event.OpenID = JSON.parse(sessionStorage.getItem('WXUser')).openid || ''
    this.fetchEventTypes()
  },
  methods: {
    onReturn () {
      this.$router.go(-1)
    },
    onNext () {
      this.event.StreetID = sessionStorage.getItem('StreetID')
      this.event.CommunityID = sessionStorage.getItem('CommunityID')
      this.event.XQID = this.xq.ID
      this.event.Type = this.selectedEventType.label
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
      if (this.event.StreetID === '' || !this.event.StreetID) {
        this.warn = '没有选择街道，请返回重新选择'
        return false
      }
      if (this.event.CommunityID === '' || !this.event.CommunityID) {
        this.warn = '没有选择社区，请返回重新选择'
        return false
      }
      if (this.event.XQID === '' || !this.event.XQID) {
        this.warn = '没有选择小区，请返回重新选择'
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
          this.eventTypes = (body.data || []).map(item => {
            return {
              value: item.Type,
              label: item.Type
            }
          })
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
        color: #565656;
      }
      .tip{
        font-size: 15px;
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
      .title{
        padding: 15px 0;
        color: #999;
      }
      .value{
        margin: 5px;
        padding: 15px 10px;
        border: solid 1px #999;
        flex: 1;
        font-size: 25px;
        background-color: #fff;
      }
    }
    .contentWrap{
      font-size:25px;
      margin-top: 20px;
      padding: 10px;
      .title{
        font-size: 25px;
        color: #999;
        margin-bottom: 5px;
      }
      .value{
        textarea{
          box-sizing: border-box;
          width: 96%;
          min-width: 50px;
          background-color: #fff;
          border: solid 1px #999;
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
.option{
  word-break: break-all;
}
.select{
  flex: 1;
  border: solid 1px #999;
  margin: 5px;
  font-size: 25px;
}
</style>