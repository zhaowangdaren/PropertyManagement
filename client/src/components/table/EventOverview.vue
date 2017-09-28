<template>
  <div :class="s.wrap">
    <div :class='s.tab'>
      <img src="~@/res/images/earth.png">
      投诉数据导出<!-- 物业信息管理 -->
    </div>
    <div :class='s.selectWrap'>
      <div :class="s.key">年份</div>
      <el-select v-model="queryEventOverview.Year" placeholder="请选择年份">
        <el-option
          v-for="item in years"
          :key="item.value"
          :label="item.label"
          :value="item.value">
        </el-option>
      </el-select>
      <div :class="s.key">季度</div>
      <el-select v-model="selectQuarter" placeholder="请选择年份">
        <el-option
          :key='0'
          label='全部'
          :value='0'>
        </el-option>
        <el-option
          v-for="item in quarters"
          :key="item"
          :label="item"
          :value="item">
        </el-option>
      </el-select>
      <div :class="s.key">月份</div>
      <el-select v-model="queryEventOverview.Month" placeholder="请选择月份">
        <el-option
          :key='0'
          label='全部'
          :value='0'>
        </el-option>
        <el-option
          v-for="item in months"
          :key="item"
          :label="item"
          :value="item">
        </el-option>
      </el-select>
      <el-button :class='s.searchBtn' type="primary" icon='search' @click='onSearch'>搜 索</el-button>
      <el-button type="success" @click='onExport'>数据导出</el-button>
    </div>
    <table>
      <caption v-if='tableTitle !== ""'>{{tableTitle}}</caption>
      <tr>
        <th>年份</th>
        <th>季度</th>
        <th>月份</th>
        <th>街道名</th>
        <th>总投诉</th>
        <th>未处理</th>
        <th>未处理占比</th>
        <th>已处理</th>
        <th>已处理占比</th>
        <th>未解决</th>
        <th>未解决占比</th>
        <th>已解决</th>
        <th>已解决占比</th>
      </tr>
      <tr v-for='eventOverview in eventOverviews'>
        <td>{{eventOverview.Year}}</td>
        <td>{{eventOverview.Quarter}}</td>
        <td>{{eventOverview.Month}}</td>
        <td>{{eventOverview.StreetName}}</td>
        <!-- <td v-if='user.type === 2 || user.type === 1'>{{eventOverview.StreetName}}</td> -->
        <!-- <td v-else>{{queryEventOverview.StreetName}}</td> -->
        <td>{{eventOverview.Sum}}</td>
        <td>{{eventOverview.Unhandle}}</td>
        <td v-if='eventOverview.Sum > 0'>{{Number(( eventOverview.Unhandle / eventOverview.Sum ) * 100).toFixed(2) + '%'}}</td>
        <td v-else>0%</td>
        <td>{{eventOverview.Handled}}</td>
        <td v-if='eventOverview.Sum > 0'>{{Number(( eventOverview.Handled / eventOverview.Sum ) * 100).toFixed(2) + '%'}}</td>
        <td v-else>0%</td>
        <td>{{eventOverview.Unsolved}}</td>
        <td v-if='eventOverview.Sum > 0'>{{Number(( eventOverview.Unsolved / eventOverview.Sum ) * 100).toFixed(2) + '%'}}</td>
        <td v-else>0%</td>
        <td>{{eventOverview.Solved}}</td>
        <td v-if='eventOverview.Sum > 0'>{{Number(( eventOverview.Solved / eventOverview.Sum ) * 100).toFixed(2) + '%'}}</td>
        <td v-else>0%</td>
      </tr>
      <tr>
        <td colspan="13" v-if='isLoading'>加载中...</td>
      </tr>
    </table>
    <el-pagination
      layout="total, prev, pager, next"
      @current-change='onChangePage'
      :page-size='queryEventOverview.PageSize'
      :total="sum">
    </el-pagination>
    <el-dialog
      title="投诉数据导出"
      :visible.sync='showExportDailog'>
      <div :class='s.loading' v-if='exportFile.isExporting'>数据导出中...</div>
      <div :class='s.loading' v-else>
        <a :href="exportFile.path">{{exportFile.fileName}}</a>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
import { Message } from 'element-ui'
import calculatQuart from '@/filters/calculatQuart'
export default {
  filters: { calculatQuart },
  props: {
    identity: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      // host: 'https://www.maszfglzx.com:3000',
      host: '//localhost:3000',
      user: {},
      tableTitle: '',
      years:[],
      streets: [],
      isLoading: false,
      eventOverviews: [
        {
          Year: 2017,
          Quarter: 3,
          Month: 9,
          StreetName: '街道名',
          Sum: 100, //总投诉
          Unhandle: 10, // 未处理
          Handled: 80, // 已处理
          Unsolved: 5, // 未解决
          Solved: 5 // 已解决
        }
      ],
      sum: 0,
      quarters:[ 1, 2, 3, 4],
      selectQuarter: 0,
      queryEventOverview: {
        PageNo: 0,
        PageSize: 10,
        StreetID: '',
        StreetName: '',
        Year: 0,
        Quarter: 3,
        Month: 9,
        UserType: 0
      },
      selectStreetIndex: 0,
      exportFile:{
        isExporting: false,
        path: '',
        fileName: ''
      },
      showExportDailog: false
    }
  },
  watch: {
    selectQuarter: function (value) {
      this.queryEventOverview.Month = 0
    }
  },
  computed: {
    months: function () {
      if(this.selectQuarter === 0) {
        return []
      }
      var months = []
      months.push(this.selectQuarter * 3 - 2)
      months.push(this.selectQuarter * 3 - 1)
      months.push(this.selectQuarter * 3)
      return months
    }
  },
  mounted () {
    this.user = JSON.parse(sessionStorage.getItem('user') || '{}') || {}
    this.queryEventOverview.UserType = this.user.type
    this.queryEventOverview.StreetID = this.user.StreetID || ''
    var today = new Date()
    this.queryEventOverview.Year = today.getFullYear()
    for (var i = 0; i < 20; i++) {
      this.years.push({
        value: this.queryEventOverview.Year - i,
        label: (this.queryEventOverview.Year - i) + '年'
      })
    }
    this.selectQuarter = calculatQuart(today.getMonth() + 1)
    this.onSearch()
  },
  methods: {
    initMonths(quarter) {
      this.months = []
      this.months.push(quarter * 3 - 2)
      this.months.push(quarter * 3 - 1)
      this.months.push(quarter * 3)
    },
    onSearch () {
      this.queryEventOverview.Quarter = this.selectQuarter
      this.fetchEventOverviews(this.queryEventOverview)
      // if (this.user.type === 2 || this.user.type === 1) this.onChangePage(this.selectStreetIndex + 1)
      // else {
      //   this.()
      // }
    },
    onExport() {
      this.showExportDailog = true
      this.exportFile.isExporting = true
      fetchpm(this, true, '/pm/event/overview/export/street', {
        method: 'POST',
        body: this.queryEventOverview
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onExport', body)
        if (body.error === 0) {
          this.exportFile.fileName = body.data.file
          this.exportFile.path = this.host + '/open/export/' + body.data.file
        } else {
          Message({type: "error", message: body.data})
        }
        this.exportFile.isExporting = false
      })
    },
    onChangePage(curPage) {
      this.queryEventOverview.PageNo = curPage - 1
      this.fetchEventOverviews(this.queryEventOverview)
    },
    fetchEventOverviews(query) {
      fetchpm(this, true, '/pm/event/overview', {
        method: 'POST',
        body: query
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventOverviews', body)
        if (body.error === 0) {
          this.eventOverviews = body.data.list || []
          console.info('eventOverviews', this.eventOverviews)
          this.sum = body.data.sum || 0
        } else {
          Message({type:'error', message: body.data})
        }
        this.isLoading = false
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  border: solid 1px #4c87b9;
  background-color: #fff;
  .tab{
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
}
.selectWrap{
  margin: 5px;
  display: flex;
  align-items: center;
  font-size: 20px;
  .key{
    margin: 0 10px;
  }
}
.searchBtn{
  margin-left: 10px;
}
.loading{
  text-align: center;
}
</style>
