<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        代办事件<!-- 居民物业纠纷处理 -->
      </div>
      <div :class='s.body'>
        <!-- search result -->
        <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>警告类型</th>
            <th>事件编号</th>
            <th>开始时间</th>
            <th>所在社区</th>
            <th>所在小区</th>
            <th>事件状态</th>
            <th>事件等级</th>
            <th>事件类型</th>
            <th>操作</th>
          </tr>
          <tr v-for='(handle, index) in events'>
            <td>
              <i class="iconfont icon-gaojing"></i>
            </td>
            <td v-text='handle.Index'></td>
            <td>{{ handle.Time | filterTime}}</td>
            <td>{{communities[index] ? communities[index].Name : ''}}</td>
            <td>{{xqs[index] ? xqs[index].Name : ''}}</td>
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
  import filterTime from '@/filters/filterTime'
  import fetchpm from '@/fetchpm'
  export default {
    props: {
      editable: Boolean
    },
    filters: {filterEventStatus, filterEventLevel, filterTime},
    data () {
      return {
        userStreetID: '',
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
        xqs:[],
        allCommunities: [],
        allXQs: []
      }
    },
    mounted () {
      this.fetchEvents('')
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      this.userStreetID = user.StreetID
      this.fetchAllCommunitiesByStreetID(this.userStreetID)
    },
    watch: {
      inputCommunityID: function (value) {
        this.fetchAllXQByCommunityID(value)
      }
    },
    methods: {
      fetchAllCommunitiesByStreetID (streetID) {
        if ( !streetID || streetID == '') return null
        fetchpm( this, true, '/pm/community/kvs', {
          method: 'POST',
          body: {streetID: streetID}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchAllCommunitiesByStreetID', body)
          this.allCommunities = body.data
        })
      },
      fetchAllXQByCommunityID (communityID) {
        if ( !communityID || communityID == '') return null
        fetchpm( this, true, '/pm/xq/kvs', {
          method: 'POST',
          body: {communityID: communityID}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchAllXQByCommunityID', body)
          this.allXQs = body.data
        })
      },
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
      fetchCommunities (ids) {
        fetchpm(this, true, '/pm/community/ids', {
          method: 'POST',
          body: {values: ids}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchCommunitiesByStreetID', body)
          this.communities = body.data || []
        })
      },
      fetchXQs (ids) {
        if (!ids) return
        fetchpm(this, true, '/pm/xq/ids', {
          method: 'POST',
          body: {values: ids}
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchXQs', body)
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
          this.events = body.data || {}
          let communityIDs = this.events.map((item) => {
            return item.CommunityID
          })
          this.fetchCommunities(communityIDs)
          let xqIDs = this.events.map((item) => {
            return item.XQID
          })
          this.fetchXQs(xqIDs)
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
      .key{
        background-color: #ddd;
      }
    }    
  }
}
</style>