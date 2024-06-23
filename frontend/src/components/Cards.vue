<script setup>
const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <n-card class="card" title="带封面的卡片">
    <template #cover>
      <img src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" style="height:100px; border-radius: 1.5em" />
    </template>
    卡片内容
  </n-card>
</template>

<style scoped>
.card:hover{
  box-shadow: rgb(0 0 0 / 30%) -20px 20px 30px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid #cee4e5;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #cee4e5;
}
.card{
  box-shadow: rgb(0 0 0 / 30%) 0px 20px 30px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid #cee4e5;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #cee4e5;
}
</style>