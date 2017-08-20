<template>
	<div>
		<div :class='s.new'>
			<span>新增事件类型</span>
			<el-input v-module='newEventType'></el-input>
			<el-button type='primary' icon='plus' @click='onAdd'>新增</el-button>
		</div>
		<table>
			<tr>
				<th>事件类型</th>
				<th class="flex1">操作</th>
			</tr>
			<tr v-for='eventType in eventTypes'>
				<td>{{eventType.Type}}</td>
				<td class="flex1">
					<el-button type='danger' icon='delete'>删除</el-button>
				</td>
			</tr>
			<tr v-if='eventTypes.length === 0'>
				<td colspan="2">无记录</td>
			</tr>
		</table>
	</div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
	data () {
		return {
			eventTypes: [],
			newEventType: ''
		}
	},
	mounted () {
		this.fetchEventTypes()
	},
	methods: {
		onAdd () {

		},
		fetchEventTypes (){
			fetchpm(this, false, '/open/wx/event/types').then(resp => {
				return resp.json()
			}).then(body => {
				console.info('fetchEventTypes', body)
				if (body.error !== 1){
					this.eventTypes = body.data || []
				}
			})
		}
	}
}
</script>

<style lang="less" modules='s'>
.new{
	display: flex;
}
</style>