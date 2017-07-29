<template>
  <div :class='s.wrap' :style='{backgroundImage: "url(" + bgImg +  ")"}'>
    <div :class='s.formWrap'>
      <div :class='s.title' v-text='sourceParams.title'></div>
      <div :class='s.body'>
        <div :class='s.row'>
          <div :class='s.icon'>
            <i class='iconfont icon-user'></i>
          </div>
          <input type="text" name="" :class='s.input' v-model='login.username'>
        </div>
        <div :class='s.row'>
          <div :class='s.icon'>
            <i class='iconfont icon-lock'></i>
          </div>
          <input type="text" name="" :class='s.input' v-model='login.password'>
        </div>
        <div :class='s.bottom'>
          <el-button :class='s.login' @click='onLogin'>登录</el-button>
          <el-button @click='onCancel'>取消</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="text/javascript">

export default {
  data () {
    return {
      host:'//localhost:3000',
      sourceParams: {
        title: 'Title',
        target: '' // /admin  /street  /gov
      },
      login: {
        username:'',
        password:''
      },
      bgImg: require('@/res/images/bottom_bg.png')
    }
  },
  mounted () {
    this.sourceParams.title = this.$route.query.title
    this.sourceParams.target = this.$route.query.target
  },
  methods: {
    onLogin () {
      var path = this.host
      switch (this.sourceParams.target) {
        case '/admin':
          path += '/admin/login'
          break;
        case '/street':
          path += '/street/login'
          break;
        case '/gov':
          path += '/gov/login'
          break;
      }
      fetch(path, {
        method: 'POST',
        body: JSON.stringify(this.login)
      }).then(resp => {
        return resp.json()
      }).then( body => {
        console.info('onLogin',body)
        if (body.error === 0)  this.$router.push({path: this.sourceParams.target})
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
      width: 40%;
    .title{
      font-size: 30px;
      text-align: center;
      padding: 10px 0;
      background-color: #fff;
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
        }
      }
      .bottom{
        text-align: center;
        margin-top: 10px;
        .login{
          background-color: #1ea5f3;
          border: solid 1px transparent;
          &:hover{
            color: #fff;
          }
        }
      }
    }
  }
}
</style>