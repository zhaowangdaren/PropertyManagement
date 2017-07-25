<template>
  <div>
    <!-- 新增小区信息 -->
    <basic-dialog 
      title='Add Country'
      :onSave='onSave'
      :onCancel='onCancel'
      :warn='warn'
    >
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        Street
        <!-- 街道名称 -->
        <search-select v-model='country.street' :values='streets'/>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        Community
        <!-- 社区名称 -->
        <search-select v-model='country.community' :values='communities'/>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        Country
        <!-- 小区名称 -->
        <input type="text" v-model='country.name' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        Address
        <!-- 地理位置 -->
        <input type="text" v-model='country.address' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        charger
        <!-- 联系人 -->
        <input type="text" v-model='country.contactName' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        tel
        <!-- 电话号码 -->
        <input type="text" v-model='country.tel' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        intro
        <!-- 描述 -->
        <input type="text" v-model='country.intro' @focus='onFocus'>
      </div>
    </basic-dialog>
  </div>
</template>

<style lang="less" module='s'>
  .item{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    padding: 10px 40px;
    font-size: 18px;
    .red{
      color: red;
    }
    input {
      flex: 1;
      margin-left: 10px;
      height: 30px;
    }
  }
  
</style>

<script>
import BasicDialog from '@/components/dialog/index'
import SearchSelect from '@/components/Select'
export default {
  components: { BasicDialog, SearchSelect },
  data () {
    return {
      host: 'http://10.176.118.61:3000',
      warn: '',
      streets: [],
      communities: [],
      country: {
        name: '',
        address: '',
        street: '',
        community: '',
        contactName: '',//联系人
        tel: '',
        intro:''
      }
    }
  },
  mounted () {
    this.fetchAllStreetName()
    this.fetchAllCommunityName()
  },
  methods: {
    onFocus () {
      this.warn = ''
    },
    onSave () {
      if (!this.checkCountry()) return
      this.addCountry()
      console.info('onSave')
    },
    onCancel () {
      this.$emit('close')
    },
    checkCountry () {
      if (this.country.name !== ''
        && this.country.address !== ''
        && this.country.street !== ''
        && this.country.community !== ''
        && this.country.contactName !== ''
        && this.country.tel !== ''
        && this.country.intro !== '') return true
      return false
    },
    addCountry () {
      fetch(this.host + '/xq/add', {
        method: 'POST',
        body:JSON.stringify(this.country)
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('addCountry', body)
        if(body.error === 1) this.warn = body.data
        else {
          this.warn = 'Add Succ'
          this.country.name = ''
          this.country.address = ''
          this.country.street = ''
          this.country.community = ''
          this.country.contactName = ''
          this.country.tel = ''
          this.country.intro = ''
        }
      })
    },
    fetchAllStreetName () {//获取所有街道名称
      fetch( this.host + '/street/key/name', {
        method: 'POST',
        body: '{}'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        this.streets = body.data
      })
    },
    fetchAllCommunityName () {
      fetch( this.host + '/community/key/name', {
        method: 'POST',
        body: '{}'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchAllCommunityName', body)
        this.communities = body.data
      })
    }
  }
}
</script>