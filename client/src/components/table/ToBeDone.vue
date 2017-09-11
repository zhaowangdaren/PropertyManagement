<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        待办事件<!-- 居民物业纠纷处理 -->
      </div>
      <div :class='s.body'>
        <!-- search result -->
        <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>警告类型</th>
            <th>事件编号</th>
            <th>开始时间</th>
            <th>所在社区</th>
            <th>所在小区</th>
            <th>事件状态</th>
            <th>事件等级</th>
            <th>事件类型</th>
            <th>操作</th>
          </tr>
          <tr v-for='(handle, index) in events'>
            <td>
              <i :class='"iconfont icon-gaojing " + s.lv1' v-if='handle.EventLevel === 1' ></i>
              <i :class='"iconfont icon-gaojing " + s.lv2' v-else-if='handle.EventLevel === 2' ></i>
              <i :class='"iconfont icon-gaojing " + s.lv3' v-else-if='handle.EventLevel === 3' ></i>
              <i class="iconfont icon-gaojing" v-else></i>
            </td>
            <td v-text='handle.Index'></td>
            <td>{{ handle.Time | filterTime}}</td>
            <td>{{communities[index] ? communities[index].Name : ''}}</td>
            <td>{{xqs[index] ? xqs[index].Name : ''}}</td>
            <td>{{handle.Status | filterEventStatus }}</td>
            <td>{{handle.EventLevel | filterEventLevel }}</td>
            <td v-text='handle.Type'></td>
            <td>
              <el-button type='text' :plain="true" v-if='handle.RequestClose === 1' >已申请关闭</el-button>
              <el-button type="primary" icon="search" @click='toDetails(handle)' class='view'>查看</el-button>
            </td>
          </tr>
          <tr v-if='events.length == 0'>
            <td colspan="9">没有记录</td>
          </tr>
        </table>
        <div :class='s.pageWrap'>
          <el-pagination
            layout="total, prev, pager, next"
            @current-change='onChangePage'
            :page-size='pageSize'
            :total="sum">
          </el-pagination>
          <!-- <div :class='s.pageTip'>其中<span>{{pageSize - events.length}}</span>条事件已关闭或已被用户撤销</div> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import filterEventStatus from '@/filters/filterEventStatus'
import filterEventLevel from '@/filters/filterEventLevel'
import filterTime from '@/filters/filterTime'
import fetchpm from '@/fetchpm'
import { Message } from 'element-ui'
export default {
  props: {
    editable: Boolean
  },
  filters: {filterEventStatus, filterEventLevel, filterTime},
  data () {
    return {
      userStreetID: '',
      userType: 0,
      //user Input

      events: [],
      pageNo: 0,
      pageSize: 10,
      sum: 0,
      communities: [],
      xqs:[]
    }
  },
  mounted () {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.userStreetID = user.StreetID
    this.userType = user.type || 0
    this.fetchEvents()
  },
  methods: {
    fetchEvents () {
      switch (this.userType) {
        case 1:
          this.onAdminFetch()
          break
        case 2:
          this.onGov()
          break
        case 3:
          this.onStreet()
          break
      }
    },
    onAdminFetch () {
      var kvs = {}
      this.fetchEventsByKVsPage(kvs, this.pageNo, this.pageSize)
    },
    onGov () {
      var kvs = {
        NoticeGov: 1
      }
      this.fetchEventsByKVsPage(kvs, this.pageNo, this.pageSize)
    },
    onStreet () {
      var kvs = {
        StreetID: this.userStreetID
      }
      this.fetchEventsByKVsPage(kvs, this.pageNo, this.pageSize)
    },
    onChangePage (curPage) {
      this.pageNo = curPage - 1
      this.fetchEvents()
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
    fetchEventsByKVsPage(kvs, pageNo, pageSize) {
      fetchpm(this, true, '/pm/event/kvs/page', {
        method: 'POST',
        body: {KVs: kvs, PageNo: pageNo, PageSize: pageSize}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0) {
          this.events = (body.data.events || []).sort((a, b) => {
            return b.Time -a.Time
          })
          this.sum = body.data.sum || 0
          let communityIDs = this.events.map((item) => {
            return item.CommunityID
          })
          this.fetchCommunities(communityIDs)
          let xqIDs = this.events.map((item) => {
            return item.XQID
          })
          this.fetchXQs(xqIDs)
        } else {
          Message({type: 'error', message: body.data})
        }
      })
    },
    toDetails (event) {
      console.info('toDetails', event.Index)
      this.$router.push({name: 'GovDetailEvent', params:{index: event.Index}})
      // this.$router.push({path:'/street/detail/' + event.Index})
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  .content{
    border: solid 1px #4c87b9;
    margin: 10px;
    background-color: #fff;
    .title{
      color: #fff;
      font-size: 20px;
      background: #4c87b9;
      padding: 10px;
      display: flex;
      align-items: center;

      color: #fff;
      font-size: 30px;
      font-family: "华文行楷";
      background: #4c87b9;
      background: -webkit-gradient(linear,center top, center bottom,from(#ff0000), to(#000000));
      
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
.lv3{
  color: #4c87b9;
}
.lv2{
  color: orange;
}
.lv1{
  color: red;
}
.pageWrap{
  display: flex;
  align-items: center;
}
.pageTip{
  color: #999;
  span{
    color: #222;
    font-size: 20px;
  }
}
</style>