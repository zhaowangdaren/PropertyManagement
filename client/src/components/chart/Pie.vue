<template>
	<div :class='s.wrap'>
		<div ref='pie' :class='s.pie'></div>
	</div>
</template>

<script>
import echarts from 'echarts/lib/echarts'
import 'echarts/lib/chart/pie'
import pieOption from '@/components/chart/pieOption.js'
export default {
	props: {
		dataProps: {
			type: Object,
			default: function () {
				return {
					title: '不同类型的投诉数据'
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
	.pie{
		height: 180px;
		width: 270px;
	}
}
</style>