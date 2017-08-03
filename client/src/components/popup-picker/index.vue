<template>
	<div>

		<div @click="togglePicker">
			<slot></slot>
		</div>

		<div v-transfer-dom="isTransferDom">
			<popup v-model="showValue">
				<div :class="$style.contaier">
					<div :class="$style.header">
						<div :class="$style.cancel" @click="showValue = false">取消</div>
						<div :class="$style.text">{{ title }}</div>
						<div :class="$style.confirm" @click="confirm">确定</div>

					</div>

					<picker
						:data="data"
						:columns="columns"
						v-on:on-change='onScroll'
						v-model="tempValue"/>
				</div>
			</popup>
		</div>
	</div>
</template>

<script>
	
	import Picker from './picker'
	import Popup from './popup'
	import array2string from '@/filters/array2String'
	import value2name from '@/filters/value2name'
	import TransferDom from './transfer-dom'

	const getObject = function (obj) {
		return JSON.parse(JSON.stringify(obj))
	}

	export default {

		props: {
			isTransferDom: {
				type: Boolean,
				default: true
			},
			value: {
				type: Array,
				default () {
					return []
				}
			},
			title: {
				type: String,
				default: ''
			},
			columns: {
				type: Number,
				default: 0
			},
			data: {
				type: Array,
				default () {
					const arr = []
					return [arr]
				}
			}
		},

		data () {
			return {
				showValue: false,
				tempValue: getObject(this.value)
			}
		},

		watch: {
			value (val) {
				this.tempValue = getObject(val)
			}
		},

		directives: {
			TransferDom
		},

		components: {
			Picker,
			Popup
		},

		filters: {
			array2string,
			value2name
		},

		methods: {
			value2name,
			togglePicker () {
				this.showValue = true
			},
			confirm () {
				this.$emit('change', array2string(this.tempValue))
				this.showValue = false
			},
			onScroll(val){
				// console.info('',val);
				this.$emit('onScroll',val);
			}
		}
	}
</script>

<style module lang="less">
	
	.header {
		padding: 10px;
		position: relative;
		display: -webkit-box;
		background-color: #fbf9fe;

		&::after {
			position: absolute;
			bottom: 0;
			left: 0;
			right: 0;
			content: '';
			width: 100%;
			height: 1px;
			background-color: #e5e5e5;
			transform: scaleY(0.5);	
		}
	}

	.cancel, .confirm, .text {
		-webkit-box-flex: 1;
		font-size: 30px;
		color: #345678;
	}

	.text {
		text-align: center;
		color: #666;
	}

	.confirm {
		text-align: right;
	}

	.show-text {
		margin: 0;
		padding: 7px 8px;
		background-color: #FFF;

		p {
			margin: 0;
		}

		&:active {
			background-color: #ECECEC;
		}
	}
</style>