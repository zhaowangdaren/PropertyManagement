<template>
  <div :class='s.wrap' :style='{backgroundColor: bgColor, color: color, fontSize: fontSize}' @click='onClick()'>
    <img :src="img" v-if='img !== ""'>
    {{text}}
    <div :class='s.cover'></div>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 8px 10px;
    img{
      width: 20px;
      margin-right: 5px;
    }
    &:hover .cover{
      display: block;
    }
    .cover{
      display: none;
      position: absolute;
      top: 0;
      right: 0;
      width: 100%;
      height: 100%;
      background: rgba(0,0,0,0.1);
    }
  }
</style>

<script>
  
  export default {
    props: {
      text: String,
      img: {
        type: String,
        default: ''
      },
      bgColor: {
        type: String,
        default: ''
      },
      color: {//字体颜色
        type: String,
        default: '#fff'
      },
      fontSize: {
        type: String,
        default: '15px'
      },
      clickMethod: Function
    },
    methods: {
      onClick () {
        if (!this.clickMethod) return
        this.clickMethod()
      }
    }
  }
</script>