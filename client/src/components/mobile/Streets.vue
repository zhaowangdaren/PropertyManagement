<!-- onSelect 告诉父组件选择 -->
<template>
  <div>
    <div :class='s.item' v-for='street in streets' @click='onStreet(street)'>
      <div v-text='street.Name' :class='s.name'></div>
      <i class='iconfont icon-next'></i>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  data () {
    return {
      streets: []
    }
  },
  mounted () {
    this.fetchStreets()
  },
  methods: {
    fetchStreets () {
      fetchpm(this, false, '/open/street',{
        method: 'POST'
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.streets = data.data.streets || []
        }
      })
    },
    onStreet (street) {
      this.$emit('onSelect', street)
    }
  }
}
</script>

<style lang="less" module='s'>
.item {
  color: #20a0ff;
  background-color: #fff;
  box-shadow: 0 1px 1px 1px #ddd;
  padding: 20px 10px;
  margin-top: 2px;
  display: flex;
  justify-content: center;
  align-items: center;
  .name{
    flex: 1;
  }
}
</style>