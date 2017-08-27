<template>
  <div :class='s.wrap'>
    <div :class='s.warn' v-text='warn'></div>
    <el-select v-model="inputEventLevel" placeholder="请选择">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value">
      </el-option>
    </el-select>
    <div :class='s.btns'>
      <el-button type='primary' @click='onSave'>确 认</el-button>
      <el-button @click='onCancel'>取 消</el-button>
    </div>
  </div>
</template>

<script type="text/javascript">
import { Message } from 'element-ui'
import fetchpm from '@/fetchpm'
export default {
  props: {
    event: Object
  },
  data() {
    return {
      warn: '',
      options: [
        {
          value: 0,
          label: '未定'
        },
        {
          value: 1,
          label: '特急'
        }, {
          value: 2,
          label: '加急'
        }, {
          value: 3,
          label: '急'
        }
      ],
      inputEventLevel: '',
      updateEvent: {}
    }
  },
  mounted () {
    Object.assign(this.updateEvent, this.event)
    this.inputEventLevel = this.updateEvent.EventLevel
  },
  methods: {
    onCancel () {
      this.$emit('cancel')
    },
    onSave () {
      console.info('onSave')
      if (this.inputEventLevel == '') return
      this.updateEvent.EventLevel = this.inputEventLevel
      fetchpm(this, true, '/pm/event/update', {
        method: 'POST',
        body: this.updateEvent
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onSave', body)
        if (body.error === 0) {
          // this.warn = '审核成功'
          Message({message:'恭喜，审核成功', type:'success'})
          this.onCancel()
          this.$emit('succ', this.inputEventLevel)
        }
        else this.warn = body.data
      })
    }
  }
}
</script>

<style lang='less' module='s'>
.wrap{
  .warn{
    color: red;
    font-size: 18px;
    margin: 10;
  }
  .btns{
    margin: 10px 20px;
    display: flex;
    justify-content: flex-end;
  }
}
</style>