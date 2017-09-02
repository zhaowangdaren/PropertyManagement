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