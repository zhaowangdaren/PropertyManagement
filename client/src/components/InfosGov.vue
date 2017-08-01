<template>
  <div :class='s.wrap'>
    <div :class='s.addDel'>
      <image-button :class='s.bt' @click='showAddDialog = true'
        text='新增'
        :img='require("@/res/images/add.png")'
        bgColor='#3598dc'
        fontSize='20px'
        />
    </div>
    <table>
      <tr>
      <!-- 用户名 真实姓名  密码  电话  操作 -->
        <th>UserName</th>
        <th>RealName</th>
        <th>Password</th>
        <th>OfficeNumber</th>
        <th>操作</th>
      </tr>
      <tr v-for='gov in govs' :class='s.gov'>
        <td v-text='gov.UserName'></td>
        <td v-text='gov.RealName'></td>
        <td v-text='gov.Password'></td>
        <td v-text='gov.Tel'></td>
        <td :class='s.operations' align="center">
          <image-button :class='s.bt'
            text='编辑'
            :img='require("@/res/images/edit.png")'
            bgColor='#26a69a'
          />
          <image-button :class='s.bt'
            text='删除'
            :img='require("@/res/images/delete.png")'
            bgColor='#cb5a5e'
            fontSize='20px'
          />
        </td>
      </tr>
    </table>
    <el-dialog 
      title='Add User'
      :visible.sync='showAddDialog'
      size='small'>
      <add-user v-if='showAddDialog' @cancel='showAddDialog = false' @addSucc='onAddSucc' :userType='userType'></add-user>
    </el-dialog>
    <el-dialog 
      title='Edit Community'
      :visible.sync='showEditDialog'
      size='small'>
      <edit-user v-if='showEditDialog' :country='editingXQ' @cancel='showEditDialog = false'></edit-user>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="showDelConfirm"
      size="tiny">
      <div>请确认删除</div>
      <div v-text='delUser.UserName'></div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="showDelConfirm = false">取 消</el-button>
        <el-button type="primary" @click="onDelConfirm">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import AddUser from '@/components/dialog/AddUser'
  import EditUser from '@/components/dialog/EditUser'

  import fetchpm from '@/fetchpm'
  export default {
    components: {ImageButton, AddUser, EditUser},
    data () {
      return {
        userType: 2,
        govs:[],
        showAddDialog: false,
        showEditDialog: false,
        showDelConfirm: false,
        delUser: {}
      }
    },
    mounted () {
      this.fetchGovs()
    },
    methods: {
      onAddSucc () {
        this.fetchGovs()
      },
      onDelConfirm () {

      },
      fetchGovs () {
        fetchpm(this, true, '/pm/users', {
          method: "POST",
          body: {type : 2}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchGovs', body)
          if(body.error !== 1) this.govs = body.data
        })
      },
      onAdd () {
        console.info('onAdd')
        this.showDialog = AddGov
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  .addDel{
    display: flex;
    align-items: center;
    margin-top: 10px;
    .bt {
      margin: 5px;
    }
  }
  table{
    width: 99%;
    font-size: 15px;
    color: #555;
    margin: auto;
    margin-top: 10px;
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
    .gov{
      &:hover {
        background-color: #ddd;
      }
      .operations{
        .bt{
          margin-left: 5px;
          display: inline-block;
        }
      }
    }
  }
}
</style>