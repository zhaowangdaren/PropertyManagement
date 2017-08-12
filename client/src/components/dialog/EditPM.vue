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
        <div :class='s.red'>*</div>
        街道
        <!-- 用户名 -->
        <el-select :class='s.elInput' v-model="inputStreetID" placeholder="请选择">
            <el-option
              v-for="item in streets"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
      </div>
      <div :class='s.item' >
        <div :class='s.red'>*</div>
        社区
        <!-- 用户名 -->
        <el-select :class='s.elInput' v-model="inputCommunityID" placeholder="请选择">
          <el-option
            v-for="item in communities"
            :key="item.ID"
            :label="item.Name"
            :value="item.ID">
          </el-option>
        </el-select>
      </div>
      <div :class='s.item' >
        <div :class='s.red'>*</div>
        小区
        <!-- 用户名 -->
        <el-select :class='s.elInput' v-model="inputXQID" placeholder="请选择">
          <el-option
            v-for="item in xqs"
            :key="item.ID"
            :label="item.Name"
            :value="item.ID">
          </el-option>
        </el-select>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        物业公司
        <!-- 用户名 -->
        <el-input :class='s.elInput' v-model="pm.Name" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        独立法人
        <!-- 独立法人 -->
        <el-input :class='s.elInput' v-model="pm.LegalPerson" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        物业资质
        <!-- 独立法人 -->
        <el-input :class='s.elInput' v-model="pm.WuYeZiZhi" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        物业性质
        <!-- 独立法人 -->
        <el-input :class='s.elInput' v-model="pm.WuYeXinZhi" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        联系电话
        <!-- 独立法人 -->
        <el-input :class='s.elInput' v-model="pm.Tel" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
    </div>
    <!-- Tab 2 -->
    <!-- Tab 2 -->
    <div :class='s.content' v-if='curView == 1'>
      <div :class='s.item'>
        小区环境
        <!-- 独立法人 -->
        <el-input 
          type='textarea'
          :class='s.elInput' 
          v-model="pm.XQEnv" 
          placeholder="请输入" 
          @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        小区保洁
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.XQCleaning" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        小区绿化
        <!-- 绿化 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.GreenVegetatio" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        管养保护
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.GuanYangBaoHu" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        消防检查
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.XiaoFanJianCha" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        电梯保养
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.DianTiBaoYang" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        车辆停放有序
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.CarParkInOrder" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        业主委员会
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.YeZhuCommunity" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        业主委员会联系方式
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.YeZhuCommunityTel" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        上一年物业等级
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.PastYearLevel" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
      <div :class='s.item'>
        建议整改措施
        <!-- 独立法人 -->
        <el-input type='textarea' :class='s.elInput' v-model="pm.JianYiZhengGaiCuoShi" placeholder="请输入" @focus='onFocus'></el-input>
      </div>
    </div>
    <div :class='s.bts'>
      <el-button @click='onSave' type='primary'>提交</el-button>
      <el-button @click='onCancel'>取消</el-button>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  props: {
    PM: Object
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
        }
      ],
      curView: 0,
      pm: {
        StreetID: '',
        CommunityID: '',
        XQID: '',
        Name:'',
        LegalPerson:''
      },

      streets: [],
      communities: [],
      xqs: [],
      inputStreetID: '',
      inputCommunityID:'',
      inputXQID: ''
    }
  },
  mounted () {
    Object.assign(this.pm, this.PM)
    this.inputStreetID = this.pm.StreetID
    this.inputCommunityID = this.pm.CommunityID
    this.inputXQID = this.pm.XQID
    this.fetchAllStreets()
  },
  watch: {
    inputStreetID: function (val) {
      this.fetchAllCommunitiesByStreetID(val)
    },
    inputCommunityID: function (val) {
      this.fetchAllXQByCommunityID(val)
    }
  },
  computed: {
  },
  methods: {
    onTab (tab) {
      this.curView = tab.view
    },
    onFocus () {
      this.warn = ''
    },
    onSave () {
      this.pm.StreetID = this.inputStreetID
      this.pm.CommunityID = this.inputCommunityID
      this.pm.XQID = this.inputXQID
      if(!this.checkPM()) return
      fetchpm(this, true, '/pm/pm/update', {
        method: 'POST',
        body: this.pm
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onSave', body)
        if (body.error == 1) {
          this.warn = body.data
        } else {
          this.warn = '修改成功'
          this.$emit('editSucc')
        }
      })
    },
    checkPM () {
      if (this.pm.StreetID == ''){
        this.warn = '请选择Street'
        return false
      }
      if (this.pm.CommunityID == ''){
        this.warn = 'Community'
        return false
      }
      if (this.pm.XQID == ''){
        this.warn = '请选择XQ'
        return false
      }
      if (this.pm.Name == ''){
        this.warn = '请输入名称'
        return false
      }
      if (this.pm.LegalPerson == ''){
        this.warn = '请输入法人代表'
        return false
      }
      return true
    },
    onCancel () {
      this.$emit('cancel')
    },
    fetchAllStreets () {
      fetchpm(this, true, '/pm/street',{
          method: 'POST'
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.streets = data.data
        }
      })
    },
    fetchAllCommunitiesByStreetID (streetID) {
      if ( !streetID || streetID == '') return null
      fetchpm( this, true, '/pm/community/kvs', {
        method: 'POST',
        body: {streetID: streetID}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunities', body)
        this.communities = body.data
      })
    },
    fetchAllXQByCommunityID (communityID) {
      if ( !communityID || communityID == '') return null
      fetchpm( this, true, '/pm/xq/kvs', {
        method: 'POST',
        body: {communityID: communityID}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllXQByCommunityID', body)
        this.xqs = body.data
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