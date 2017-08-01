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
        <th v-if='USER_TYPE == 3' >StreetName</th>
        <th>Password</th>
        <th>OfficeNumber</th>
        <th>操作</th>
      </tr>
      <tr v-for='(user, index) in users' :class='s.gov'>
        <td v-text='user.UserName'></td>
        <td v-text='user.RealName'></td>
        <td v-if='USER_TYPE == 3' v-text='streetNames[index]'></td>
        <td v-text='user.Password'></td>
        <td v-text='user.Tel'></td>
        <td :class='s.operations' align="center">
          <image-button :class='s.bt'
            text='编辑'
            :img='require("@/res/images/edit.png")'
            bgColor='#26a69a'
            @click='onEdit(user)'
          />
          <image-button :class='s.bt'
            text='删除'
            :img='require("@/res/images/delete.png")'
            bgColor='#cb5a5e'
            fontSize='20px'
            @click='onDel(user)'
          />
        </td>
      </tr>
    </table>
    <el-dialog 
      title='Add User'
      :visible.sync='showAddDialog'
      size='small'>
      <add-user v-if='showAddDialog' @cancel='showAddDialog = false' @addSucc='onAddSucc' :USER_TYPE='USER_TYPE'></add-user>
    </el-dialog>
    <el-dialog 
      title='Edit Community'
      :visible.sync='showEditDialog'
      size='small'>
      <edit-user v-if='showEditDialog' :USER='editingUser' @cancel='showEditDialog = false' @editSucc='onEditSucc'></edit-user>
    </el-dialog>
    <el-dialog
      title="提示"
      :visible.sync="showDelConfirm"
      size="tiny">
      <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
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
    props: {
      USER_TYPE: Number//用户类型
    },
    data () {
      return {
        users:[],
        showAddDialog: false,
        showEditDialog: false,
        showDelConfirm: false,
        delUser: {},
        editingUser: {},
        warn:'',
        streets:[]
      }
    },
    computed: {
      streetNames: function () {
        if (this.USER_TYPE !== 3) return []
        return this.users.map(user => {
          for (let i = 0; i < this.streets.length; i++) {
            if (this.streets[i].ID == user.StreetID) return this.streets[i].Name
          }
          return ''
        })
      }
    },
    mounted () {
      this.fetchUsers()
      if (this.USER_TYPE == 3) this.fetchAllStreets()
    },
    methods: {
      onDel (user) {
        this.delUser = user
        this.showDelConfirm = true
      },
      onEdit (user) {
        this.editingUser = user
        this.showEditDialog = true
      },
      onEditSucc () {
        this.fetchUsers()
      },
      onAddSucc () {
        this.fetchUsers()
      },
      onDelConfirm () {
        fetchpm(this, true, '/pm/users/del', {
          method: 'POST',
          body: {values: [this.delUser.ID]}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('onDelConfirm', body)
          if (body.error == 1) this.warn = body.data
          this.fetchUsers()
          this.showDelConfirm = false
        })
      },
      fetchUsers () {
        fetchpm(this, true, '/pm/users', {
          method: "POST",
          body: {type : this.USER_TYPE}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchUsers', body)
          if(body.error !== 1) this.users = body.data
        })
      },
      fetchAllStreets () {
        fetchpm(this, true, '/pm/street',{
          method: 'POST'
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then( data => {
          // console.info('fetechStreets', data)
          if (data.error === 0) {
            console.info ('fetchAllStreets',data)
            this.streets = data.data
          }
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  .warn{
    text-align: center;
    color: red;
  }
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