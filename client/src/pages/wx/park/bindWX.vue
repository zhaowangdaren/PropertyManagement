<!-- 停车管理员申请绑定微信 -->
<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div v-if='warn !== ""' :class='s.warn'>{{warn}}</div>
      <div :class='s.item'>
        <div :class='s.key'>{{$route.query.xqName}}</div>
      </div>
      <div :class='s.item'>
        <div :class='s.key'>真实姓名:</div>
        <input 
          :class='s.value' 
          type="text" 
          v-model='parkManager.ActualName'
          @focus='onFocus'
          placeholder="请填写真实姓名" />
      </div>
      <div :class='s.item'>
        <div :class='s.key'>手机号码:</div>
        <input 
          :class='s.value' 
          type="text" 
          v-model='parkManager.Tel'
          @focus='onFocus'
          placeholder="请填写手机号码"/>
      </div>
      <div :class='s.tip'>
        注意：此功能只针对停车管理员，普通居民不可绑定。请输入真实信息，以便审核
      </div>
      <div
        @click='onBind'
        :class='s.btn' 
        :style='{backgroundColor: bindSucc ? "rgb(25, 163, 24)": "#20a0ff"}'>{{btnWord}}
      </div>
    </div>
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
        title: '停车管理员绑定',
        rightBtns: []
      },
      warn: '',
      btnWord: '申请绑定',
      bindSucc: false,
      code: '',
      isLoading: true,
      parkManager: {
        OpenID: '',
        ActualName: '',
        Tel: '',
        XQID: ''
      }
    }
  },
  mounted () {
    var WXUser = JSON.parse(sessionStorage.getItem('WXUser') || "{openid: ''}")
    this.parkManager.OpenID = WXUser.openid
    this.parkManager.XQID = this.$route.query.xqid
  },
  methods: {
    onFocus () {
      this.warn = ''
      this.bindSucc = false
      this.btnWord = '申请绑定'
    },
    checkInfo () {
      if (this.parkManager.PMID === '') {
        this.warn = '物业公司不存在或填写错误'
        return false
      }
      if (this.parkManager.ActualName === '') {
        this.warn = '真实姓名不能为空'
        return false
      }

      if (this.parkManager.Tel === '' || this.parkManager.Tel.length < 11) {
        this.warn = '手机号码填写错误'
        return false
      }

      if (this.parkManager.OpenID === '' || !this.parkManager.OpenID) {
        this.warn = '用户信息获取失败，请退出重新进行操作'
        return false
      }
      return true
    },
    onBind () {
      if (!this.checkInfo()) return
      if (this.bindSucc) {
        this.$router.go(-4)
        return
      }
      fetchpm(this, false, '/open/park/manager/bind', {
        method: 'POST',
        body: this.parkManager
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