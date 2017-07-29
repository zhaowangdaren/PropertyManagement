<template>
  <div :class='s.wrap'>
    <div :class='s.street'>Street</div>
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
              <el-input v-model="eventHandle.Index" placeholder="请输入事件编号"></el-input>
            </td>
            <td :class='s.key'>EventLevel</td>
            <td>
              <el-select v-model="eventHandle.eventLevel" filterable placeholder="全部">
                <el-option
                  v-for="item in eventLevels"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>EventType</td>
            <td>
              <el-select v-model="eventHandle.eventType" filterable placeholder="全部">
                <el-option
                  v-for="item in eventTypes"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
          </tr>
          <tr>
            <td :class='s.key'>StartTime</td>
            <td>
              <el-date-picker
                v-model="eventHandle.startTime"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
            </td>
            <td :class='s.key'>To</td>
            <td>
              <el-date-picker
                v-model="eventHandle.endTime"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
            </td>
            <td :class='s.key'>EventStatus</td>
            <td>
              <el-select v-model="eventHandle.eventStatus" filterable placeholder="全部">
                <el-option
                  v-for="item in eventStatus"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
          </tr>
          <tr>
            <td :class='s.key'>Community</td>
            <td>
              <el-select v-model="eventHandle.community" filterable placeholder="全部">
                <el-option
                  v-for="item in communities"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>XQs</td>
            <td>
              <el-select v-model="eventHandle.xq" filterable placeholder="全部">
                <el-option
                  v-for="item in xqs"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
          </tr>   
        </table>
        <!-- search result -->
        <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>Index</th>
            <th>AuthorCategory</th>
            <th>AuthorName</th>
            <th>Time</th>
            <th>HandleInfo</th>
            <th>SubmitIndex</th>
            <th>操作</th>
          </tr>
          <tr v-for='handle in searchResult'>
            <td v-text='handle.Index'></td>
            <td v-text='handle.AuthorCategory'></td>
            <td v-text='handle.AuthorName'></td>
            <td v-text='handle.Time'></td>
            <td v-text='handle.HandleInfo'></td>
            <td v-text='handle.SubmitIndex'></td>
            <td>
              <el-button type="primary" icon="search" @click='toDetails(handle)'>查看</el-button>
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
    data () {
      return {
        host: '//localhost:3000',
        eventHandle: {
          eventLevel:'',
          eventType: '',
          startTime:'',
          endTime:''
        },
        searchResult:[{}],
        eventHandles: [],
        eventLevels:["全部", "特急", "加急", "急"],
        eventTypes: ["全部", "type1", "type2", 'type3'],
        eventStatus: ["全部", 'status1', 'status2', 'status3', 'status4'],
        communities: [],
        xqs:[]
      }
    },
    mounted () {
      this.fetchEvents('')
    },
    methods: {
      fetchEvents (index) {
        if (!index ) {
          console.info('index', index)
        }
        fetch(this.host + '/eventHandle/kvs', {
          method: 'POST',
          body: JSON.stringify({AuthorName: 'jack'})
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEvents', body)
          this.searchResult = body.data
        })
      },
      toDetails (event) {
        console.info('toDetails', event.Index)
        // this.$router.push({name: 'eventDetail', params:{index: event.Index}})
        // this.$router.push({path:'/street/detail/' + event.Index})
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
            // background-color: #ddd;
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