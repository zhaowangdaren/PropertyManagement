<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      用户管理<!--   用户管理 -->
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
  background: #fff;
  .title{
    color: #fff;
    font-size: 20px;
    background: #4c87b9;
    padding: 10px;
    display: flex;
    align-items: center;
    color: #fff;
    font-size: 30px;
    font-family: "华文行楷";
    background: #4c87b9;
    background: -webkit-gradient(linear,center top, center bottom,from(#ff0000), to(#000000));
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
        border: 1px solid green;
        padding: 10px;
        background: #fff;
        font-size: 17px;
        background-color: #6D6E8C;
        background: -webkit-gradient(linear,left top,  right bottom,from(rgb(255,255,255)), to(rgb(166, 202, 240)));
        color: black;
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
import InfosGov from '@/components/InfosGov'
import InfosCourt from '@/components/InfosCourt'
import InfosStreetUser from '@/components/InfosStreetUser'
  export default {
    components: { InfosGov, InfosStreetUser, InfosCourt},
    data () {
      return {
        tabs: [
          {
            name: '区政府人员',//区政府人员
            view: InfosGov
          },
          {
            name: '街道用户',//街道用户
            view: InfosStreetUser
          },
          {
            name: '法官',//街道用户
            view: InfosCourt
          }
        ],
        curView: InfosGov
      } 
    },
    methods: {
      onTab (tab) {
        this.curView = tab.view
      }
    }
  }
</script>