<template>
  <div :class='s.wrap'>
    <div :class='s.addDel'>
      <image-button :class='s.bt' :clickMethod='onAdd'
        text='新增'
        :img='require("@/res/images/add.png")'
        bgColor='#3598dc'
        fontSize='20px'
        />
    </div>
    <table>
      <tr >
        <th>UserName</th>
        <!-- 用户名 -->
        <th>RealName</th>
        <!-- 真实姓名 -->
        <th>Street</th>
        <!-- 街道 -->
        <th>Password</th>
        <!-- 密码 -->
        <th>Tel</th>
        <!-- 电话 -->
        <th>操作</th>
      </tr>
      <tr v-for='user in users' :class='s.users'>
        <td v-text='user.UserName'></td>
        <td v-text='user.RealName'></td>
        <td v-text='user.Street'></td>
        <td v-text='user.Password'></td>
        <td v-text='user.Tel'></td>
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
      <tr v-if='users.length == 0'>
        <td colspan="6">无记录</td>
      </tr>
    </table>
    <component :is='showDialog' @close='showDialog = ""' />
  </div>
</template>

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
    margin: 10px auto;
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
    .users{
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

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import AddStreetUser from '@/components/dialog/AddStreetUser'
  import fetchpm from '@/fetchpm'
  export default {
    components: {ImageButton, AddStreetUser},
    data () {
      return {
        users:[],
        showDialog: ''
      }
    },
    mounted () {
      this.fetechUsers()
    },
    methods: {
      fetechUsers () {
        fetchpm(this, true, '/pm/users', {
          method: 'POST',
          body: {type: 3}
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then( data => {
          console.info('fetechusers', data)
          if (data.error === 0 && data.data !== null) {
            this.users = data.data
          }
        })
      },
      onAdd () {
        console.info('onAdd')
        this.showDialog = AddStreetUser
      }
    }
  }
</script>