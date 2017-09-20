<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
      <div :class='s.up'>
        <div :class='s.warn' v-text='warn'></div>
        <div :class='s.item'>
          <div :class='s.key'>进度：</div>
          <div :class='s.value'>{{event.Status | filterEventStatus }}</div>
          <template v-if='identity === "user"'>
            <el-button 
              v-if='(event.Status >= 0 && event.Status < 2) || event.Status === 4'
              type='danger' 
              :class='s.revoke'
              @click='onRevoke'>撤销事件</el-button>
            <el-button 
              v-if='event.Status === 2 || event.Status === 4'
              type='success' 
              :class='s.revoke'
              @click='onSolved'>确认已解决</el-button>
            <el-button 
              v-if='event.Status === 2'
              type='success' 
              :class='s.revoke'
              @click='onUnsolved'>未解决</el-button>
            <el-button
              v-if='event.Status === 3 || event.Status === 4 || event.Status === -2'
              type='success'
              :class='s.revoke'
              @click='showAssess = true'>评 价</el-button>
          </template>
          <template v-if='identity === "pmuser"'>
            <el-button 
              v-if='event.Status !== 3'
              type='success' 
              :class='s.revoke'
              @click='onHandle'>已处理</el-button>
          </template>
        </div>
      </div>
      <div :class='s.event' v-for='handle in eventHandles'>
        <div :class='s.left'>
          <div :class='s.identity'>{{handle.AuthorCategory | filterUserIdentity}}</div>
          <div><span :class='s.key'>提交时间:</span><span>{{handle.Time | filterTime}}</span></div>
          <div><span :class='s.key'>提交内容:</span><span v-text='handle.HandleInfo'></span></div>
          <div><span :class='s.key'>提交类型:</span><span>{{ handle.HandleType | filterHandleType }}</span></div>
          <div v-if='handle.Imgs.length > 0'><span :class='s.key' >提交图片:</span><span class='iconfont icon-pic' @click='onShowImgs(handle.Imgs)'>点击查看图片</span></div>
          <div v-if='handle.HandleType === 3 && identity === "user" && event.Status !== 3'><div :class='s.replyBtn' @click='showReply = true'>回 复</div></div>
        </div>
      </div>
      
      <div :class='s.event' v-if='event.Index !== ""'>
        <div :class='s.left'>
          <div :class='s.identity'>居民</div>
          <div><span :class='s.key'>提交时间:</span><span>{{event.Time | filterTime}}</span></div>
          <div><span :class='s.key'>提交内容:</span><span v-text='event.Content'></span></div>
          <div><span :class='s.key'>联系电话:</span><span v-text='event.Tel'></span></div>
          <div v-if='event.Imgs.length > 0'><span :class='s.key'>提交图片:</span><span class='iconfont icon-pic' @click='onShowImgs(event.Imgs)'>点击查看图片</span></div>
        </div>
      </div>
    </div>
    <div>
      <div :class="s.callBtn" @click='showAdd = true' v-if='event.Status !== 3 && event.Status !== -1 && identity === "user"'>补 充</div>
      <div :class="s.callBtn" @click='showAdd = true' v-if='event.Status !== 3 && event.Status !== -1 && identity === "court"'>询 问</div>
      <y-dialog
        :visible.sync='showAdd'>
        <div :class='s.title'>补充信息</div>
        <div :class='s.textarea'><textarea v-model='addContent' placeholder="输入补充内容"></textarea></div>
        <div :class='s.replyBtn' @click='onAdd'>{{ adding ? "提交中。。。" : "提 交"}}</div>
      </y-dialog>
    </div>
    <y-dialog
      :visible.sync='showImgs'>
      <imgs :imgs='showingImgs' v-if='showImgs'></imgs>
    </y-dialog>
    <y-dialog
      :visible.sync='showReply'>
      <div :class='s.title'>回复</div>
      <div :class='s.textarea'><textarea v-model='replyContent' placeholder="输入回复内容"></textarea></div>
      <div :class='s.replyBtn' @click='onReply'>{{ replying ? "提交中。。。" : "提 交"}}</div>
    </y-dialog>
    <y-dialog
      :visible.sync='showAssess'>
      <assess @submit='onAssess'></assess>
    </y-dialog>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import YDialog from '@/components/mobile/YDialog'
import Imgs from '@/components/dialog/Imgs'
import Assess from '@/components/mobile/Assess'
import filterEventStatus from '@/filters/filterEventStatus'
import filterHandleType from '@/filters/filterHandleType'
import filterTime from '@/filters/filterTime'
import fetchpm from '@/fetchpm'
import filterUserIdentity from '@/filters/filterUserIdentity'
import { Message } from 'element-ui'
export default {
  components: { ActionBar, Imgs, YDialog, Assess },
  filters: { filterEventStatus, filterTime, filterUserIdentity, filterHandleType },
  data () {
    return {
      headerOptions: {
        title: '查看进度详情'
      },
      identity: '',
      eventIndex: '',
      event: {
        Status: 0,
        Imgs:'',
        OpenID: '',
        Index: ''
      },
      eventHandles: [],
      showImgs: false,
      showingImgs: '',
      showReply: false,
      replying: false,
      showAssess: false,
      assessing: false,
      showAdd: false,
      adding: false,
      addContent: '',
      warn:'',
      replyContent: ''
    }
  },
  mounted () {
    var WXUser = JSON.parse(sessionStorage.getItem('WXUser')) || {openid: null}
    this.event.OpenID = WXUser.openid
    if (this.event.OpenID === null || this.event.OpenID === '') this.fetchOpenID(this.$route.query.code).then(openid => {
      if (openid === null) {
        this.warn = '获取用户身份失败，请返回重试'
        return
      }
      this.event.OpenID = openid
      this.initData()
    })
    else this.initData()
    this.identity = this.$route.query.identity || ''
  },
  methods: {
    onAssess (assess) {
      if (this.assessing) return
      this.assessing = true
      var eventHandle = {
        Index: this.event.Index,
        XQID: this.event.XQID,
        OpenID: this.event.OpenID,
        HandleType: 6,
        HandleInfo: '评分：' + assess.score + '星。服务意见：' + assess.content
      }
      fetchpm(this, false, "/open/eventHandle/user/add", {
        method: 'POST',
        body: eventHandle
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onReply', body)
        if (body.error === 0) {
          Message({type: "success", message: '提交成功'})
          this.eventHandles.unshift(body.data)
          this.showAssess = false
        } else {
          Message({type: "error", message: body.data})
        }
        this.assessing = false
      })
    },
    onAdd () {
      if (this.addContent === '' || this.adding) return
      if (this.identity === 'user') this.onUserAdd()
      if (this.identity === 'court') this.onCourtAsk()
    },
    onCourtAsk () {
      this.adding = true
      var eventHandle = {
        Index: this.event.Index,
        XQID: this.event.XQID,
        OpenID: this.event.OpenID,
        HandleInfo: this.addContent
      }
      fetchpm(this, false, '/open/eventHandle/court/ask', {
        method: 'POST',
        body: eventHandle
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onCourtAsk', body)
        if (body.error === 0) {
          Message({type: "success", message: '提交成功'})
          this.eventHandles.unshift(body.data)
          this.showAdd = false
          this.addContent = ''
        } else {
          Message({type: "error", message: body.data})
        }
        this.adding = false
      })
    },
    onUserAdd () {
      this.adding = true
      var eventHandle = {
        Index: this.event.Index,
        XQID: this.event.XQID,
        OpenID: this.event.OpenID,
        HandleInfo: this.addContent
      }
      fetchpm(this, false, "/open/eventHandle/user/add", {
        method: 'POST',
        body: eventHandle
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onReply', body)
        if (body.error === 0) {
          Message({type: "success", message: '提交成功'})
          this.eventHandles.unshift(body.data)
          this.showAdd = false
          this.addContent = ''
        } else {
          Message({type: "error", message: body.data})
        }
        this.adding = false
      })
    },
    onReply () {
      if (this.replyContent === '' || this.replying) return
      this.replying = true
      var eventHandle = {
        Index: this.event.Index,
        XQID: this.event.XQID,
        OpenID: this.event.OpenID,
        HandleInfo: this.replyContent
      }
      fetchpm(this, false, "/open/eventHandle/user/reply", {
        method: 'POST',
        body: eventHandle
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onReply', body)
        if (body.error === 0) {
          Message({type: "success", message: '提交成功'})
          this.eventHandles.unshift(body.data)
          this.showReply = false
          this.replyContent = ''
        } else {
          Message({type: "error", message: body.data})
        }
        this.replying = false
      })
    },
    initData () {
      this.eventIndex = this.$route.query.index
      this.fetchEvent()
      this.fetchEventHandle()
    },
    onShowImgs (imgs) {
      this.showingImgs = imgs
      this.showImgs = true
      console.info('onShowImgs',this.showingImgs)
    },
    onHandle () {
      this.$router.push({name:'PMEvneHandle', query: {index: this.event.Index}})
    },
    onUnsolved () {
      var temp = Object.assign({}, this.event)
      temp.Status = 4
      fetchpm(this, false,'/open/event/update', {
        method: 'POST',
        body: temp
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onRevoke', body)
        if (body.error === 0 ) {
          this.event.Status = 4
        }
        else this.warn = body.data
      })
    },
    onSolved () {
      var temp = Object.assign({}, this.event)
      temp.Status = 3
      fetchpm(this, false,'/open/event/update', {
        method: 'POST',
        body: temp
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onRevoke', body)
        if (body.error === 0 ) {
          this.event.Status = 3
          this.warn = '投诉已解决'
        }
        else this.warn = body.data
      })
    },
    onRevoke () {
      this.event.Status = -1
      fetchpm(this, false,'/open/event/update', {
        method: 'POST',
        body: this.event
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onRevoke', body)
        if (body.error === 0 ) this.warn = '撤销成功'
        else this.warn = body.data
      })
    },
    fetchEvent () {
      fetchpm(this, false, '/open/event/id/' + this.eventIndex, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEvent', body)
        if (body.error === 0 ) {
          this.event = body.data.events[0] || {}
        }
      })
    },
    fetchEventHandle () {
      fetchpm(this, false, '/open/eventHandle/kvs', {
        method: 'POST',
        body: {Index: this.eventIndex}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventHandle', body)
        if (body.error === 0){
          this.eventHandles = (body.data || []).sort((a, b) => {
            return b.Time - a.Time
          })
        }
      })
    },
    fetchOpenID () {
      return fetchpm(this, false, '/open/wx/openid/' + this.$route.query.code, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchOpenID', body)
        if (body.error === 0) {
          sessionStorage.setItem('WXUser', JSON.stringify(body.data))
          return body.data.openid
        } else {
          return null
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.warn{
  text-align: center;
  font-size: 20px;
}
.content{
  padding-top: 80px;
  padding-bottom: 80px;
  font-size: 25px;
  .up{
    background-color: #fff;
    box-shadow: 0 1px 2px 2px rgba(0, 0, 0, 0.1);
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
  .event{
    margin: 10px;
    background-color: #fff;
    border-radius: 2px;
    padding: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    box-shadow: 1px 1px 2px 0 rgba(0, 0, 0, .25);
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
}
.replyBtn{
  background-color: #1eb504;
  color: #fff;
  width: 80%;
  text-align: center;
  padding: 10px;
  border-radius: 5px;
  margin: 10px auto;
  font-size: 25px;
}
.title{
  text-align: center;
  font-size: 30px;
  font-weight: bold;
}
.textarea{
  text-align: center;
  textarea{
    width: 90%;
    min-height: 200px;
    background-color: #f1f3f6;
    border: solid 1px #ddd;
    font-size: 25px;
  }

}
.callBtn{
  position: fixed;
  bottom: 10px;
  left: 50%;
  transform: translateX(-50%);
  padding: 10px 20px;
  font-size: 20px;
  font-weight: 500;
  background-color: rgb(25, 163, 24);
  border-radius: 100px;
  box-shadow: 0 2px 5px rgba(0,0,0,.25);
  color: #fff;
}
</style>