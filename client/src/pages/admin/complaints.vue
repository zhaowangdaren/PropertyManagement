<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      投诉数据
    </div>
    <div :class="s.content">
      <div :class="s.left">
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
    </div>
    <p :class="s.red">警告：清除无用数据后，无法恢复，请确认维护人员已经备份，清除操作一般半年为一周期</p>
    <el-button type='danger'>执行清除</el-button>
  </div>
</template>

<script>
import Gauge from '@/components/gauge'
export default {
  components: { Gauge },
  data () {
    return {
      allEvents: {
        value: 10,
        name:'可清除事件数'
      },
      allHandles: {
        value: 50,
        name: '可清除记录数'
      }
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  margin: 10px;
  border: solid 1px #4c87b9;
  width: 100%;
  background: #fff;
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
    justify-content: center;
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
