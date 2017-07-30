<template>
  <div>
    <div :class='s.searchWrap'>
    <!-- 街道名 -->
      <div :class='s.title'>name</div>
      <div :class='s.inputWrap'>
        <input type="text" v-model='inputStreetName'>
        <div :class='s.dropList'>
          <div v-for='streetName in showNames' v-text='streetName' :class='s.streetName' @click='inputStreetName = streetName'></div>
        </div>
      </div>
      <!-- 小区名 -->
      <div :class='s.title'>country</div>
      <div :class='s.inputWrap'>
        <input type="text" v-model='inputXQName'>
        <div :class='s.dropList'>
          <div v-for='item in showXQNames' v-text='item' :class='s.streetName' @click='inputXQName = item'></div>
        </div>
      </div>
      <image-button :class='s.searchBt' :clickMethod='onSearch'
        text='查询'
        :img='require("@/res/images/ic_serach.png")'
        bgColor='#4c87b9'
        color='#fff'
      />
    </div>
    <div :class='s.addDel'>
      <image-button :class='s.bt' :clickMethod='onAdd'
        text='新增'
        :img='require("@/res/images/add.png")'
        bgColor='#3598dc'
        fontSize='20px'
        />
      <image-button :class='s.bt'
        text='删除'
        :img='require("@/res/images/delete.png")'
        bgColor='#cb5a5e'
        fontSize='20px'
        />
    </div>
    <table>
      <tr >
        <th>选择</th>
        <!-- 街道名 -->
        <th>name</th>
        <!-- 社区名 -->
        <th>name2</th>
        <!-- 小区名 -->
        <th>country</th>
        <!-- 位置 -->
        <th>address</th>
        <!-- 负责人 -->
        <th>charger</th>
        <!-- 电话号码 -->
        <th>tel</th>
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
          <image-button
            text='编辑'
            :img='require("@/res/images/edit.png")'
            bgColor='#26a69a'
          />
        </td>
      </tr>
    </table>
    <component :is='showDialog' @close='showDialog = ""' />
  </div>
</template>

<style lang="less" module='s'>
  .searchWrap{
    border: solid 1px #ddd;
    display: flex;
    align-items: center;
    width: 100%;
    position: relative;
    .title{
      background: #e0e0e0;
      padding: 10px;
      font-size: 20px;
    }
    .inputWrap{
      border: solid 1px #ddd;
      position: relative;
      font-size: 18px;
      flex: 0.4;
      margin: 0 10px;
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
        z-index: 2;
        .streetName{
          padding: 5px;
          &:hover{
            background: #ddd;
          }
        }
      }
    }
    .searchBt{
      position: absolute;
      right: 10px;
    }
  }
  .addDel{
    display: flex;
    align-items: center;
    margin-top: 10px;
    .bt{
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
  }
</style>

<script type="text/javascript">
  import ImageButton from '@/components/ImageButton'
  import AddCountry from '@/components/dialog/AddCountry'
  import fetchpm from '@/fetchpm'

  export default {
    components: {ImageButton, AddCountry},
    data () {
      return {
        host:'http://10.176.118.61:3000',
        showDialog: '',
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
      this.fetechXQs()
      this.fetechAllStreetName()
      this.fetechAllXQName()
    },
    methods: {
      onAdd () {
        this.showDialog = AddCountry
      },
      onSearch () {
        if(this.inputStreetName === '' && this.inputStreetName === '') return
        var params = {}
        if (this.inputStreetName !== '') params.street = this.inputStreetName
        if (this.inputXQName !== '') params.name = this.inputXQName
        fetchpm(this, true, '/pm/xq/kvs', {
          method: 'POST',
          body: params
        }).then( resp => {
          return resp.json()
        }).then( data => {
          console.info('fetchXQKVs', data)
          if (data.error !== 1) this.xiaoQus = data.data
        })
      },
      fetechXQs () {
        // var request = {
        //   method: 'POST',
        //   url: this.host + '/xq'
        // }
        // Ajax(request, data => {
        //   this.xiaoQus = JSON.parse(data).XiaoQus
        // })
        fetchpm(this,true,'/pm/xq', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(data => {
          console.info(data)
          this.xiaoQus = data.data
        })
      },
      fetechAllStreetName () {
        // var request = {
        //   method: 'POST',
        //   url:'http://10.176.118.61:3000/communityInfo/key',
        //   query: {
        //     key: 'streetname'
        //   }
        // }
        // Ajax(request, data => {
        //   if (data === null || data === 'succ') return
        //   this.streetNames = JSON.parse(data)
        // console.info(data)
        //   if (this.streetNames === null) return
        // })
        fetchpm(this, true,'/pm/street/key/name', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(data => {
          console.info('fetechAllStreetName',data)
          this.streetNames = data.data
        })
      },
      fetechAllXQName () {
        // var request = {
        //   method: 'POST',
        //   url:'http://10.176.118.61:3000/xiaoQu/key',
        //   query: {
        //     key: 'name'
        //   }
        // }
        // Ajax(request, data => {
        //   if (data === null) return
        //     console.info(data)
        //   this.xqNames = JSON.parse(data)
        // })
        fetchpm(this, true,'/pm/xq/key/name', {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then(data => {
          console.info('fetechAllXQName', data)
          this.xqNames = data.data
        })
      }
    }
  }
</script>