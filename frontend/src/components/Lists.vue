<script setup>
import { ref, onMounted } from 'vue';
import { NGrid, NGi } from 'naive-ui';
import Cards from "@/components/Cards.vue";

const items = ref([1, 2, 3, 4, 5, 6, 7, 8, 9]);  // 示例数组，可以根据需要修改
const dealt = ref(false);

const dealCards = () => {
  dealt.value = false; // Reset the dealt state
  setTimeout(() => {
    dealt.value = true; // Trigger the animation
  }, 100); // Small delay to ensure reset
};

const getCardStyle = (index) => {
  const delay = (index + 1) * 0.1
  return dealt.value ? { transitionDelay: `${delay}s` }:{}
};

onMounted(() => {
  dealCards();
});

</script>

<template>
  <n-grid :x-gap="20" :y-gap="20" cols="4" style="width: 60%">
    <n-gi v-for="(item, index) in items" :key="index"
          :class="['grid-item', {'dealt': dealt}]" :style="getCardStyle(index)">
      <Cards :item="item" />
    </n-gi>
  </n-grid>
</template>

<style scoped>
.grid-item {
  width: 150%;
  height: 200px;
  transform: perspective(300px) rotateX(180deg);
  opacity: 0;
  transition: transform 0.8s ease, opacity 100s ease;
}
.grid-item.dealt {
  transForm: translateY(0);
  opacity: 1;
}

.card-content {
  background-color: lightblue;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 5%;
  width: 100%;
  height: 100%;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}
</style>