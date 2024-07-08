<script setup>
import Form from "../components/Form.vue"
</script>

<template>
  <n-flex class="container">
    <n-space size="large" class="formlist">
      <Form />
    </n-space>
  </n-flex>
</template>

<style scoped>
.container {
  position: flex;
  width: 100%;
  height: 100vh;
  justify-content: center;
}
.formlist {
  background-color: aliceblue;
  width: 100%;
  align-items: center;
  padding-left: 35%;
}
</style>