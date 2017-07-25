<template>
  <div>
  <!-- 新增社区信息 -->
    <basic-dialog 
      title='Community'
      :onSave='onSave'
      :onCancel='onCancel'
      :warn='warn'
    >
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        name
        <!-- 街道名称 -->
        <search-select v-model='community.streetName' :values='values'/>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        name2
        <!-- 社区名称 -->
        <input type="text" v-model='community.name' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        charger
        <!-- 负责人 -->
        <input type="text" v-model='community.personInCharge' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        tel
        <!-- 电话号码 -->
        <input type="text" v-model='community.tel' @focus='onFocus'>
      </div>
      <div :class='s.item'>
        <div :class='s.red'>*</div>
        intro
        <!-- 描述 -->
        <input type="text" v-model='community.intro' @focus='onFocus'>
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
import SearchSelect from '@/components/SearchSelect'
export default {
  components: { BasicDialog, SearchSelect },
  data () {
    return {
      host: 'http://10.176.118.61:3000',
      inputWord: '',
      values: [],
      community: {
        name: '',
        psersonInCharge:'',
        tel:'',
        streetName:'',
        intro: ''
      },
      warn:''
    }
  },
  mounted () {
    this.fetchAllStreetName()
  },
  methods: {
    onSave () {
      console.info(this.community)
      if (!this.checkStreet()) return
      this.addStreet()
      console.info('onSave')

    },
    onCancel () {
      console.info('onCancel')
      this.$emit('close')
    },
    checkStreet () {
      if( this.community.name !== ''
        && this.community.personInCharge !== ''
        && this.community.tel !== ''
        && this.community.streetName !== ''
        && this.community.intro !== '') return true
      return false
    },
    addStreet () {
      fetch(this.host + '/community/add', {
        method: 'POST',
        body: JSON.stringify(this.community)
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then(body => {
        if(body.error ==1) this.warn = body.data
        else {
          this.warn = 'Add succ' //街道新增成功
          this.community.name = '',
          this.community.personInCharge = ''
          this.community.tel = ''
          this.community.streetName = ''
          this.community.intro = ''
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
        this.values = body.data
      })
    },
    onFocus () {
      this.warn = ''
    }
  }
}
</script>