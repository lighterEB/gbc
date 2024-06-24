<script setup>
import {ref} from 'vue'

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})
const isHovered = ref(false);
const text = ref("*****************************************************")
</script>

<template>
  <n-card class="card" title="JetBrains" @mouseenter="isHovered = true" @mouseleave="isHovered = false">
    <n-avatar
        :size="48"
        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
    />
    <n-space class="text-area" v-if="!isHovered">
      {{ text }}
    </n-space>
      <n-button class="btn-copy" text color="#8a2be2" v-if="isHovered">
        复制
      </n-button>
  </n-card>
</template>

<style scoped>
.card:hover {
  box-shadow: rgb(255, 255, 255) -10px 20px 30px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid #9a49e5;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #9a49e5;
}

.card {
  box-shadow: rgb(255, 255, 255) 0px 10px 20px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid #9a49e5;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #9a49e5;
}

.btn-copy {
  border:0.2rem solid #9a49e5;
  border-radius: 1.5rem;
  width: 100%;
  height: 50%;
  padding: 10px;
  opacity: 1;
  color: black;
}

.btn-copy:hover {
  border:0.2rem solid #661dab;
  border-radius: 1.5rem;
  width: 100%;
  height: 50%;
  color: black;
  opacity: 1;
}
</style>