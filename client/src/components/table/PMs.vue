<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        物业信息<!--  -->
      </div>
      <div :class='s.searchWrap'>
        <table>
          <tr>
            <td :class='s.key' v-if='STREET_ID === ""'>街道名</td>
            <td v-if='STREET_ID === ""'>
              <el-select v-model="inputStreetID" filterable placeholder="选择街道">
                <el-option
                  v-for="item in streets"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>社区名</td>
            <td>
              <el-select v-model="inputCommunityID" filterable placeholder="选择社区">
                <el-option
                  v-for="item in communities"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>小区名</td>
            <td>
              <el-select v-model="inputXQID" filterable placeholder="选择小区">
                <el-option
                  v-for="item in xqs"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
          </tr>
        </table>
      </div>
      <div :class='s.addDel'>
        <el-button type="primary" icon="plus" @click='onAdd' v-if='EDITABLE'>新增</el-button>
        <el-button type="primary" icon="search" @click='onSearch'>查询</el-button>
      </div>
      <div :class='s.table'>
        <table>
          <tr >
            <th>小区名</th>
            <!-- 小区名 -->
            <th>物业公司</th>
            <!-- 物业公司 -->
            <th>独立法人</th>
            <!-- 独立法人 -->
            <th>物业资质</th>
            <!-- 物业资质 -->
            <th>物业性质</th>
            <!-- 物业性质 -->
            <th>操作</th>
          </tr>
          <tr v-for='(item, index) in pms'>
            <td v-text='xqNames[index]'></td>
            <td v-text='item.Name'></td>
            <td v-text='item.LegalPerson'></td>
            <td v-text='item.WuYeZiZhi'></td>
            <td v-text='item.WuYeXinZhi'></td>
            <td align="center">
              <el-button v-if='EDITABLE' type="primary" icon="edit" @click='onEdit(item)'>编辑</el-button>
              <el-button v-if='EDITABLE' type="danger" icon="delete" @click='onDel(item)'>删除</el-button>
              <el-button v-if='!EDITABLE' type='primary' @click='onDetails(item)'>查看</el-button>
            </td>
          </tr>
        </table>
      </div>
    </div>
    <el-dialog
      title='新增物业'
      :visible.sync='showAddDialog'
      size='small'>
      <add-p-m @cancel='showAddDialog = false' @addSucc='onAddSucc'></add-p-m>
    </el-dialog>
    <el-dialog
      title='编辑物业信息'
      :visible.sync='showEditDialog'
      size='small'>
      <edit-p-m v-if='showEditDialog' :PM='editingPM' @cancel='showEditDialog = false' @editSucc='onEditSucc'></edit-p-m>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="showDelConfirm"
      size="tiny">
      <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
      <div>请确认删除</div>
      <div v-text='delPM.Name'></div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="showDelConfirm = false">取 消</el-button>
        <el-button type="primary" @click="onDelConfirm">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title='物业详细信息'
      :visible.sync='showDetailDialog'
      size='small'>
      <details-p-m :pm='showingPM'></details-p-m>
    </el-dialog>
  </div>
</template>

<script>
import AddPM from '@/components/dialog/AddPM'
import EditPM from '@/components/dialog/EditPM'
import DetailsPM from '@/components/dialog/DetailsPM'
import fetchpm from '@/fetchpm'
export default {
  components: {AddPM, EditPM, DetailsPM},
  props: {
    EDITABLE: Boolean,//是否可编辑
    STREET_ID: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      userStreetID: '',
      streets: [],
      communities: [],
      xqs: [],
      inputStreetID:'',
      inputCommunityID: '',
      inputXQID:'',
      isLoadingInput: false,
      pms:[],
      showAddDialog: false,
      showEditDialog: false,
      showDelConfirm: false,
      showDetailDialog: false,
      delPM: {},
      warn:'',
      editingPM: {},
      showingPM: {}
    }
  },
  computed: {
    xqNames: function () {
      return this.pms.map(pm => {
        for (var i = 0; i < this.xqs.length; i++) {
          if (this.xqs[i].ID == pm.XQID) return this.xqs[i].Name
        }
      })
    }
  },
  mounted () {

    this.fetchPMs()
    this.fetchAllXQs()
    this.inputStreetID = this.STREET_ID
    this.fetechAllStreets()
  },
  watch: {
    inputStreetID: function (val) {
      this.inputCommunityID = ''
      if (this.isLoadingInput ) return
      this.fetchCommunitiesByStreetID(val)
    },
    inputCommunityID: function (val) {
      this.inputXQID = ''
      console.info('inputCommunityID', val)
      if (this.isLoadingInput ) return
      this.fetchXQByCommunityID(val)
    }
  },
  methods: {
    onDetails (pm) {
      this.showingPM = pm
      this.showDetailDialog = true
    },
    onDel (pm) {
      this.delPM = pm
      this.showDelConfirm = true
    },
    onDelConfirm () {
      fetchpm(this, true, '/pm/pms/del', {
        method: 'POST',
        body: {values: [this.delPM.ID]}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onDelConfirm', body)
        if (body.error == 1) {
          this.warn = body.data
        } else {
          this.showDelConfirm = false
          this.delPM = {}
          this.fetchPMs()
        }
      })
    },
    onEdit (item) {
      // this.editingPM = item
      Object.assign(this.editingPM, item)
      this.showEditDialog = true
    },
    onEditSucc () {
      this.fetchPMs()
    },
    onAdd () {
      this.showAddDialog = true
    },
    onAddSucc () {
      this.fetchPMs()
    },
    onSearch () {
      //小区》社区》街道
      if (this.inputXQID !== '') {
        this.fetchPMKVs({XQID: this.inputXQID})
        return
      }
      if (this.inputCommunityID !== '') {
        this.fetchPMKVs({communityID: this.inputCommunityID})
        return
      }
      if (this.inputStreetID !== '') {
        this.fetchPMKVs({streetID: this.inputStreetID})
        return
      }
    },
    fetchPMs () {
      fetchpm(this, true, '/pm/pm', {
        method: 'POST'
      }).then(resp =>{
        return resp.json()
      }).then(body => {
        console.info('fetchPMs', body)
        if (body.error !== 1) this.pms = body.data
      })
    },
    fetchPMKVs (kvs) {
      fetchpm(this, true, '/pm/pm/kvs', {
        method: 'POST',
        body: kvs
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPMKVs', body)
        if (body.error !== 1) this.pms = body.data
      })
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
          this.streets = data.data
        }
      })
    },
    fetchAllXQs () {
      fetchpm( this, true, '/pm/xq', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllXQs', body)
        this.xqs = body.data
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
          this.communities = body.data
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
          this.xqs = body.data
        }
        this.isLoadingInput = false
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
    margin: 10px;
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
    .searchWrap{
      display: flex;
      align-items: center;
      position: relative;
      margin: 10px;
      table{
        width: 100%;
        background-color: #fff;
        td{
          text-align: center;
          padding: 2px;
          border: solid 1px #ddd;
        }
        .key{
          background-color: #ddd;
          font-size: 20px;
        }
      }
    }
    .addDel {
      display: flex;
      justify-content: flex-end;
      margin: 10px;
    }
    .table{
      margin: 10px;
      table{
        width: 100%;
        font-size: 15px;
        color: #555;
        margin: 10px auto;
        background-color: #fff;
        text-align: center;
        th{
          text-align: center;
          padding: 5px;
          border: solid 1px #ddd;
          background-color: #f0f0f0;
        }
        tr{
          &:hover {
            background-color: #f0f0f0;
          }
        }
        td{
          padding: 5px;
          border: solid 1px #ddd;
        }
      }
    }
  }
}
</style>
