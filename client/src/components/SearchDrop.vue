<template>
  <div :class='s.inputWrap'>
    <input type="text" v-model='inputName'>
    <div :class='s.dropList'>
      <div v-for='item in showDatas' v-text='item.text' :class='s.streetName' @click='inputName = streetName'></div>
    </div>
  </div>
</template>

<style lang="less" module='s'>
.inputWrap{
  border: solid 1px #ddd;
  position: relative;
  font-size: 18px;
  flex: 0.5;
  margin-left: 10px;
  input{
    width: 95%;
    border: solid 0px transparent;
    padding: 5px;
  }
  &:hover .dropList{
    display: block;
  }
  .dropList{
    position: absolute;
    width: 100%;
    background: #fff;
    box-shadow: 1px 1px 1px 1px #ddd;
    display: none;
    .streetName{
      padding: 5px;
      &:hover{
        background: #ddd;
      }
    }
  }
}
</style>