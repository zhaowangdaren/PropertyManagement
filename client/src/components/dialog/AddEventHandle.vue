<template>
  <div>
    <div :class='s.warn' v-text='warn'></div>
    <el-upload
      action="//www.maszfglzx.com:3000/open/upload"
      list-type="picture-card"
      :on-success="handleUploadSucc"
      :on-remove="handleRemove">
      <i class="el-icon-plus"></i>
    </el-upload>
    <el-input
      :class='s.upload'
      type='textarea'
      placeholder='请输入'
      v-model='eventHandle.HandleInfo'></el-input>

    <div :class='s.btns'>
      <el-button type='primary' @click='onSave'>确 认</el-button>
      <el-button @click='onCancel'>取 消</el-button>
    </div>
  </div>
</template>

<script>
import { Message } from 'element-ui'
import BasicDialog from '@/components/dialog/BasicDialog'
import fetchpm from '@/fetchpm'
export default {
  components: { BasicDialog },
  props: {
    eventIndex: String
  },
  data () {
    return {
      warn: '',
      imgs: [],
      eventHandle: {
        Index: '',
        AuthorCategory: 0,
        AuthorName: '',
        HandleInfo: '',
        HandleType: 0,
        Imgs: ''
      }
    }
  },
  mounted () {
    this.eventHandle.Index = this.eventIndex
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.eventHandle.AuthorCategory = user.type
    this.eventHandle.AuthorName = user.UserName
  },
  methods: {
    onSave () {
      this.eventHandle.HandleType = 3
      if (!this.checkImgs()) return
      fetchpm(this, true, '/pm/eventHandle/add', {
        method: 'POST',
        body: this.eventHandle
      }).then(resp => {
        return resp.json()
      }).then(body => {
        if (body.error === 0) {
          Message({message:'恭喜，成功提交事件处理', type:'success'})
          this.$emit('succ', body.data)
          this.$emit('cancel')
        } else {
          this.warn = body.data
        }
      })
    },
    onCancel () {
      this.$emit('cancel')
    },
    checkImgs () {
      if (this.imgs.length > 9) {
        Message({message:'图片数量超过9张', type:'warning'})
        return false
      }
      for (var i = 0; i < this.imgs.length; i++) {
        // this.event['Img' + (i + 1)] = this.imgs[i].response.data.md5
        if (i == (this.imgs.length - 1)) this.eventHandle.Imgs += this.imgs[i].response.data.md5
        else this.eventHandle.Imgs += this.imgs[i].response.data.md5 + ','
      }
      return true
    },
    handleUploadSucc (response, file, fileList) {
      console.info('handleUploadSucc response ',response)
      if (fileList.length > 9) {
        Message({message:'图片数量超过9张', type:'warning'})
      } else {
        this.imgs = fileList
      }
    },
    handleRemove (file, fileList) {
      console.info('handleRemove', file)
      console.info('fileList', fileList)
      this.imgs = fileList
    }
  }
}
</script>

<style lang="less" module='s'>
.warn{
  color: red;
  font-size: 18px;
  margin: 10;
}
.upload{
  margin: 10px auto;
}
.imgDialog{
  width: 100%;
}
.btns{
  margin: 10px 20px;
  display: flex;
  justify-content: flex-end;
}
</style>