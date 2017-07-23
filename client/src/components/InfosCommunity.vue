<template>
  <div>
    <div :class='s.searchWrap'>
      <div :class='s.title'>name</div>
      <!-- 街道名 -->
      <div :class='s.inputWrap'>
        <input type="text" v-model='inputName'>
        <div :class='s.dropList'>
          <div :class='s.streetName + " " + s.all' @click='fetchCommunities'>全部</div>
          <div v-for='streetName in showNames' v-text='streetName' :class='s.streetName' @click='inputName = streetName'></div>
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
      <tr>
        <th>选择</th>
        <th>name</th>
        <!-- 街道名 -->
        <th>name2</th>
        <!-- 社区名 -->
        <th>charger</th>
        <!-- 负责人 -->
        <th>tel</th>
        <!-- 电话号码 -->
        <th :class='s.descr'>描述</th>
        <th>操作</th>
      </tr>
      <tr v-for='item in communities' :class='s.street'>
        <td>
          <input type="checkbox" v-model='item.checked'>
        </td>
        <td v-text='item.StreetName'></td>
        <td v-text='item.Name'></td>
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
    <!-- <add-community /> -->
  </div>
</template>

<style lang="less" module='s'>
  .searchWrap{
    border: solid 1px #ddd;
    display: flex;
    align-items: center;
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
        z-index: 2;
        .streetName{
          padding: 5px;
          color: #555;
          &:hover{
            background: #0593f5;
          }
        }
        .all{
          background: #ddd;
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
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  import AddCommunity from '@/components/dialog/AddCommunity'
  export default {
    components: {ImageButton, AddCommunity},
    data () {
      return {
        communities:[],
        streetNames: [],
        inputName:'',
        showDialog: ''
      }
    },
    mounted () {
      this.fetchCommunities('')
      this.fetchAllStreetName()
    },
    computed: {
      showNames: function () {
        return this.streetNames.filter(value => {
          var index = value.indexOf(this.inputName) > -1
          return index
        })
      }
    },
    methods: {
      onSearch () {
        if (!this.inputName && this.inputName === '') return
        this.fetchKVs()
      },
      fetchCommunities () {
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/communityInfo'
        }
        Ajax(request, data => {
          if (data === null) return
          this.communities = JSON.parse(data).Communities
          console.info(this.communities)
        })
      },
      fetchAllStreetName () {//获取所有街道名称
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/communityInfo/key',
          query: {
            key: 'streetname'
          }
        }
        Ajax(request, data => {
          if (data === null) return
          if (data === 'succ') return
          this.streetNames = JSON.parse(data)
          if (this.streetNames === null) return
        })
      },
      fetchKVs () {
        console.info('fetchKVs')
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/communityKVs',
          query: {
            query: JSON.stringify({streetName: this.inputName})
          }
        }
        Ajax(request, data => {
          if (data === null || data === 'succ' || data === '') return
          this.communities = JSON.parse(data).Communities
        })
      },
      onAdd () {
        console.info('onAdd')
        this.showDialog = AddCommunity
      }
    }
  }
</script>