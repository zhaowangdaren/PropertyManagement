<template>
  <div>
    <table :class='s.searchWrap'>
      <tr>
        <td :class='s.title'>Street Name</td>
        <td>
          <search-select v-model='inputStreetName' :values='streetNames' />
        </td>
        <td :class='s.title'>Community</td>
        <td>
          <search-select v-model='inputCommunityName' :values='communityNames' />
        </td>
        <td :class='s.title'>Country</td>
        <td>
          <search-select v-model='inputXiaoQu' :values='xqNames' />
        </td>
      </tr>
      <tr>
        <td :class='s.title'>HouseBuildNo</td>
        <td>
          <div :class='s.inputWrap'>
            <input type="text" v-model='inputHouseBuildNo' :class=''>
          </div>
        </td>
        <td :class='s.title'>HouseNo</td>
        <td>
          <div :class='s.inputWrap'>
            <input type="text" v-model='inputHouseNo'>
          </div>
        </td>
        <td :class='s.title'>Owner</td>
        <td>
          <div :class='s.inputWrap'>
            <input type="text" v-model='inputOwner'>
          </div>
        </td>
      </tr>
    </table>
    <div :class='s.addDel'>
      <image-button :class='s.bt' :clickMethod='onAdd'
        text='新增'
        :img='require("@/res/images/add.png")'
        bgColor='#3598dc'
        />
      
      <image-button :class='s.searchBt' :clickMethod='onSearch'
        text='查询'
        :img='require("@/res/images/ic_serach.png")'
        bgColor='#4c87b9'
        color='#fff'
      />
    </div>
    <table>
      <tr >
        <!-- 房产登记人 -->
        <th>Owner</th>
        <!-- 所在小区 -->
        <th>XQ</th>
        <!-- 楼栋号 -->
        <th>HouseBuildNo</th>
        <!-- 门牌号 -->
        <th>HouseNo</th>
        <!-- 房屋类型 -->
        <th>HouseType</th>
        <th>操作</th>
      </tr>
      <tr v-for='item in houses' :class='s.street'>
        <td v-text='item.Owner'></td>
        <td v-text='item.XQ'></td>
        <td v-text='item.HouseBuildNo'></td>
        <td v-text='item.HouseNo'></td>
        <td v-text='item.HouseType'></td>
        <td align="center">
          <image-button :class='s.bt'
            text='编辑'
            :img='require("@/res/images/edit.png")'
            bgColor='#26a69a'
          />
          <image-button :class='s.bt'
            text='删除'
            :img='require("@/res/images/delete.png")'
            bgColor='#cb5a5e'
            />
        </td>
      </tr>
    </table>
    <el-dialog 
      title='Add Build'
      size='large'
      :visible.sync="showAdd"
    >
      <add-build @close='showAdd = false'
      @save='showAdd = false'></add-build>
    </el-dialog>
  </div>
</template>

<style lang="less" module='s'>
  .searchWrap{
    border: solid 1px #ddd;
    .title{
      background: #f0f0f0;
      padding: 10px;
      font-size: 20px;
    }
    .inputWrap {
        position: relative;
        font-size: 18px;
        flex: 1;
        margin-left: 10px;
      input{
        width: 100%;
        border: solid 1px #ddd;
        height: 30px;
      }
    } 
    .searchBt{
      position: absolute;
      right: 10px;
    }
  }
  .addDel{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    margin-top: 10px;
    margin-right: 10px;
    .bt{
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
      background: #f0f0f0;
    }
    td{
      padding: 5px;
      border: solid 1px #ddd;
    }
    .street{
      &:hover {
        background-color: #ddd;
      }
      .bt{
        display: inline-block;
      }
    }
  }
</style>

<script type="text/javascript">
  import ImageButton from '@/components/ImageButton'
  import SearchSelect from '@/components/SearchSelect'
  import AddBuild from '@/components/dialog/AddBuild'
  export default {
    components: {ImageButton, AddBuild, SearchSelect},
    data () {
      return {
        host:'http://10.176.118.61:3000',
        showAdd: false,
        streetNames: [],
        communityNames: [],
        xqNames: [],
        inputStreetName:'',
        inputCommunityName: '',
        inputXiaoQu:'',
        inputHouseBuildNo:'',
        inputHouseNo:'',
        inputOwner:'',
        isLoadingInput: false,
        houses:[]
      }
    },
    computed: {
    },
    mounted () {
      this.fetchHouses()
      this.fetechAllStreetName()
    },
    watch: {
      inputStreetName: function (val) {
        if (this.isLoadingInput ) return
        this.fetchCommunitiesByStreetName(val)
      },
      inputCommunityName: function (val) {
        if (this.isLoadingInput ) return
        this.fetchXQByCommunityName(val)
      }
    },
    methods: {
      onAdd () {
        this.showAdd = true
      },
      onSearch () {
        //小区》社区》街道
        if (this.inputXiaoQu !== '') {
          this.fetchHouseKVs({xiaoQu: this.inputXiaoQu})
          return
        }
        if (this.inputCommunityName !== '') {
          this.fetchHouseKVs({community: this.inputCommunityName})
          return
        }
        if (this.inputStreetName !== '') {
          this.fetchHouseKVs({street: this.inputStreetName})
          return
        }
      },
      fetchHouses () {
        fetch(this.host + '/house', {
          method: 'POST',
          body: '{}'
        }).then(resp =>{
          return resp.json()
        }).then(body => {
          console.info('fetchHouses', body)
          if (body.error !== 1) this.houses = body.data
        })
      },
      fetchHouseKVs (kvs) {
        fetch(this.host + '/pm/kvs', {
          method: 'POST',
          body: JSON.stringify(kvs)
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchHouseKVs', body)
          if (body.error !== 1) this.houses = body.data
        })
      },
      fetechAllStreetName () {
        fetch(this.host + '/street/key/name', {
          method: 'POST',
          body: '{}'
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetechAllStreetName',body)
          if (body.error !== 1) this.streetNames = body.data
        })
      },
      fetchCommunitiesByStreetName (streetName) {
        this.isLoadingInput = true
        if (!streetName) return
        fetch(this.host + '/community/streetName/'+streetName, {
          method: 'POST'
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchCommunitiesByStreetName', body)
          if(body.error !== 0) console.error("Error: search CommunitiesByStreetName. Reason:" + body.data)
          else if (body.data !== null ) {
            this.communityNames = body.data.map(item => {
              return item.Name
            })
          }
          this.isLoadingInput = false
        })
      },
      fetchXQByCommunityName (communityName) {
        this.isLoadingInput = true
        if (!communityName) return
        fetch(this.host + '/xq/kvs', {
          method: 'POST',
          body: JSON.stringify({community: communityName})
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchXQByCommunityName', body)
          if(body.error !== 0) console.error("Error: search fetchXQByCommunityName. Reason:" + body.data)
          else if (body.data !== null ) {
            this.xqNames = body.data.map(item => {
              return item.Name
            })
          }
          this.isLoadingInput = false
        })
      }
    }
  }
</script>