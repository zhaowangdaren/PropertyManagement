<template>
  <div :class='s.wrap'>
    <div :class="s.up">
      <i class="iconfont icon-flower"></i>
    </div>
    <div :class="s.down">
      <div :class='s.item'>评分:<i v-for="star in startNum" :class="'iconfont icon-pingjia'" :style="{color: star <= (score - 1 ) ? '#f7ba2a' : '#aaa'}" @click="score = star + 1"></i></div>
      <div :class='s.item'>
        服务意见
      </div>
      <div :class='s.item'>
        <textarea :class='s.textarea' v-model='content'></textarea>
      </div>
      <div :class='s.item'>
        <div :class='s.btn' @click='onSubmit'>提交</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      score: 0, //1~5
      startNum: 5,
      content: ''
    }
  },
  methods: {
    onSubmit () {
      if (this.content == '' && this.score == 0) return
      var assess = {
        score: this.score, 
        content: this.content
      }
      this.$emit('submit', assess)
    }
  }
}
</script>

<style lang="less" module="s">
.wrap{
  overflow: auto;
}
.up{
  // background-color: rgb(25, 163, 24);
  // background-image: url("../../res/images/upbg5.jpg");
  background-color: #9266f9;
  background-size: cover;
  height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
  i{
    color: #fff;
    font-size: 50px;
  }
}
.item{
  font-size: 25px;
  padding: 10px;
  display: flex;
  align-items: center;
  i{
    font-size: 50px;
    flex: 1;
    margin-left: 10px;
  }
}
.textarea{
  width: 100%;
  height: 200px;
  // border: none;
  border: solid 1px #999;
  resize: none;
}
.btn{
  background-color: rgb(25, 163, 24);
  color: #fff;
  padding: 20px;
  text-align: center;
  margin: 20px;
  font-size: 25px;
  border-radius: 5px;
  flex: 1;
}
</style>