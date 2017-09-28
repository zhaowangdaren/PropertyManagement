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
    <div :class='s.reason'>
      <textarea placeholder="请输入理由" v-model='eventHandle.HandleInfo'></textarea>
    </div>
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
    event: {
      type: Object,
      default: function () {
        return {
          EventLevel: 0
        }
      }
    }
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
      inputEventLevel: 0,
      updateEvent: {},
      eventHandle: {
        Index: '',
        AuthorCategory: 0,
        AuthorName: '',
        HandleInfo: '',
        HandleType: 7,
        EventLevel: 0,
        Imgs: ''
      }
    }
  },
  mounted () {
    // Object.assign(this.updateEvent, this.event)
    this.inputEventLevel = this.event.EventLevel

    this.eventHandle.Index = this.event.Index
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.eventHandle.AuthorCategory = user.type
    this.eventHandle.AuthorName = user.UserName
    this.eventHandle.EventLevel = this.event.EventLevel
  },
  methods: {
    onCancel () {
      this.$emit('cancel')
    },
    onSave () {
      this.eventHandle.EventLevel = parseInt(this.inputEventLevel)
      // this.eventHandle.HandleInfo = '审核投诉等级：' + (this.options.find(item => { return item.value === this.eventHandle.EventLevel})).label
      fetchpm(this, true, '/pm/eventHandle/audit/level', {
        method: 'POST',
        body: this.eventHandle
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0) {
          Message({message:'恭喜，成功提交事件处理', type:'success'})
          this.eventHandle.HandleInfo = ''  
          this.$emit('succ', this.eventHandle.EventLevel)
          this.$emit('cancel')
        } else {
          this.warn = body.data
          Message({message:body.data, type:'error'})
        }
      })
      // console.info('onSave')
      // if (this.inputEventLevel == '') return
      // this.updateEvent.EventLevel = this.inputEventLevel
      // if (this.updateEvent.Status === 0) this.updateEvent.Status = 1
      // fetchpm(this, true, '/pm/event/update', {
      //   method: 'POST',
      //   body: this.updateEvent
      // }).then(resp => {
      //   return resp.json()
      // }).then(body => {
      //   console.info('onSave', body)
      //   if (body.error === 0) {
      //     // this.warn = '审核成功'
      //     Message({message:'恭喜，审核成功', type:'success'})
      //     this.onCancel()
      //     this.$emit('succ', this.inputEventLevel)
      //   }
      //   else this.warn = body.data
      // })
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
.reason{
  margin-top: 10px; 
  text-align: center;
  textarea{
    margin: 0 auto;
    width: 80%;
    min-height: 200px;
  }
}
</style>