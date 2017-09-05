<template>
  <div :class="s.wrap">
    <router-view></router-view>
  </div>
</template>

<style lang="less" module="s">
.wrap{
  height: 100%;
}
</style>