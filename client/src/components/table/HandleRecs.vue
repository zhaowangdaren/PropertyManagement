<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      居民物业纠纷处理记录<!-- 居民物业纠纷处理 -->
    </div>
    <div :class='s.body'>
      <table>
        <!-- 提交人类别  提交人用户名  提交时间  提交文本说明  操作 -->
        <tr>
          <th>提交人类别</th>
          <th>提交人用户名</th>
          <th>提交时间</th>
          <th>提交内容</th>
          <th>操作</th>
        </tr>
        <tr v-for='handle in eventHandles'>
          <td v-text='handle.AuthorCategory'></td>
          <td v-text='handle.AuthorName'></td>
          <td>{{ handle.Time | filterTime }}</td>
          <td v-text='handle.HandleInfo' class="descr"></td>
          <td>
            <el-button type="primary" icon="search" @click='onHandleDetails(handle)'>详情</el-button>
          </td>
        </tr>
        <tr v-if='eventHandles.length === 0'>
          <td colspan="5">无记录</td>
        </tr>
      </table>
      <el-dialog 
        v-model="showHandleDetails" 
        size="tiny" 
        title='物业纠纷处理详情'>
        <details-event-handle 
          v-if='showHandleDetails'
          :eventHandle='showingHandle'></details-event-handle>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import DetailsEventHandle from '@/components/dialog/DetailsEventHandle'
import fetchpm from '@/fetchpm'
import filterTime from '@/filters/filterTime'
export default {
  components: { DetailsEventHandle },
  filters: { filterTime },
  data () {
    return {
      showHandleDetails: false,
      eventHandles: [],
      showingHandle:{}
    }
  },
  mounted () {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.fetchEventHandles(user.UserName, user.type)
  },
  methods: {
    onHandleDetails (eventHandle) {
      this.showingHandle = eventHandle
      this.showHandleDetails = true
    },
    fetchEventHandles (AuthorName, AuthorCategory) {
      fetchpm(this, true, '/pm/eventHandle/kvs', {
        method: 'POST',
        body: {AuthorName: AuthorName, AuthorCategory: AuthorCategory}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventHandles', body)
        if (body.error === 0) {
          this.eventHandles = body.data || []
        }
        
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  margin: 10px;
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
    img{
      width: 20px;
      margin-right: 10px;
    }
  }
  .body{
    margin: 10px;
  }
}
</style>