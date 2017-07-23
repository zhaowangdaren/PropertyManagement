<template>
  <div :class='s.inputWrap'>
    <input type="text" @input='updateValue($event.target.value)' :value='inputValue'>
    <div :class='s.dropList'>
      <div :class='s.streetName + " " + s.all' @click='fetchAll'>全部</div>
      <div v-for='value in values' v-text='value' :class='s.streetName' @click='updateValue(value)'></div>
    </div>
  </div>
</template>

<style lang="less" module='s'>
  .inputWrap{
    border: solid 1px #ddd;
    position: relative;
    font-size: 18px;
    flex: 1;
    margin-left: 10px;
    // width: 100%;
    input{
      // width: 95%;
      border: solid 0px transparent;
      padding: 5px;
      height: 30px;
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
      z-index: 2;
      .streetName{
        padding: 5px;
        color: #555;
        &:hover{
          background: #0593f5;
        }
      }
      .all{
        background: #ddd;
      }
    }
  }
</style>

<script>
export default {
  props: {
    values: Array
  },
  data () {
    return {
      inputValue: ''
    }
  },
  methods: {
    updateValue (value) {
      this.inputValue = value
      this.$emit('input', value)
    },
    fetchAll () {
      this.$emit('all')
    }
  }
}
</script>