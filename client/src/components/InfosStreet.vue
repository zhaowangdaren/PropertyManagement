<template>
  <div :class='s.wrap'>
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
        <th>街道名</th>
        <!-- 街道名 -->
        <th>负责人</th>
        <!-- 负责人 -->
        <th>电话号码</th>
        <!-- 电话号码 -->
        <th>授权码</th>
        <!-- 授权码 -->
        <th :class='s.descr'>描述</th>
        <!-- 描述 -->
        <th>操作</th>
      </tr>
      <tr v-for='street in streets' :class='s.street'>
        <td>
          <input type="checkbox" v-model='street.checked'>
        </td>
        <td v-text='street.Name'></td>
        <td v-text='street.PersonInCharge'></td>
        <td v-text='street.Tel'></td>
        <td v-text='street.AuthCode'></td>
        <td v-text='street.Intro'></td>
        <td>
          <el-button type="primary" icon="edit" @click="onEdit(street)">编辑</el-button>
        </td>
      </tr>
    </table>
    <el-dialog
      title='新增街道'
      :visible.sync='showAddDialog'
      size='small'>
      <add-street v-if='showAddDialog' @cancel='showAddDialog = false' @addSucc='onAddSucc'></add-street>
    </el-dialog>
    <el-dialog
      title='修改街道'
      :visible.sync='showEditDialog'
      size='small'>
      <edit-street v-if='showEditDialog' :street='editingStreet' @cancel='showEditDialog = false'></edit-street>
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
  import AddStreet from '@/components/dialog/AddStreet'
  import EditStreet from '@/components/dialog/EditStreet'
  import fetchpm from '@/fetchpm'
  export default {
    components: {ImageButton, AddStreet, EditStreet},
    data () {
      return {
        streets:[],
        showAddDialog: false,
        showEditDialog: false,
        showDelConfirm: false,
        dels: [],
        editingStreet: null
      }
    },
    mounted () {
      this.fetchStreets()
    },
    methods: {
      fetchStreets () {
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
      onAdd () {
        console.info('onAdd')
        this.showAddDialog = true
      },
      onAddSucc () {
        this.fetchStreets()
      },
      onEdit (street) {
        this.editingStreet = street
        this.showEditDialog = true
      },
      onDel () {
        this.dels = this.streets.filter(item => {
          return item.checked
        })
        console.info('Del', this.dels)
        if (this.dels.length == 0) return
        this.showDelConfirm = true
      },
      onDelCancel () {
        this.showDelConfirm = false
        this.dels = []
      },
      onDelConfirm () {
        if (this.dels.length ===0) return
        fetchpm(this, true, '/pm/street/del', {
          method: 'POST',
          body: {values: this.dels.map(item => {
                            return item.ID
                          })
          }
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onDel', body)
          if (body.error === 0 ) {
            this.fetchStreets()
            this.showDelConfirm = false
          }
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  // margin: 10px;
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
      min-width: 50px;
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
}
</style>
