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
        <td v-text='gov.OfficeNumber'></td>
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

<script type="text/javascript">
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import AddGov from '@/components/dialog/AddGov'
  export default {
    components: {ImageButton, AddGov},
    data () {
      return {
        host:'http://10.176.118.61:3000',
        govs:[],
        showDialog: ''
      }
    },
    mounted () {
      this.fetchGovs()
    },
    methods: {
      fetchGovs () {
        fetch(this.host + '/gov', {
          method: "POST",
          body: '{}'
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