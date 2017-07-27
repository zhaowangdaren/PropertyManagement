<template>
  <div :class='s.wrap'>
    <div :class='s.street'>Event Details</div>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        Waiting<!-- 待办事件 -->
      </div>
      <div :class='s.body'>
        <div :class='s.left'>
          First time Image
          <div>
            <img src="~@/res/images/test.png" @click='imgVisible = true' :class='s.eventImg'>
            <el-dialog v-model="imgVisible" size="tiny">
              <img width="100%" src="~@/res/images/test.png" alt="">
            </el-dialog>
          </div>
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
              <td :class='s.key'>WarnType</td>
              <td :class='s.value'><i class="iconfont icon-gaojing"></i></td>
            </tr>
            <tr>
              <td :class='s.key'>Index</td>
              <td :class='s.value' v-text='event.Index'></td>
            </tr>
            <tr>
              <td :class='s.key'>Street</td>
              <td :class='s.value' v-text='event.Street'></td>
            </tr>
            <tr>
              <td :class='s.key'>Community</td>
              <td :class='s.value' v-text='event.Community'></td>
            </tr>
            <tr>
              <td :class='s.key'>XQ</td>
              <td :class='s.value' v-text='event.XQ'></td>
            </tr>
            <tr>
              <td :class='s.key'>Time</td>
              <td :class='s.value' v-text='event.Time'></td>
            </tr>
            <tr>
              <td :class='s.key'>Type</td>
              <td :class='s.value' v-text='event.Type'></td>
            </tr>
            <tr>
              <td :class='s.key'>Event Status</td>
              <td :class='s.value'>{{event.Status | filterEventStatus }}</td>
            </tr>
            <tr>
              <td :class='s.key'>Event Level</td>
              <td :class='s.value'>{{event.EventLevel | filterEventLevel }}</td>
            </tr>
          </table>
        </div>
        
      </div>
      
    </div>
    <div :class='s.handleContent'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        EventHandle<!-- 待办事件 -->
      </div>
      <div :class='s.body'>
        <table>
          <!-- 提交人类别  提交人用户名  提交时间  提交文本说明  操作 -->
          <tr>
            <th>AuthorCategory</th>
            <th>AuthorName</th>
            <th>Time</th>
            <th>HandleInfo</th>
            <th>操作</th>
          </tr>
          <tr v-for='handle in eventHandles'>
            <td v-text='handle.AuthorCategory'></td>
            <td v-text='handle.AuthorName'></td>
            <td v-text='handle.Time'></td>
            <td v-text='handle.HandleInfo'></td>
            <td>
              <el-button type="primary" icon="search" @click='showHandleDetails = true'>详情</el-button>
              <el-dialog v-model="showHandleDetails" size="tiny" title='Event Handle Details'>
                <div v-text='handle.HandleInfo'></div>
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

  export default {
    filters: {filterEventStatus, filterEventLevel},
    props: {
      eventIndex: String
    },
    data () {
      return {
        host: '//localhost:3000',
        event: {},
        eventHandles:[{}],
        imgVisible: false,
        showHandleDetails: false
      }
    },
    mounted () {
      console.info(this.$route)
      this.fetchEvents(this.$route.params.index)
    },
    methods: {
      fetchEvents (index) {
        if (!index ) {
          console.info('index', index)
          index = ''
        }
        fetch(this.host + '/eventHandle/' + index, {
          method: 'POST',
          body: JSON.stringify({index: index})
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEventHandle', body)
          this.eventHandles = body.data
        })
      }
    }
  }
</script>
<style lang="less" module='s'>
.wrap{
  margin: 10px;
  width: 100%;
  .street{
    background-color: #ddd;
    width: 100%;
    font-size: 25px;
    padding: 5px;
  }
  .content{
    border: solid 1px #4c87b9;
    margin-top: 50px;
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
      .left{
        padding: 20px;
        text-align: center;
        flex: 1;
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
      }
      .right{
        flex: 1;
        margin: 10px;
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
              background-color: #ddd;
            }
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
        th, td {
          padding: 5px;
          border: solid 1px #ddd;
          text-align: center;
        }
        th{
          background-color: #ddd;
        } 
        tr{
          &:hover {
            background-color: #ddd;
          }
        }
      }
    }
  }
}
</style>