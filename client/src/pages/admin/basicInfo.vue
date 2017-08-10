<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      基础信息维护
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
  border: solid 1px #4c87b9;
  width: 100%;
  background: #fff;
  .title{
    color: #fff;
    font-size: 20px;
    background: #4c87b9;
    padding: 10px;
    display: flex;
    align-items: center;
    img{
      width: 20px;
      margin-right: 10px;
    }
  }
  .content{
    padding: 10px;
    .tabs{
      color: #5b9bd1;
      font-size: 15px;
      display: flex;
      align-items: center;
      cursor: default;
      .tab{
        padding: 10px;
        background: #fff;
      }
      .active{
        border-top: solid 2px red;
        border-left: solid 1px #ddd;
        border-right: solid 1px #ddd;
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
  export default {
    components: { InfosStreet, InfosCommunity, InfosCountry},
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