<!-- 微信用户绑定 -->
<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div v-if='warn !== ""' :class='s.warn'>{{warn}}</div>
    <div :class='s.item'>
      <div :class='s.key'>真实姓名:</div>
      <input 
        :class='s.value' 
        type="text"
        v-model='courtUser.ActualName'
        @focus='warn = "", showSearchResults = false'
        placeholder="请填写真实姓名" />
    </div>
    <div :class='s.item'>
      <div :class='s.key'>手机号码:</div>
      <input 
        :class='s.value' 
        type="text" 
        v-model='courtUser.Tel'
        @focus='warn = "", showSearchResults = false'
        placeholder="请填写手机号码"/>
    </div>
    <div :class='s.tip'>
      注意：此功能只针对法官，普通居民不可绑定。请输入真实信息，以便审核
    </div>
    <div
      @click='onBind'
      :class='s.btn' 
      :style='{backgroundColor: bindSucc ? "rgb(25, 163, 24)": "red"}'>{{btnWord}}</div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import fetchpm from '@/fetchpm'
export default {
  components: { ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: [],
        title: '用户名绑定',
        rightBtns: []
      },
      warn: '',
      courtUser: {
        ActualName: '',
        OpenID: '',
        Tel: ''
      },
      showSearchResults: false,
      btnWord: '申请绑定',
      bindSucc: true,
      isLoading: true
    }
  },
  mounted () {
    var WXUser = JSON.parse(sessionStorage.getItem('WXUser') || '{"openid": ""}')
    if (!WXUser.openid || WXUser.openid === "") this.fetchOpenID()
  },
  methods: {
    checkInfo () {
      if (this.courtUser.ActualName === '') {
        this.warn = '真实姓名不能为空'
        return false
      }

      if (this.courtUser.Tel === '' || this.courtUser.Tel.length < 11) {
        this.warn = '手机号码填写错误'
        return false
      }

      if (this.courtUser.OpenID === '' || !this.courtUser.OpenID) {
        this.warn = '用户信息获取失败，请退出重新进行操作'
        return false
      }
      return true
    },
    onBind () {
      if (!this.checkInfo()) return
      fetchpm(this, false, '/open/court/wx/bind', {
        method: 'POST',
        body: this.courtUser
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
    },
    fetchOpenID () {
      fetchpm(this, false, '/open/wx/openid/' + this.$route.query.code, {
        method: 'POST',
        body: JSON.stringify({})
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchOpenID', body)
        if (body.error === 0) {
          sessionStorage.setItem('courtUser', JSON.stringify(body.data))
          this.courtUser.OpenID = body.data.openid
          this.isLoading = false
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  margin-top: 90px;
  .warn{
    text-align: center;
    color: red;
  }
  .item{
    color: #555;
    margin: 20px 10px;
    position: relative;
    .key{
      font-size: 25px;
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
  .results{
    position: absolute;
    z-index: 2;
    width: 100%;
    font-size: 26px;
    background-color: #f5f5f5;
    box-shadow: 1px 0 2px 1px rgba(0, 0, 0, 0.3);
    .pm{
      padding: 10px;
      background: #fff;
      margin-bottom: 2px;
    }
  }
  .tip{
    color: red;
    padding: 5px;
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