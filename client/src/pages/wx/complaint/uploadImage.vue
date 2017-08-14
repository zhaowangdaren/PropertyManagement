<template>
  <div :class='s.wrap'>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <div :class='s.red'>注意：图片大小不要超过5M，最多上传9张</div>
      <el-upload
        :class='s.upload'
        action="//47.94.7.154:3000/open/upload"
        list-type="picture-card"
        :on-success="handleUploadSucc"
        :on-preview="handlePictureCardPreview"
        :on-remove="handleRemove">
        <i class="el-icon-plus"></i>
      </el-upload>
      <el-dialog v-model="dialogVisible" size="large">
        <img width="100%" :src="dialogImageUrl" alt="">
      </el-dialog>
      <div :class='s.btn' @click='onNext'>下一步</div>
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import { Message } from 'element-ui'
import fetchpm from '@/fetchpm'
export default {
  components: { ActionBar },
  data () {
    return {
      dialogVisible: false,
      dialogImageUrl: '',
      headerOptions: {
        title: '上传图片',
        rightBtns: ['完成']
      },
      event:{
        Imgs: ''//以逗号为分隔符
      },
      imgs:[],
      warn:''
    }
  },
  mounted () {
    this.fetchEvent()
  },
  methods: {
    fetchEvent () {
      fetchpm(this, false, '/open/event/id/' + this.$route.query.id, {
        method: 'POST'
      }).then(response => {
        return response.json()
      }).then(body => {
        console.info('fetchEvent', body)
        if (body.error === 0 && body.data && body.data.length >0 ) this.event = body.data[0]
      })
    },
    onNext () {
      if(!this.checkImgs()) return
      fetchpm(this, false, '/open/event/update', {
        method: 'PUT',
        body: this.event
      }).then(response => {
        return response.json()
      }).then(body => {
        console.info('onNext', body)
        if (body.error === 0) this.$router.push({path:'/wx/complaint/finish'})
        else {
          Message({message: body.data, type: 'error'})
        }
      })
    },
    checkImgs () {
      if (this.imgs.length > 9) {
        Message({message:'图片数量超过9张', type:'warning'})
        return false
      }
      for (var i = 0; i < this.imgs.length; i++) {
        // this.event['Img' + (i + 1)] = this.imgs[i].response.data.md5
        this.event.Imgs += this.imgs[i].response.data.md5 + ','
      }
      return true
    },
    handleUploadSucc (response, file, fileList) {
      console.info('handleUploadSucc response ',response)
      console.info(file)
      console.info(fileList)
      if (fileList.length > 9) {
        Message({message:'图片数量超过9张', type:'warning'})
      } else {
        this.imgs = fileList
      }
    },
    handlePictureCardPreview (file) {
      console.info('handlePictureCardPreview', file)
      if (file.response.error !== 0) return
      this.dialogImageUrl = this.host + '/open/image/' + file.response.data.md5
      this.dialogVisible = true
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
.body{
  margin: 80px 5px 5px 5px;
  font-size: 25px;
  .red{
    color: red;
  }
  .upload{
    margin-top: 10px;
  }
  .btn{
    background-color: #20a0ff;
    font-size: 30px;
    color: #fff;
    text-align: center;
    padding: 15px 0;
    margin: 20px;
  }
}
</style>