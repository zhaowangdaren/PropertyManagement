<template>
  <div :class='s.wrap'>
  <!-- 新增社区信息 -->
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      街道
      <!-- 街道名称 -->
      <el-select :class='s.elInput' v-model="community.StreetID" placeholder="请选择街道">
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
      社区名
      <!-- 社区名称 -->
      <el-input :class='s.elInput' v-model="community.Name" placeholder="请输入社区名" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      负责人
      <!-- 负责人 -->
      <el-input :class='s.elInput' v-model="community.PersonInCharge" placeholder="请输入负责人" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      电话号码
      <!-- 电话号码 -->
      <el-input :class='s.elInput' v-model="community.Tel" placeholder="请输入电话号码" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      描述
      <!-- 描述 -->
      <el-input 
        :class='s.elInput'
        type="textarea"
        autosize
        v-model="community.Intro" 
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
import { Message } from 'element-ui'
import SearchSelect from '@/components/SearchSelect'
import fetchpm from '@/fetchpm' 
export default {
  components: { SearchSelect },
  props: {
    community: Object
  },
  data () {
    return {
      host: 'http://10.176.118.61:3000',
      inputWord: '',
      streets: [],
      warn:''
    }
  },
  mounted () {
    this.fetchAllStreetName()
  },
  methods: {
    onSave () {
      console.info(this.community)
      if (!this.checkCommunity()) return
      this.addCommunity()
      console.info('onSave')

    },
    onCancel () {
      console.info('onCancel')
      this.$emit('cancel')
    },
    checkCommunity () {
      if( this.community.name !== ''
        && this.community.personInCharge !== ''
        && this.community.tel !== ''
        && this.community.streetName !== ''
        && this.community.intro !== '') return true
      return false
    },
    addCommunity () {
      fetchpm(this, true, '/pm/community/update', {
        method: 'POST',
        body: this.community
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then(body => {
        if(body.error ==1) this.warn = body.data
        else {
          // this.warn = '修改成功' //
          Message({message:'恭喜，修改成功', type:'success'})
          this.onCancel()
          this.$emit('editSucc')
        }
      })
    },
    fetchAllStreetName () {//获取所有街道名称
      fetchpm(this, true, '/pm/street', {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        this.streets = body.data
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
