<template>
  <div :class='s.wrap'>
    <div ref='chart' :class='s.chart'></div>
  </div>
</template>

<script>
import echarts from 'echarts/lib/echarts'
import 'echarts/lib/chart/bar'
import 'echarts/lib/component/tooltip'
import 'echarts/lib/component/title'
import 'echarts/lib/component/legend'
import barOption from '@/components/chart/barOption'
export default {
  props: {
    dataProps: {
      type: Object,
      default: function () {
        return {
          title: '',
          data: []
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
      this.chart.setOption(barOption(this.dataProps))
    }
  }
}
</script>
<style lang='less' module='s'>
.wrap{
  width: 100%;
  height: 100%;
  background-color: #fff;
  border: solid 1px #ddd;
  border-radius: 5px;
  box-shadow: 1px 1px 4px 0px #ddedf7,-1px -1px 4px 0px #ddedf7;
}
.chart{
  width: 90%;
  height: 600px;
}
</style>