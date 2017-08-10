<template>
  <div :class='s.wrap'>
    <!-- 新增小区信息 -->
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      街道名称
      <!-- 街道名称 -->
      <el-select :class='s.elInput' v-model="searchStreetID" placeholder="请选择街道名称">
        <el-option
          v-for="item in streets"
          :key="item.ID"
          :label="item.Name"
          :value="item.ID">
        </el-option>
      </el-select>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      社区名称
      <!-- 社区名称 -->
      <el-select :class='s.elInput' v-model="searchCommunityID" placeholder="请选择">
        <el-option
          v-for="item in communities"
          :key="item.ID"
          :label="item.Name"
          :value="item.ID">
        </el-option>
      </el-select>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      小区名称
      <!-- 小区名称 -->
      <el-input :class='s.elInput' v-model="country.Name" placeholder="请输入"小区名称 @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      地理位置
      <!-- 地理位置 -->
      <el-input :class='s.elInput' v-model="country.Address" placeholder="请输入地理位置" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      联系人
      <!-- 联系人 -->
      <el-input :class='s.elInput' v-model="country.ContactName" placeholder="请输入联系人" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      电话号码
      <!-- 电话号码 -->
      <el-input :class='s.elInput' v-model="country.Tel" placeholder="请输入电话号码" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      描述
      <!-- 描述 -->
      <el-input 
        :class='s.elInput'
        type="textarea"
        autosize
        v-model="country.Intro" 
        placeholder="请输入描述" 
        @focus='onFocus'>
      </el-input>
    </div>
    <div :class='s.item'>
      <el-button @click='onSave' type='primary'>提交</el-button>
      <el-button @click='onCancel'>取消</el-button>
    </div>
  </div>
</template>

<script>
import BasicDialog from '@/components/dialog/index'
import SearchSelect from '@/components/SearchSelect'
import fetchpm from '@/fetchpm'
export default {
  components: { BasicDialog, SearchSelect },
  props: {
    country: Object
  },
  data () {
    return {
      warn: '',
      streets: [],
      communities: [],
      searchStreetID: '',
      searchCommunityID:''
    }
  },
  mounted () {
    this.searchStreetID = this.country.StreetID

    this.fetchAllStreets()
    this.fetchAllCommunities()
    // this.fetchAllCommunitiesByStreetID(this.country.StreetID)
  },
  watch: {
    searchStreetID: function (value) {
      if (this.searchStreetID == this.country.StreetID) this.searchCommunityID = this.country.CommunityID
      else this.searchCommunityID = ''
      
      this.fetchAllCommunitiesByStreetID(value)
    }
  },
  methods: {
    onFocus () {
      this.warn = ''
    },
    onSave () {
      this.country.StreetID = this.searchStreetID
      this.country.CommunityID = this.searchCommunityID
      console.info(this.country)
      if (!this.checkCountry()) return
      this.updateCountry()
      console.info('onSave')
    },
    onCancel () {
      this.$emit('cancel')
    },
    checkCountry () {
      if (this.country.Name !== ''
        && this.country.Address !== ''
        && this.country.StreetID !== ''
        && this.country.CommunityID !== ''
        && this.country.ContactName !== ''
        && this.country.Tel !== ''
        && this.country.Intro !== '') return true
      return false
    },
    updateCountry () {
      fetchpm(this,true,'/pm/xq/update', {
        method: 'POST',
        body:this.country
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('addCountry', body)
        if(body.error === 1) this.warn = body.data
        else {
          this.$emit('editSucc')
          this.warn = 'Edit Succ'
        }
      })
    },
    fetchAllStreets () {//获取所有街道名称
      fetchpm( this, true, '/pm/street', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        this.streets = body.data
      })
    },
    fetchAllCommunities () {
      fetchpm( this, true, '/pm/community', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunities', body)
        this.communities = body.data
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
</style>
