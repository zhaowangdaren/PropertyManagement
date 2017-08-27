<template>
  <div :class='s.wrap'>
    <div :class='s.start' @click='onRules'>物业条例</div>
    <div :class='s.start' @click='onManageMethod'>管理办法</div>
    <div :class='s.start' @click='onQA'>答疑解惑</div>
    <div :class='s.start' @click='onKPI'>物业考核标准</div>

  </div>
</template>

<script>
export default {
  methods: {
    onRules () {
      this.$router.push({path:'/wx/pm/rules'})
    },
    onManageMethod () {
      this.$router.push({path:'/wx/pm/rules/mp'})
    },
    onQA () {
      this.$router.push({path:'/wx/pm/rules/qa'})
    },
    onKPI () {
      this.$router.push({path:'/wx/pm/rules/kpi'})
    }
  }
}
</script>
<style lang="less" module='s'>
.wrap{
  background-color: #20a0ff;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .start{
    border:solid 2px #fff;
    border-radius: 5px;
    padding: 20px;
    margin: 5px;
    color: #fff;
    font-size: 30px;
    font-weight: bold;
    text-align: center;
    width: 60%;
    &:active {
      background-color: rgba(254,254,254, 0.1);
    }
  }
}
</style>