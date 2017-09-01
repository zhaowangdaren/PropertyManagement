<template>
  <div :class='s.wrap'>
    <div :class='s.start' @click='onStart'>新的投诉</div>
    <div :class='s.start' @click='onProgress'>投诉进度</div>
    <div :class='s.start' @click='onMethod'>投诉流程</div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  data () {
    return {
    }
  },
  methods: {
    onStart () {
      this.$router.push({path: '/wx/complaint/selectStreet'})
    },
    onProgress () {
      this.$router.push({path: '/wx/viewProgress'})
    },
    onMethod () {
      this.$router.push({path: '/wx/complaint/method'})
    },
    fetchOpenID () {
      fetchpm(this, false, '/open/wx/openid/' + this.$route.query.code, {
        method: 'POST'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchOpenID', body)
        if (body.error === 0) {
          sessionStorage.setItem('WXUser', JSON.stringify(body.data))
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  background-color: #20a0ff;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .start{
    text-align: center;
    width: 60%;
    border:solid 2px #fff;
    border-radius: 5px;
    padding: 20px;
    margin: 5px;
    color: #fff;
    font-size: 30px;
    font-weight: bold;
    &:active {
      background-color: rgba(254,254,254, 0.1);
    }
  }
}
</style>