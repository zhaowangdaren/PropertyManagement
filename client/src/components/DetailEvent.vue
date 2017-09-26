<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        事件详情<!-- 待办事件 -->
      </div>
      <div :class='s.body'>
        <div :class='s.left'>
          用户上传
          <div :class='s.imgs'>
            <img v-for='img in eventImgs' :src='host + "/open/image/" + img' @click='onImg(img)' :class='s.eventImg'>
            <div v-show='eventImgs.length === 0'>暂无图片</div>
          </div>
          <el-dialog v-model="imgVisible" size="tiny">
            <img width="100%" :src='host + "/open/image/" + showingImg' alt="">
          </el-dialog>
        </div>
        <div :class='s.right'>
          <table>
          <!-- 警告类型：
  事件编号： 071320171902051386
  所在街道： 江东街道
  所在社区： 金瑞社区
  投诉小区： 金瑞新城
  投诉时间： 2017-07-13 19:02:05
  事件类型： 物业管理费
  事件状态： 已审核待处理
  事件等级： 特急 -->
            <tr>
              <td class='searchKey'>警告类型</td>
              <td :class='s.value'>
                <i :class='"iconfont icon-gaojing " + s.lv1' v-if='event.EventLevel === 1' ></i>
                <i :class='"iconfont icon-gaojing " + s.lv2' v-if='event.EventLevel === 2' ></i>
                <i :class='"iconfont icon-gaojing " + s.lv3' v-if='event.EventLevel === 3' ></i>
              </td>
            </tr>
            <tr>
              <td class='searchKey'>事件编号</td>
              <td :class='s.value' v-text='event.Index'></td>
            </tr>
            <tr>
              <td class='searchKey'>街道</td>
              <td :class='s.value' v-text='streetName'></td>
            </tr>
            <tr>
              <td class='searchKey'>社区</td>
              <td :class='s.value' v-text='communityName'></td>
            </tr>
            <tr>
              <td class='searchKey'>小区</td>
              <td :class='s.value' v-text='xqName'></td>
            </tr>
            <tr>
              <td class='searchKey'>投诉时间</td>
              <td :class='s.value'>{{event.Time | filterTime }}</td>
            </tr>
            <tr>
              <td class='searchKey'>投诉内容</td>
              <td :class='s.value'>{{event.Content }}</td>
            </tr>
            <tr>
              <td class='searchKey'>事件类型</td>
              <td :class='s.value' v-text='event.Type'></td>
            </tr>
            <tr>
              <td class='searchKey'>事件状态</td>
              <td :class='s.value'>{{event.Status | filterEventStatus }}</td>
            </tr>
            <tr>
              <td class='searchKey'>事件等级</td>
              <td :class='s.value'>{{event.EventLevel | filterEventLevel }}</td>
            </tr>
            <tr>
              <td class='searchKey'>操作</td>
              <td :class='s.value'>
                <div :class='s.btnWrap'>
                  <el-button type='success' v-show='user.type === 3' @click='showAddReason = true;reasonType = 1'>推送至物业公司</el-button>
                  <el-button type='primary' @click='onAsk' v-show="user.type !== 5">询问</el-button>
                </div>
                <div :class='s.btnWrap'>
                  <el-button
                    type='success'
                    v-if='user.type === 3'
                    :disabled='event.NoticeGov === 1'
                    @click='showAddReason = true;reasonType = 10'>推送至政府部门</el-button>
                  <el-button type='primary'
                    v-if='user.type === 3'
                    @click='showAduitEventLevel = true'>审核等级</el-button>
                  <el-dialog
                    v-if='user.type === 3'
                    title='审核等级'
                    size='tiny'
                    :visible.sync='showAduitEventLevel'>
                    <aduit-event-level
                      :event='event'
                      @cancel='showAduitEventLevel = false'
                      @succ='onAduitLevelSucc' />
                  </el-dialog>
                </div>
                <div :class='s.btnWrap'>
                  <el-button
                    v-if='user.type == 3'
                    type='danger'
                    :disabled='(event.RequestClose === 1 ? true : false)'
                    @click='showAddReason = true;reasonType = 8'>申请关闭</el-button>
                  <el-button
                    v-if='user.type == 2'
                    type='success'
                    :disabled='(event.TalkAbout === 1 ? true: false)'
                    @click='onTalkAbout'>
                    约谈物业公司
                  </el-button>
                  <el-dialog
                    :visible.sync='showTalkAboutDialog'
                    title='填写约谈内容'>
                    <textarea :class='s.textarea' v-model='editingTalkAboutContent'></textarea>
                    <div>
                      <el-button type='primary' @click='onTalkAboutSure'>确认</el-button>
                      <el-button @click='showTalkAboutDialog = false'>取消</el-button>
                    </div>
                  </el-dialog>
                  <el-button
                    v-if='user.type == 2'
                    type='success'
                    :disabled='((event.RequestClose === 1 && event.Status !== -2)? false : true)'
                    @click='onAgreeClose'>
                    同意关闭
                  </el-button>
                </div>
                <div :class='s.btnWrap'>
                  <el-button
                    v-if='user.type === 2'
                    type='success'
                    :disabled='event.ToCourt === 1 ? true : false'
                    :loading='toCourting'
                    @click='showAddReason = true;reasonType = 11'>
                    推送至法官
                  </el-button>
                </div>
                <template v-if='user.type === 5'>
                  <div :class="s.btnWrap">
                    <el-button
                      type='success'
                      :disabled='event.CourtAccepted === 1 ? true : false'
                      @click='onCourtAccepted'>
                      {{event.CourtAccepted === 1 ? "已受理" : "受理"}}
                    </el-button>
                  </div>
                  <div :class="s.btnWrap">
                    <el-button
                      type='success'
                      @click='onCourtAskPM'>
                      询问物业公司
                    </el-button>
                    <el-button
                      type='success'
                      @click='onCourtAskUser'>
                      询问投诉人
                    </el-button>
                  </div>
                </template>

              </td>
            </tr>
          </table>
        </div>
      </div>
    </div>
    <div :class='s.handleContent'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        事件处理<!-- 待办事件 -->
      </div>
      <div :class='s.body'>
        <table>
          <!-- 提交人类别  提交人用户名  提交时间  提交文本说明  操作 -->
          <tr>
            <th>事件编号</th>
            <th>人员类别</th>
            <th>姓名</th>
            <th>操作时间</th>
            <th>操作类型</th>
            <th>内容</th>
            <th>说明</th>
          </tr>
          <tr v-for='handle in eventHandles' v-if='eventHandles && eventHandles.length > 0'>
            <td>{{handle.Index}}</td>
            <td v-text=''>{{ handle.AuthorCategory | filterUserIdentity}}</td>
            <td >{{handle.AuthorCategory === 0 ? "居民" : handle.AuthorName}}</td>
            <td >{{handle.Time | filterTime }}</td>
            <td>{{handle.HandleType | filterHandleType}}</td>
            <td>{{handle.HandleInfo}}</td>
            <td>
              <el-button type="primary" icon="search" @click='onShowHandelDetail(handle)'>详情</el-button>
              <el-dialog :visible.sync="showHandleDetails" size="tiny" title='事件处理详情'>
                <details-event-handle :eventHandle='handleDetail'></details-event-handle>
              </el-dialog>
            </td>
          </tr>
          <tr v-if='!eventHandles || eventHandles.length === 0'>
            <td colspan="7">无记录</td>
          </tr>
        </table>
        <el-pagination
          layout="total, prev, pager, next"
          @current-change='onChangePage'
          :page-size='pageSize'
          :total="sum">
        </el-pagination>
      </div>
    </div>
    <el-dialog
      title='询问事件'
      :visible.sync='showAddEventHandle'>
      <add-event-handle
        :eventIndex='event.Index'
        :handleType='addEventHandleType'
        @succ='onAddEventHandleSucc'
        @cancel='showAddEventHandle = false'>
      </add-event-handle>
    </el-dialog>
    <el-dialog
      title='理由'
      :visible.sync='showAddReason'>
      <div :class='s.reason'>
        <textarea v-model='reasonContent' placeholder="请输入理由"></textarea>
      </div>
      <div :class='s.bottom'>
        <el-button type='primary' @click='onReasonSure'>确 定</el-button>
        <el-button @click='onReasonCancel'>取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import filterUserIdentity from '@/filters/filterUserIdentity'
  import filterHandleType from '@/filters/filterHandleType'
  import filterTime from '@/filters/filterTime'
  import fetchpm from '@/fetchpm'
  import BasicDialog from '@/components/dialog/BasicDialog'
  import AduitEventLevel from '@/components/dialog/AduitEventLevel'
  import AddEventHandle from '@/components/dialog/AddEventHandle'
  import DetailsEventHandle from '@/components/dialog/DetailsEventHandle'
  import { Message } from 'element-ui'
  export default {
    components: {
      BasicDialog, AduitEventLevel, AddEventHandle, DetailsEventHandle
    },
    filters: {filterEventStatus, filterEventLevel, filterTime, filterUserIdentity, filterHandleType},
    props: {
      eventIndex: String
    },
    data () {
      return {
        host:'//www.maszfglzx.com:3000',
        user: {
          type: 0,
        },
        event: {
          Index: '',
          StreetID: '',
          CommunityID: '',
          XQID: '',
          Time: '',
          Type: '',
          Status: 0,
          EventLevel: 0,
          RequestClose: 0,
          NoticeGov: 0,
          NoticePM: 0,
          ToCourt: 0,
          TalkAbout: 0,
          TalkAboutContent: '', //约谈内容
          Imgs:''//以逗号为分隔符
        },
        eventImgs: [],
        streetName:'',
        communityName: '',
        xqName: '',
        showingImg: '',
        eventHandles:[],
        pageNo: 0,
        pageSize: 10,
        sum: 0,
        imgVisible: false,
        showHandleDetails: false,//查看详情
        showAduitEventLevel: false,
        showAddEventHandle: false,//询问
        showTalkAboutDialog: false,
        showAddReason: false,
        reasonContent: '',//理由
        reasonType: 0, // 1
        editingTalkAboutContent: '',
        toCourting: false,
        addEventHandleType: 3, // 事件处理类型
        handleDetail: {}
      }
    },
    watch: {
      showAddReason: function (value) {
        if (!value) {
          this.reasonType = 0
          this.reasonContent = ''
        }
      }
    },
    mounted () {
      console.info(this.$route)
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      this.user = user
      this.fetchEvent(this.$route.params.index)
      this.fetchEventHandles(this.$route.params.index)
    },
    methods: {
      onShowHandelDetail (handle) {
        this.showHandleDetails = true
        this.handleDetail = handle
      },
      onReasonSure () {
        switch (this.reasonType) {
          case 1:
            this.onNoticePM()
            break
          case 10:
            this.onNoticeGov()
            break
          case 8:
            this.onClose()
            break
          case 11:
            this.onToCourt()
            break
        }
      },
      onReasonCancel () {
        this.showAddReason = false
        this.reasonType = 0
        this.reasonContent = ''
      },
      onCourtAccepted () {
        var eventHandle = {
          Index: this.event.Index,
          XQID: this.event.XQID,
          AuthorCategory: this.user.type,
          AuthorName: this.user.UserName,
          HandleInfo: '法官受理',
          handleType: 16,
          Imgs: ''
        }
        fetchpm(this, false, '/open/eventHandle/court/accept', {
          method: 'POST',
          body: eventHandle
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0) {
            Message({message: '受理成功', type:'success'})
            this.event.CourtAccepted = 1
          } else {
            Message({type: 'error', message: body.data})
          }
        }).catch(error => {
          Message({type: 'error', message: error})
        })
      },
      onAsk () { //普通的询问
        this.addEventHandleType = 3
        this.showAddEventHandle = true
      },
      onCourtAskPM () { // 法官询问物业
        this.addEventHandleType = 12
        this.showAddEventHandle = true
      },
      onCourtAskUser () { // 法官询问投诉人
        this.addEventHandleType = 13
        this.showAddEventHandle = true
      },
      onToCourt () {
        if (this.toCourting) return
        this.toCourting = true
        var eventHandle = {
          Index: this.event.Index,
          XQID: this.event.XQID,
          AuthorCategory: this.user.type,
          AuthorName: this.user.UserName,
          HandleInfo: this.reasonContent,
          Imgs: ''
        }

        fetchpm(this, true, '/pm/eventHandle/notice/court', {
          method: 'POST',
          body: eventHandle
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error == 0) {
            this.event.NoticeGov = 1
            Message({type:'success', message: '推送成功'})
            this.fetchEventHandles(this.event.Index)
          }
          else Message({type: 'error', message: body.data})
          this.toCourting = false
          this.showAddReason = false
        })
      },
      onTalkAbout () {
        // Message({type: "warning", message: '该功能正在开发中'})
        this.showTalkAboutDialog = true
      },
      onTalkAboutSure () {
        var eventHandle = {
          Index: this.event.Index,
          XQID: this.event.XQID,
          AuthorCategory: 2,
          AuthorName: this.user.UserName,
          HandleInfo: this.editingTalkAboutContent,
          Imgs: ''
        }
        this.govTalkAboutPM(eventHandle).then(body => {
          if (body.error === 0) {
            this.event.TalkAbout = 1
            Message({type:'success', message: '约谈成功'})
            this.fetchEventHandles(this.event.Index)
            this.showTalkAboutDialog = false
          } else {
            Message({type: 'error', message: body.data})
          }
        })
      },
      govTalkAboutPM (eventHandle) {
        return fetchpm(this, true, '/pm/eventHandle/govtalkpm', {
          method: 'POST',
          body: eventHandle
        }).then(resp => {
          return resp.json()
        })
      },
      onNoticeGov () {
        var eventHandle = {
          Index: this.event.Index,
          XQID: this.event.XQID,
          AuthorCategory: this.user.type,
          AuthorName: this.user.UserName,
          HandleInfo: this.reasonContent,
          Imgs: ''
        }

        fetchpm(this, true, '/pm/eventHandle/notice/gov', {
          method: 'POST',
          body: eventHandle
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error == 0) {
            this.event.NoticeGov = 1
            Message({type:'success', message: '推送成功'})
            this.fetchEventHandles(this.event.Index)
          }
          else Message({type: 'error', message: body.data})
          this.showAddReason = false
        })
      },
      onNoticePM () {
        var eventHandle = {
          Index: this.event.Index,
          XQID: this.event.XQID,
          AuthorCategory: this.user.type,
          AuthorName: this.user.UserName,
          HandleInfo: this.reasonContent,
          Imgs: ''
        }

        fetchpm(this, true, '/pm/eventHandle/noticepm', {
          method: 'POST',
          body: eventHandle
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error == 0) {
            this.event.ToCourt = 1
            Message({type:'success', message: '推送成功'})
            this.fetchEventHandles(this.event.Index)
          }
          else Message({type: 'error', message: body.data})
          this.showAddReason = false
        })
      },
      onAgreeClose () {
        var eventHandle = {
          Index: this.event.Index,
          XQID: this.event.XQID,
          AuthorCategory: this.user.type,
          AuthorName: this.user.UserName,
          HandleInfo: '同意关闭',
          HandleType: 9,
          EventLevel: this.event.EventLevel,
          Imgs: ''
        }
        fetchpm(this, true, '/pm/eventHandle/agree/close', {
          method: 'POST',
          body: eventHandle
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0) {
            Message({type: 'success', message: '关闭成功'})
            this.event.Status = -2
            this.fetchEventHandles(this.event.Index)
          } else {
            Message({type: 'error', message: body.data})
          }
        })
      },
      onClose () {
        var eventHandle = {
          Index: this.event.Index,
          AuthorCategory: this.user.type,
          AuthorName: this.user.UserName,
          HandleInfo: this.reasonContent,
          HandleType: 8,
          EventLevel: parseInt(this.event.EventLevel),
          Imgs: ''
        }
        fetchpm(this, true, '/pm/eventHandle/request/close', {
          method: 'POST',
          body: eventHandle
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0) {
            Message({type: 'success', message: '申请成功'})
            this.event.RequestClose = 1
            this.fetchEventHandles(this.event.Index)
          } else {
            Message({type: 'error', message: body.data})
          }
          this.showAddReason = false
        })
        // var temp = Object.assign({}, this.event)
        // temp.RequestClose = 1
        // fetchpm(this, true, '/eventHandle/request/close', {
        //   method: 'POST',
        //   body: temp
        // }).then(resp => {
        //   return resp.json()
        // }).then(body => {
        //   if (body.error === 0) {
        //     Message({type: 'success', message: '申请成功'})
        //     this.event.RequestClose = 1
        //   } else {
        //     Message({type: 'error', message: body.data})
        //   }
        // })
      },
      onAddEventHandleSucc (eventHandle) {
        this.eventHandles.unshift(eventHandle)
        if (this.event.Status >= 0) this.event.Status = 1 //处理中
        this.updateEvent(this.event)
      },
      updateEvent (event) {
        fetchpm(this, true, '/pm/event/update', {
          method: 'POST',
          body: event
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onSave', body)
          if (body.error !== 0) {
            Message({type: 'error', message: body.data})
            this.event = event
          }
        })
      },
      onAduitLevelSucc (eventLevel) {
        this.event.EventLevel = parseInt(eventLevel)
        this.fetchEventHandles(this.event.Index)
      },
      onHandle (eventIndex) {//事件处理

      },
      onImg (img) {
        this.showingImg = img
        this.imgVisible = true
      },
      fetchEvent (eventIndex) {
        fetchpm(this, true, '/pm/event/kvs', {
          method: 'POST',
          body: {index: eventIndex}
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchEvent', body)
          if (body.error == 0 && body.data.length > 0) {
            this.event = body.data[0]
            this.initEventImgs(this.event.Imgs)
            this.fetchStreet(this.event.StreetID)
            this.fetchCommunity(this.event.CommunityID)
            this.fetchXQ(this.event.XQID)
          }
        })
      },
      initEventImgs (imgs) {
        this.eventImgs = imgs.split(',')
        if (this.eventImgs[this.eventImgs.length - 1] == '') this.eventImgs = this.eventImgs.slice(0, this.eventImgs.length - 1)
      },
      onChangePage (curPage) {
        this.pageNo = curPage - 1
        this.fetchEventHandles(this.event.Index)
      },
      fetchEventHandles (index) {
        if (!index ) {
          console.info('index', index)
          index = ''
        }
        fetchpm(this, true, '/pm/eventHandle/kvs/page', {
          method: 'POST',
          body: { KVs:{ Index: index}, PageNo: this.pageNo, PageSize: this.pageSize }
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEventHandles', body)
          if (body.error === 0) {
            this.eventHandles = body.data.eventHandles || []
            this.sum = body.data.sum || 0
          }
        })
      },
      fetchStreet (id) {
        fetchpm(this, true, '/pm/street/ids', {
          method: 'POST',
          body: {values: [id]}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchStreet', body)
          if (body.error == 0 && body.data && body.data.length > 0) {
            this.streetName = body.data[0].Name
          }
        })
      },
      fetchXQ (id) {
        fetchpm(this, true, '/pm/xq/ids', {
          method: 'POST',
          body: {values: [id]}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchXQ', body)
          if (body.error == 0 && body.data && body.data.length > 0) {
            this.xqName = body.data[0].Name
          }
        })
      },
      fetchCommunity (id) {
        fetchpm(this, true, '/pm/community/ids', {
          method: 'POST',
          body: {values: [id]}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchCommunity', body)
          if (body.error == 0 && body.data && body.data.length > 0) {
            this.communityName = body.data[0].Name
          }
        })
      }
    }
  }
</script>
<style lang="less" module='s'>
.reason{
  text-align: center;
  textarea{
    margin: 0 auto;
    width: 80%;
    min-height: 200px;
  }
}
.wrap{
  margin: 10px;
  // width: 100%;
  .street{
    background-color: #ddd;
    width: 100%;
    font-size: 25px;
    padding: 5px;
  }
  .content{
    border: solid 1px #4c87b9;
    margin-top: 10px;
    .title{
      color: #fff;
      font-size: 20px;
      background: #4c87b9;
      padding: 10px;
      display: flex;
      align-items: center;
      color: #fff;
      font-size: 30px;
      font-family: "华文行楷";
      background: #4c87b9;
      background: -webkit-gradient(linear,center top, center bottom,from(#ff0000), to(#000000));
      img{
        width: 20px;
        margin-right: 10px;
      }
    }
    .body{
      display: flex;
      width: 100%;
      font-size: 18px;
      background-color: #fff;
      .left{
        padding: 20px;
        text-align: center;
        flex: 1;
        .imgs{
          overflow-y: scroll;
          height: 300px;
          .eventImg{
            overflow: hidden;
            background-color: #fff;
            border: 1px solid #c0ccda;
            border-radius: 6px;
            box-sizing: border-box;
            width: 148px;
            height: 148px;
            margin: 10px 8px 8px 0;
            display: inline-block;
          }
          div{
            color: #999;
            margin: 10px;
          }
        }
      }
      .right{
        flex: 2;
        margin: 10px;
        display: flex;
        align-items: center;
        table{
          color: #555;
          width: 100%;
          td {
            padding: 5px;
            border: solid 1px #ddd;
            text-align: center;
          }
          tr{
            &:hover {
              background-color: #f0f0f0;
            }
          }
          .key{
            background-color: #f0f0f0;
          }
          .value{
            width: 60%;
          }
        }
      }
    }
  }
  .handleContent{
    border: solid 1px #4c87b9;
    margin-top: 50px;
    background-color: #fff;
    .title{
      color: #fff;
      font-size: 20px;
      background: #4c87b9;
      padding: 10px;
      display: flex;
      align-items: center;
      color: #fff;
      font-size: 30px;
      font-family: "华文行楷";
      background: #4c87b9;
      background: -webkit-gradient(linear,center top, center bottom,from(#ff0000), to(#000000));
      img{
        width: 20px;
        margin-right: 10px;
      }
    }
    .body{
      margin: 10px;
      table{
        width: 100%;
        font-size: 15px;
        color: #555;
        background-color: #fff;
        th, td {
          padding: 5px;
          border: solid 1px #ddd;
          text-align: center;
        }
        th{
          background-color: #f0f0f0;
        }
        tr{
          &:hover {
            background-color: #f0f0f0;
          }
        }
      }
    }
  }
}
.lv3{
  color: #4c87b9;
}
.lv2{
  color: orange;
}
.lv1{
  color: red;
}
.btnWrap{
  margin: 5px auto;
}
.textarea{
  min-width: 50%;
  max-width: 90%;
  min-height: 100px;
}
.bottom{
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
