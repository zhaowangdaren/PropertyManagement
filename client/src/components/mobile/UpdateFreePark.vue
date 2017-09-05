<template>
  <div>
    <div :class='s.succ'>
      <i class="iconfont icon-succ" v-if='succ'></i>
    </div>
    <div :class='s.item'>
      <div>空闲车位</div>
      <div :class='s.inputWrap'><input :class='s.value' type="number" name="" v-model='park.FreeNum' @focus='succ = false'></div>
    </div>
    <div :class='s.item'>
      <div>总车位</div>
      <div  :class='s.inputWrap'><input :class='s.value' type="number" name="" v-model='park.Sum' @focus='succ = false'></div>
    </div>
    <div :class='s.btn' @click='onUpdate'>
      发布
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  props: {
    xqId: String,
    xqName: String
  },
  data () {
    return {
      succ: false,
      park: {
        Sum: 0,
        FreeNum: 0
      }
    }
  },
  mounted () {
    this.fetchPark(this.xqId)
  },
  methods: {
    onUpdate () {
      this.park.Sum = parseInt(this.park.Sum)
      this.park.FreeNum = parseInt(this.park.FreeNum)
      fetchpm(this, false, '/open/park/update', {
        method: 'POST',
        body: this.park
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onUpdate', body)
        if (body.error === 0) this.succ = true
      })
    },
    fetchPark (xqId) {
      fetchpm(this, false, '/open/park/query/xqid/' + xqId, {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPark', body)
        if (body.error ===0 ) this.park = body.data || {Sum: 0, FreeNum: 0}
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.succ{
  margin-top: 10px;
  text-align: center;
  i{
    color: rgb(25, 163, 24);
    font-size: 40px;
  }
}
.item{
  font-size: 25px;
  margin: 10px 30px;
}
.inputWrap{
  border: solid 1px #999;
  border-radius: 5px;
  input{
    border: none;
  }
}
.value{
  width: 100%;
  // box-shadow: 0px 1px 1px 1px rgba(0,0,0,0.2) inset;
  padding: 10px;
  background-color: #fff;
  box-sizing: border-box;
  font-size: 25px;
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
</style>