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
      <template v-if='user.type === 2 || user.type === 1'>
        <div :class="s.key">街道</div>
        <el-select v-model="selectStreetIndex" placeholder="请选择街道">
          <el-option
            v-for="(item, index) in streets"
            :key="item.ID"
            :label="item.Name"
            :value="index">
          </el-option>
        </el-select>
      </template>
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
        <td>{{eventOverview.Month | calculatQuart}}</td>
        <td>{{eventOverview.Month}}</td>
        <td v-if='user.type === 2 || user.type === 1'>{{streets[queryEventOverview.curPage].Name}}</td>
        <td v-else>{{queryEventOverview.StreetName}}</td>
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
      v-if='user.type === 2 || user.type === 1'
      layout="total, prev, pager, next"
      @current-change='onChangePage'
      :page-size='12'
      :total="12 * streets.length">
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
      host: 'https://www.maszfglzx.com:3000',
      user: {},
      tableTitle: '',
      years:[],
      streets: [],
      isLoading: false,
      eventOverviews: [
        {
          Year: 2017,
          Quart: 3,
          Month: 9,
          StreetName: '街道名',
          Sum: 100, //总投诉
          Unhandle: 10, // 未处理
          Handled: 80, // 已处理
          Unsolved: 5, // 未解决
          Solved: 5 // 已解决
        }
      ],
      queryEventOverview: {
        curPage: 0,
        StreetID: '',
        StreetName: '',
        Year: 0,
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
  mounted() {
    this.user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.queryEventOverview.UserType = this.user.type
    var todayYear = new Date().getFullYear()
    this.queryEventOverview.Year = todayYear
    for (var i = 0; i < 20; i++) {
      this.years.push({
        value: todayYear - i,
        label: (todayYear - i) + '年'
      })
    }
    if (this.user.type === 2 || this.user.type === 1) this.fetchAllStreets()
    else this.fetchStreet(this.user.StreetID)
  },
  methods: {
    onSearch () {
      if (this.user.type === 2 || this.user.type === 1) this.onChangePage(this.selectStreetIndex + 1)
      else {
        this.fetchStreetEventOverview()
      }
    },
    onExport() {
      this.showExportDailog = true

      this.exportFile.isExporting = true
      fetchpm(this, true, '/pm/event/overview/export', {
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
      this.queryEventOverview.curPage = curPage - 1
      this.queryEventOverview.StreetID = this.streets[this.queryEventOverview.curPage].ID
      this.queryEventOverview.StreetName = this.streets[this.queryEventOverview.curPage].Name
      this.fetchStreetEventOverview()
    },
    fetchStreetEventOverview() {
      this.isLoading = true
      this.eventOverviews = []
      for (var i = 0; i < 12; i++) {
        this.queryEventOverview.Month = i + 1
        let query = Object.assign({}, this.queryEventOverview)
        this.fetchEventOverviews(query)
      }
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
          this.eventOverviews.push(body.data || {})
        } else {
          Message({type:'error', message: body.data})
        }
        if (query.Month === 12) {
          this.eventOverviews = this.eventOverviews.sort((a, b) => {
            return a.Month - b.Month
          })
          this.isLoading = false
        }
      })
    },
    fetchAllStreets() {//获取所有街道名称
      fetchpm( this, true, '/pm/street', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        this.streets = body.data.streets || []
        this.onChangePage(1)
      })
    },
    fetchStreet(streetID) {
      fetchpm(this, true, '/pm/street/ids', {
        method: 'POST',
        body: {Values: [streetID]}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0) {
          this.queryEventOverview.StreetID = body.data[0].ID
          this.queryEventOverview.StreetName = body.data[0].Name
          this.fetchStreetEventOverview()
        }
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
