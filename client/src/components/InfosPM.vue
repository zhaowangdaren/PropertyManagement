<template>
  <div>
    <div :class='s.searchWrap'>
    <!-- 街道名 -->
      <div :class='s.title'>Street Name</div>
      <div :class='s.inputWrap'>
        <search-select v-model='inputStreetName' :values='streetNames' />
      </div>
      <!-- 社区名 -->
      <div :class='s.title'>Community</div>
      <div :class='s.inputWrap'>
        <search-select v-model='inputCommunityName' :values='communityNames' />
      </div>
      <!-- 小区名 -->
      <div :class='s.title'>country</div>
      <div :class='s.inputWrap'>
        <search-select v-model='inputXiaoQu' :values='xqNames' />
      </div>
    </div>
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
      <tr v-for='item in pms' :class='s.street'>
        <td v-text='item.XiaoQu'></td>
        <td v-text='item.Name'></td>
        <td v-text='item.LegalPerson'></td>
        <td v-text='item.WuYeZiZhi'></td>
        <td v-text='item.WuYeXinZhi'></td>
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
      position: relative;
      font-size: 18px;
      flex: 1;
      margin-right: 5px;
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
  import AddPM from '@/components/dialog/AddPM'
  export default {
    components: {ImageButton, AddPM, SearchSelect},
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
        fetch(this.host + '/pm', {
          method: 'POST',
          body: '{}'
        }).then(resp =>{
          return resp.json()
        }).then(body => {
          console.info('fetchPMs', body)
          if (body.error !== 1) this.pms = body.data
        })
      },
      fetchPMKVs (kvs) {
        fetch(this.host + '/pm/kvs', {
          method: 'POST',
          body: JSON.stringify(kvs)
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchPMKVs', body)
          if (body.error !== 1) this.pms = body.data
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