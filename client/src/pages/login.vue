<template>
  <div :class='s.wrap' :style='{backgroundImage: "url(" + bgImg +  ")"}'>
    <div :class='s.formWrap'>
      <div :class='s.title' v-text='sourceParams.title'></div>
      <form :class='s.body' @submit.prevent='onLogin'>
        <p :class="s.warn" v-text='warn'></p>
        <div :class='s.row'>
          <div :class='s.icon'>
            <i class='iconfont icon-user'></i>
          </div>
          <input type="text" name="" :class='s.input' v-model='login.username' autofocus>
        </div>
        <div :class='s.row'>
          <div :class='s.icon'>
            <i class='iconfont icon-lock'></i>
          </div>
          <input type="password" name="" :class='s.input' v-model='login.password'>
        </div>
        <div :class='s.bottom'>
          <input type="submit" name="" value='登 录' @click='onLogin' :style='{visibility: "hidden"}' ></span>
          <el-button type='primary' @click='onLogin' :loading='isLogining'>登 录</el-button>
          <el-button type='success' @click='onRegist'>注册</el-button>
          <!-- <span :class='s.cancel' @click='onCancel'>取 消</span> -->
          <el-button @click='onCancel'>取 消</el-button>
        </div>
      </form>
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
      sourceParams: {
        title: 'Title',
        target: '' // /admin /gov /street
      },
      login: {
        username:'',
        password:'',
        type: 0
      },
      bgImg: require('@/res/images/bottom_bg.png'),
      isLogining: false
    }
  },
  mounted () {
    this.sourceParams.title = this.$route.query.title
    this.sourceParams.target = this.$route.query.target
    console.info(this.sourceParams)
    switch (this.sourceParams.target) {
      case '/admin':
        this.login.type = 1
        break
      case '/gov':
        this.login.type = 2
        break
      case '/street':
        this.login.type = 3
        break
      case '/court':
        this.login.type =5
        break
    }
  },
  methods: {
    onRegist () {
      this.$router.push({name: 'Regist', query: {target: this.sourceParams.target}})
    },
    checkInput () {
      if (this.login.username === '') {
        this.warn = '请输入用户名'
        return false
      }
      if (this.login.password === '') {
        this.warn = '请输入密码'
        return false
      }
      return true
    },
    onLogin () {
      if (!this.checkInput() || this.isLogining) return
      this.isLogining = true
      fetchpm(this, false, '/login', {
        method: 'POST',
        body: this.login
      }).then(resp => {
        return resp.json()
      }).then( body => {
        console.info('onLogin',body)
        if (body.error == 0 && body.data.token && body.data.token !== "")  {
          body.data.type = this.login.type
          sessionStorage.setItem('user', JSON.stringify(body.data))
          this.$router.push({path: this.sourceParams.target})
        } else {
          this.warn = body.data.message
        }
        this.isLogining = false
      }).catch(error => {
        Message({type:'error', message: error})
        this.isLogining = false
      })
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
  .formWrap{
    // width: 40%;
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
      background-color: rgba(254,254,254, 0.3);
      padding: 20px 50px;
      .row{
        display: flex;
        margin: 5px;
        .icon{
          text-align:center;
          background-color: #fff;
          padding: 0 10px;
          height: 48px;
          display: flex;
          justify-content: center;
          align-items: center;
          i{
            color: #666;
          }
        }
        .input{
          flex: 1;
          height:46px;
          border: solid 1px #ddd;
          background-color: rgb(250, 255, 189);
          font-size: 20px;
        }
      }
      .bottom{
        text-align: center;
        margin-top: 10px;
        display: flex;
        justify-content: center;
        .login{
          display: inline-block;
          line-height: 1;
          white-space: nowrap;
          cursor: pointer;
          border: 1px solid #bfcbd9;
          color: #fff;
          -webkit-appearance: none;
          text-align: center;
          box-sizing: border-box;
          outline: none;
          margin: 0;
          -moz-user-select: none;
          -webkit-user-select: none;
          -ms-user-select: none;
          padding: 10px 15px;
          font-size: 18px;
          border-radius: 4px;
          background-color: #20a0ff;
          border-color: #20a0ff;
        }
        .cancel{
          background-color: #fff;
          border: solid 1px #4db3ff;
          padding: 5px 15px;
          font-size: 18px;
          border-radius: 4px;
          color: #4db3ff;
          margin-left: 10px;
        }
      }
    }
  }
}
</style>
