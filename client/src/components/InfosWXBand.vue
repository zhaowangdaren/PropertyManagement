<template>
  <div :class='s.wrap'>
    <div :class='s.searchWrap'>
      微信用户绑定名称(包含)
      <!-- 微信用户绑定名称（包含） -->
      <search-select v-model='wxName' :values='wxNames' :class='s.searchSelect'/>
      <!-- <image-button :class='s.searchBt' :clickMethod='onSearch'
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
    <table>
      <tr >
        <th>序号</th>
        <th>微信绑定名</th>
        <!-- 用户名 -->
        <th>OpenID</th>
        <!-- OPenID -->
        <th>电话</th>
        <!-- 电话 -->
        <th>操作</th>
      </tr>
      <tr v-for='(user, index) in users' :class='s.users'>
        <td v-text='index + 1'></td>
        <td v-text='user.WXName'></td>
        <td v-text='user.OpenID'></td>
        <td v-text='user.Tel'></td>
        <td align="center">
          <el-button type="danger" icon="delete">强制解绑</el-button>
        </td>
      </tr>
    </table>
    <component :is='showDialog' @close='showDialog = ""' />
  </div>
</template>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import SearchSelect from '@/components/SearchSelect'
  import fetchpm from '@/fetchpm'

  export default {
    components: {ImageButton, SearchSelect},
    data () {
      return {
        host:'http://10.176.118.61:3000',
        users:[],
        showDialog: '',
        wxName: '',
        wxNames: []
      }
    },
    mounted () {
      this.fetechWXUsers('')
      this.fetechAllWXUsersName()
    },
    methods: {
      fetechWXUsers (name) {
        fetchpm(this, true, '/pm/wxUser', {
          method: 'POST',
          body: {name:name, pageNo: 1, pageSize: 1}
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then( data => {
          console.info('fetechusers', data)
          if (data.error === 0 ) {
            this.users = data.data
          }
        })
      },
      fetechAllWXUsersName () {
        fetchpm(this, true, '/pm/wxUser/key/wxName', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetechAllWXUsersName', body)
          if (body.error !== 1) this.wxNames = body.data
        })
      },
      onSearch () {
        if (this.wxName === '') return
        this.fetechWXUsers(this.wxName)
      },
      onAdd () {
        console.info('onAdd')
        this.showDialog = AddStreetUser
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  .searchWrap{
    display: flex;
    align-items: center;
    margin-top: 10px;
    width: 100%;
    position: relative;
    .searchSelect {
      flex: 0.8;
    }
    .searchBt{
      position: absolute;
      right: 10px;
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
    }
  }
}
</style>
