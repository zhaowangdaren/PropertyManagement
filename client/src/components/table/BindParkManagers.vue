<template>
  <div :class='s.wrap'>
    <div :class='s.searchWrap'>
      <div :class='s.key'>真实姓名</div>
      <!-- 微信用户绑定名称（包含） -->
      <!-- <search-select v-model='pmName' :values='wxNames' :class='s.searchSelect'/> -->
      <el-select v-model="searchActualName" placeholder="请输入姓名">
        <el-option @click='onAll'>全部</el-option>
        <el-option 
          v-for='user in filterNameResult'
          :key='user._id'
          :label='user.ActualName'
          :value='user.ActualName'>
        </el-option>
      </el-select>
      <el-button
        :class='s.searchBt'
        @click='onSearch'
        type='primary'
        icon='search'>查询</el-button>
    </div>
    <table>
      <tr >
        <th>序号</th>
        <th>真实姓名</th>
        <th>小区</th>
        <!-- 用户名 -->
        <th>OpenID</th>
        <!-- OPenID -->
        <th>电话</th>
        <!-- 电话 -->
        <th>操作</th>
      </tr>
      <tr v-for='(user, index) in users' :class='s.users'>
        <td v-text='index + 1'></td>
        <td v-text='user.ActualName'></td>
        <td><span v-if='xqs[index]'>{{xqs[index].Name}}</span></td>
        <td v-text='user.OpenID'></td>
        <td v-text='user.Tel'></td>
        <td align="center">
          <el-button
            type="info"
            v-if='user.Bind === -1'
            :loading='isBinding'
            :disabled='user.Bind === -1'
            >已解绑</el-button>
          <el-button
            @click='onBind'
            type="primary" 
            :loading='isBinding'
            :disabled='user.Bind === 1'
            >允许绑定</el-button>
          <el-button 
            @click='onUnbind'
            :type='user.Bind === 1 ? "danger" : ""' 
            icon="delete"
            :loading='isUnbinding'
            :disabled='user.Bind === 0'>强制解绑</el-button>
        </td>
      </tr>
      <tr v-if='users.length=== 0'>
        <td colspan="6">无记录</td>
      </tr>
    </table>
    <component :is='showDialog' @close='showDialog = ""' />
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
        managers:[],
        showDialog: '',
        wxName: '',
        wxNames: [],
        xqs: [],
        isUnbinding: false,
        isBinding: false,
        query: {
          ActualName: '',
          PageNo: 1,
          PageSize: 20
        }
      }
    },
    mounted () {
      this.fetchParkManagers('')
    },
    methods: {
      onUnbind (manager) {
        this.isUnbinding = true
        var tempManager = Object.assign({}, manager)
        tempManager.Bind = -1
        fetchpm(this, true, '/pm/park/manager/update', {
          method: 'POST',
          body: tempManager
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onUnbind', body)
          if (body.error === 0) {
            manager.Bind = -1
          } else {
            this.$message({type:'error', message: body.data})
          }
          this.isUnbinding = false
        })
      },
      onBind (user) {
        this.isBinding = true
        var tempManager = Object.assign({}, manager)
        tempManager.Bind = -1
        fetchpm(this, true, '/pm/park/manager/update', {
          method: 'POST',
          body: tempManager
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onBind', body)
          if (body.error === 0) {
            manager.Bind = -1
          } else {
            Message({type:'error', message: body.data})
          }
          this.isBinding = false
        })
      },
      fetchParkManagers (name) {
        this.query.ActualName = name
        fetchpm(this, true, '/pm/park/managers', {
          method: 'POST',
          body: this.query
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then( data => {
          console.info('fetchusers', data)
          if (data.error === 0 ) {
            this.users = data.data.parkManagers || []
            let ids = this.users.map(item => {
              return item.PMID
            })
            this.fetchPMByID(ids)
          }
        })
      },
      fetchXQs (ids){
        fetchpm(this, false, '/open/xq/ids', {
          method: 'POST',
          body:{values: ids}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          if (body.error === 0){
            this.xqs = body.data || []
          }
        })
      },
      onSearch () {
        
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  .searchWrap{
    display: none;
    align-items: center;
    margin-top: 10px;
    width: 100%;
    position: relative;
    .key{
      font-size: 18px;
      margin-left: 5px;
    }
    .searchSelect {
      flex: 0.8;
    }
    .searchBt{
      position: absolute;
      right: 10px;
    }
  }
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
    .users{
      &:hover {
        background-color: #ddd;
      }
    }
  }
}
</style>
