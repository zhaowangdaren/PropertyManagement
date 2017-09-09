<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        居民物业纠纷处理<!-- 居民物业纠纷处理 -->
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
            <td class='searchKey'>事件编号</td>
            <td>
              <el-input v-model="inputIndex" placeholder="请输入事件编号"></el-input>
            </td>
            <td class='searchKey'>事件等级</td>
            <td>
              <el-select v-model="inputEventLevel" filterable placeholder="全部" :class='s.elSelect'>
                <el-option
                  v-for="item in eventLevels"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </td>
            <td class='searchKey'>事件类型</td>
            <td>
              <el-select v-model="inputType" filterable placeholder="全部" :class='s.elSelect'>
                <el-option
                  key=''
                  label='全部'
                  value=''>
                </el-option>
                <el-option
                  v-for="item in eventTypes"
                  :key="item.Type"
                  :label="item.Type"
                  :value="item.Type">
                </el-option>
              </el-select>
            </td>
          </tr>
          <tr>
            <td class='searchKey'>选择开始时间</td>
            <td>
              <el-date-picker
                :class='s.elSelect'
                v-model="inputStartTime"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
            </td>
            <td class='searchKey'>选择结束时间</td>
            <td>
              <el-date-picker
                :class='s.elSelect'
                v-model="inputEndTime"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
            </td>
            <td class='searchKey'>事件状态</td>
            <td>
              <el-select v-model="inputEventStatus" filterable placeholder="全部" :class='s.elSelect'>
                <el-option
                  key=''
                  label='全部'
                  value=''>
                </el-option>
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
            <td class='searchKey'>社区</td>
            <td>
              <el-select v-model="inputCommunityID" filterable placeholder="全部" :class='s.elSelect'>
                <el-option
                  v-for="item in allCommunities"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td class='searchKey'>小区</td>
            <td>
              <el-select v-model="inputXQID" filterable placeholder="全部" :class='s.elSelect'>
                <el-option
                  v-for="item in allXQs"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td colspan="2" style="text-align:right;">
              <el-button type='primary' icon='search' @click='onSearch' class='search'>查询</el-button>
            </td>
          </tr>   
        </table>
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
              <i :class='"iconfont icon-gaojing " + s.lv1' v-if='handle.EventLevel === 1' ></i>
              <i :class='"iconfont icon-gaojing " + s.lv2' v-else-if='handle.EventLevel === 2' ></i>
              <i :class='"iconfont icon-gaojing " + s.lv3' v-else-if='handle.EventLevel === 3' ></i>
              <i v-else class="iconfont icon-gaojing"></i>
            </td>
            <td v-text='handle.Index'></td>
            <td>{{ handle.Time | filterTime}}</td>
            <td>{{communities[index] ? communities[index].Name : ''}}</td>
            <td>{{xqs[index] ? xqs[index].Name : ''}}</td>
            <td :style='{color: handle.Status === 3 ? "#1eb504" : ""}'>{{handle.Status | filterEventStatus }}</td>
            <td>{{handle.EventLevel | filterEventLevel }}</td>
            <td v-text='handle.Type'></td>
            <td>
              <el-button type="primary" icon="search" @click='toDetails(handle)' class='view'>查看</el-button>
            </td>
          </tr>
          <tr v-if='events.length == 0'>
            <td colspan="9">没有记录</td>
          </tr>
        </table>
        <div :class='s.pageWrap'>
          <el-pagination
            layout="total, prev, pager, next"
            @current-change='onChangePage'
            :page-size='pageSize'
            :total="sum">
          </el-pagination>
          <div :class='s.pageTip'>其中<span>{{sum - events.length}}</span>条事件已关闭或已被用户撤销</div>
        </div>
        
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
        userType: 0,
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
        pageNo: 0,
        pageSize: 10,
        sum: 0,
        eventLevels:[{value: 0, label:"全部"}, {value: 1, label:"特急"}, {value: 2, label: "加急"}, {value: 3, label: "急"}],
        eventTypes: [],
        eventStatus: [{value: 0, label:"居民提交"}, {value: -2, label: '已关闭'}, {value: -1, label:'用户撤销'}, {value: 1, label: '已审核待处理'}, {value: 2, label:'已处理待确认'}, {value: 3, label:'已解决'}],
        communities: [],
        xqs:[],
        allCommunities: [],
        allXQs: []
      }
    },
    mounted () {
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      this.userStreetID = user.StreetID
      this.userType = user.type
      this.onSearch()
      this.fetchAllCommunitiesByStreetID(this.userStreetID)
      this.fetchEventTypes()
    },
    watch: {
      inputCommunityID: function (value) {
        this.fetchAllXQByCommunityID(value)
      }
    },
    methods: {
      onChangePage (curPage) {
        this.pageNo = curPage - 1
        this.onSearch()
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
      },
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
        fetchpm(this, true, '/pm/event/kvs/page', {
          method: 'POST',
          body: {KVs: search, PageNo: this.pageNo, PageSize: this.pageSize}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onSearch', body)
          if (body.error === 0) {
            this.events = (body.data.events || []).sort((a, b) => {
              return b.Time - a.Time
            })
            // this.events = (body.data.events || []).filter(item => {
            //   return item.Status >= 0
            // })
            this.sum = body.data.sum || 0
            var cids = this.events.map(item => {
              return item.CommunityID
            })
            this.fetchCommunities(cids)
            var xids = this.events.map(item => {
              return item.XQID
            })
            this.fetchXQs(xids)
          }
          else Message({type: 'error', message: body.data})
        })
      },
      formatInputData () {
        var result = {}
        if (this.userType === 3) result = {StreetID: this.userStreetID}
        if (this.inputIndex !== '') result.Index = this.inputIndex
        if (this.inputEventLevel !== '' && this.inputEventLevel !== 0) result.EventLevel = this.inputEventLevel
        if (this.inputType !== '' && this.inputType !== 0) result.Type = this.inputType
        if (this.inputEventStatus !== '' && this.inputEventStatus !== 0) result.Status = this.inputEventStatus
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
          console.info('fetchCommunities', body)
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
      .key{
        background-color: #f1f3f6;
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
.pageWrap{
  display: flex;
  align-items: center;
}
.pageTip{
  display: none;
  color: #999;
  span{
    color: #222;
    font-size: 20px;
  }
}
.elSelect{
  width: 100%;
}
</style>