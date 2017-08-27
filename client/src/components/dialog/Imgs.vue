<template>
  <div :class='s.wrap'>
    <img v-for='img in sImgs' :src='"//www.maszfglzx.com:3000/open/image/" + img'>
  </div>
</template>

<script>
export default {
  props: {
    imgs: String
  },
  data () {
    return {
      sImgs: []
    }
  },
  mounted () {
    this.formatImgs()
  },
  methods: {
    formatImgs () {
      this.sImgs = this.imgs.split(',')
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  text-align: center;
  img {
    width: 90%;
    overflow-y: scroll;
    height: 100%;
    background-color: #ddd;
    padding-bottom: 5px;
  }
}
</style>