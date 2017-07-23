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
      <image-button :class='s.searchBt' :clickMethod='onSearch'
        text='查询'
        :img='require("@/res/images/ic_serach.png")'
        bgColor='#4c87b9'
        color='#fff'
      />
    </div>
    <div :class='s.addDel'>
      <image-button :class='s.bt'
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
          <image-button
            text='编辑'
            :img='require("@/res/images/edit.png")'
            bgColor='#26a69a'
          />
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
  import Ajax from '@/Ajax'
  import ImageButton from '@/components/ImageButton'
  export default {
    components: {ImageButton},
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
      onSearch () {
        if(this.inputStreetName === '' && this.inputStreetName === '') return
        var params = {}
        if (this.inputStreetName !== '') params.street = this.inputStreetName
        if (this.inputXQName !== '') params.name = this.inputXQName
        var request = {
          method: 'POST',
          url: 'http://10.176.118.61:3000/xiaoQuKVs',
          query: {
            query: JSON.stringify(params)
          }
        }
        Ajax(request, data => {
          if (data === null || data === 'succ') return
          this.xiaoQus = JSON.parse(data).XiaoQus
        })
      },
      fetechStreets () {
        var request = {
          method: 'POST',
          url:'http://10.176.118.61:3000/xiaoQu'
        }
        Ajax(request, data => {
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
          if (data === null || data === 'succ') return
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