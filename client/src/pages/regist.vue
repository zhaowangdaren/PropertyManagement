<template>
  <div :class='s.wrap' :style='{backgroundImage: "url(" + bgImg +  ")"}'>
    <div :class='s.formWrap'>
      <div :class='s.title' v-text='title'></div>
        <p :class="s.warn" v-text='warn'></p>
        <div :class='s.body'>
          <table :class='s.table'>
            <tr>
              <td :class='s.left'><span :class='s.red'>*</span>用户名</td>
              <td :class='s.mid'><el-input type="text" name="" v-model='regist.UserName'></el-input></td>
              <td :class='s.right'>由字母、数字或者下划线组成</td>
            </tr>
            <tr>
              <td :class='s.left'><span :class='s.red'>*</span>真实姓名</td>
              <td :class='s.mid'><el-input type="text" name="" v-model='regist.RealName'></el-input></td>
              <td :class='s.right'>真实中文名</td>
            </tr>
            <tr>
              <td :class='s.left'><span :class='s.red'>*</span>密码</td>
              <td :class='s.mid'><el-input type="text" name="" v-model='regist.Password'></el-input></td>
              <td :class='s.right'>由6-16个英文字母、数字或下划线组成</td>
            </tr>
            <tr>
              <td :class='s.left'><span :class='s.red'>*</span>联系电话</td>
              <td :class='s.mid'><el-input type="text" name="" v-model='regist.Tel'></el-input></td>
              <td :class='s.right'>手机号码或者固定电话</td>
            </tr>
            <tr v-if='regist.Type === 3'>
              <td :class='s.left'><span :class='s.red'>*</span>所属街道</td>
              <td :class='s.mid'>
                <!-- <el-select v-model="regist.StreetID" filterable placeholder="请选择" >
                  <el-option
                    v-for="item in streets"
                    :key="item.ID"
                    :label="item.Name"
                    :value="item.ID">
                  </el-option>
                </el-select> -->
                <el-select
                  :class='s.selectStreet'
                  v-model="regist.StreetID"
                  filterable
                  remote
                  placeholder="请输入关键词"
                  :remote-method="searchStreet"
                  :loading="isSearching">
                  <el-option
                    v-for="item in streets"
                    :key="item.ID"
                    :label="item.Name"
                    :value="item.ID">
                  </el-option>
                </el-select>
              </td>
              <td :class='s.right'>选择所属街道</td>
            </tr>
            <tr v-if='regist.Type === 3'>
              <td :class='s.left'><span :class='s.red'>*</span>授权码</td>
              <td :class='s.mid'>
                <el-input v-model='regist.Code'></el-input>
              </td>
              <td :class='s.right'>输入指定授权码</td>
            </tr>
          </table>
        <div  :class='s.bottom'>
          <el-button type='primary' @click='onSubmit'>提交</el-button>
          <el-button @click='onCancel'>返回</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="text/javascript">
import fetchpm from '@/fetchpm'
import { Message } from 'element-ui'
export default {
  data () {
    return {
      warn: '',
      target: '',
      title: '注册',
      regist: {
        UserName: '',
        Password: '',
        RealName: '',
        Tel: '',
        StreetID: '',
        Code: '', // 授权码
        Type: 0,
      },
      streets: [],
      isSearching: false,
      bgImg: require('@/res/images/bottom_bg.png'),
      isLogining: false
    }
  },
  mounted () {
    this.target = this.$route.query.target
    switch (this.target) {
      case '/admin':
        this.title = '管理员注册'
        this.regist.Type = 1
        break
      case '/gov':
        this.title = '政府工作人员注册'
        this.regist.Type = 2
        break
      case '/street':
        this.title = '街道工作人员注册'
        this.regist.Type = 3
        break
      case '/court':
        this.title = '法官注册'
        this.regist.Type = 5
        break
    }
  },
  methods: {
    searchStreet (name) {
      if (name === '') return
      this.isSearching = true
      fetchpm(this, false, '/open/streets/search/name/' + name, {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0) this.streets = body.data || []
        else Message({type: 'error', message: body.data})
        this.isSearching = false
      }).catch(error => {
        Message({type: 'error', message: error})
        this.isSearching = false
      })
    },
    onSubmit () {
      fetchpm(this, false, '/open/regist', {
        method: 'POST',
        body: this.regist
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onSubmit', body)
        if (body.error === 0) Message({type:'success', message: '恭喜，注册成功，请等待管理员审核'})
        else this.warn = body.data
      })
    }, 
    onRegist () {
      this.$router.push({name: 'Regist', query: {target: this.sourceParams.target}})
    },
    checkInput () {
      if (this.regist.UserName === '') {
        this.warn = '请输入用户名'
        return false
      }
      // .search(/\ /g)
      if (this.regist.UserName.search(/\s/g) > -1) {
        this.warn = '用户名中不能包含空格'
        return false
      }
      if (this.regist.RealName === '') {
        this.warn = '请输入真实用户名'
        return false
      }
      if (this.regist.Password === '') {
        this.warn = '请输入密码'
        return false
      }
      if (this.regist.Password.search(/\s/g) > -1) {
        this.warn = '密码中不能包含空格'
        return false
      }
      if (this.regist.Password.length < 6 || this.regist.Password.length > 16) {
        this.warn = '密码长度为' + this.regist.Password + ',不符合要求'
        return false
      }
      if (this.regist.Tel === '') {
        this.warn = '请输入联系电话'
        return false
      }

      if (this.regist.Type === 3 && this.regist.StreetID === '') {
        this.warn = '请选择所属街道'
        return false
      }

      if (this.regist.Type === 3 && this.regist.Code === '') {
        this.warn = '请输入授权码'
        return false
      }
      return true
    },
    onCancel () {
      this.$router.go(-1);
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  background-color: #379ad3;
  background-position: bottom;
  background-repeat: repeat-x;
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: auto;
    // min-width: 40%;
  .title{
    font-size: 30px;
    text-align: center;
    padding: 10px 0;
    background-color: #fff;
  }
  .warn{
    color: red;
    font-size: 20px;
    font-weight: bold;
    text-align: center;
  }
  .body{
    background-color: rgba(255,255,255, 0.5);
    padding: 20px 50px;
    .bottom{
      text-align: center;
      margin-top: 10px;
      display: flex;
      justify-content: center;
    }
  }
}
.table{
  font-size: 16px;
  color: #000;
  box-sizing: border-box;
  tr{
    display: flex;
  }
}
.red{
  color: red;
}
.left{
  flex: 1;
  text-align: right;
  min-width: 100px;
}
.mid{
  flex: 2;
  // min-width: 300px;
  input{
    box-sizing: border-box;
    width: 100%;
  }
}
.right{
  flex: 1;
  text-align: left;
  
  // overflow: hidden;
  // text-overflow:ellipsis;
  // white-space: nowrap;
}
.selectStreet{
  width: 100%;
}
</style>
