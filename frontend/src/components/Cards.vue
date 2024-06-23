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
  <head>
    {{ item }}
  </head>
</n-card>
</template>

<style scoped>
.card:hover{
  box-shadow: rgb(0 0 0 / 30%) -20px 60px 30px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid lightblue;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  width: 100%;
  height: 100%;
  background-color: lightblue;
}
.card{
  box-shadow: rgb(0 0 0 / 30%) 0px 20px 30px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid lightblue;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  width: 100%;
  height: 100%;
  background-color: lightblue;
}
</style>