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
        title: '房屋工程',
        rightBtns: []
      },
      list: [
        {
          label: '家具装饰公司',
          value: 0
        },
        {
          label: '家具装饰公司',
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