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
              <td v-text='item.Name'></td>
              <td>{{ xqs[index] ? xqs[index].Name : ""}}</td>
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
              <td colspan="7">无记录</td>
            </tr>
          </template>
        </table>
        <el-pagination
          layout="total, prev, pager, next"
          @current-change='onChangePage'
          :page-size='queryPM.PageSize'
          :total="sum">
        </el-pagination>
      </div>
    </div>
  </div>
</template>

<script type="text/javascript">
import filterEventStatus from '@/filters/filterEventStatus'
import filterEventLevel from '@/filters/filterEventLevel'
import calculatPMKPI from '@/filters/calculatPMKPI'
import fetchpm from '@/fetchpm'
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
      loading: false,
      quarters: [{value:1, label:'1'},{value:2,label:'2'},{value:3, label:'3'},{value:4, label: '4'}],
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

      queryPM: {
        Name: '',
        PageNo: 0,
        PageSize: 10
      },
      showEditDialog: false,
      kpiOther: 0,
      editingKPI: null
    }
  },
  mounted () {
    this.selectedYear = new Date()
    var thisMonth = new Date().getMonth()
    this.selectedQuarter = parseInt((thisMonth + 1) / 3) + (((thisMonth + 1) % 3) > 0 ? 1 : 0)
    // this.onInputSearch()
    // this.fetechAllStreets()
    
    var query = {
        Name: '',
        PageNo: 0,
        PageSize: 10
      }
    this.onSearch(query)
  },
  watch: {
    selectedStreetID: function (value) {
      this.selectedCommunityID = ''
      this.fetchCommunitiesByStreetID(value)
    },
    selectedCommunityID: function (value) {
      this.selectedXQID = ''
      this.fetchXQByCommunityID(value)
    },
    KPIs: function (value) {
      this.fetchPMs(this.KPIs.map(item => { return item.XQID}))
    }
  },
  methods: {
    onBTSearch () {
      var queryPM = {
        Name: '',
        PageNo: 0,
        PageSize: 10
      }
      this.onSearch(queryPM)
    },
    sortKPIs (kpis) {
      return kpis.sort((a, b) => {
        return a.PMName > b.Name ? 1 : -1
      })
    },
    onEdit (item) {
      this.kpiOther = item.Other
      this.showEditDialog = true
      this.editingKPI = item
    },
    onEditSure () {
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
      this.fetchPMsV2(query)
      // this.fetchXQs(query)
    },
    fetchPMsV2 (query){
      fetchpm(this, true, '/pm/pm', {
        method: 'POST',
        body: query
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPMsV2', body)
        if (body.error === 0) {
          this.KPIs = body.data.pms || []
          this.sum = body.data.sum
          this.fetchKPIs(this.KPIs.map(item => {return item.XQID}), this.selectedYear.getFullYear(), this.selectedQuarter)
          this.fetchXQsByIds(this.KPIs.map(item => {return item.XQID}))
        }
        setTimeout(() => {
          this.loading = false
        }, 1000)
      })
    },
    fetchXQsByIds (ids) {
      fetchpm(this, true, '/pm/xq/ids', {
        method: 'POST',
        body: {Values: ids}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetcXQs', body)
        if (body.error === 0) {
          this.xqs = body.data || []
        }
      })
    },
    fetchKPIs (ids, year, quarter) {
      for (let i = 0; i < ids.length; i++) {
        let params = {
          XQID: ids[i],
          Year: year,
          Quarter: quarter
        }
        fetchpm(this, true, '/pm/pmkpi/query', {
          method: 'POST',
          body: params
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0) {
            let pmkpi = body.data || {}
            this.setPMKPI(pmkpi)
            // this.KPIs[i] = Object.assign(this.KPIs[i], pmkpi)
           // this.$set( this.KPIs[i], 'Score' ,(100 - pmkpi.YWNo * 0.5 - pmkpi.RWNo * 1 - pmkpi.IWNo * 3 - pmkpi.Other))
            console.info('KPIs', this.KPIs)
          }
        })
      }
    },
    setPMKPI (pmkpi) {
      var index = this.KPIs.findIndex((item, i, array) => {
        console.info('findIndex', item.XQID + '----' + pmkpi.XQID)
        return item.XQID === pmkpi.XQID
      })
      console.warn('pmkpi index', index)
      this.KPIs[index] = Object.assign(this.KPIs[index], pmkpi)
    },
    fetchPMs (ids) {
      fetchpm(this, true, '/pm/pm/xqids', {
        method: 'POST',
        body: {Values: ids}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPMs', body)
        if (body.error === 0) {
          this.pms = body.data || []
          for (let i = 0; i < this.pms.length; i++) {
            this.$set(this.KPIs[i], 'PMName', this.pms[i].Name)
          }
        }
      })
    },
    onChangePage (curPage) {
      this.queryPM.PageNo = curPage - 1
      this.onSearch(this.queryPM)
    },
    onCurQuarter () {
      this.selectedYear = new Date()
      this.selectedQuarter = parseInt((this.selectedYear.getMonth() + 1 ) / 3) + (((this.selectedYear.getMonth() + 1) % 3) > 0 ? 1 : 0)
      var queryPM = {
        Name: '',
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
    },
    fetchCommunitiesByStreetID (streetID) {
      if (!streetID) return
      this.isLoadingInput = true
      this.communities = []
      this.inputCommunityID = ''
      fetchpm(this, true, '/pm/community/kvs', {
        method: 'POST',
        body: {streetID: streetID}
      }).then(resp => {
        return resp.json()
      }).then( body => {
        console.info('fetchCommunitiesByStreetName', body)
        if(body.error !== 0) {
          console.error("Error: search CommunitiesByStreetName. Reason:" + body.data)
        } else if (body.data !== null ) {
          this.communities = body.data || []
        }
        this.isLoadingInput = false
      })
    },
    fetchXQByCommunityID (communityID) {
      if (!communityID) return
      this.isLoadingInput = true
      fetchpm(this, true, '/pm/xq/kvs', {
        method: 'POST',
        body: {communityID: communityID}
      }).then(resp => {
        return resp.json()
      }).then( body => {
        console.info('fetchXQByCommunityName', body)
        if(body.error !== 0) console.error("Error: search fetchXQByCommunityName. Reason:" + body.data)
        else if (body.data !== null ) {
          this.xqs = body.data || []
        }
        this.isLoadingInput = false
      })
    },
    formatSearchData () {
      var searchData = {}
      if (this.selectedYear !== -1 && this.selectedYear !== '') searchData.Year = this.selectedYear.getFullYear()
      if (this.selectedQuarter !== 0) searchData.Quarter = this.selectedQuarter
      return searchData
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