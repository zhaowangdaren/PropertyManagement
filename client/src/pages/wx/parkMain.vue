<template>
  <div :class='s.wrap'>
    <div :class='s.searchWrap'>
      <div :class='s.searchIcon'><i class="iconfont icon-ic_serach"></i></div>
      <div :class='s.inputWrap'>
        <input type="text" name="" placeholder="输入小区名">
      </div>
    </div>
    <div :class='s.searchResults'>
      <div v-for='xq in xqs' :class='s.item'>{{xq.Name}}</div>
    </div>
    <div :class='s.parkIcon'>
      <i class="iconfont icon-tingchewei"></i>
    </div>
    <div :class='s.manager' @click='onManager'>我是管理员</div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      xqs: [
        {Name: 'xq1'},
        {Name: 'xq1'},
        {Name: 'xq1'},
        {Name: 'xq1'},
        {Name: 'xq1'},
        {Name: 'xq1'},
        {Name: 'xq1'}
      ]
    }
  },
  methods: {
    onCheck () {

    },
    onManager () {
      
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  background-color: #20a0ff;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.searchWrap{
  position: relative;
  display: flex;
  margin: 20px 20px 0 20px;
  background-color: #fff;
  justify-content: center;
  align-content: center;
  font-size: 23px;
  height: 50px;
  border-radius: 25px;
  .searchIcon{
    color: #999;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    border: solid 1px transparent;
    i{
      font-size: 23px;
      margin: 5px;
    }
  }
  .inputWrap{
    flex: 1;
    display: flex;
    align-items: center;
    input{
      border: none;
      width: 96%;
      box-sizing: border-box;
    }
  }
  
}
.searchResults{
  background-color: #fff;
  color: #565656;
  margin: 0 auto;
  width: 86%;
  font-size: 22px;
  .item{
    padding: 10px;
    border-bottom: solid 1px #ccc;
  }
}
.parkIcon{
  color: #fff;
  position: fixed;
  top: 50%;
  left: 0;
  right: 0;
  text-align: center;
  i{
    font-size: 100px;
  }
}
.manager{
  position: fixed;
  bottom: 10px;
  left: 0;
  right: 0;
  text-align: center;
  color: #fff;
  font-size: 22px;
  &:active {
    color: rgba(255,255,255, 0.5);
  }
}
</style>