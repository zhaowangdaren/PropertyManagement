<template lang="html">
  <div :class="s.wrap">
    <div :class="s.warn" v-text='warn'></div>
    <table>
      <tr>
        <td>
          <span :class="s.red">*</span>
          旧密码
        </td>
        <td>
          <el-input
            type='password'
            v-model='user.Password'
            @focus='onFocus(1)'></el-input>
        </td>
        <td>
          <span :class='s.tip' :style='{color: errorIndex == 1 ? "red" : ""}'>请输入旧密码</span>
        </td>
      </tr>
      <tr>
        <td>
          <span :class="s.red">*</span>
          新密码
        </td>
        <td>
          <el-input
            type='password'
            v-model='user.NewPassword'
            @focus='onFocus(2)'></el-input>
        </td>
        <td>
          <span :class='s.tip' :style='{color: errorIndex == 2 ? "red" : ""}'>由8-20个英文字母、数字或下划线组成,且至少包含2种</span>
        </td>
      </tr>
      <tr>
        <td>
          <span :class="s.red">*</span>
          确认新密码
        </td>
        <td>
          <el-input
            type='password'
            v-model='reNewPassword'
            @focus='onFocus(3)'></el-input>
        </td>
        <td>
          <span :class='s.tip' :style='{color: errorIndex == 3 ? "red" : ""}'>请再输入一遍新密码</span>
        </td>
      </tr>
    </table>
    <div :class="s.btns">
      <el-button type='primary' @click='onSave'>确 认</el-button>
      <el-button @click='onCancel'>取 消</el-button>
    </div>
  </div>
</template>

<script>
import { Message } from 'element-ui'
import fetchpm from '@/fetchpm'
export default {
  data () {
    return {
      user: {
        Type: 0,
        UserName: '',
        Password: '',
        NewPassword: ''
      },
      reNewPassword: '',
      errorIndex: 0,
      warn: ''
    }
  },
  mounted () {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.user.Type = user.type
    this.user.UserName = user.UserName
  },
  methods: {
    onSave () {
      if (!this.checkInput()) return
      fetchpm(this, true, '/pm/user/changePassword', {
        method: 'POST',
        body: this.user
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('changePassword', body)
        if (body.error == 1) {
          this.warn = body.data
        } else {
          // this.warn = '密码修改成功'
          Message({message:'恭喜，密码修改成功', type:'success'})
          this.onCancel()
          this.user.Password = '',
          this.user.NewPassword = ''
          this.reNewPassword = ''
        }
      })
    },
    checkInput () {
      if (this.reNewPassword !== this.user.NewPassword) {
        this.warn = '两次输入的密码不一致'
        return false
      }
      if (this.user.Password.length == 0) {
        this.errorIndex = 1
        return false
      }
      if (this.user.NewPassword.length == 0) {
        this.errorIndex = 2
        return false
      }
      if (this.reNewPassword.length == 0) {
        this.errorIndex = 3
        return false
      }
      return true
    },
    onCancel () {

    },
    onFocus (index) {
      if (this.errorIndex === index) this.errorIndex = 0
      this.warn=''
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  font-size: 18px;
  .warn{
    color: red;
    font-size: 16px;
    text-align: center;
    margin: 10px;
  }
  td{
    padding: 0 5px;
  }
  .btns{
    margin: 10px 20px;
    display: flex;
    justify-content: flex-end;
  }
  .red{
    color: red;
  }
  .tip{
    color: #999;
    font-size: 12px;
  }
}
</style>
