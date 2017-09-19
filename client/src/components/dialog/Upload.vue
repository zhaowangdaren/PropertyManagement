<template>
  <div>
    <el-upload
      :class='s.upload'
      drag
      :action='host + "/open/upload/file"'
      multiple
      :on-success='handleUploadSucc'
      :on-remove='handleRemove'>
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
    </el-upload>
    <div :class='s.bottomBtns'>
      <el-button type='primary' @click='onSure' :loading='uploading'>确 定</el-button>
      <el-button @click='onCancel'>取 消</el-button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    uploading: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      host: 'https://www.maszfglzx.com:3000',
      files:[]
    }
  },
  methods: {
    handleUploadSucc (response, file, fileList) {
      this.files = fileList
    },
    handleRemove () {
      this.files = fileList
    },
    onSure () {
      this.$emit('sure', this.files)
    },
    onCancel () {
      this.$emit('cancel')
    }
  }
}
</script>

<style lang="less" module='s'>
.upload{
  width: 80%;
}
</style>