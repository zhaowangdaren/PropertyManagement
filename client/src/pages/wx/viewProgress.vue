<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.content'>
      <div :class='s.item'>
        Evnet Status
        <select v-model='selectedStatus'>
          <option v-for='status in eventStatus' 
            :value='status.value'
            v-text='status.label'></option>
        </select>
      </div>
      <div :class='s.item'>
        Time
        <select v-model='selectedTime'>
          <option v-for='time in times' 
            :value='time.value'
            v-text='time.label'></option>
        </select>
      </div>
    </div>
    
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: { ActionBar },
  data () {
    return {
      headerOptions: {
        title: '查看进度'
      },
      events: [],
      host:'http://localhost:3000',
      selectedStatus: 0,
      eventStatus:[{value: 0, label:"全部"}, {value: 4, label: '已解决'}, {value: 1, label:'未解决'}],
      selectedTime: 0,
      times:[{value: 0, label:"近一个月内"}, {value: 1, label: '近三个月内'}, {value: 2, label:'近一年内'}, {value: 3, label: '所有'}]
    }
  }
}
</script>

<style lang="less" module='s'>
.content{
  margin-top: 80px;
  font-size: 25px;
  .item{
    padding: 10px;
    display: flex;
    align-items: center;
    select{
      flex: 1;
      padding: 10px;
      border-radius: 5px;
      margin: 10px;
    }
  }
}

</style>