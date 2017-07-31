<template>
  <div :class='s.wrap'>
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
  <!-- 新增街道信息 -->
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      name
      <!-- 街道名 -->
      <el-input :class='s.elInput' v-model="street.name" placeholder="请输入street name" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      charger
      <!-- 负责人 -->
      <el-input :class='s.elInput' v-model="street.personInCharge" placeholder="请输入" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      tel
      <!-- 电话号码 -->
      <el-input
        :class='s.elInput'
        v-model="street.tel" 
        placeholder="请输入" 
        @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      authCode
      <!-- 授权码 -->
      <el-input :class='s.elInput' v-model="street.authCode" placeholder="请输入authCode" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      intro
      <!-- 描述 -->
      <el-input 
        :class='s.elInput'
        type="textarea"
        autosize
        v-model="street.intro" 
        placeholder="请输入intro" 
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
import fetchpm from '@/fetchpm'
  export default {
    components: {},
    data () {
      return {
        street: {
          name: '',
          psersonInCharge:'',
          tel:'',
          authCode:'',
          intro: ''
        },
        warn:''
      }
    },
    methods:{
      onSave () {
        if (!this.checkStreet()) return
        this.addStreet()
        console.info('onSave')

      },
      onCancel () {
        console.info('onCancel')
        this.$emit('cancel')
        this.warn = ''
      },
      checkStreet () {
        if( this.street.name !== ''
          && this.street.personInCharge !== ''
          && this.street.tel !== ''
          && this.street.authCode !== ''
          && this.street.intro !== '') return true
        return false
      },
      addStreet () {
        fetchpm(this, true, '/pm/street/add', {
          method: 'POST',
          body: this.street
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then(body => {
          if(body.error ==1) this.warn = body.data
          else {
            this.$emit('addSucc')
            this.warn = '街道新增成功'
            this.street.name = '',
            this.street.personInCharge = ''
            this.street.tel = ''
            this.street.authCode = ''
            this.street.intro = ''
          }
        })
      },
      onFocus () {
        this.warn = ''
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