<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        PM infos<!--  -->
      </div>
      <div :class='s.searchWrap'>
        <table>
          <tr>
            <td :class='s.key' v-if='editable'>Street Name</td>
            <td v-if='editable'>
              <el-select v-model="inputStreetName" filterable placeholder="select Street">
                <el-option
                  v-for="item in streetNames"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>Community</td>
            <td>
              <el-select v-model="inputCommunityName" filterable placeholder="select community">
                <el-option
                  v-for="item in communityNames"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>country</td>
            <td>
              <el-select v-model="inputXiaoQu" filterable placeholder="select xqNames">
                <el-option
                  v-for="item in xqNames"
                  :key="item"
                  :label="item"
                  :value="item">
                </el-option>
              </el-select>
            </td>
          </tr>
        </table>
      </div>
      <div :class='s.addDel'>
        <el-button type="primary" icon="plus" @click='onAdd' v-if='editable'>新增</el-button>
        <el-button type="primary" icon="search" @click='onSearch'>查询</el-button>
      </div>
      <div :class='s.table'>
        <table>
          <tr >
            <th>Country Name</th>
            <!-- 小区名 -->
            <th>PM Name</th>
            <!-- 物业公司 -->
            <th>LegalPerson</th>
            <!-- 独立法人 -->
            <th>WuYeZiZhi</th>
            <!-- 物业资质 -->
            <th>WuYeXinZhi</th>
            <!-- 物业性质 -->
            <th>操作</th>
          </tr>
          <tr v-for='item in pms'>
            <td v-text='item.XQ'></td>
            <td v-text='item.Name'></td>
            <td v-text='item.LegalPerson'></td>
            <td v-text='item.WuYeZiZhi'></td>
            <td v-text='item.WuYeXinZhi'></td>
            <td align="center">
              <el-button v-if='editable' type="primary" icon="edit" @click='onEdit(item)'>编辑</el-button>
              <el-button v-if='editable' type="primary" icon="delete" @click='onDel(item)'>删除</el-button>
              <el-button type='primary'>查看</el-button>
            </td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import AddPM from '@/components/dialog/AddPM'
import fetchpm from '@/fetchpm'
export default {
  components: {AddPM},
  props: {
    editable: Boolean//是否可编辑
  },
  data () {
    return {
      host:'http://10.176.118.61:3000',
      showDialog: '',
      streetNames: [],
      communityNames: [],
      xqNames: [],
      inputStreetName:'',
      inputCommunityName: '',
      inputXiaoQu:'',
      isLoadingInput: false,
      pms:[]
    }
  },
  computed: {
  },
  mounted () {
    this.fetchPMs()
    this.editable ? this.fetechAllStreetName() : this.fetchCommunitiesByStreetName('1')
  },
  watch: {
    inputStreetName: function (val) {
      if (this.isLoadingInput ) return
      this.fetchCommunitiesByStreetName(val)
    },
    inputCommunityName: function (val) {
      console.info('inputCommunityName', val)
      if (this.isLoadingInput ) return
      this.fetchXQByCommunityName(val)
    }
  },
  methods: {
    onEdit (item) {

    },
    onAdd () {
      this.showDialog = AddCountry
    },
    onSearch () {
      //小区》社区》街道
      if (this.inputXiaoQu !== '') {
        this.fetchPMKVs({xiaoQu: this.inputXiaoQu})
        return
      }
      if (this.inputCommunityName !== '') {
        this.fetchPMKVs({community: this.inputCommunityName})
        return
      }
      if (this.inputStreetName !== '') {
        this.fetchPMKVs({street: this.inputStreetName})
        return
      }
    },
    fetchPMs () {
      fetchpm(this, true, '/pm/pm', {
        method: 'POST'
      }).then(resp =>{
        return resp.json()
      }).then(body => {
        console.info('fetchPMs', body)
        if (body.error !== 1) this.pms = body.data
      })
    },
    fetchPMKVs (kvs) {
      fetchpm(this, true, '/pm/pm/kvs', {
        method: 'POST',
        body: kvs
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPMKVs', body)
        if (body.error !== 1) this.pms = body.data
      })
    },
    fetechAllStreetName () {
      fetchpm(this, true, '/pm/street/key/name', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetechAllStreetName',body)
        if (body.error !== 1) this.streetNames = body.data
      })
    },
    fetchCommunitiesByStreetName (streetName) {
      if (!streetName) return
      this.isLoadingInput = true
      this.communityNames = []
      this.inputCommunityName = ''
      fetchpm(this, true, '/pm/community/streetName/'+streetName, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then( body => {
        console.info('fetchCommunitiesByStreetName', body)
        if(body.error !== 0) {
          console.error("Error: search CommunitiesByStreetName. Reason:" + body.data)
        } else if (body.data !== null ) {
          this.communityNames = body.data.map(item => {
            return item.Name
          })
        }
        this.isLoadingInput = false
      })
    },
    fetchXQByCommunityName (communityName) {
      if (!communityName) return
      this.isLoadingInput = true
      fetchpm(this, true, '/pm/xq/kvs', {
        method: 'POST',
        body: {community: communityName}
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

<style lang='less' module='s'>
.wrap{
  .content{
    border: solid 1px #4c87b9;
    margin-top: 50px;
    .title{
      color: #fff;
      font-size: 20px;
      background: #4c87b9;
      padding: 10px;
      display: flex;
      align-items: center;
      img{
        width: 20px;
        margin-right: 10px;
      }
    }
    .searchWrap{
      display: flex;
      align-items: center;
      position: relative;
      margin: 10px;
      table{
        width: 100%;
        background-color: #fff;
        td{
          text-align: center;
          padding: 2px;
          border: solid 1px #ddd;
        }
        .key{
          background-color: #ddd;
        }
      }
    }
    .addDel {
      display: flex;
      justify-content: flex-end;
      margin: 10px;
    }
    .table{
      margin: 10px;
      table{
        width: 100%;
        font-size: 15px;
        color: #555;
        margin: 10px auto;
        background-color: #fff;
        th{
          text-align: center;
          padding: 5px;
          border: solid 1px #ddd;
          background-color: #ddd;
        }
        td{
          padding: 5px;
          border: solid 1px #ddd;
        }
      }
    }
  }
}
</style>