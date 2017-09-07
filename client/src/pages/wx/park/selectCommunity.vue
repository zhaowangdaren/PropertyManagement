<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div :class='s.item' v-for='community in communities' @click='onCommunity(community)'>
        <div v-text='community.Name' :class='s.name'></div>
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
        title: '选择社区'
      },
      communities: []
    }
  },
  mounted () {
    this.headerOptions.leftBtns[0].event = this.onReturn
    this.fetchCommunities(this.$route.query.streetID)
    sessionStorage.setItem('StreetID', this.$route.query.streetID)
  },
  methods: {
    onReturn () {
      this.$router.go(-1)
    },
    onCommunity (community) {
      this.$router.push({path: '/wx/park/manager/select/xq', query: {communityID: community.ID}})
    },
    fetchCommunities (streetID) {
      fetchpm(this, false, '/open/community/kvs',{
        method: 'POST',
        body: {streetID: streetID}
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.communities = data.data
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