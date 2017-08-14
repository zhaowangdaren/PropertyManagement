<template>
  <div :class='s.wrap'>
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      用户名
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.UserName" placeholder="请输入" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item' v-if='user.Type == 3'>
      <div :class='s.red'>*</div>
      街道
      <!-- 用户名 -->
        <el-select :class='s.elInput' v-model="selectedStreetID" placeholder="请选择街道">
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
      真实姓名
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.RealName" placeholder="请输入真实姓名" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      密码
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.Password" placeholder="请输入密码" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      办公电话
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.Tel" placeholder="请输入办公电话" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <el-button @click='onSave' type='primary'>提 交</el-button>
      <el-button @click='onCancel'>取 消</el-button>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
import { Message } from 'element-ui'
export default {
  props: {
    USER: Object
  },
  data () {
    return {
      warn: '',
      user: {
        UserName: '',
        RealName: '',
        Password: '',
        StreetID: '',
        Tel: ''
      },
      selectedStreetID: '',
      streets: []
    }
  },
  mounted () {
    Object.assign(this.user, this.USER)
    this.getAllStreets()
    this.selectedStreetID = this.user.StreetID
  },
  methods: {
    onSave () {
      if (this.selectedStreetID != '') this.user.StreetID = this.selectedStreetID
      if(!this.checkUser()) return
      fetchpm(this, true, '/pm/user/update', {
        method: 'POST',
        body: this.user
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onSave', body)
        if (body.error == 1) this.warn = body.data
        else {
          // this.warn = '修改成功'
          this.$emit('editSucc')
          Message({message:'恭喜，修改成功', type:'success'})
          this.onCancel()
        }
      })
    },
    onCancel () {
      this.$emit('cancel')
    },
    onFocus () {
      this.warn = ''
    },
    checkUser () {
      if (this.user.UserName == '') {
        this.warn = '用户名不能为空'
        return false
      } 
      if (this.user.Password == '') {
        this.warn = '密码不能为空'
        return false
      }
      return true
    },
    getAllStreets () {
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