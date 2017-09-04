<!-- 工程服务列表 -->
<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <list-menu 
      :class="s.listMenu"
      :list='list'
      @item='onItem'>
    </list-menu>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
import ListMenu from '@/components/mobile/ListMenu'
export default {
  components: { ListMenu, ActionBar },
  data () {
    return {
      headerOptions: {
        leftBtns: [],
        title: '房屋维修',
        rightBtns: []
      },
      list: [
        {
          label: '房屋维修公司',
          value: 0
        },
        {
          label: '房屋维修公司',
          value: 1
        }
      ]
    }
  },
  methods: {
    onItem (item) {
      console.info('onItem', item)
      this.$router.push({name: 'CSCompany'})
    } 
  }
}
</script>

<style lang="less" module="s">
.listMenu{
  margin-top: 88px;
}
</style>