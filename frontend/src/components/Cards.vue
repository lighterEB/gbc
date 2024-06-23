<script setup>
const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})
</script>

<template>
<n-card class="card">
  {{ item }}
</n-card>
</template>

<style scoped>
.card{
  box-shadow: rgb(0 0 0 / 30%) 0 8px 40px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid lightblue;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  width: 100%;
  height: 100%;
  position: relative;
  background-color: lightblue;
}
</style>