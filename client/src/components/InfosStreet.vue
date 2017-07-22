<template>
  <div>
    <div :class='s.addDel'>
      <button :class='s.add'>
        <img src="~@/res/images/add.png">新增
      </button>
      <button :class='s.del'>
        <img src="~@/res/images/delete.png">删除
      </button>
    </div>
    <table>
      <tr >
        <th>选择</th>
        <th>街道名</th>
        <th>负责人</th>
        <th>电话号码</th>
        <th>授权码</th>
        <th :class='s.descr'>描述</th>
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
          <button :class='s.operation'>
            <img src="~@/res/images/edit.png"/>编辑
          </button>
        </td>
      </tr>
    </table>
  </div>
</template>

<style lang="less" module='s'>
  .addDel{
    display: flex;
    align-items: center;
    margin-top: 10px;
    button{
      border: transparent;
      padding: 5px 10px;
      margin: 10px;
      font-size: 20px;
      color: #fff;
      img{
        width: 15px;
        margin-right: 5px;
      }
    }
    .add{
      background: #3598dc;
    }
    .del{
      background: #cb5a5e;
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
    .operation{
      display: flex;
      background: #26a69a;
      color: #fff;
      min-width: 60px;
      font-size: 15px;
      align-items: center;
      margin: auto;
      img{
        width: 20px;
        margin-right: 5px;
      }
      &:hover{
        background: #29847b;
      }
    }
  }
</style>

<script type="text/javascript">
  import Ajax from '@/Ajax'
  export default {
    data () {
      return {
        streets:[]
      }
    },
    mounted () {
      this.fetechStreets()
    },
    methods: {
      fetechStreets () {
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/streetInfo'
        }
        Ajax(request, data => {
          if (data === null) return
          this.streets = JSON.parse(data).Streets
        })
      }
    }
  }
</script>