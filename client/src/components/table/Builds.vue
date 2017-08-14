<template>
  <div :class='s.wrap'>
    <div :class='s.tab'>
      <img src="~@/res/images/earth.png">
      建筑信息管理<!-- 物业信息管理 -->
    </div>
    <table :class='s.searchWrap'>
      <tr>
        <td :class='s.title'>街道</td>
        <td>
          <el-select v-model="inputStreetID" filterable placeholder="全部">
            <el-option
              v-for="item in streets"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </td>
        <td :class='s.title'>社区</td>
        <td>
          <el-select v-model="inputCommunityID" filterable placeholder="全部">
            <el-option
              v-for="item in communities"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </td>
        <td :class='s.title'>小区</td>
        <td>
          <el-select v-model="inputXQID" filterable placeholder="全部">
            <el-option
              v-for="item in xqs"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </td>
      </tr>
      <tr>
        <td :class='s.title'>房屋楼栋号</td>
        <td>
        <!-- 房屋楼栋号 -->
          <el-input v-model="inputHouseBuildNo" placeholder="请输入号"></el-input>
        </td>
        <td :class='s.title'>房屋门牌号</td>
        <td>
          <el-input v-model="inputHouseNo" placeholder="请输入HouseNo"></el-input>
        </td>
        <td :class='s.title'>房产登记人</td>
        <td>
          <el-input v-model="inputOwner" placeholder="请输入owner"></el-input>
        </td>
      </tr>
    </table>
    <div :class='s.addDel'>
      <el-button v-if='EDITABLE' type="primary" icon="plus" @click='showAddDialog = true'>新增</el-button>
      <el-button type="primary" icon="search" @click='onSearch'>查询</el-button>
    </div>
    <table>
      <tr >
        <!-- 房产登记人 -->
        <th>房产登记人</th>
        <!-- 所在小区 -->
        <th>所在小区</th>
        <!-- 楼栋号 -->
        <th>房屋楼栋号</th>
        <!-- 门牌号 -->
        <th>房屋门牌号</th>
        <!-- 房屋类型 -->
        <th>房屋类型</th>
        <th>操作</th>
      </tr>
      <tr v-for='item in houses'>
        <td v-text='item.Owner'></td>
        <td v-text='item.XQ'></td>
        <td v-text='item.HouseBuildNo'></td>
        <td v-text='item.HouseNo'></td>
        <td v-text='item.HouseType'></td>
        <td align="center">
          <el-button v-if='EDITABLE' type="primary" icon="edit" @click='onEdit(item)'>编辑</el-button>
          <el-button v-if='EDITABLE' type="danger" icon="delete" @click='onDel(item)'>删除</el-button>
          <el-button type='primary'>查看</el-button>
        </td>
      </tr>
      <tr v-if='houses.length === 0'>
        <td colspan="6">无记录</td>
      </tr>
    </table>
    <el-dialog
      title='新增房屋信息'
      size='large'
      :visible.sync="showAddDialog">
      <add-build
        @close='showAddDialog = false'
        @addSucc='onAddSucc'>
      </add-build>
    </el-dialog>
    <el-dialog
      title='编辑房屋信息'
      size='large'
      :visible.sync="showEditDialog">
      <edit-build
        :HOUSE='editingHouse'
        @close='showEditDialog = false'
        @editSucc='onEditSucc'>
      </edit-build>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="showDelConfirm"
      size="tiny">
      <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
      <div>请确认删除</div>
      <div v-text='delHouse.BuildNo'></div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="showDelConfirm = false">取 消</el-button>
        <el-button type="primary" @click="onDelConfirm">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script type="text/javascript">
  import SearchSelect from '@/components/SearchSelect'
  import AddBuild from '@/components/dialog/AddBuild'
  import EditBuild from '@/components/dialog/EditBuild'
  import fetchpm from '@/fetchpm'
  export default {
    props: {
      EDITABLE: Boolean
    },
    components: {AddBuild, EditBuild},
    data () {
      return {
        warn: '',
        showAddDialog: false,
        showEditDialog: false,
        showDelConfirm: false,
        streets: [],
        communities: [],
        xqs: [],
        inputStreetID:'',
        inputCommunityID: '',
        inputXQID:'',
        inputHouseBuildNo:'',
        inputHouseNo:'',
        inputOwner:'',
        isLoadingInput: false,
        houses:[],
        editingHouse: {},
        delHouse: {}
      }
    },
    computed: {
    },
    mounted () {
      this.fetchHouses()
      this.fetechAllStreets()
    },
    watch: {
      inputStreetID: function (val) {
        this.inputCommunityID = ''
        if (this.isLoadingInput ) return
        this.fetchAllCommunitiesByStreetID(val)
      },
      inputCommunityID: function (val) {
        this.inputXQID = ''
        if (this.isLoadingInput ) return
        this.fetchAllXQByCommunityID(val)
      }
    },
    methods: {
      onAddSucc () {
        this.fetchHouses()
      },
      onEdit (house) {
        this.editingHouse = house
        this.showEditDialog = true
      },
      onEditSucc () {
        this.fetchHouses()
      },
      onDel (delHouse) {
        this.delHouse = delHouse
        this.showDelConfirm = true
      },
      onDelConfirm () {
        fetchpm(this, true, '/pm/houses/del', {
          method: 'POST',
          body: {values: [this.delHouse.ID]}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onDelConfirm', body)
          if (body.error == 1) {
            this.warn = body.data
          } else {
            this.fetchHouses()
            this.showDelConfirm = false
          }
        })
      },
      onSearch () {
        //小区》社区》街道
        if (this.inputXQID !== '') {
          this.fetchHouseKVs({XQID: this.inputXQID})
          return
        }
        if (this.inputCommunityID !== '') {
          this.fetchHouseKVs({communityID: this.inputCommunityID})
          return
        }
        if (this.inputStreetID !== '') {
          this.fetchHouseKVs({streetID: this.inputStreetID})
          return
        }
      },
      fetchHouses () {
        fetchpm(this, true, '/pm/house', {
          method: 'POST'
        }).then(resp =>{
          return resp.json()
        }).then(body => {
          console.info('fetchHouses', body)
          if (body.error !== 1) this.houses = body.data
        })
      },
      fetchHouseKVs (kvs) {
        fetchpm(this, true, '/pm/house/kvs', {
          method: 'POST',
          body: kvs
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchHouseKVs', body)
          if (body.error !== 1) this.houses = body.data
        })
      },
      fetechAllStreets () {
        fetchpm(this, true, '/pm/street', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetechAllStreetName',body)
          if (body.error !== 1) this.streets = body.data
        })
      },
      fetchAllCommunitiesByStreetID (streetID) {
        if ( !streetID || streetID == '') return null
        fetchpm( this, true, '/pm/community/kvs', {
          method: 'POST',
          body: {streetID: streetID}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchAllCommunities', body)
          this.communities = body.data
        })
      },
      fetchAllXQByCommunityID (communityID) {
        if ( !communityID || communityID == '') return null
        fetchpm( this, true, '/pm/xq/kvs', {
          method: 'POST',
          body: {communityID: communityID}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchAllXQByCommunityID', body)
          this.xqs = body.data
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  .tab{
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
  .searchWrap{
    background-color: #fff;
    border: solid 1px #ddd;
    .title{
      background: #eee;
      padding: 10px;
      font-size: 20px;
    }
    .inputWrap {
        position: relative;
        font-size: 18px;
        flex: 1;
        margin-left: 10px;
      input{
        width: 100%;
        border: solid 1px #ddd;
        height: 30px;
      }
    }
    .searchBt{
      position: absolute;
      right: 10px;
    }
  }
  .addDel{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    margin-top: 10px;
    margin-right: 10px;
    .bt{
      margin: 5px;
    }
  }
  table{
    width: 99%;
    font-size: 15px;
    color: #555;
    margin: 10px auto;
    background-color: #fff;
    text-align: center;
    th{
      text-align: center;
      padding: 5px;
      border: solid 1px #ddd;
      background: #f0f0f0;
    }
    td{
      padding: 5px;
      border: solid 1px #ddd;
    }
    tr{
      &:hover {
        background-color: #ddd;
      }
      .bt{
        display: inline-block;
      }
    }
  }
  .warn{
    color: red;
    text-align: center;
  }
}
</style>
