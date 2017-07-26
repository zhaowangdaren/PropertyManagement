<template>
  <div :class='s.inputWrap'>
    <input type="text" @input='updateValue($event.target.value)' :value='inputValue'>
    <div :class='s.dropList'>
      <div :class='s.streetName + " " + s.all' @click='fetchAll'>全部</div>
      <div v-for='value in showValues' v-text='value' :class='s.streetName' @click='updateValue(value)'></div>
    </div>
  </div>
</template>

<style lang="less" module='s'>
  .inputWrap{
    position: relative;
    font-size: 18px;
    flex: 1;
    margin-left: 10px;
    // width: 100%;
    input{
      width: 100%;
      border: solid 1px #ddd;
      // padding-left: 5px;
      height: 30px;
    }
    &:hover .dropList{
      display: block;
    }
    .dropList{
      position: absolute;
      width: 100%;
      background: #fff;
      box-shadow: 1px 1px 2px 2px #ddd;
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
  watch: {
    values: function function_name(val) {
      if (val.length == 0) {
        this.updateValue('')
      }
    }
  },
  computed: {
    showValues: function () {
      return this.values.filter(value => {
        var index = value.indexOf(this.inputValue) > -1
        return index
      })
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