<script setup>
const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <n-card class="card" title="JetBrains">
    <n-avatar
        :size="48"
        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
    />
  </n-card>
</template>

<style scoped>
.card:hover{
  box-shadow: rgb(255, 255, 255) -10px 20px 30px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid #9a49e5;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #9a49e5;
}
.card{
  box-shadow: rgb(255, 255, 255) 0px 10px 20px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid #9a49e5;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #9a49e5;
}
</style>