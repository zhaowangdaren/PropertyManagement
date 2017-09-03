<template>
  <div :class='s.wrap'>
    <div style="text-align: center;">
      <span :class='s.freeNum' :style='{color: park.FreeNum > 0 ? "green" : "red"}'>{{park.FreeNum}}</span>
      <span :class='s.tip'>可用</span>
    </div>
    <div :class='s.tip'>总车位:<span :class='s.sum'>{{park.Sum}}</span></div>
    <div :class='s.tip'>{{xqName}}</div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  props: {
    xqID: String,
    xqName: String
  },
  data () {
    return {
      park: {
        XQID: '',
        Sum: 0,
        FreeNum: 0
      }
    }
  },
  mounted () {
    this.fetchPark(this.xqID)
  },
  watch: {
    xqID: function () {
      this.fetchPark(this.xqID)
    }
  },
  methods: {
    fetchPark (xqid) {
      console.info('fetchPark', xqid)
      fetchpm(this, false, '/open/park/query/xqid/'+ xqid, {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchPark', body)
        if (body.error === 0) this.park = body.data || {}
      })
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  box-shadow: 1px 1px 2px 0 rgba(0,0,0,0.2), -1px -1px 2px 0 rgba(0,0,0,0.2);
  background-color: #fff;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  padding: 5px;
  // align-items: center;
}
.tip{
  color: #999;
  padding: 5px;
  font-size: 23px;
  overflow: hidden;
  text-overflow:ellipsis;
  white-space: nowrap;
}
.freeNum{
  font-size: 60px;
  text-align: center;
  font-weight: bold;
}
.sum{
  font-size: 30px;
  margin-left: 10px;
}
</style>