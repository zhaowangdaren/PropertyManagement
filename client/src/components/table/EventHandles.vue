<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        Event handles<!-- 居民物业纠纷处理 -->
      </div>
      <div :class='s.body'>
        <table>
        <!-- 
          事件编号
          事件等级
          事件类型
          开始时间。从XXX到XXX
          事件状态
          所在社区
          所在小区
           -->
          <tr>
            <td :class='s.key'>Index</td>
            <td>
              <el-input v-model="inputIndex" placeholder="请输入事件编号"></el-input>
            </td>
            <td :class='s.key'>EventLevel</td>
            <td>
              <el-select v-model="inputEventLevel" filterable placeholder="全部">
                <el-option
                  v-for="item in eventLevels"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>EventType</td>
            <td>
              <el-select v-model="inputType" filterable placeholder="全部">
                <el-option
                  v-for="item in eventTypes"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </td>
          </tr>
          <tr>
            <td :class='s.key'>StartTime</td>
            <td>
              <el-date-picker
                v-model="inputStartTime"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
            </td>
            <td :class='s.key'>To</td>
            <td>
              <el-date-picker
                v-model="inputEndTime"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
            </td>
            <td :class='s.key'>EventStatus</td>
            <td>
              <el-select v-model="inputEventStatus" filterable placeholder="全部">
                <el-option
                  v-for="item in eventStatus"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </td>
          </tr>
          <tr>
            <td :class='s.key'>Community</td>
            <td>
              <el-select v-model="inputCommunityID" filterable placeholder="全部">
                <el-option
                  v-for="item in communities"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>XQs</td>
            <td>
              <el-select v-model="inputXQID" filterable placeholder="全部">
                <el-option
                  v-for="item in xqs"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td colspan="2" style="text-align:right;">
              <el-button type='primary' icon='search' @click='onSearch'>查询</el-button>
            </td>
          </tr>   
        </table>
        <!-- search result -->
        <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>Warning</th>
            <th>Index</th>
            <th>Time</th>
            <th>Community</th>
            <th>XQ</th>
            <th>Status</th>
            <th>EventLevel</th>
            <th>Type</th>
            <th>操作</th>
          </tr>
          <tr v-for='handle in events'>
            <td>
              <i class="iconfont icon-gaojing"></i>
            </td>
            <td v-text='handle.Index'></td>
            <td v-text='handle.Time'></td>
            <td v-text='handle.CommunityID'></td>
            <td v-text='handle.XQID'></td>
            <td>{{handle.Status | filterEventStatus }}</td>
            <td>{{handle.EventLevel | filterEventLevel }}</td>
            <td v-text='handle.Type'></td>
            <td>
              <el-button type="primary" icon="search" @click='toDetails(handle)'>查看</el-button>
            </td>
          </tr>
          <tr v-if='events.length == 0'>
            <td colspan="9">没有记录</td>
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
  export default {
    props: {
      editable: Boolean
    },
    filters: {filterEventStatus, filterEventLevel},
    data () {
      return {
        userStreetID: '',

        host: '//localhost:3000',
        //user Input
        inputCommunityID: '',
        inputXQID: '',
        inputIndex: '',
        inputType:'',
        inputEventLevel:'',
        inputStartTime:'',
        inputEndTime: '',
        inputEventStatus: '',

        events: [],
        eventLevels:[{value: 0, label:"全部"}, {value: 1, label:"特急"}, {value: 2, label: "加急"}, {value: 3, label: "急"}],
        eventTypes: [{value: 0,label: "全部"}, {value: 1, label:"type1"}, {value: 2, label:"type2"}, {value: 3, label: 'type3'}],
        eventStatus: [{value: 0, label:"全部"}, {value: 1, label: 'status1'}, {value: 2, label:'status2'}, {value: 3, label: 'status3'}, {value: 4, label:'status4'}],
        communities: [],
        xqs:[]
      }
    },
    mounted () {
      this.fetchEvents('')
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      this.userStreetID = user.StreetID
      this.fetchCommunitiesByStreetID(this.userStreetID)
    },
    watch: {
      inputCommunityID: function (value) {
        this.fetchXQByCommunityID(value)
      }
    },
    methods: {
      onSearch () {
        var search = this.formatInputData()
        fetchpm(this, true, '/pm/event/kvs', {
          method: 'POST',
          body: search
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onSearch', body)
          if (body.error !== 1) this.events = body.data || []
        })
      },
      formatInputData () {
        console.info('inputCommunityID', this.inputCommunityID)
        console.info('inputXQID', this.inputXQID)
        console.info('inputIndex', this.inputIndex)
        console.info('inputType', this.inputType)
        console.info('inputStartTime', this.inputStartTime)
        var result = {}
        if (this.inputIndex !== '') result.Index = this.inputIndex
        if (this.inputEventLevel !== '' && this.inputEventLevel !== 0) result.EventLevel = this.inputEventLevel
        if (this.inputType !== '' && this.inputType !== 0) result.Type = this.inputType
        if (this.inputEventStatus !== '' && this.inputEventStatus !== 0) result.EventStatus = this.inputEventStatus
        if (this.inputStartTime !== '') result.StartTime = this.inputStartTime.getTime() / 1000
        if (this.inputEndTime !== '') result.EndTime = this.inputEndTime.getTime() / 1000
        if (this.inputCommunityID !== '') result.CommunityID = this.inputCommunityID
        if (this.inputXQID !== '') result.XQID = this.inputXQID
        return result
      },
      fetchCommunitiesByStreetID (streetID) {
        fetchpm(this, true, '/pm/community/kvs', {
          method: 'POST',
          body: {streetID: streetID}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchCommunitiesByStreetID', body)
          this.communities = body.data
        })
      },
      fetchXQByCommunityID (communityID) {
        if (!communityID) return
        fetchpm(this, true, '/pm/xq/kvs', {
          method: 'POST',
          body: {communityID: communityID}
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchXQByCommunityName', body)
          if(body.error !== 0) console.error("Error: search fetchXQByCommunityName. Reason:" + body.data)
          else if (body.data !== null ) {
            this.xqs = body.data
          }
        })
      },
      fetchEvents (index) {
        if (!index ) {
          console.info('index', index)
        }
        fetchpm(this, true, '/pm/event', {
          method: 'POST',
          body: {index: index}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEvents', body)
          this.events = body.data
        })
      },
      toDetails (event) {
        console.info('toDetails', event.Index)
        this.$router.push({name: 'eventDetail', params:{index: event.Index}})
        // this.$router.push({path:'/street/detail/' + event.Index})
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  .content{
    border: solid 1px #4c87b9;
    margin: 10px;
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
        margin-top: 10px;
        th, td {
          padding: 5px;
          border: solid 1px #ddd;
          text-align: center;
        }
        th{
          background-color: #eee;
        } 
        tr{
          &:hover {
            background-color: #ddd;
          }
        }
        .key{
          background-color: #ddd;
        }
      }
    }    
  }
}
</style>