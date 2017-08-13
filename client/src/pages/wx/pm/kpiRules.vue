<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <div :class='s.line'></div>
      花山区物业管理信息平台评分标准
      <div :class='s.line'></div>
    </div>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;一、每个住宅小区每季度设定基础分100分。</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;二、小区内的投诉处理情况，对应考核标准，对承办单位（小区的物业公司或托管社区）进行考核。</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;三、为督促各承办单位做好微信平台在线回复工作，特开设红黄督办栏目。超过3天未回复的帖子用黄榜提醒，超过7天的用红榜警示。上黄榜的小区，一次扣0.5分；上红榜的小区，一次扣1分。</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;四、重点投诉问题处理</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;出现重点投诉问题范围内的一次扣3分</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;①同一投诉人就同一问题投诉三次以上的；</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;②同一问题被不同投诉人投诉三次以上的；</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;③投诉群体数量较大的问题；</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;④其他影响较大的问题。</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;属于以上范围的问题，由区住房管理中心物业管理科约谈问题承办单位，承办单位作出情况说明或整改。三天之内能够给予投诉人答复或问题整改完成的，不扣分；对涉及面广、办理难度较大，在规定时限内难以办结的，承办单位应在规定的时间内向所在街道办事处说明情况，并书面报区住房管理中心物业管理科同意后，方可延期办理，最长延期时间不得超过30天，能够给予投诉人答复或问题整改完成的，不扣分；延时到期后仍未能给投诉人答复的或问题未能整改完成的，给承办单位一次扣5分。</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;五、每季度各小区得分占区住房管理中心每季度物业考核评分的50％</p>
  </div>
</template>

<style lang="less" module='s'>
.wrap{
  font-size: 22px;
  color: #555;
  padding: 10px;
  .title{
    display: flex;
    align-items: center;
    font-size: 28px;
    margin: 10px;
    color: #000;
    .line{
      flex: 1;
      margin: 5px;
      border-top: solid 1px #aaa;
    }
  }
  p{
    margin: 5px auto;
  }
}
</style>