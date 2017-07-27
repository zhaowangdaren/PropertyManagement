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
          <tr>
            <td :class='s.key'>Community</td>
            <td>
              <el-select v-model="searchCommunity" filterable placeholder="全部">
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
              <el-select v-model="searchXQ" filterable placeholder="全部">
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
          <tr v-for='item in pms'>
            <td>
              <i class="iconfont icon-gaojing"></i>
            </td>
            <td v-text='item.Index'></td>
            <td v-text='item.Time'></td>
            <td v-text='item.Community'></td>
            <td v-text='item.XQ'></td>
            <td>{{item.Status | filterEventStatus }}</td>
            <td>{{item.EventLevel | filterEventLevel }}</td>
            <td v-text='item.Type'></td>
            <td>
              <el-button type="primary" icon="search" @click='toDetails(item)'>查看</el-button>
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
        pms:[{}],
        communities: ['1','2'],
        xqs:['1','2'],
        searchCommunity:'',
        searchXQ: ''
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
        fetch(this.host + '/event', {
          method: 'POST',
          body: JSON.stringify({index: index})
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEvents', body)
          this.events = body.data
        })
      },
      toDetails (event) {
        console.info('toDetails', event.Index)
        // this.$router.push({path:'/street/detail/', params:{index: event.Index}})
        this.$router.push({path:'/street/detail/' + event.Index})
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
        margin-top: 10px;
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