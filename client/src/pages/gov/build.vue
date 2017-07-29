<template>
  <div>
    <builds></builds>
  </div>
</template>

<script type="text/javascript">
import Builds from '@/components/table/Builds'
export default {
  components: { Builds },
  props: {
    editable: Boolean
  }
}
</script>

<style lang="less" module='s'>
.wrap{

}
</style>