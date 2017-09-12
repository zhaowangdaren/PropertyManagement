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
                  <el-button type='success' v-show='user.type === 3' @click='onNoticePM'>推送至物业公司</el-button>
                  <el-button type='primary' @click='showAddEventHandle = true'>询问</el-button>
                  <el-dialog
                    title='询问事件'
                    :visible.sync='showAddEventHandle'>
                    <add-event-handle
                      :eventIndex='event.Index' 
                      @succ='onAddEventHandleSucc'
                      @cancel='showAddEventHandle = false'>
                    </add-event-handle>
                  </el-dialog>
                </div>
                <div :class='s.btnWrap'>
                  <el-button 
                    type='success' 
                    v-if='user.type === 3'
                    :disabled='event.NoticeGov === 1'
                    @click='onNoticeGov'>推送至政府部门</el-button>
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
                    @click='onClose'>申请关闭</el-button>
                  <el-button
                    v-if='user.type == 2'
                    type='success'
                    @click='onTalkAbout'>
                    约谈物业公司
                  </el-button>
                  <el-button
                    v-if='user.type == 2'
                    type='success'
                    :disabled='(event.Status === -2 ? true : false)'
                    @click='onAgreeClose'>
                    同意关闭
                  </el-button>
                </div>
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
            <th>提交人类别</th>
            <th>提交人名</th>
            <th>提交时间</th>
            <th>处理内容</th>
            <th>操作</th>
          </tr>
          <tr v-for='handle in eventHandles' v-if='eventHandles && eventHandles.length > 0'>
            <td v-text=''>{{ handle.AuthorCategory | filterUserIdentity}}</td>
            <td v-text='handle.AuthorName'></td>
            <td >{{handle.Time | filterTime }}</td>
            <td v-text='handle.HandleInfo'></td>
            <td>
              <el-button type="primary" icon="search" @click='showHandleDetails = true'>详情</el-button>
              <el-dialog v-model="showHandleDetails" size="tiny" title='事件处理详情'>
                <details-event-handle :eventHandle='handle'></details-event-handle>
              </el-dialog>
            </td>
          </tr>
          <tr v-if='!eventHandles || eventHandles.length === 0'>
            <td colspan="5">无记录</td>
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
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import filterUserIdentity from '@/filters/filterUserIdentity'
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
    filters: {filterEventStatus, filterEventLevel, filterTime, filterUserIdentity},
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
        showAddEventHandle: false//询问
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
      onTalkAbout () {
        Message({type: "warning", message: '该功能正在开发中'})
      },
      onNoticeGov () {
        this.event.NoticeGov = 1
        this.updateEvent(this.event)
      },
      onNoticePM () {
        var eventHandle = {
          Index: this.event.Index,
          XQID: this.event.XQID,
          AuthorCategory: this.user.type,
          AuthorName: this.,
          HandleInfo: 'push2PM',
          Imgs: ''
        }

        fetchpm(this, true, '/pm/eventHandle/noticepm?index=' + this.event.Index + '&xqid=' + this.event.XQID, {
          method: 'GET'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error == 0) Message({type:'success', message: '推送成功'})
          else Message({type: 'error', message: body.data})
        })
      },
      onAgreeClose () {
        var temp = Object.assign({}, this.event)
        temp.Status = -2
        fetchpm(this, true, '/pm/event/update', {
          method: 'POST',
          body: temp
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0) {
            Message({type: 'success', message: '成功关闭'})
            this.event.Status = -2
          } else {
            Message({type: 'error', message: body.data})
          }
        })
      },
      onClose () {
        var temp = Object.assign({}, this.event)
        temp.RequestClose = 1
        fetchpm(this, true, '/pm/event/update', {
          method: 'POST',
          body: temp
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0) {
            Message({type: 'success', message: '申请成功'})
            this.event.RequestClose = 1
          } else {
            Message({type: 'error', message: body.data})
          }
        })

      },
      onAddEventHandleSucc (eventHandle) {
        this.eventHandles.push(eventHandle)
        this.event.Status = 2
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
          }
        })
      },
      onAduitLevelSucc (eventLevel) {
        this.event.EventLevel = eventLevel
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
          body: {KVs:{ Index: index}, PageNo: this.pageNo, PageSize: this.pageSize}
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
</style>
