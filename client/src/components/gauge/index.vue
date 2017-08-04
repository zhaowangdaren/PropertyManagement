<template lang="html">
  <div :class="s.wrap">
    <div :class="s.chart" ref='chart'>

    </div>
  </div>
</template>

<script>
import echarts from 'echarts/lib/echarts'
import 'echarts/lib/chart/gauge'
import chartOption from './option'
export default {
  props: {
    DATA: {
      type: Object,
      default: function () {
        return {
          value: 0,
          name: '可清除事件数'
        }
      }
    }
  },
  data () {
    return {
      chart: null
    }
  },
  mounted () {
    this.initChart()
  },
  methods: {
    initChart () {
      this.chart = echarts.init(this.$refs.chart)
      this.chart.setOption(chartOption(this.DATA))
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  .chart{
    width: 400px;
    height: 400px;
    background-color: skyblue;
  }
}
</style>
