<template>
  <div :class='s.wrap'>
    <pmkpis></pmkpis>
  </div>
</template>

<script>
import PMKPIs from '@/components/table/PMKPIs'
  export default {
    components: { 'pmkpis': PMKPIs},
    data () {
      return {
        
      }
    },
    mounted () {
    },
    methods: {
    }
  }
</script>
<style lang="less" module='s'>
.wrap{
  margin: 10px;
  // width: 100%;
  .street{
    background-color: #ddd;
    width: 100%;
    font-size: 25px;
    padding: 5px;
  }
  .content{
    border: solid 1px #4c87b9;
    margin-top: 50px;
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
    .body{
      margin: 10px;
      .searchWrap{
        display: flex;
        justify-content: flex-end;
        margin: 10px;
      }
      table{
        width: 100%;
        margin-top: 10px;
        font-size: 15px;
        color: #555;
        background-color: #fff;
        th, td {
          padding: 5px;
          border: solid 1px #ddd;
          text-align: center;
          font-size: 15px;
        }
        th{
          background-color: #ddd;
        } 
        tr{
          &:hover {
            // background-color: #ddd;
          }
        }
        .key{
          background-color: #ddd;
        }
      }
    }    
  }
}
</style>