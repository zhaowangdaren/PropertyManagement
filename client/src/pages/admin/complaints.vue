<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      投诉数据
    </div>
    <div :class="s.content">
      <!-- <div :class="s.left">
        <gauge :DATA='allEvents'></gauge>
        <div :class="s.chartTitle">
          投诉总事件数{{allEvents.value}}
        </div>
      </div>
      <div :class="s.right">
        <gauge :DATA='allHandles'></gauge>
        <div :class="s.chartTitle">
          提交总记录数{{allHandles.value}}
        </div>
      </div>
    <p :class="s.red">警告：清除无用数据后，无法恢复，请确认维护人员已经备份，清除操作一般半年为一周期</p>
    <el-button type='danger'>执行清除</el-button> -->
      <div :class='s.item'>
        <num-event-status :class='s.status'></num-event-status>
        <pie :dataProps='diffTypeNum' v-if='hasLoadTypes'></pie>
      </div>
      <div :class='s.item'>
        <bar :dataProps='barData' v-if='barData.data.x.length !== 0'></bar>
      </div>
    </div>
  </div>
</template>

<script>
import NumEventStatus from '@/components/chart/NumEventStatus'
import Pie from '@/components/chart/Pie'
import Bar from '@/components/chart/Bar'
import fetchpm from '@/fetchpm'
export default {
  components: { Pie, NumEventStatus, Bar },
  data () {
    return {
      allEvents: {
        value: 10,
        name:'可清除事件数'
      },
      allHandles: {
        value: 50,
        name: '可清除记录数'
      },
      diffTypeNum: {
        title: '不同类型的投诉数据',
        data: []
      },
      hasLoadTypes: false,
      eventXQs:[],
      barData:{
        title: '各小区投诉数量',
        data: {
          y: [],
          x: []
        }
      }
    }
  },
  mounted () {
    this.fetchDiffTypeNum()
    this.fetchEventXQ()
  },
  methods: {
    fetchDiffTypeNum () {
      fetchpm(this, false, '/open/event/types/num', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchDiffTypeNum', body)
        this.diffTypeNum.data = this.formatTypes(body.data || [])
        this.hasLoadTypes = true
        console.info('diffTypeNum', this.diffTypeNum)
      })
    },
    /**
     * 转化type
     * @param  {Array} types [{"_id":{"type":"type3"},"num":2},{"_id":{"type":"Type1"},"num":3}]
     * @return {Array}       [{"value": 50, "name": "type3"}]
     */
    formatTypes (types) {
      var result = []
      for (var i = 0; i < types.length; i++) {
        result.push({name: types[i]._id.type, value: types[i].num})
      }
      return result
    },
    fetchEventXQ () {
      fetchpm(this, true, '/pm/event/key/nums/xqid?recs=20', {
        method: 'GET'
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventXQ', body)
        if (body.error === 0) this.eventXQs = body.data || []
        console.info('eventXQs', this.eventXQs)
        this.formatEventXQ(this.eventXQs).then(data => {
          this.barData.data = data
          console.info('fetchEventXQ then',data)
        })
      })
    },
    formatEventXQ (events) {
      events.sort((a, b) => {
        return a.num - b.num
      })
      var result = []
      var ids = []
      var x = []
      for (var i = 0; i < events.length; i++) {
        ids.push(events[i]._id.xqid)
        x.push({value: events[i].num})
      }
      console.info('ids', ids)
      var y = []
      return fetchpm(this, true, '/pm/xq/ids', {
        method: 'POST',
        body: {values: ids}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('xqids', body)
        if (body.error === 0) {
          y = body.data || []
          y = y.map(item => {
            return item.Name
          })
        }
        return {x: x, y: y}
      })
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  margin: 10px;
  border: solid 1px #4c87b9;
  text-align: center;
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
  .content{
    padding: 10px;
    font-size: 20px;
    text-align: center;
    color: #999;
    // display: flex;
    align-items: flex-start;
    .item{
      flex: 1;
      margin: 5px;
    }
  }

}

.status{
  margin-bottom: 5px;
}
</style>
