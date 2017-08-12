<template>
  <div :class='s.wrap'>
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
  <!-- 新增街道信息 -->
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      街道名
      <!-- 街道名 -->
      <el-input :class='s.elInput' v-model="street.Name" placeholder="请输入街道名" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      负责人
      <!-- 负责人 -->
      <el-input :class='s.elInput' v-model="street.PersonInCharge" placeholder="请输入负责人" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      电话号码
      <!-- 电话号码 -->
      <el-input
        :class='s.elInput'
        v-model="street.Tel" 
        placeholder="请输入电话号码" 
        @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      授权码
      <!-- 授权码 -->
      <el-input :class='s.elInput' v-model="street.AuthCode" placeholder="请输入授权码" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      描述
      <!-- 描述 -->
      <el-input 
        :class='s.elInput'
        type="textarea"
        autosize
        v-model="street.Intro" 
        placeholder="请输入描述" 
        @focus='onFocus'>
      </el-input>
    </div>
    <div :class='s.item'>
      <el-button @click='onSave' type='primary'>保存</el-button>
      <el-button @click='onCancel'>取消</el-button>

    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
  export default {
    props: {
      street: {
        type: Object,
        default: function () {
          return {

          }
        }
      }
    },
    components: {},
    data () {
      return {
        warn:''
      }
    },
    methods:{
      onSave () {
        if (!this.checkStreet()) return
        this.editStreet()
        console.info('onSave')
      },
      onCancel () {
        console.info('onCancel')
        this.$emit('cancel')
        this.warn = ''
      },
      checkStreet () {
        if( this.street.Name !== ''
          && this.street.PersonInCharge !== ''
          && this.street.Tel !== ''
          && this.street.AuthCode !== ''
          && this.street.Intro !== '') return true
        return false
      },
      editStreet () {
        fetchpm(this, true, '/pm/street/update', {
          method: 'POST',
          body: this.street
        }).then(resp => {
          console.info(resp)
          return resp.json()
        }).then(body => {
          if(body.error ==1) this.warn = body.data
          else {
            this.warn = '修改成功'
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