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
          label: '家具装饰',
          value: 0,
          event: this.onFurniture
        },
        {
          label: '房屋维修',
          value: 1,
          event: this.onFix
        }
      ]
    }
  },
  methods: {
    onItem (item) {
      console.info('onItem', item)
      item.event()
    },
    onFurniture () {
      this.$router.push({name: 'CSCFurniture'})
    },
    onFix () {
      this.$router.push({name: 'CSCFix'})
    }
  }
}
</script>

<style lang="less" module="s">
.listMenu{
  margin-top: 88px;
}
</style>