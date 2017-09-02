<!-- 微信用户绑定 -->
<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div v-if='warn !== ""' :class='s.warn'>{{warn}}</div>
    <div :class='s.item'>
      <div :class='s.key'>物业公司名:</div>
      <div :class='s.searchPMs'>
        <input 
          :class='s.value'
          type="text"
          v-model='pmName'
          placeholder="请填写物业公司名称"
          @blur='onPMNameBlur'
          @focus='showSearchResults = true'>
        <div :class='s.results' v-if='showSearchResults'>
          <div v-for='pm in pms' :class='s.pm' @click='onPM(pm)'>{{pm.Name}}</div>
        </div>
      </div>
    </div>
    <div :class='s.item'>
      <div :class='s.key'>真实姓名:</div>
      <input 
        :class='s.value' 
        type="text" 
        v-model='pmUser.ActualName'
        @focus='warn = "", showSearchResults = false'
        placeholder="请填写真实姓名" />
    </div>
    <div :class='s.item'>
      <div :class='s.key'>手机号码:</div>
      <input 
        :class='s.value' 
        type="text" 
        v-model='pmUser.Tel'
        @focus='warn = "", showSearchResults = false'
        placeholder="请填写手机号码"/>
    </div>
    <div :class='s.tip'>
      注意：此功能只针对物业公司人员，普通居民不可绑定。请输入真实信息，以便审核
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
      pmUser: {
        PMID: '',
        ActualName: '',
        OpenID: '',
        Request: 1,
        Tel: ''
      },
      pms: [],
      showSearchResults: false,
      pmName: '',
      btnWord: '申请绑定',
      bindSucc: true,
      code: '',
      isLoading: true
    }
  },
  watch: {
    pmName: function (val) {
      if (val.length >= 2) this.fetchPMs()
      else this.pms = []
    }
  },
  mounted () {
    this.code = this.$route.query.code
    this.fetchOpenID()
  },
  methods: {
    onPM (pm) {
      this.warn = ''
      this.pmUser.PMID = pm.ID
      this.pmName = pm.Name
      this.showSearchResults = false
    },
    onPMNameBlur () {
      // this.showSearchResults = false
    },
    checkInfo () {
      if (this.pmUser.PMID === '') {
        this.warn = '物业公司不存在或填写错误'
        return false
      }
      if (this.pmUser.ActualName === '') {
        this.warn = '真实姓名不能为空'
        return false
      }

      if (this.pmUser.Tel === '' || this.pmUser.Tel.length < 11) {
        this.warn = '手机号码填写错误'
        return false
      }

      if (this.pmUser.OpenID === '' || !this.pmUser.OpenID) {
        this.warn = '用户信息获取失败，请退出重新进行操作'
        return false
      }
      return true
    },
    onBind () {
      if (!this.checkInfo()) return
      fetchpm(this, false, '/open/pm/bind', {
        method: 'POST',
        body: this.pmUser
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
      fetchpm(this, false, '/open/wx/openid/' + this.code, {
        method: 'POST',
        body: JSON.stringify({})
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchOpenID', body)
        if (body.error === 0) {
          sessionStorage.setItem('pmUser', JSON.stringify(body.data))
          this.pmUser.OpenID = body.data.openid
          this.isLoading = false
        }
      })
    },
    fetchPMs () {
      fetchpm(this, false, '/open/pms/name/' + this.pmName, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPMs', body)
        if (body.error === 0 ) this.pms = body.data || []
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
    box-shadow: 1px 0 2px 0px rgba(0, 0, 0, 0.2);
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