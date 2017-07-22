<template>
  <div>
    <div :class='s.searchWrap'>
      <div :class='s.title'>街道名</div>
      <div :class='s.inputWrap'>
        <input type="text" v-model='inputStreetName'>
        <div :class='s.dropList'>
          <div v-for='streetName in showNames' v-text='streetName' :class='s.streetName' @click='inputStreetName = streetName'></div>
        </div>
      </div>
      <div :class='s.title'>小区名</div>
      <div :class='s.inputWrap'>
        <input type="text" v-model='inputXQName'>
        <div :class='s.dropList'>
          <div v-for='item in showXQNames' v-text='item' :class='s.streetName' @click='inputXQName = item'></div>
        </div>
      </div>
    </div>
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
        <th>社区名</th>
        <th>小区名</th>
        <th>位置</th>
        <th>负责人</th>
        <th>电话号码</th>
        <th :class='s.descr'>描述</th>
        <th>操作</th>
      </tr>
      <tr v-for='item in xiaoQus' :class='s.street'>
        <td>
          <input type="checkbox" v-model='item.checked'>
        </td>
        <td v-text='item.Street'></td>
        <td v-text='item.Community'></td>
        <td v-text='item.Name'></td>
        <td v-text='item.Address'></td>
        <td v-text='item.PersonInCharge'></td>
        <td v-text='item.Tel'></td>
        <td v-text='item.Intro'></td>
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
  .searchWrap{
    border: solid 1px #ddd;
    display: flex;
    align-items: center;
    .title{
      background: #e0e0e0;
      padding: 10px;
      font-size: 20px;
    }
    .inputWrap{
      border: solid 1px #ddd;
      position: relative;
      font-size: 18px;
      flex: 0.5;
      margin-left: 10px;
      input{
        width: 95%;
        border: solid 0px transparent;
        padding: 5px;
      }
      &:hover .dropList{
        display: block;
      }
      .dropList{
        position: absolute;
        width: 100%;
        background: #fff;
        box-shadow: 1px 1px 1px 1px #ddd;
        display: none;
        .streetName{
          padding: 5px;
          &:hover{
            background: #ddd;
          }
        }
      }

    }
    select{
      border: solid 1px #ddd;
      min-width: 200px;
      height: 30px;
      margin-left: 10px;
      font-size: 15px;
      background: #fff;
      border-radius: 0;
    }
  }
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
        xiaoQus:[],
        streetNames: [],
        xqNames: [],
        inputStreetName:'',
        inputXQName:''
      }
    },
    computed: {
      showNames: function () {
        return this.streetNames.filter(value => {
          var index = value.indexOf(this.inputStreetName) > -1
          return index
        })
      },
      showXQNames: function () {
        return this.xqNames.filter(value => {
          var index = value.indexOf(this.inputXQName) > -1
          return index
        })
      }
    },
    mounted () {
      this.fetechStreets()
      this.fetechAllStreetName()
      this.fetechAllXQName()
    },
    methods: {
      fetechStreets () {
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/xiaoQu'
        }
        Ajax(request, data => {
          console.info(data)
          this.xiaoQus = JSON.parse(data).XiaoQus
        })
      },
      fetechAllStreetName () {
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/communityInfo/key',
          query: {
            key: 'streetname'
          }
        }
        Ajax(request, data => {
          if (data === null) return
          this.streetNames = JSON.parse(data)
        console.info(data)
          if (this.streetNames === null) return
        })
      },
      fetechAllXQName () {
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/xiaoQu/key',
          query: {
            key: 'name'
          }
        }
        Ajax(request, data => {
          if (data === null) return
            console.info(data)
          this.xqNames = JSON.parse(data)
        })
      }
    }
  }
</script>