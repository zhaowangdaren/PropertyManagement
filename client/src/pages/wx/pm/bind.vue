<!-- 微信用户绑定 -->
<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.item'>
      <div :class='s.key'>绑定用户名</div>
      <input :class='s.value' type="text" v-module='wxUser.WXName' placeholder="请填写用户名">
    </div>
    <div :class='s.item'>
      <div :class='s.key'>手机号码</div>
      <input :class='s.value' type="text" v-module='wxUser.WXName' placeholder="请填写手机号码"/>
    </div>
    <div :class='s.tip'>
      注意：此功能只针对物业公司人员，普通居民不可绑定
    </div>
    <div 
      @click='onBind'
      :class='s.btn' 
      :style='{backgroundColor: bindSucc ? "rgb(25, 163, 24)": "red"}'>{{btnWord}}</div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: { ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: [],
        title: '用户名绑定',
        rightBtns: []
      },
      host:'http://10.176.118.61:3000',
      wxUser: {
        WXName: '',
        OpenID: '123',
        Tel: ''
      },
      btnWord: '申请绑定',
      bindSucc: true
    }
  },
  methods: {
    onBind () {
      fetch(this.host + '/open/pm/bind', {
        method: 'POST',
        body: JSON.stringify(this.wxUser)
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onBind', body)
        if (body.error !== 0) {
          console.error(body.error)
          this.bindSucc = false
          this.btnWord = body.data
        } else {
          this.bindSucc = true
          this.btnWord = '申请成功'
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  margin-top: 90px;
  .item{
    color: #555;
    margin: 20px 10px;
    position: relative;
    .key{
      font-size: 28px;
    }
    .value{
      width: 100%;
      // box-shadow: 0px 1px 1px 1px rgba(0,0,0,0.2) inset;
      padding: 10px;
      background-color: #fff;
      box-sizing: border-box;
      font-size: 25px;
    }
  }
  .tip{
    color: red;
  }
  .btn{
    background-color: rgb(25, 163, 24);
    color: #fff;
    padding: 20px;
    text-align: center;
    margin: 20px;
    font-size: 25px;
    border-radius: 5px;
  }
}
</style>