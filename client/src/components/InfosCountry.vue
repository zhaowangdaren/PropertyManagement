<template>
  <div :class='s.wrap'>
    <div :class='s.searchWrap'>
    <!-- 街道名 -->
      <div :class='s.title'>街道名</div>
      <div :class='s.inputWrap'>
        <el-select :class='s.elInput' v-model="searchStreetID" placeholder="请选择">
          <el-option
            v-for="item in streets"
            :key="item.ID"
            :label="item.Name"
            :value="item.ID">
          </el-option>
        </el-select>
      </div>
      <!-- 小区名 -->
      <div :class='s.title'>社区名</div>
      <div :class='s.inputWrap'>
        <el-select :class='s.elInput' v-model="searchCommunityID" placeholder="请选择">
          <el-option
            v-for="item in searchCommunities"
            :key="item.ID"
            :label="item.Name"
            :value="item.ID">
          </el-option>
        </el-select>
      </div>
      <!-- <image-button :class='s.searchBt' @click='onSearch'
        text='查询'
        :img='require("@/res/images/ic_serach.png")'
        bgColor='#4c87b9'
        color='#fff'
      /> -->
      <el-button
        :class='s.searchBt'
        @click='onSearch'
        type='primary'
        icon='search'>查询</el-button>
    </div>
    <div :class='s.addDel'>
      <!-- <image-button :class='s.bt' @click='onAdd'
        text='新增'
        :img='require("@/res/images/add.png")'
        bgColor='#3598dc'
        fontSize='20px'
        /> -->
      <el-button
        @click='onAdd'
        type='primary'
        icon='plus'>新增</el-button>
      <!-- <image-button :class='s.bt'
        text='删除'
        :img='require("@/res/images/delete.png")'
        bgColor='#cb5a5e'
        fontSize='20px'
        @click='onDel'
        /> -->
      <el-button
        @click='onDel'
        type='danger'
        icon='delete'>删除</el-button>
    </div>
    <table>
      <tr >
        <th>选择</th>
        <!-- 街道名 -->
        <th>街道名</th>
        <!-- 社区名 -->
        <th>社区名</th>
        <!-- 小区名 -->
        <th>小区名</th>
        <!-- 位置 -->
        <th>地址</th>
        <!-- 负责人 -->
        <th>负责人</th>
        <!-- 电话号码 -->
        <th>电话号码</th>
        <th>描述</th>
        <th>操作</th>
      </tr>
      <tr v-for='(item, index) in xqs' :class='s.street'>
        <td>
          <input type="checkbox" v-model='item.checked'>
        </td>
        <td v-text='streetNames[index]'></td>
        <td v-text='communityNames[index]'></td>
        <td v-text='item.Name'></td>
        <td v-text='item.Address' class='descr'></td>
        <td v-text='item.ContactName'></td>
        <td v-text='item.Tel'></td>
        <td v-text='item.Intro' class='descr'></td>
        <td>
          <!-- <image-button
            text='编辑'
            :img='require("@/res/images/edit.png")'
            bgColor='#26a69a'
            @click='onEdit(item)'
          /> -->
          <el-button
            @click='onEdit(item)'
            type='primary'
            icon='edit'>编辑</el-button>
        </td>
      </tr>
      <tr v-if='!xqs || xqs.length == 0'>
        <td colspan="9">无记录</td>
      </tr>
    </table>
    <el-pagination v-if='showPage'
      layout="total, prev, pager, next"
      @current-change='onChangePage'
      :page-size='pageSize'
      :total="sum">
    </el-pagination>
    <el-dialog
      title='新增小区'
      :visible.sync='showAddDialog'
      size='small'>
      <add-country v-if='showAddDialog' @cancel='showAddDialog = false' @addSucc='onAddSucc'></add-country>
    </el-dialog>
    <el-dialog
      title='编辑小区'
      :visible.sync='showEditDialog'
      size='small'>
      <edit-country v-if='showEditDialog' :country='editingXQ' @cancel='showEditDialog = false'></edit-country>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="showDelConfirm"
      size="tiny">
      <div>请确认删除</div>
      <div v-for='item in dels' v-text='item.Name' :class='s.delItem'></div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="onDelCancel">取 消</el-button>
        <el-button type="primary" @click="onDelConfirm">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script type="text/javascript">
  import ImageButton from '@/components/ImageButton'
  import AddCountry from '@/components/dialog/AddCountry'
  import EditCountry from '@/components/dialog/EditCountry'

  import fetchpm from '@/fetchpm'

  export default {
    components: {ImageButton, AddCountry, EditCountry},
    data () {
      return {
        showDialog: '',
        streets: [],
        communities: [],
        searchCommunities: [],
        xqs: [],
        searchStreetID:'',
        searchCommunityID: '',

        showAddDialog: false,
        showEditDialog: false,
        editingXQ: null,
        showDelConfirm: false,
        dels: [],
        pageNo: 0,
        pageSize: 10,
        sum: 0,
        showPage: true
      }
    },
    watch: {
      searchStreetID: function (val) {
        this.searchCommunityID = ''
        this.fetchCommunitiesByStreetID(val)
      }
    },
    computed: {
      streetNames: function () {
        if (this.xqs.length <= 0 ) return
        return this.xqs.map(xq => {
          let name = ''
          for (var i = 0; i < this.streets.length; i++) {
            if(xq.StreetID == this.streets[i].ID){
              name = this.streets[i].Name
              break
            }
          }
          return name
        })
      },
      communityNames: function () {
        if (this.xqs.length <= 0 ) return
        return this.xqs.map(xq => {
          let name = ''
          for (var i = 0; i < this.communities.length; i++) {
            if(xq.CommunityID == this.communities[i].ID){
              name = this.communities[i].Name
              break
            }
          }
          return name
        })
      }
    },
    mounted () {
      this.fetchXQs()
      this.fetchAllStreets()
      this.fetchAllCommunities()
    },
    methods: {
      onEdit (xq) {
        this.editingXQ = xq
        this.showEditDialog = true
      },
      onAdd () {
        this.showAddDialog = true
      },
      onAddSucc () {
        this.fetchXQs()
      },
      onDel () {
        this.dels = this.xqs.filter(item => {
          return item.checked
        })
        if (!this.dels || this.dels.length == 0) return
        this.showDelConfirm = true
      },
      onDelConfirm () {
        if (this.dels.length <= 0) return
        fetchpm(this, true, '/pm/xq/del', {
          method: 'POST',
          body: {values: this.dels.map(item => {
            return item.ID
          })}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onDelConfirm', body)
          if (body.error === 0) {
            this.fetchXQs()
            this.showDelConfirm = false
          }
        })
      },
      onDelCancel () {
        this.showDelConfirm = false
      },
      onSearch () {
        if(this.searchStreetID === '' && this.searchCommunityID === '') return
        var params = {}
        if (this.searchIdStreet !== '') params.StreetID = this.searchStreetID
        if (this.searchCommunityID !== '') params.CommunityID = this.searchCommunityID
        console.info('onSearch', this.searchIdStreet, this.searchIdXQ)
        fetchpm(this, true, '/pm/xq/kvs', {
          method: 'POST',
          body: params
        }).then( resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchXQKVs', body)
          if (body.error !== 1) this.xqs = body.data || []
          this.showPage = false
        })
      },
      setStreets () {
        if (this.streets.length <= 0 || this.xqs.length <= 0) return
        for(let j = 0; j < this.xqs.length; j++) {
          for (var i = 0; i < this.streets.length; i++) {
            if (this.streets[i].ID == this.xqs[j].StreetID) {
              this.xqs[j].StreetName = this.streets[i].Name
              break;
            }
          }
        }
      },
      onChangePage (curPage) {
        this.pageNo = curPage - 1
        this.fetchXQs()
      },
      fetchXQs () {
        fetchpm(this,true,'/pm/xq', {
          method: 'POST',
          body: {pageNo: this.pageNo, pageSize: this.pageSize}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info(body)
          this.xqs = body.data.xqs || []
          this.sum = body.data.sum
          this.showPage = true
        })
      },
      fetchAllStreets () {
        fetchpm(this, true,'/pm/street', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetechAllStreetName',body)
          this.streets = body.data.streets || []
        })
      },
      fetchAllCommunities () {
        fetchpm(this, true,'/pm/community', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchAllCommunities', body)
          this.communities = body.data.communities || []
        })
      },
      fetchCommunitiesByStreetID (streetID) {
        fetchpm(this, true, '/pm/community/kvs', {
          method: 'POST',
          body: {streetID: streetID}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchCommunitiesByStreetID', body)
          this.searchCommunities = body.data || []
          this.showPage = false
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  .searchWrap{
    border-bottom: solid 1px #ddd;
    display: flex;
    align-items: center;
    width: 100%;
    position: relative;
    .title{
      background: #e0e0e0;
      padding: 10px;
      font-size: 20px;
    }
    .inputWrap{
      position: relative;
      font-size: 18px;
      flex: 0.4;
      margin: 0 10px;
      .elInput{
        width: 100%;
      }
    }
    .searchBt{
      position: absolute;
      right: 10px;
    }
  }
  .addDel{
    display: flex;
    align-items: center;
    margin: 10px;
    .bt{
      margin: 5px;
    }
  }
  .delItem{
    padding: 10px;
    border-radius: 5px;
    border: solid 1px rgba(0,0,0,0.2);
    box-shadow: 1px 1px 1px 0 rgba(0,0,0,0.2);
    text-align: center;
    font-size: 20px;
    font-weight: bold;
  }
}
</style>
