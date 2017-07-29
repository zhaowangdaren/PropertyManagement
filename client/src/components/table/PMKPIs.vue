<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        PM KPI<!-- 物业考核查询-->
      </div>
      <div :class='s.body'>
        <table>
          <tr>
            <td :class='s.key'>Year</td>
            <td>
              <el-date-picker
                v-model="selectedYear"
                type="year"
                placeholder="选择年份">
              </el-date-picker>
            </td>
            <!-- 季度 -->
            <td :class='s.key'>Quarter</td>
            <td>
              <el-select v-model="selectedQuarter" filterable placeholder="select Quarter">
                <el-option
                  v-for="item in quarters"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>Company Name</td>
            <td>
              <el-input v-model="searchCompanyName" placeholder="请输入Company Name"></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.key'>Street</td>
            <td>
              <el-select v-model="selectedStreet" filterable placeholder="select Street">
                <el-option
                  v-for="item in streets"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>Community</td>
            <td>
              <el-select v-model="selectedCommunity" filterable placeholder="select Community">
                <el-option
                  v-for="item in communities"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>XQ</td>
            <td>
              <el-select v-model="selectedXQ" filterable placeholder="select XQ">
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
        <div :class='s.searchWrap'>
          <el-button type="primary">Cur Quarter</el-button>
          <el-button type="primary" icon="search">查询</el-button>
        </div>
        <!-- search result -->
        <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>Year</th>
            <th>Quarter</th>
            <th>Company Name</th>
            <th>Red Warning Num</th>
            <th>Yellow Warning Num</th>
            <th>Important Warning Num</th>
            <th>Score</th>
          </tr>
          <tr v-for='item in KPIs'>
            <td v-text='item.Year'></td>
            <td v-text='item.Quarter'></td>
            <td v-text='item.Name'></td>
            <td v-text='item.YWNo'></td>
            <td v-text='item.RWNo'></td>
            <td v-text='item.IWNo'></td>
            <td v-text='item.Score'></td>
          </tr>
          <tr v-if='KPIs.length === 0'>
            <td colspan="7">无记录</td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script type="text/javascript">
import filterEventStatus from '@/filters/filterEventStatus'
import filterEventLevel from '@/filters/filterEventLevel'

export default {
  filters: {filterEventStatus, filterEventLevel},
  data () {
    return {
      host: '//localhost:3000',
      selectedYear: '',
      quarters: ['全部', 1,2,3,4],
      selectedQuarter: '',
      searchCompanyName: '',
      selectedStreet:'',
      selectedCommunity: '',
      selectedXQ: '',
      xqs:[],
      communities: [],
      streets: [],
      KPIs:[],
      
    }
  },
  mounted () {
    this.fetchKPIs('')
  },
  methods: {
    fetchKPIs (name) {
      if (!name ) {
        console.info('name', name)
      }
      fetch(this.host + '/pmkpi', {
        method: 'POST',
        body: JSON.stringify({name: name})
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchKPIs', body)
        this.KPIs = body.data
      })
    },
    toDetails (event) {
      event.index = 1
      console.info('toDetails', event.Index)
      this.$router.push({path:'/street/detail/', params:{index: '1'}})
      // this.$router.push({path:'/street/detail/' + event.Index})
    }
  }
}
</script>

<style lang='less' module='s'>
.wrap{
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
      .searchWrap{
        display: flex;
        justify-content: flex-end;
        margin: 10px;
      }
      table{
        width: 100%;
        margin-top: 10px;
        font-size: 15px;
        color: #555;
        background-color: #fff;
        th, td {
          padding: 5px;
          border: solid 1px #ddd;
          text-align: center;
          font-size: 15px;
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