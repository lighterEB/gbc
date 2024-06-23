<template>
  <n-space>
    <n-card>
      <h2>Page 2</h2>
      <p>This is the second page.</p>
      <n-button @click="$emit('update:currentPage', 'page1')">Go to Page 1</n-button>
    </n-card>
  </n-space>
</template>

<script>
export default {
  name: 'Page2',
  props: {
    currentPage: String
  }
};
</script>
