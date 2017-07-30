<template>
  <div>
    
  </div>
</template>

<script>
  export default {
    data () {
      return {
        host: '//localhost:3000',
        token: ''
      }
    },
    mounted () {
      // this.login()
      // this.fetchStreets()
      // this.delStreet()
      // this.fetchXQKVs()
      this.onLogin()
    },
    methods: {
      login(){
        fetch('//localhost:3000/login', {
          method: 'POST',
          body:JSON.stringify({user:'admin', password: '123456'})
        }).then(resp => {
          return resp.json()
        }).then(data => {
          console.info('login',data)
        })
      },
      fetchStreets () {
        fetch(this.host + '/street', {
          method: 'POST',
          body: JSON.stringify({name:"", pageSize:10, pageNo: 0})
        }).then(resp => {
          return resp.json()
        }).then(data => {
          console.info('fetchStreets', data)
        })
      },
      delStreet () {
        fetch(this.host + '/street/del', {
          method: 'POST',
          body: JSON.stringify({values: ['1']})
        }).then(resp => {
          return resp.json()
        }).then(data => {
          console.info('delStreet', data)
        })
      },
      fetchXQKVs () {
        fetch(this.host + '/xq/kvs', {
          method: 'POST',
          body:JSON.stringify({name: '2', street: '2'})
        }).then( resp => {
          return resp.json()
        }).then( data => {
          console.info('fetchXQKVs', data)
        })
      },
      onLogin() {
        fetch(this.host + '/login', {
          method: 'POST',
          body: JSON.stringify({username: 'admin', password: 'admin'})
        }).then( resp => {
          return resp.json()
        }).then( body => {
          console.info('onLogin', body)
          this.token = body.token
          this.onHello()
        })
      },
      onHello () {
        fetch(this.host + '/auth/hello', {
          method: 'GET',
          headers: {
            'Authorization': 'Bearer ' + this.token
          }
        }).then( resp => {
          return resp.json()
        }).then( body => {
          console.info(body)
        })
      }
    }
  }
</script>