<template>
	<div>
		<div :class='s.new'>
			<div class='searchKey'>新增事件类型:</div>
			<el-input v-model='newEventType' :class='s.input' placeholder='请输入新增类型'></el-input>
			<el-button type='primary' icon='plus' @click='onAdd' :loading='isAdding' class='add'>新增</el-button>
		</div>
		<table>
			<tr>
				<th>事件类型</th>
				<th class="flex1">操作</th>
			</tr>
			<tr v-for='eventType in eventTypes'>
				<td>{{eventType.Type}}</td>
				<td>
					<el-button type='danger' icon='delete' @click='onShowDel(eventType)' :loading='isDeling' class='del'>删除</el-button>
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
import { Message } from 'element-ui'
export default {
	data () {
		return {
			eventTypes: [],
			newEventType: '',
			isAdding: false,
			isDeling: false,
			delingType: '',
			showDel: false
		}
	},
	mounted () {
		this.fetchEventTypes()
	},
	methods: {
		onShowDel (eventType) {
			this.$confirm('将删除事件类型：' + eventType.Type + ', 是否继续？', '提示', {
				confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
			}).then(() => {
				this.onDel(eventType)
			}).catch(() => {
				Message({type: 'info', message: '已取消删除'})
			})
		},
		onDel (eventType) {
			if (this.isDeling) return
			this.isDeling = true
			fetchpm(this, true, '/pm/event/type/del', {
				method: 'POST',
				body: eventType
			}).then(resp => {
				return resp.json()
			}).then(body => {
				if (body.error !== 0) {
					Message({message: body.data, type:'error'})
				} else {
					this.fetchEventTypes()
					Message({type: 'success', message: '删除成功'})
				}
				this.isDeling = false
			})
		},
		onAdd () {
			if (this.newEventType === '' || this.isAdding) return
			this.isAdding = true
			fetchpm(this, true, '/pm/event/type/add', {
				method: 'POST',
				body: {type: this.newEventType}
			}).then(resp => {
				return resp.json()
			}).then(body => {
				console.info('onAdd', body)
				if (body.error !== 0) {
					Message({message: body.data, type:'error'})
				} else {
					this.newEventType = ''
					Message({message: '新增成功', type:'success'})
					this.fetchEventTypes()
				}
				this.isAdding = false
			})
		},
		fetchEventTypes () {
			fetchpm(this, false, '/open/wx/event/types',{
				method: 'GET'
			}).then(resp => {
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

<style lang="less" module='s'>
.new{
	display: flex;
	align-items: center;
	margin: 5px;
	.title{
		font-size: 20px;
		margin: auto 5px;
	}
	.input{
		flex: 1;
		margin: auto 5px;
	}
}
</style>