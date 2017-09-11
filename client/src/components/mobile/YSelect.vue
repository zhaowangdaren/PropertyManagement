<template>
  <div :class='s.wrap'>
    <div
      :class='[s.selected, {[s.active]: showOptions}]'
      @click='showOptions = !showOptions' 
      v-html='selected.label'></div>
    <div 
      :class='s.options'
      v-if='showOptions'>
      <div
        v-for='option in options'
        :class='s.option'
        @click='onOption(option)'
         v-html='option.label'></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    selected: {
      type: Object,
      default: function () {
        return {
          value: '',
          label: 'asdfasdf'
        }
      }
    },
    options: Array // [{value: '', label: ''}]
  },
  data () {
    return {
      showOptions: false
    }
  },
  methods: {
    onOption (option) {
      this.selected = option
      this.showOptions = false
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  position: relative;
}
.options{
  position: absolute;
  left: 0;
  // top: 0;
  background-color: #fff;
  overflow: scroll;
  margin-top: 2px;
}
.option{
  box-shadow: 0 -1px 1px 0 rgba(0,0,0,.2) inset;
  padding: 10px;
}
.selected{
  height: 100%;
  background-color: #fff;
  margin: auto 0;
  display: flex;
  align-items: center;
  border: solid 1px transparent;
}
.active{
  border: solid 1px #20a0ff;
}
</style>