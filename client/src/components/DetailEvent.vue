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
            <div v-if='eventImgs.length === 0'>暂无图片</div>
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
              <td :class='s.key'>警告类型</td>
              <td :class='s.value'><i class="iconfont icon-gaojing"></i></td>
            </tr>
            <tr>
              <td :class='s.key'>事件编号</td>
              <td :class='s.value' v-text='event.Index'></td>
            </tr>
            <tr>
              <td :class='s.key'>街道</td>
              <td :class='s.value' v-text='streetName'></td>
            </tr>
            <tr>
              <td :class='s.key'>社区</td>
              <td :class='s.value' v-text='communityName'></td>
            </tr>
            <tr>
              <td :class='s.key'>小区</td>
              <td :class='s.value' v-text='xqName'></td>
            </tr>
            <tr>
              <td :class='s.key'>投诉时间</td>
              <td :class='s.value' v-text='event.Time'></td>
            </tr>
            <tr>
              <td :class='s.key'>事件类型</td>
              <td :class='s.value' v-text='event.Type'></td>
            </tr>
            <tr>
              <td :class='s.key'>事件状态</td>
              <td :class='s.value'>{{event.Status | filterEventStatus }}</td>
            </tr>
            <tr>
              <td :class='s.key'>事件等级</td>
              <td :class='s.value'>{{event.EventLevel | filterEventLevel }}</td>
            </tr>
            <tr>
              <td :class='s.key'>操作</td>
              <td :class='s.value'>
                <el-button type='success' v-if='userType === 3' @click='onNoticePM'>推送至物业公司</el-button>
                <el-button type='primary' 
                  v-if='userType === 3'
                  @click='showAduitEventLevel = true'>审核等级</el-button>
                <el-dialog
                  v-if='userType === 3'
                  title='审核等级'
                  size='tiny'
                  :visible.sync='showAduitEventLevel'>
                  <aduit-event-level 
                    :event='event'
                    @cancel='showAduitEventLevel = false'
                    @succ='onAduitLevelSucc' />
                </el-dialog>
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
                <el-button
                  v-if='userType == 3'
                  type='danger'
                  @click='onClose'>申请关闭</el-button>
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
          <tr v-for='handle in eventHandles'>
            <td v-text='handle.AuthorCategory'></td>
            <td v-text='handle.AuthorName'></td>
            <td v-text='handle.Time'></td>
            <td v-text='handle.HandleInfo'></td>
            <td>
              <el-button type="primary" icon="search" @click='showHandleDetails = true'>详情</el-button>
              <el-dialog v-model="showHandleDetails" size="tiny" title='事件处理详情'>
                <details-event-handle :eventHandle='handle'></details-event-handle>
              </el-dialog>
            </td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import fetchpm from '@/fetchpm'
  import BasicDialog from '@/components/dialog/BasicDialog'
  import AduitEventLevel from '@/components/dialog/AduitEventLevel'
  import AddEventHandle from '@/components/dialog/AddEventHandle'
  import DetailsEventHandle from '@/components/dialog/DetailsEventHandle'
  export default {
    components: {
      BasicDialog, AduitEventLevel, AddEventHandle, DetailsEventHandle
    },
    filters: {filterEventStatus, filterEventLevel},
    props: {
      eventIndex: String
    },
    data () {
      return {
        host:'//www.maszfglzx.com/#',
        userType: 0,
        event: {
          Index: '',
          StreetID: '',
          CommunityID: '',
          XQID: '',
          Time: '',
          Type: '',
          Status: 0,
          EventLevel: 0,
          Imgs:''//以逗号为分隔符
        },
        eventImgs: [],
        streetName:'',
        communityName: '',
        xqName: '',
        showingImg: '',
        eventHandles:[],
        imgVisible: false,
        showHandleDetails: false,//查看详情
        showAduitEventLevel: false,
        showAddEventHandle: false//询问
      }
    },
    mounted () {
      console.info(this.$route)
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      this.userType = user.type
      this.fetchEvent(this.$route.params.index)
      this.fetchEventDetails(this.$route.params.index)
    },
    methods: {
      onNoticePM () {
        fetchpm(this, true, '/pm/eventHandle/noticepm?index=' + this.event.Index + '&xqid=' + this.event.XQID, {
          method: 'GET'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error == 0) Message({type:'success', message: '推送成功'})
        })
      },
      onClose () {
      },
      onAddEventHandleSucc (eventHandle) {
        this.eventHandles.push(eventHandle)
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
      fetchEventDetails (index) {
        if (!index ) {
          console.info('index', index)
          index = ''
        }
        fetchpm(this, true, '/pm/eventHandle/detail/' + index, {
          method: 'POST',
          body: {index: index}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEventHandle', body)
          if (body.error === 0) {
            this.eventHandles = body.data || []
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
          if (body.error == 0 && body.data.length > 0) {
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
          if (body.error == 0 && body.data.length > 0) {
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
          if (body.error == 0 && body.data.length > 0) {
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
</style>
