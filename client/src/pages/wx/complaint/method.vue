<template>
  <div :class='s.wrap'>
    <div :class='s.header'>投诉流程</div>
    <div :class='s.item'>
      <div :class='s.tip'>1. 进入【马鞍山市花山区物业管理平台】微信公众号，点击【物业管理】->【投诉与建议】。</div>
      <img src="~@/res/images/cp1.png">
    </div>
    <div :class='s.item'>
      <div :class='s.tip'>2. 进入投诉页面，点击【新的投诉】。</div>
      <img src="~@/res/images/cp2.png">
    </div>
    <div :class='s.item'>
      <div :class='s.tip'>3. 选择街道、社区、小区。</div>
      <img src="~@/res/images/cp3.png">
      <img src="~@/res/images/cp4.png">
      <img src="~@/res/images/cp5.png">
    </div>
    <div :class='s.item'>
      <div :class='s.tip'>4. 填写投诉内容。填写完成后点击右上角的下一步。</div>
      <img src="~@/res/images/cp6.png">
    </div>
    <div :class='s.item'>
      <div :class='s.tip'>5. 上传相关图片。图片上传完成后点击下一步。</div>
      <img src="~@/res/images/cp7.png">
    </div>
    <div :class='s.item'>
      <div :class='s.tip'>6. 完成投诉。</div>
      <img src="~@/res/images/cp8.png">
    </div>
  </div>
</template>

<style lang="less" module='s'>
.wrap{
  // text-align: center;
  font-size: 22px;
  color: #555;
  margin: 10px;
}
.header{
  color: #222;
  font-size: 25px;
  text-align: center;
}
.item{
  text-align: center;
  .tip{
    text-align: left;
  }
  img{
    width: 80%;
    box-shadow: 1px 1px 2px 0 rgba(0, 0, 0, .3), -1px -1px 2px 0 rgba(0, 0, 0, .3);
  }
}
</style>