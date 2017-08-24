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
      <pie :dataProps='diffTypeNum' v-if='hasLoadTypes'></pie>
    </div>
  </div>
</template>

<script>
// import Gauge from '@/components/gauge'
import Pie from '@/components/chart/Pie'
import fetchpm from '@/fetchpm'
export default {
  components: { Pie },
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
      hasLoadTypes: false
    }
  },
  mounted () {
    this.fetchDiffTypeNum()
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
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  margin: 10px;
  border: solid 1px #4c87b9;
  width: 100%;
  text-align: center;
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
  .content{
    padding: 10px;
    font-size: 20px;
    text-align: center;
    color: #999;
    display: flex;
    // justify-content: center;
    .left, .right {
      position: relative;
      .chartTitle{
        margin-top: -80px;
        color: #000;
      }
    }
  }
  .red{
    color: red;
    margin: 20px;
  }

}
</style>
