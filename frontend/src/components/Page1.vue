<template>
  <n-space>
    <n-card>
      <h2>Page 1</h2>
      <p>This is the first page.</p>
      <n-button @click="$emit('update:currentPage', 'page2')">Go to Page 2</n-button>
    </n-card>
  </n-space>
</template>

<script>
export default {
  name: 'Page1',
  props: {
    currentPage: String
  }
};
</script>