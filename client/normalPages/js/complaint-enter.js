new Vue({
  el:'#app',
  data: {
    isLoading: true,
    code: '',
    openid: '',
    host: 'https://www.maszfglzx.com'
  },
  mounted () {
    this.code = getURLParam('code')
    this.fetchOpenID()
  },
  methods: {
    onStart () {
      window.location = this.host + '/#/wx/complaint/selectStreet?openid=' + this.openid
    },
    onProgress () {
      window.location = this.host + '/#/wx/viewProgress?openid=' + this.openid
    },
    onMethod () {
      window.location = this.host + '/#/wx/complaint/method'
    },
    fetchOpenID () {
      fetch(this.host + ':3000/open/wx/openid/' + this.code, {
        method: 'POST',
        body: JSON.stringify({})
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchOpenID', body)
        if (body.error === 0) {
          sessionStorage.setItem('WXUser', JSON.stringify(body.data))
          this.openid = body.data.openid
          this.isLoading = false
        }
      })
    }
  }
})