<template>
	<div :class='s.wrap'>
		<div ref='pie' :class='s.pie'></div>
	</div>
</template>

<script>
import echarts from 'echarts/lib/echarts'
import 'echarts/lib/chart/pie'
import 'echarts/lib/component/tooltip'
import 'echarts/lib/component/title'
import 'echarts/lib/component/legend'


import pieOption from '@/components/chart/pieOption.js'
export default {
	props: {
		dataProps: {
			type: Object,
			default: function () {
				return {
					title: '不同类型的投诉数据',
					data: []
				}
			}
		}
	},
	data () {
		return {
			pie: null
		}
	},
	mounted () {
		this.initPie()
	},
	methods: {
		initPie () {
			this.pie = echarts.init(this.$refs.pie)
			this.pie.setOption(pieOption(this.dataProps))
		}
	}
}
</script>
<style lang="less" module='s'>
.wrap{
	background-color: #fff;
	border: solid 1px #ddd;
	border-radius: 5px;
	padding: 5px;
	box-shadow: 1px 1px 4px -2px rgba(0, 0, 0, 0.5),-1px -1px 4px -2px rgba(0, 0, 0, 0.5);
	.pie{
		height: 360px;
		width: 480px;
	}
}
</style>