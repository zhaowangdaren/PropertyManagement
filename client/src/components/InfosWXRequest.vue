<template>
  <div>
    Request
    <table>
      <tr >
        <th>序号</th>
        <th>WXName</th>
        <!-- 用户名 -->
        <th>OpenID</th>
        <!-- OPenID -->
        <th>Tel</th>
        <!-- 电话 -->
        <th>操作</th>
      </tr>
      <tr v-for='(user, index) in users' :class='s.users'>
        <td v-text='index + 1'></td>
        <td v-text='user.UserName'></td>
        <td v-text='user.OpenID'></td>
        <td v-text='user.Tel'></td>
        <td :class='s.operations'>
          <image-button :class='s.bt'
            text='强制解绑'
            :img='require("@/res/images/delete.png")'
            bgColor='##c23f44'
            fontSize='20px'
          />
        </td>
      </tr>
      <tr v-if='users.length === 0' :class='s.noData'>
        <td colspan="5" align="center">没有记录</td>
      </tr>
    </table>
    <component :is='showDialog' @close='showDialog = ""' />
  </div>
</template>

<style lang="less" module='s'>
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
    th{
      text-align: center;
      padding: 5px;
      border: solid 1px #ddd;
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
        display: flex;
        justify-content: center;
        .bt{
          margin-left: 5px;
        }
      }
    }
    .noData{
      background: #eee;
    }
  }
</style>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  export default {
    components: {ImageButton},
    data () {
      return {
        host:'http://10.176.118.61:3000',
        users:[],
        showDialog: ''
      }
    },
    mounted () {
      this.fetechWX()
    },
    methods: {
      fetechWX () {
        // fetch(this.host + '/streetUser', {
        //   method: 'POST',
        //   body: JSON.stringify({name:'', pageNo: 1, pageSize: 1})
        // }).then(resp => {
        //   console.info(resp)
        //   return resp.json()
        // }).then( data => {
        //   console.info('fetechusers', data)
        //   if (data.error === 0) {
        //     this.users = data.data
        //   }
        // })
      },
      onAdd () {
        console.info('onAdd')
      }
    }
  }
</script>