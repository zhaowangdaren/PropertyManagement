<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div :class='s.item' v-for='xq in xqs' @click='onXQ(xq)'>
        <div v-text='xq.Name' :class='s.name'></div>
        <i class='iconfont icon-next'></i>
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
        leftBtns: [{text:'上一步', event: null}],
        title: '选择小区',
        rightBtns: []
      },
      xqs: []
    }
  },
  mounted () {
    this.headerOptions.leftBtns[0].event = this.onReturn
    this.fetchXQs(this.$route.query.communityID)
    sessionStorage.setItem('CommunityID', this.$route.query.communityID)
  },
  methods: {
    onReturn () {
      this.$router.go(-1)
    },
    onXQ (xq) {
      this.$router.push({path: '/wx/complaint/addEvent', query: {XQID: xq.ID, XQName: xq.Name}})
    },
    fetchXQs (communityID) {
      fetchpm(this, false, '/open/xq/kvs',{
        method: 'POST',
        body: {communityID: communityID}
      }).then(resp => {
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.xqs = data.data
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.body{
  margin-top: 80px;
  font-size: 25px;
  .item {
    color: #20a0ff;
    background-color: #fff;
    box-shadow: 0 1px 1px 1px #ddd;
    padding: 15px 10px;
    margin-top: 2px;
    display: flex;
    justify-content: center;
    .name{
      flex: 1;
    }
  }
}
</style>