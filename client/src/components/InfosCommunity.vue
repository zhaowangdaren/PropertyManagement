<template>
  <div>
    <div :class='s.searchWrap'>
      <div :class='s.title'>name</div>
      <!-- 街道名 -->
      <div :class='s.inputWrap'>
        <el-select :class='s.elInput' v-model="searchStreetID" placeholder="请选择">
          <el-option
            v-for="item in showStreets"
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
      <tr>
        <th>选择</th>
        <th>name</th>
        <!-- 街道名 -->
        <th>name2</th>
        <!-- 社区名 -->
        <th>charger</th>
        <!-- 负责人 -->
        <th>tel</th>
        <!-- 电话号码 -->
        <th :class='s.descr'>描述</th>
        <th>操作</th>
      </tr>
      <tr v-for='item in communities' :class='s.street'>
        <td>
          <input type="checkbox" v-model='item.checked'>
        </td>
        <td v-text='item.StreetName'></td>
        <td v-text='item.Name'></td>
        <td v-text='item.PersonInCharge'></td>
        <td v-text='item.Tel'></td>
        <td v-text='item.Intro'></td>
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
      <tr v-if='communities == null || communities.length == 0'>
        <td colspan="7">无记录</td>
      </tr>
    </table>
    <el-dialog
      title='Add Community'
      :visible.sync='showAddDialog'
      size='small'>
      <add-community
        v-if='showAddDialog'
        @cancel='onAddCancel'
        @addSucc='onAddSucc'></add-community>
    </el-dialog>
    <el-dialog
      title='Edit Community'
      :visible.sync='showEditDialog'
      size='small'>
      <edit-community v-if='showEditDialog'
        :community='editingCommunity'
        @cancel='showEditDialog = false'
        @editSucc='fetchCommunities'></edit-community>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="showDelConfirm"
      size="tiny">
      <div>请确认删除</div>
      <div v-for='item in dels' v-text='item.Name'></div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="onDelCancel">取 消</el-button>
        <el-button type="primary" @click="onDelConfirm">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import AddCommunity from '@/components/dialog/AddCommunity'
  import EditCommunity from '@/components/dialog/EditCommunity'

  import fetchpm from '@/fetchpm'
  export default {
    components: {ImageButton, AddCommunity, EditCommunity},
    data () {
      return {
        host: 'http://localhost:3000',
        communities:[],
        streets: [],
        dels: [],
        searchStreetID:'',
        showAddDialog: false,
        showEditDialog: false,
        showDelConfirm: false,
        editingCommunity: null
      }
    },
    mounted () {
      this.fetchCommunities('')
      this.fetchAllStreets()
    },
    computed: {
      showStreets: function () {
        if (!this.streets ) return []
        return this.streets.filter(street => {
          var index = street.Name.indexOf(this.inputName) > -1
          return index
        })
      }
    },
    methods: {
      onAdd () {
        this.showAddDialog = true
      },
      onAddCancel () {
        console.info('onAddCancel')
        this.showAddDialog = false
      },
      onAddSucc () {
        this.fetchCommunities()
        this.fetchAllStreets()
      },
      onDel () {
        this.dels = this.communities.filter(item => {
          return item.checked
        })
        if (this.dels.length <= 0) return
        this.showDelConfirm = true
      },
      onDelConfirm () {
        if (this.dels.length <= 0) return
        fetchpm(this, true, '/pm/community/del', {
          method: 'POST',
          body: {values: this.dels.map(item => {
            return item.ID
          })}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onDelConfirm', body)
          if (body.error === 0) {
            this.fetchCommunities()
            this.showDelConfirm = false
          }
        })
      },
      onDelCancel () {
        this.showDelConfirm = false
        this.dels = []
      },
      onEdit (community) {
        this.editingCommunity = community
        this.showEditDialog = true
      },
      onSearch (streetID) {
        if (streetID && streetID !== '') {
          this.searchStreetID = streetID
        }
        console.info('onSearch', this.searchStreetID)
        if (!this.searchStreetID && this.searchStreetID === '') return
        this.fetchCommunitiesByStreetID(this.searchStreetID)
      },
      fetchCommunities () {
        this.inputName = ''
        fetchpm(this, true, '/pm/community', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(data => {
          console.info(data)
          if (data.error == 1) return
          this.communities = data.data
          this.setCommunitiesName()
        })
      },
      setCommunitiesName () {
        console.info(this.streets, this.communities)
        if (this.communities.length <= 0 || this.streets.length == 0) return
        for(let j = 0; j < this.communities.length; j++) {
          for (var i = 0; i < this.streets.length; i++) {
            if (this.streets[i].ID == this.communities[j].StreetID) {
              this.communities[j].StreetName = this.streets[i].Name
              break;
            }
          }
        }
      },
      fetchAllStreets () {//获取所有街道名称
        fetchpm( this, true, '/pm/street', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          this.streets = body.data
          this.setCommunitiesName()
        })
      },
      fetchCommunitiesByStreetID(streetID) {
        if (!streetID) return
        fetchpm(this, true, '/pm/community/kvs', {
          method: 'POST',
          body: {
            streetID: streetID
          }
        }).then(resp => {
          return resp.json()
        }).then( data => {
          if(data.error !== 0) console.error("Error: search CommunitiesByStreetName. Reason:" + data.data)
          this.communities = data.data
          this.setCommunitiesName()
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
  .searchWrap{
    border: solid 1px #ddd;
    display: flex;
    align-items: center;
    position: relative;
    .title{
      background: #e0e0e0;
      padding: 10px;
      font-size: 20px;
    }
    .inputWrap{
      position: relative;
      font-size: 18px;
      flex: 0.5;
      margin-left: 10px;
      .elInput {
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
    .street{
      &:hover {
        background-color: #ddd;
      }
    }
  }
</style>
