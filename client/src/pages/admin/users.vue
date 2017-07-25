<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      Home<!-- 基础信息维护 -->
    </div>
    <div :class='s.content'>
      <div :class='s.tabs'>
        <div :class='s.tab + " " + (tab.view === curView ? s.active : "")' v-for='tab in tabs' @click='onTab(tab)'>
          {{tab.name}}
        </div>
      </div>
      <component :is='curView' :class='s.view'/>
    </div>
  </div>
</template>