<template>
  <div :class="s.wrap">
    <div :class="s.row">
      <div :class="s.item" @click='onConstruct'>
        <div><i class="iconfont icon-gongcheng" style="color: #20a0ff"></i></div>
        <div>工程</div>
      </div>
      <div :class='s.item'>
        <div><i class="iconfont icon-canyin" style="color: rgb(212, 25, 41)"></i></div>
        <div>餐饮</div>
      </div>
      <div :class='s.item'>
        <div><i class="iconfont icon-repairs" style="color: rgb(251,170,83)"></i></div>
        <div>维修</div>
      </div>
    </div>
    <div :class="s.row">
      <div :class="s.item">
        <div><i class="iconfont icon-lvyoujiaotong" style="color: rgb(67,165,235)"></i></div>
        <div>旅游交通</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    onConstruct () {
      this.$router.push({path: '/wx/community/constructs'})
    }
  }
}
</script>
<style lang="less" module="s">
.wrap{
  // background-color: #20a0ff;
  background-color: #f1f3f6;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.row{
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.item{
  margin: 10px;
  width: 80px;
  height: 80px;
  background-color: #fff;
  box-shadow: 1px 1px 1px 0px rgba(0,0,0,0.2);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
  font-size: 20px;
  border-radius: 5px;
  color: #565656;
  i{
    font-size: 40px;
  }
}
</style>