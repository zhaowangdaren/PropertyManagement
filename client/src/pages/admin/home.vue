<template>
  <div>
    <h1 :class='s.title'>数据总览</h1>
    <div :class='s.datas'>
      <div :class='s.item'>
        <div :class='s.spec'><span >999</span>条</div>
        <div :class='s.text'>今日新增投诉</div>
      </div>
      <div :class='s.item'>
        <div :class='s.spec'><span>999</span>个</div>
        <div :class='s.text'>今日被投诉小区数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.spec'><span>999</span>条</div>
        <div :class='s.text'>今日新增事件处理</div>
      </div>
    </div>
    <div :class='s.datas'>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>街道用户数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>区政府用户数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>物业信息</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>建筑信息</div>
      </div>
    </div>
    <div :class='s.datas'>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>投诉记录数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>社区总数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>街道总数</div>
      </div>
      <div :class='s.item'>
        <div :class='s.value'><span>999</span>条</div>
        <div :class='s.text'>小区总数</div>
      </div>
    </div>
  </div>
</template>

<style lang="less" module='s'>
.title{
  font-size: 18px;
  color: #555;
}
.datas{
  display: flex;
  background-color: #fff;
  padding: 5px;
}
.item{
  flex: 1;
  margin: 5px;
  border-radius: 5px;
  padding: 20px 20px;
  font-size: 18px;
  text-align: center;
  background-color: #eee;
  color: #666;
  .value{
    span{
      color: #212138;
      font-size: 25px;
    }
    font-weight: bold;
  }
  .spec{
    span{
      color: #e82d2d;
      font-size: 25px;
    }
    font-weight: bold;
  }
}


</style>