<template>
  <div :class='s.wrap'>
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
    <div :class='s.tabs'>
      <div :class='s.tab + " " + (tab.view === curView ? s.active : "")' 
            v-for='tab in tabs' 
            @click='onTab(tab)'>
          {{tab.name}}
      </div>
    </div>
    <div :class='s.content' v-if='curView == 0'>
      <div :class='s.item' >
        街道名
        <!-- 用户名 -->
        <el-input :class='s.elInput' disabled v-model="street.Name" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item' >
        社区名
        <!-- 用户名 -->
        <el-input :class='s.elInput' disabled v-model="community.Name" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item' >
        小区名
        <el-input :class='s.elInput' disabled v-model="xq.Name" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        物业公司名称
        <!-- 用户名 -->
        <el-input :class='s.elInput' disabled v-model="pm.Name" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        独立法人
        <!-- 独立法人 -->
        <el-input :class='s.elInput' disabled v-model="pm.LegalPerson" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        物业资质
        <!-- 独立法人 -->
        <el-input :class='s.elInput' disabled v-model="pm.WuYeZiZhi" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        物业性质
        <!-- 独立法人 -->
        <el-input :class='s.elInput' disabled v-model="pm.WuYeXinZhi" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        联系电话
        <!-- 独立法人 -->
        <el-input :class='s.elInput' disabled v-model="pm.Tel" @focus='onFocus'></el-input>
      </div>
    </div>
    <!-- Tab 2 -->
    <!-- Tab 2 -->
    <div :class='s.content' v-if='curView == 1'>
      <div :class='s.item'>
        小区环境
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.XQEnv" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        小区保洁
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.XQCleaning" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        小区绿化
        <!-- 绿化 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.GreenVegetatio" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        管养保护
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.GuanYangBaoHu" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        消防检查
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.XiaoFanJianCha" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        电梯保养
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.DianTiBaoYang" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        车辆有序停放
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.CarParkInOrder" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        业主委员会
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.YeZhuCommunity" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        业主委员会联系方式
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.YeZhuCommunityTel" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        上一年物业等级
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.PastYearLevel" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        建议整改措施
        <!-- 独立法人 -->
        <el-input type='textarea' disabled :class='s.elInput' v-model="pm.JianYiZhengGaiCuoShi" @focus='onFocus'></el-input>
      </div>
    </div>
    <div :class='s.content' v-show='curView == 2'>
      <div v-show='curView == 2'>
        <div ref='map' :class='s.map'>Map</div>
      </div>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  props: {
    pm: Object
  },
  data () {
    return {
      warn: '',
      tabs: [
        {
          name: '基本信息',//基本信息
          view: 0
        },
        {
          name: '详细信息',//详细信息
          view: 1
        },
        {
          name: '小区位置',//详细信息
          view: 2
        }
      ],
      curView: 0,
      street: {},
      community: {},
      xq: {
        Address: ''
      },
      inputStreetID: '',
      inputCommunityID:'',
      inputXQID: '',
      map: null,
      placeSearch: null
    }
  },
  mounted () {
    this.fetchStreet(this.pm.StreetID)
    this.fetchCommunity(this.pm.CommunityID)
    this.fetchXQ(this.pm.XQID)
  },
  watch: {
    pm: function (val) {
      this.fetchStreet(this.pm.StreetID)
      this.fetchCommunity(this.pm.CommunityID)
      this.fetchXQ(this.pm.XQID)
    },
    xq: function (val) {
      this.placeSearch.search(this.xq.Address, (status, result) => {
          console.info('placeSearch status'+ status, result)
        })
    }
  },
  computed: {
  },
  methods: {
    onTab (tab) {
      this.curView = tab.view
      this.$nextTick(() => {
        if (this.curView == 2) this.initMap()
      })
    },
    initMap () {
      if (this.map !== null) return 
      this.map = new AMap.Map(this.$refs.map, {
        zoom: 10
      })
      AMap.service(['AMap.PlaceSearch'], () => {
        this.placeSearch = new AMap.PlaceSearch({
          pageSize: 5,
          pageIndex: 1,
          city: '0555',
          map: this.map
        })
        this.placeSearch.search(this.xq.Address, (status, result) => {
          console.info('placeSearch status'+ status, result)
        })
      })
      
    },
    fetchStreet (id) {
      fetchpm(this, true, '/pm/street/ids',{
          method: 'POST',
          body: {values: [id]}
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.street = data.data[0] || {}
        }
      })
    },
    fetchCommunity (id) {
      if ( !id || id == '') return null
      fetchpm( this, true, '/pm/community/ids', {
        method: 'POST',
        body: {values: [id]}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunities', body)
        this.community = body.data[0] || {}
      })
    },
    fetchXQ (id) {
      if ( !id || id == '') return null
      fetchpm( this, true, '/pm/xq/ids', {
        method: 'POST',
        body: {values: [id]}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllXQByCommunityID', body)
        this.xq = body.data[0] || {}
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  width: 100%;
  .warn{
    text-align: center;
    color: red;
  }
  .tabs{
    color: #5b9bd1;
    font-size: 15px;
    display: flex;
    align-items: center;
    cursor: default;
    .tab{
      padding: 10px;
      background: #fff;
    }
    .active{
      border-top: solid 2px red;
      border-left: solid 1px #ddd;
      border-right: solid 1px #ddd;
    }
  }
  .content{
    border: solid 1px #ddd;
    margin-top: -1px;
    padding: 10px;
    .item{
      display: flex;
      justify-content: flex-end;
      align-items: center;
      padding: 10px 0px;
      font-size: 18px;
      .red{
        color: red;
      }
      .elInput{
        width: 75%;
        margin-left: 10px
      }
    }
    .map{
      width: 100%;
      height: 300px;
    }
  }
  .bts{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    padding: 10px 0px;
    font-size: 18px;
  }
}  
</style>