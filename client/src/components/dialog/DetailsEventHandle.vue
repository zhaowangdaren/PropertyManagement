<template>
  <div :class='s.wrap'>
    <div :class='s.imgs'>
      <img :src='host + "/open/image/" +img' v-for='img in imgs'>
    </div>
    <div v-text='eventHandle.HandleInfo'></div>
  </div>
</template>

<script>
export default {
  props: {
    eventHandle:Object
  },
  data () {
    return {
      imgs: [],
      host: '//47.94.7.154:3000'
    }
  },
  mounted () {
    this.initEventImgs(this.eventHandle.Imgs)
  },
  methods: {
    initEventImgs (imgs) {
      this.imgs = imgs.split(',')
      if (this.imgs[this.imgs.length - 1] == '') this.imgs = this.imgs.slice(0, this.imgs.length - 1)
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  text-align: center;
  .imgs{
    margin: 10px;
    img{
      width: 148px;
      height: 148px;
      border-radius: 5px;
      border: solid 1px #c0ccda;
    }
  }
}
</style>