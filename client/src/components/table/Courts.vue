<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        申请法院调解请求<!-- 推送法院调解请求 -->
      </div>
      <div :class='s.body'>
        <!-- search result -->
        <table>
        <!-- 事件编号 开始时间  所在街道  所在社区  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>事件编号</th>
            <th>开始时间</th>
            <th>所在街道</th>
            <th>所在社区</th>
            <th>所在小区</th>
            <th>事件状态</th>
            <th>事件等级</th>
            <th>事件类别</th>
            <th>操作</th>
          </tr>
          <tr v-for='handle in searchResult'>
            <td v-text='handle.Index'></td>
            <td v-text='handle.Time'></td>
            <td v-text='handle.Street'></td>
            <td v-text='handle.Community'></td>
            <td v-text='handle.XQ'></td>
            <td>{{handle.Status | filterEventStatus }}</td>
            <td>{{handle.EventLevel | filterEventLevel }}</td>
            <td v-text='handle.Type'></td>
            <td>
              <el-button type="primary" icon="search" @click='toDetails(handle)'>查看</el-button>
            </td>
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
        searchResult:[]
      }
    },
    mounted () {
      var params = {
        toCourt: 1
      }
      this.fetchEventKVs(params)
    },
    methods: {
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
          this.searchResult = body.data
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
  width: 100%;
  .content{
    border: solid 1px #4c87b9;
    margin-top: 20px;
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
      table{
        width: 100%;
        font-size: 15px;
        color: #555;
        background-color: #fff;
        th, td {
          padding: 5px;
          border: solid 1px #ddd;
          text-align: center;
        }
        th{
          background-color: #eee;
        } 
        tr{
          &:hover {
            background-color: #ddd;
          }
        }
        .key{
          background-color: #ddd;
        }
      }
    }    
  }
}
</style>