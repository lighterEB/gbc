<template>
  <div class="toolbar">
    <n-tooltip trigger="hover">
      <template #trigger>
        <n-button @click="minimizeWindow">最小化</n-button>
      </template>
      <span>Minimize</span>
    </n-tooltip>
    <n-tooltip trigger="hover">
      <template #trigger>
        <n-button @click="maximizeWindow">最大化</n-button>
      </template>
      <span>Maximize</span>
    </n-tooltip>
    <n-tooltip trigger="hover">
      <template #trigger>
        <n-button @click="closeWindow">关闭</n-button>
      </template>
      <span>Close</span>
    </n-tooltip>
  </div>
</template>

<script>
export default {
  name: 'Toolbar',
  methods: {
    minimizeWindow() {
      window.backend.Window.Minimize();
    },
    maximizeWindow() {
      window.backend.Window.Maximize();
    },
    closeWindow() {
      window.backend.Window.Close();
    }
  }
};
</script>

<style scoped>
.toolbar {
  display: flex;
  justify-content: flex-end;
  background-color: #f5f5f5;
  padding: 10px;
}
</style>
