<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      基础信息维护<!-- 基础信息维护 -->
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

<style lang="less" module='s'>
.wrap{
  margin: 10px;
  .title{
    color: #fff;
    font-size: 25px;
    font-family: "微软雅黑";
    font-size: 30px;
    font-family: "华文行楷";
    background: #4c87b9;
    background: -webkit-gradient(linear,center top, center bottom,from(#ff0000), to(#000000));
    padding: 10px;
    display: flex;
    align-items: center;
    box-shadow: 1px 0px 4px -2px rgba(0, 0, 0, 0.5),-1px 0 4px -2px rgba(0, 0, 0, 0.5), 0 -1px 4px -2px rgba(0, 0, 0, 0.5); 
    img{
      width: 20px;
      margin-right: 10px;
    }
  }
  .content{
    padding: 10px;
    background: #fff;
    border: solid 1px #4c87b9;
    box-shadow: 1px 0px 4px -2px rgba(0, 0, 0, 0.5),-1px 0 4px -2px rgba(0, 0, 0, 0.5), 0 1px 4px -2px rgba(0, 0, 0, 0.5);
    .tabs{
      color: #5b9bd1;
      font-size: 15px;
      display: flex;
      align-items: center;
      cursor: default;
      .tab{
        border: 1px solid green;
        padding: 10px;
        background: #fff;
        font-size: 19px;
        background-color: #6D6E8C;
        background: -webkit-gradient(linear,left top,  right bottom,from(rgb(255,255,255)), to(rgb(166, 202, 240)));
        color: black;
        cursor: pointer;
      }
      .active{
        border-top: solid 2px red;
        border-left: solid 1px #ddd;
        border-right: solid 1px #ddd;
        background: #cf6;
      }
    }
    .view{
      border: solid 1px #ddd;
      margin-top: -1px;
      background: #fff;
      height: 100%;
    }
  }
}
</style>

<script type="text/javascript">
import InfosStreet from '@/components/InfosStreet'
import InfosCommunity from '@/components/InfosCommunity'
import InfosCountry from '@/components/InfosCountry'
import EventTypes from '@/components/table/EventTypes'
  export default {
    components: { InfosStreet, InfosCommunity, InfosCountry, EventTypes},
    data () {
      return {
        tabs: [
          {
            name: '街道信息维护',//街道信息维护
            view: InfosStreet
          },
          {
            name: '社区信息维护',//社区信息维护
            view: InfosCommunity
          },
          {
            name: '小区信息维护',//社区信息维护
            view: InfosCountry
          },
          {
            name: '设置事件类型',
            view: EventTypes
          }
        ],
        curView: InfosStreet
      } 
    },
    methods: {
      onTab (tab) {
        this.curView = tab.view
      }
    }
  }
</script>