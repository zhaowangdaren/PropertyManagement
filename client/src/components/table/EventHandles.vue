<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        居民物业纠纷处理
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
            <template v-if='from === "gov"'>
              <td class='searchKey'>街道</td>
              <td>
                <el-select v-model="inputStreetID" filterable placeholder="全部" :class='s.elSelect'>
                  <el-option
                    key=''
                    label='全部'
                    value=''>
                  </el-option>
                  <el-option
                    v-for="item in allStreets"
                    :key="item.ID"
                    :label="item.Name"
                    :value="item.ID">
                  </el-option>
                </el-select>
              </td>
            </template>
            <td class='searchKey'>社区</td>
            <td>
              <el-select v-model="inputCommunityID" filterable placeholder="全部" :class='s.elSelect'>
                <el-option
                  key=''
                  label='全部'
                  value=''>
                </el-option>
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
                  key=''
                  label='全部'
                  value=''>
                </el-option>
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
            <th>OpenID</th>
            <th>所在社区</th>
            <th>所在小区</th>
            <th>事件状态</th>
            <th>事件等级</th>
            <th>事件类型</th>
            <th>操作</th>
          </tr>
          <tr v-for='(event, index) in events'>
            <td>
              <i :class='"iconfont icon-gaojing " + s.lv1' v-if='event.EventLevel === 1' ></i>
              <i :class='"iconfont icon-gaojing " + s.lv2' v-else-if='event.EventLevel === 2' ></i>
              <i :class='"iconfont icon-gaojing " + s.lv3' v-else-if='event.EventLevel === 3' ></i>
              <i v-else class="iconfont icon-gaojing"></i>
            </td>
            <td v-text='event.Index'></td>
            <td>{{ event.Time | filterTime}}</td>
            <td>{{event.OpenID}}</td>
            <td>{{communities[index] ? communities[index].Name : ''}}</td>
            <td>{{xqs[index] ? xqs[index].Name : ''}}</td>
            <td :style='{color: event.Status === 3 ? "#1eb504" : ""}'>{{event.Status | filterEventStatus }}</td>
            <td>{{event.EventLevel | filterEventLevel }}</td>
            <td v-text='event.Type'></td>
            <td>
              <el-button type="primary" icon="search" @click='toDetails(event)' class='view'>查看</el-button>
              <el-button v-if='from === "street" || from === ""' type="danger" icon="delete" @click='onDel(event)' class='del'>删除</el-button>
              <el-dialog
                title="请确认删除"
                :visible.sync='showDelDialog'>
                <div :class="s.delItem">事件编号:<span :class='s.value'>{{delEvent.Index}}</span></div>
                <div :class="s.delItem">事件类型:<span :class='s.value'>{{delEvent.Type}}</span></div>
                <div :class="s.delItem">投诉时间:<span :class='s.value'>{{delEvent.Time | filterTime}}</span></div>
                <div :class='s.bottomBtns'>
                  <el-button type='primary' @click='onDelSure' :loading='deling'>确 认</el-button>
                  <el-button @click='showDelDialog = false'>取 消</el-button>
                </div>
              </el-dialog>
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
  import { Message } from 'element-ui'
  export default {
    props: {
      editable: Boolean,
      from: {
        type: String,
        default: ''
      }
    },
    filters: {filterEventStatus, filterEventLevel, filterTime},
    data () {
      return {
        userStreetID: '',
        userType: 0,
        //user Input
        inputStreetID: '',
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
        eventStatus: [{value: 0, label:"居民提交"}, {value: -2, label: '已关闭'}, {value: -1, label:'已撤销'}, {value: 1, label: '处理中'}, {value: 2, label:'已处理待确认'}, {value: 3, label:'已解决'}, {value: 4, label:'未解决'}],
        communities: [],
        xqs:[],
        allStreets: [],
        allCommunities: [],
        allXQs: [],
        showDelDialog: false,
        delEvent: {
          Index: '',
          Type: '',
          Time: 0
        },
        deling: false
      }
    },
    mounted () {
      var user = JSON.parse(sessionStorage.getItem('user')) || {}
      this.userStreetID = user.StreetID
      this.userType = user.type
      this.onSearch()
      if (this.userType === 2 || this.userType === 1) this.fetchAllStreets()
      else if (this.userType === 3) this.fetchAllCommunitiesByStreetID(this.userStreetID)
      this.fetchEventTypes()
    },
    watch: {
      inputStreetID: function (value) {
        this.inputCommunityID = ''
        if (value === '') return
        if (this.from === 'gov'){
          this.fetchAllCommunitiesByStreetID(value)
        }
      },
      inputCommunityID: function (value) {
        this.inputXQID = ''
        if (value === '') return
        this.fetchAllXQByCommunityID(value)
      }
    },
    methods: {
      fetchAllStreets () {//获取所有街道
        fetchpm( this, true, '/pm/street', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          this.allStreets = body.data.streets
        })
      },
      onDel (event) {
        this.delEvent = event
        this.showDelDialog = true
      },
      onDelSure () {
        this.deling = true
        fetchpm(this, true, '/pm/event/del/index/' + this.delEvent.Index, {
          method: 'GET'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0 ) {
            this.onSearch()
            this.showDelDialog = false
            Message({type: 'success', message: '成功删除'})
          } else {
            Message({type: 'error', message: body.data})
          }
          this.deling = false
        })
      },
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
        if (this.userType === 3 && this.userStreetID !== '') result = {StreetID: this.userStreetID}
        else if ((this.userType === 1 || this.userType === 2) && this.inputStreetID !== '') result = {StreetID: this.inputStreetID}
        if (this.inputIndex !== '') result.Index = this.inputIndex
        if (this.inputEventLevel !== '' && this.inputEventLevel !== 0) result.EventLevel = this.inputEventLevel
        if (this.inputType !== '' && this.inputType !== 0) result.Type = this.inputType
        if (this.inputEventStatus !== '') result.Status = this.inputEventStatus
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
        if (this.from === 'gov') {
          this.$router.push({name: 'GovDetailEvent', params:{index: event.Index}})
          return
        }
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
.delItem{
  margin: 2px;
  display: flex;
  align-items: center;
  font-size: 20px;
  .value{
    flex: 1;
    background-color: #f1f3f6;
    border: solid 1px #ddd;
    border-radius: 2px;
    padding: 5px;
  }
}
.bottomBtns{
  margin-top: 10px;
  display: flex;
  justify-content: flex-end;
}
</style>
