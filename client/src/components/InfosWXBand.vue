<template>
  <div>
    <div :class='s.searchWrap'>
      WXName
      <!-- 微信用户绑定名称（包含） -->
      <search-select v-model='wxName' :values='wxNames' :class='s.searchSelect'/>
      <image-button :class='s.searchBt' :clickMethod='onSearch'
        text='查询'
        :img='require("@/res/images/ic_serach.png")'
        bgColor='#4c87b9'
        color='#fff'
      />
    </div>
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
        <td v-text='user.WXName'></td>
        <td v-text='user.OpenID'></td>
        <td v-text='user.Tel'></td>
        <td align="center">
          <image-button
            text='强制解绑'
            :img='require("@/res/images/delete.png")'
            bgColor='#c23f44'
            fontSize='20px'
          />
        </td>
      </tr>
    </table>
    <component :is='showDialog' @close='showDialog = ""' />
  </div>
</template>

<style lang="less" module='s'>
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
</style>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import SearchSelect from '@/components/SearchSelect'
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
        fetch(this.host + '/wxUser', {
          method: 'POST',
          body: JSON.stringify({name:name, pageNo: 1, pageSize: 1})
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
        fetch(this.host + '/wxUser/key/wxName', {
          method: 'POST',
          body: '{}'
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