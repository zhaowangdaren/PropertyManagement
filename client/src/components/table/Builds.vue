<template>
  <div :class='s.wrap'>
    <table :class='s.searchWrap'>
      <tr>
        <td :class='s.title'>Street Name</td>
        <td>
          <el-select v-model="inputStreetName" filterable placeholder="全部">
            <el-option
              v-for="item in streetNames"
              :key="item"
              :label="item"
              :value="item">
            </el-option>
          </el-select>
        </td>
        <td :class='s.title'>Community</td>
        <td>
          <el-select v-model="inputCommunityName" filterable placeholder="全部">
            <el-option
              v-for="item in communityNames"
              :key="item"
              :label="item"
              :value="item">
            </el-option>
          </el-select>
        </td>
        <td :class='s.title'>Country</td>
        <td>
          <el-select v-model="inputXiaoQu" filterable placeholder="全部">
            <el-option
              v-for="item in xqNames"
              :key="item"
              :label="item"
              :value="item">
            </el-option>
          </el-select>
        </td>
      </tr>
      <tr>
        <td :class='s.title'>HouseBuildNo</td>
        <td>
        <!-- 房屋楼栋号 -->
          <el-input v-model="inputHouseBuildNo" placeholder="请输入号"></el-input>
        </td>
        <td :class='s.title'>HouseNo</td>
        <td>
          <el-input v-model="inputHouseNo" placeholder="请输入HouseNo"></el-input>
        </td>
        <td :class='s.title'>Owner</td>
        <td>
          <el-input v-model="inputOwner" placeholder="请输入owner"></el-input>
        </td>
      </tr>
    </table>
    <div :class='s.addDel'>
      <el-button v-if='editable' type="primary" icon="plus" @click='onAdd'>新增</el-button>
      <el-button type="primary" icon="search" @click='onSearch'>查询</el-button>
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
      <tr v-for='item in houses'>
        <td v-text='item.Owner'></td>
        <td v-text='item.XQ'></td>
        <td v-text='item.HouseBuildNo'></td>
        <td v-text='item.HouseNo'></td>
        <td v-text='item.HouseType'></td>
        <td align="center">
          <el-button v-if='editable' type="primary" icon="edit" @click='onEdit(item)'>编辑</el-button>
          <el-button v-if='editable' type="primary" icon="delete" @click='onDel(item)'>删除</el-button>
          <el-button type='primary'>查看</el-button>
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

<script type="text/javascript">
  import ImageButton from '@/components/ImageButton'
  import SearchSelect from '@/components/SearchSelect'
  import AddBuild from '@/components/dialog/AddBuild'
  export default {
    props: {
      editable: Boolean
    },
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

<style lang="less" module='s'>
.wrap{
  .searchWrap{
    background-color: #fff;
    border: solid 1px #ddd;
    .title{
      background: #eee;
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
    background-color: #fff;
    th{
      text-align: center;
      padding: 5px;
      border: solid 1px #ddd;
      background: #eee;
    }
    td{
      padding: 5px;
      border: solid 1px #ddd;
    }
    tr{
      &:hover {
        background-color: #ddd;
      }
      .bt{
        display: inline-block;
      }
    }
  }
}
</style>