<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        物业考核查询<!-- 物业考核查询-->
      </div>
      <div :class='s.body'>
        <table>
          <tr>
            <td class='searchKey'>年份</td>
            <td>
              <el-date-picker
                :class='s.elSelect'
                v-model="selectedYear"
                type="year"
                placeholder="选择年份">
              </el-date-picker>
            </td>
            <!-- 季度 -->
            <td class='searchKey'>季度</td>
            <td>
              <el-select v-model="selectedQuarter" filterable placeholder="选择季度" :class='s.elSelect'>
                <el-option
                  v-for="item in quarters"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </td>
            <td class='searchKey'>街道</td>
            <td>
              <el-select v-model="selectedStreetID" filterable placeholder="请选择街道" :class='s.elSelect'>
                <el-option
                  key=''
                  value=''
                  label='全部'></el-option>
                <el-option
                  v-for="item in streets"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <div :class='s.hidden'>
              <td class='searchKey'>物业公司名</td>
              <td>
                <el-input v-model="searchCompanyName" placeholder="请输入物业公司名"></el-input>
              </td>
            </div>
          </tr>
          <div :class='s.hidden'>
            <tr>
              <td class='searchKey'>街道</td>
              <td>
                <el-select v-model="selectedStreetID" filterable placeholder="请选择街道" :class='s.elSelect'>
                  <el-option
                    v-for="item in streets"
                    :key="item.ID"
                    :label="item.Name"
                    :value="item.ID">
                  </el-option>
                </el-select>
              </td>
              <td class='searchKey'>社区</td>
              <td>
                <el-select v-model="selectedCommunityID" filterable placeholder="请选择社区" :class='s.elSelect'>
                  <el-option
                    v-for="item in communities"
                    :key="item.ID"
                    :label="item.Name"
                    :value="item.ID">
                  </el-option>
                </el-select>
              </td>
              <td class='searchKey'>小区</td>
              <td>
                <el-select v-model="selectedXQID" filterable placeholder="请选择小区" :class='s.elSelect'>
                  <el-option
                    v-for="item in xqs"
                    :key="item.ID"
                    :label="item.Name"
                    :value="item.ID">
                  </el-option>
                </el-select>
              </td>
            </tr>
          </div>
        </table>
        <div :class='s.searchWrap'>
          <el-button type="primary" @click='onExport'>数据导出</el-button>
          <el-button type="primary" @click='onCurQuarter' class='view'>当前季度</el-button>
          <el-button type="primary" icon="search" @click='onBTSearch' class='search'>查询</el-button>
        </div>
        <!-- search result -->
        <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>年份</th>
            <th>季度</th>
            <th>物业公司名</th>
            <th>小区名</th>
            <th>黄色警告次数</th>
            <th>红色警告次数</th>
            <th>重大事件警告次数</th>
            <th>其他</th>
            <th>得分</th>
            <th v-if='from === "gov"'>操作</th>
          </tr>
          <tr v-if='loading'>
            <td colspan="10" :class='s.loading'>加载中...</td>
          </tr>
          <template v-else>
            <tr v-for='(item,index) in KPIs'>
              <td v-text='item.Year'></td>
              <td v-text='item.Quarter'></td>
              <td v-text='item.PMName'></td>
              <td>{{item.XQName}}</td>
              <td v-text='item.YWNo'></td>
              <td v-text='item.RWNo'></td>
              <td v-text='item.IWNo'></td>
              <td v-text='item.Other'></td>
              <td >{{ item | calculatPMKPI }}</td>
              <td v-if='from === "gov"'>
                <el-button
                  class='edit'
                  @click='onEdit(item)'
                  type='primary'
                  icon='edit'>
                  编辑
                </el-button>
                <el-dialog
                  :visible.sync='showEditDialog'
                  title='编辑考核其他扣分项'>
                  <div :class='s.editKPIOther'>扣除物业其他项分数：<input type="" name="" v-model='kpiOther'></div>
                  <div :class='s.editBtns'>
                    <el-button type="primary" @click='onEditSure'>确定</el-button>
                    <el-button @click='showEditDialog = false'>取消</el-button>
                  </div>
                </el-dialog>
              </td>
            </tr>
            <tr v-if='KPIs.length === 0'>
              <td colspan="10">无记录</td>
            </tr>
          </template>
        </table>
        <el-pagination
          layout="total, prev, pager, next"
          @current-change='onChangePage'
          :page-size='queryPMKPI.PageSize'
          :total="sum">
        </el-pagination>
      </div>
    </div>
    <el-dialog
      title="物业考核数据导出"
      :visible.sync='showExportDailog'>
      <div :class='s.loading' v-if='exportFile.isExporting'>数据导出中...</div>
      <div :class='s.loading' v-else>
        <a :href="exportFile.path">{{exportFile.fileName}}</a>
      </div>
    </el-dialog>
  </div>
</template>

<script type="text/javascript">
import filterEventStatus from '@/filters/filterEventStatus'
import filterEventLevel from '@/filters/filterEventLevel'
import calculatPMKPI from '@/filters/calculatPMKPI'
import fetchpm from '@/fetchpm'
import { Message } from 'element-ui'

export default {
  filters: {filterEventStatus, filterEventLevel, calculatPMKPI},
  props: {
    from: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      // host: '//localhost:3000',
      host: 'https://www.maszfglzx.com:3000',
      loading: false,
      quarters: [
        {value:1, label: '1'},
        {value:2, label: '2'},
        {value:3, label: '3'},
        {value:4, label: '4'}
      ],
      selectedYear: 0,
      selectedQuarter: 0,
      searchCompanyName: '',
      selectedStreetID:'',
      selectedCommunityID: '',
      selectedXQID: '',
      KPIs: [],
      xqs:[],
      pms:[],
      communities: [],
      streets: [],
      sum: 0,
      queryPMKPI: {
        StreetID: '',
        XQID: '',
        Year: 0,
        Quarter: 0,
        PageNo: 0,
        PageSize: 10
      },
      showEditDialog: false,
      kpiOther: 0,
      editingKPI: null,
      showExportDailog: false,
      exportFile: {
        isExporting: false,
        path: '',
        fileName: ''
      }

    }
  },
  mounted () {
    this.selectedYear = new Date()
    var thisMonth = new Date().getMonth()
    this.selectedQuarter = parseInt((thisMonth + 1) / 3) + (((thisMonth + 1) % 3) > 0 ? 1 : 0)
    this.fetechAllStreets()
    this.queryPMKPI.StreetID = this.selectedStreetID
    this.queryPMKPI.Year = this.selectedYear.getFullYear()
    this.queryPMKPI.Quarter = this.selectedQuarter
    this.onSearch(this.queryPMKPI)
  },
  methods: {
    onExport() {
      this.exportFile.isExporting = true
      this.showExportDailog = true
      fetchpm(this, true, '/pm/pm/kpi/export/street', {
        method: 'POST',
        body: {
          Year: this.selectedYear.getFullYear(),
          Quarter: this.selectedQuarter,
          StreetID: this.selectedStreetID
        }
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0) {
          this.exportFile.fileName = body.data.file
          this.exportFile.path = this.host + '/open/export/' + body.data.file
        } else {
          Message({type: "error", message: body.data})
        }
        this.exportFile.isExporting = false
      })
    },
    onBTSearch () {
      var queryPM = {
        StreetID: this.selectedStreetID,
        XQID: '',
        Year: this.selectedYear.getFullYear(),
        Quarter: this.selectedQuarter,
        PageNo: 0,
        PageSize: 10
      }
      this.onSearch(queryPM)
    },
    onEdit (item) {
      this.kpiOther = item.Other
      this.showEditDialog = true
      this.editingKPI = item
    },
    onEditSure () {
      if (this.kpiOther.length == 0 || isNaN(this.kpiOther)) {
        return
      }
      this.editingKPI.Other = parseInt(this.kpiOther)
      fetchpm(this, true, '/pm/pmkpi/update/other', {
        method: 'POST',
        body: this.editingKPI
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onEditSure', body)
        if (body.error === 0) this.showEditDialog = false
      })
    },
    onSearch (query) {
      if (this.loading) return
      this.loading = true
      this.fetchPMKPIs(query)
    },
    fetchPMKPIs (query) {
      fetchpm(this, true, '/pm/pmkpi/query', {
        method: 'POST',
        body: query
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPMKPIs', body)
        if (body.error === 0) {
          this.KPIs = body.data.list || []
          this.sum = body.data.sum || 0
        }
        this.loading = false
      })
    },
    onChangePage (curPage) {
      this.queryPMKPI.PageNo = curPage - 1
      this.onSearch(this.queryPMKPI)
    },
    onCurQuarter () {
      this.selectedYear = new Date()
      this.selectedQuarter = parseInt((this.selectedYear.getMonth() + 1 ) / 3) + (((this.selectedYear.getMonth() + 1) % 3) > 0 ? 1 : 0)
      var queryPM = {
        StreetID: this.selectedStreetID,
        XQID: '',
        Year: this.selectedYear.getFullYear(),
        Quarter: this.selectedQuarter,
        PageNo: 0,
        PageSize: 10
      }
      this.onSearch(queryPM)
    },
    fetechAllStreets() {
      fetchpm(this, true, '/pm/street',{
          method: 'POST'
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.streets = data.data.streets || []
        }
      })
    }
  }
}
</script>

<style lang='less' module='s'>
.wrap{
  .content{
    border: solid 1px #4c87b9;
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
      .searchWrap{
        display: flex;
        justify-content: flex-end;
        margin: 10px;
      }
      .key{
        background-color: #f0f0f0;
      }
    }    
  }
}
.elSelect{
  width: 100%;
}
.hidden{
  display: none;
}
.editKPIOther{
  font-size: 30px;
  input{
    background-color: #f1f3f6;
  }
}
.editBtns{
  margin-top: 20px;
}
.loading{

}
</style>