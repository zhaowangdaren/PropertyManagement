<template>
  <div :class='s.wrap'>
    <table>
      <tr >
        <th>序号</th>
        <th>真实姓名</th>
        <!-- 用户名 -->
        <th>OpenID</th>
        <!-- OPenID -->
        <th>电话</th>
        <!-- 电话 -->
        <th>操作</th>
      </tr>
      <tr v-for='(court, index) in courts' :class='s.managers'>
        <td v-text='index + 1'></td>
        <td v-text='court.ActualName'></td>
        <td v-text='court.OpenID'></td>
        <td v-text='court.Tel'></td>
        <td align="center">
          <el-button
            type="info"
            v-if='court.Bind === -1'
            :loading='isBinding'
            :disabled='court.Bind === -1'
            >已解绑</el-button>
          <el-button
            @click='onBind(court)'
            type="primary" 
            :loading='isBinding'
            :disabled='court.Bind === 1'
            >允许绑定</el-button>
          <el-button 
            @click='onUnbind(court)'
            :type='court.Bind === 1 ? "danger" : ""' 
            icon="delete"
            :loading='isUnbinding'
            :disabled='court.Bind === 0 || court.Bind === -1'>强制解绑</el-button>
          <el-button 
            @click='onDel(court)'
            type='danger' 
            icon="delete"
            :loading='isDeling'>删除</el-button>
        </td>
      </tr>
      <tr v-if='courts.length=== 0'>
        <td colspan="6">无记录</td>
      </tr>
    </table>
    <el-pagination
      layout="total, prev, pager, next"
      @current-change='onChangePage'
      :page-size='query.PageSize'
      :total="sum">
    </el-pagination>
  </div>
</template>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import SearchSelect from '@/components/SearchSelect'
  import fetchpm from '@/fetchpm'
  import { Message } from 'element-ui'

  export default {
    components: {ImageButton, SearchSelect},
    data () {
      return {
        courts:[],
        isUnbinding: false,
        isBinding: false,
        isDeling: false,
        query: {
          Name: '', // ActualName
          PageNo: 0,
          PageSize: 10
        },
        sum: 0
      }
    },
    mounted () {
      this.fetchCourts(this.query)
    },
    methods: {
      fetchCourts (query) {
        fetchpm(this, true, '/pm/court/wx', {
          method: 'POST',
          body: query
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchCourts', body)
          if (body.error === 0){
            this.courts = body.data.list || []
            this.sum = body.data.sum
          }
        })
      },
      onChangePage () {
        this.query.PageNo = curPage - 1
        this.fetchCourts(this.query)
      },
      onDel (court) {
        this.isDeling = true
        fetchpm(this, true, '/pm/court/wx/del/openid/' + court.OpenID, {
          method: 'GET',
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0) this.fetchCourts(this.query)
          else Message({type:'error', message: body.data})
          this.isDeling = false
        }).catch(error => {
          Message({type:'error', message: error.message})
          this.isDeling = false
        })
      },
      onUnbind (court) {
        this.isUnbinding = true
        var tempCourt = Object.assign({}, court)
        tempCourt.Bind = -1
        fetchpm(this, true, '/pm/court/wx/update', {
          method: 'POST',
          body: tempCourt
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onUnbind', body)
          if (body.error === 0) {
            court.Bind = -1
          } else {
            Message({type:'error', message: body.data})
          }
          this.isUnbinding = false
        })
      },
      onBind (court) {
        this.isBinding = true
        var tempCourt = Object.assign({}, court)
        tempCourt.Bind = 1
        fetchpm(this, true, '/pm/court/wx/update', {
          method: 'POST',
          body: tempCourt
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onBind', body)
          if (body.error === 0) {
            court.Bind = 1
          } else {
            Message({type:'error', message: body.data})
          }
          this.isBinding = false
        })
      },
      onSearch () {
        
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  table{
    width: 99%;
    font-size: 15px;
    color: #555;
    margin: 10px auto;
    text-align: center;
    th{
      text-align: center;
      padding: 5px;
      border: solid 1px #ddd;
      background-color: #f0f0f0;
    }
    td{
      padding: 5px;
      border: solid 1px #ddd;
    }
    .managers{
      &:hover {
        background-color: #f1f3f6;
      }
    }
  }
}
</style>
