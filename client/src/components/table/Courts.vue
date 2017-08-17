<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        推送法院调解请求<!-- 推送法院调解请求 -->
      </div>
      <div :class='s.body'>
        <!-- search result -->
        <table>
        <!-- 事件编号 开始时间  所在街道  所在社区  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>事件编号</th>
            <th>时间</th>
            <th>街道</th>
            <th>社区</th>
            <th>小区</th>
            <th>事件状态</th>
            <th>事件等级</th>
            <th>事件类型</th>
            <th>操作</th>
          </tr>
          <tr v-for='(handle, index) in searchResult'>
            <td v-text='handle.Index'></td>
            <td v-text='handle.Time'></td>
            <td >{{streets[index] ? streets[index].Name : ''}}</td>
            <td>{{communities[index] ? communities[index].Name : ''}}</td>
            <td>{{xqs[index] ? xqs[index].Name : ''}}</td>
            <td>{{handle.Status | filterEventStatus }}</td>
            <td>{{handle.EventLevel | filterEventLevel }}</td>
            <td v-text='handle.Type'></td>
            <td>
              <el-button type="primary" icon="search" @click='toDetails(handle)'>查看</el-button>
            </td>
          </tr>
          <tr v-if='searchResult.length === 0'>
            <td colspan="9">无记录</td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import fetchpm from '@/fetchpm'
  export default {
    props: {
      editable: Boolean
    },
    filters: {filterEventStatus, filterEventLevel},
    data () {
      return {
        host: '//localhost:3000',
        searchResult:[],
        streets: [],
        communities: [],
        xqs:[]
      }
    },
    mounted () {
      var params = {
        toCourt: 1
      }
      this.fetchEventKVs(params)
    },
    methods: {
      fetchStreets (ids) {
        fetchpm(this, true, '/pm/street/ids', {
          method: 'POST',
          body: {values: ids}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchCommunitiesByStreetID', body)
          this.streets = body.data || []
        })
      },
      fetchCommunities (ids) {
        fetchpm(this, true, '/pm/community/ids', {
          method: 'POST',
          body: {values: ids}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchCommunitiesByStreetID', body)
          this.communities = body.data || []
        })
      },
      fetchXQs (ids) {
        if (!ids) return
        fetchpm(this, true, '/pm/xq/ids', {
          method: 'POST',
          body: {values: ids}
        }).then(resp => {
          return resp.json()
        }).then( body => {
          console.info('fetchXQs', body)
          if(body.error !== 0) console.error("Error: search fetchXQByCommunityName. Reason:" + body.data)
          else if (body.data !== null ) {
            this.xqs = body.data
          }
        })
      },
      fetchEventKVs (params) {
        if (!params ) {
          console.info('params', params)
          params = {}
        }
        fetchpm(this, true, '/pm/event/kvs', {
          method: 'POST',
          body: params
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEventKVs', body)
          this.searchResult = body.data || []
          let streetIDs = this.searchResult.map((item) => {
            return item.StreetID
          })
          this.fetchStreets(streetIDs)
          let communityIDs = this.searchResult.map((item) => {
            return item.CommunityID
          })
          this.fetchCommunities(communityIDs)
          let xqIDs = this.searchResult.map((item) => {
            return item.XQID
          })
          this.fetchXQs(xqIDs)
        })
      },
      toDetails (event) {
        console.info('toDetails', event.Index)
        // this.$router.push({name: 'eventDetail', params:{index: event.Index}})
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  // width: 100%;
  margin: 10px;
  .content{
    border: solid 1px #4c87b9;
    margin-top: 20px;
    background-color: #fff;
    .title{
      color: #fff;
      font-size: 20px;
      background: #4c87b9;
      padding: 10px;
      display: flex;
      align-items: center;
      img{
        width: 20px;
        margin-right: 10px;
      }
    }
    .body{
      margin: 10px;
      .key{
        background-color: #ddd;
      }
    }    
  }
}
</style>