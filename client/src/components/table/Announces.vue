<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        政府公告<!-- 推送法院调解请求 -->
      </div>
      <div :class='s.body'>
        <!-- search result -->
        <div :class='s.uploadWrap'></div>
        <div>已发布公告列表</div>
        <div v-for='announce in announces'><a href="">{{announce.fileName}}</a></div>
      </div>
    </div>
  </div>
</template>

<script>
  import fetchpm from '@/fetchpm'
  export default {
    filters: {},
    data () {
      return {
        host: "https://www.maszfglzx.com:3000",
        announces: [],
        query: {
          fileName: '',
          pageNo: 0,
          pageSize: 10
        }
      }
    },
    mounted () {
      this.fetchAnnounces(this.query)
    },
    methods: {
      fetchAnnounces(query) {
        fetchpm(this, true, '/pm/announce/search/name', {
          method: 'POST',
          body: query
        }).then(resp => {
          return resp.json()
        }).then(body => {
          console.info('fetchAnnounces', body)
        })
      }
    }
  }
</script>

<style lang="less" module='s'>
.wrap{
  // width: 100%;
  margin: 10px;
  .content{
    border: solid 1px #4c87b9;
    background-color: #fff;
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
    .body{
      margin: 10px;
    }    
  }
}
.uploadWrap{
  
}
</style>