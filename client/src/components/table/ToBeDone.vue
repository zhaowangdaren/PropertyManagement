<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        Waiting<!-- 待办事件 -->
      </div>
      <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
        <tr>
          <th>Warning</th>
          <th>Index</th>
          <th>Time</th>
          <th>XQ</th>
          <th>Status</th>
          <th>EventLevel</th>
          <th>Type</th>
          <th>操作</th>
        </tr>
        <tr v-for='event in events'>
          <td>
            <i class="iconfont icon-gaojing"></i>
          </td>
          <td v-text='event.Index'></td>
          <td v-text='event.Time'></td>
          <td v-text='event.XQ'></td>
          <td>{{event.Status | filterEventStatus }}</td>
          <td>{{event.EventLevel | filterEventLevel }}</td>
          <td v-text='event.Type'></td>
          <td>
            <el-button type="primary" icon="search" @click='toDetails(event)'>查看</el-button>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
  import filterEventStatus from '@/filters/filterEventStatus'
  import filterEventLevel from '@/filters/filterEventLevel'
  import fetchpm from '@/fetchpm'
  export default {
    filters: {filterEventStatus, filterEventLevel},
    data () {
      return {
        events: []
      }
    },
    mounted () {
      this.fetchEvents('')
    },
    methods: {
      fetchEvents (index) {
        if (!index ) {
          console.info('index', index)
        }
        fetchpm(this, true, '/pm/event', {
          method: 'POST',
          body: {index: index}
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchEvents', body)
          this.events = body.data
        })
      },
      toDetails (event) {
        console.info('toDetails', event.Index)
        // this.$router.push({path:'/street/detail/', params:{index: event.Index}})
        this.$router.push({name:'eventDetail', params: { index: event.Index}})
      }
    }
  }
</script>

<style lang='less' module='s'>
.wrap{
  .content{
    border: solid 1px #4c87b9;
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
    table{
      width: 99%;
      font-size: 15px;
      color: #555;
      margin: auto;
      margin-top: 10px;
      th, td {
        padding: 5px;
        border: solid 1px #ddd;
        text-align: center;
      }
      th{
        background-color: #ddd;
      } 
      tr{
        &:hover {
          background-color: #ddd;
        }
      }
    }
  }
}
</style>