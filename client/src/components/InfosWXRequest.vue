<template>
  <div :class='s.wrap'>
    <table>
      <tr >
        <th>序号</th>
        <th>微信名</th>
        <!-- 用户名 -->
        <th>OpenID</th>
        <!-- OPenID -->
        <th>电话</th>
        <!-- 电话 -->
        <th>操作</th>
      </tr>
      <tr v-for='(user, index) in users' :class='s.users'>
        <td v-text='index + 1'></td>
        <td v-text='user.UserName'></td>
        <td v-text='user.OpenID'></td>
        <td v-text='user.Tel'></td>
        <td :class='s.operations'>
          <el-button type='success'>允许</el-button>
          <el-button type='warning'>拒绝</el-button>
        </td>
      </tr>
      <tr v-if='users.length === 0' :class='s.noData'>
        <td colspan="5" align="center">无记录</td>
      </tr>
    </table>
    <component :is='showDialog' @close='showDialog = ""' />
  </div>
</template>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import fetchpm from '@/fetchpm'
  export default {
    components: {ImageButton},
    data () {
      return {
        users:[],
        showDialog: ''
      }
    },
    mounted () {
      this.fetechWX()
    },
    methods: {
      fetechWX () {
        fetchpm(this, true, '/pm/wxUser', {
          method: 'POST',
          body: {name:name, pageNo: 1, pageSize: 1}
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then( data => {
          console.info('fetechusers', data)
          if (data.error === 0 ) {
            this.users = data.data || []
          }
        })
      },
      onAdd () {
        console.info('onAdd')
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
    margin: 10px auto;
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
      .operations{
        display: flex;
        justify-content: center;
        .bt{
          margin-left: 5px;
        }
      }
    }
  }
}
</style>
