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
export default {
  components: { ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: ['上一步'],
        title: 'Select Community',
        rightBtns: []
      },
      communities: [],
      host:'http://10.176.118.61:3000'
    }
  },
  mounted () {
    this.fetchCommunities(this.$route.query.streetID)
  },
  methods: {
    onCommunity (community) {
      this.$router.push({path: '/wx/complaint/xq', query: {communityID: community.ID}})
    },
    fetchCommunities (streetID) {
      fetch(this.host + '/open/community/kvs',{
        method: 'POST',
        body: JSON.stringify({streetID: streetID})
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