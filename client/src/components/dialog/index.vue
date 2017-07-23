<template>
  <div :class='s.wrap'>
    <div :class='s.contentWrap'>
      <div :class='s.titleWrap'>
        <div :class='s.title' v-text='title'></div>
        <img src="~@/res/images/close.png" @click='onClose()'>
      </div>
      <div v-if='warn !== ""' v-text='warn' :class='s.warn'></div>
      <slot></slot>
      <div :class='s.btnWrap'>
        <image-button :class='s.bt' :clickMethod='onClickS'
        text='保存'
        bgColor='#26a69a'
        fontSize='20px'
        />
        <image-button :class='s.bt' :clickMethod='onClickC'
        text='取消'
        bgColor='#ddd'
        fontSize='20px'
        color='#000'
        />
      </div>
    </div>
  </div>
</template>

<style lang="less" module='s'>
  .wrap{
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0,0,0,0.5);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    .contentWrap{
      background: #fff;
      min-width: 40%;
      .titleWrap{
        background: #fff;
        padding: 10px;
        display: flex;
        align-items: center;
        border-bottom: solid 1px #ddd;
        .title{
          flex: 1;
          font-size: 19px;
          font-weight: bold;
          margin-right: 50px;
        }
        img{
          width: 20px;
        }
      }
      .warn{
        color: red;
        text-align: center;
        font-size: 12px;
        padding: 5px;
      }
    }
    .btnWrap{
      display: flex;
      justify-content: flex-end;
      align-items: center;
      padding: 20px;
      .bt{
        margin-left: 10px;
      }
    }
    
  }
</style>

<script>
  import ImageButton from '@/components/ImageButton'
  export default {
    components: { ImageButton},
    props: {
      title: String,
      onSave: Function,
      onCancel: Function,
      warn: String
    },
    methods: {
      onClickS (){
        if (!this.onSave) return
        this.onSave()
      },
      onClickC () {
        if (!this.onCancel) return
        this.onCancel()
      },
      onClose () {
        this.onCancel()
      }
    }
  }
</script>
