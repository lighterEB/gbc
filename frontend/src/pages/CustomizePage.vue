<script setup>
</script>

<template>
<n-flex class="container">
  <n-button @click="queryData">
    测试
  </n-button>
  <n-text style="text-align: center; width: 100px; height: 100px; background-color: #f1f2e5">
    {{ productsList }}
  </n-text>
  <n-button>Oops!</n-button>
  <n-button>Oops!</n-button>
  <n-button>Oops!</n-button>
  <n-button>Long! Long! Cross the line!</n-button>

</n-flex>
</template>

<style scoped>
.container {
  height: 100%;
  width: 100%;
  position: relative;
  background-color: #f1f2e5;
}
</style>