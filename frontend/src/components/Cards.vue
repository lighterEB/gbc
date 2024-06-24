<script setup>
import {ref} from 'vue'

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})
const title = ref(props.item["name"].split('.')[0]);
const pic = ref(props.item["pic"]);
const isHovered = ref(false);
const text = ref("************************************")
console.log(pic.value)
console.log(isHovered.value)
</script>

<template>
  <n-card class="card" :title="title">
    <n-avatar :size="64" style="background-color: #9a49e5;">
      <div v-html="pic"></div>
    </n-avatar>
    <!-- :size="48"
    src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
/> -->
    <div @mouseenter="isHovered = true" @mouseleave="isHovered = false" style="border: 1em lightgreen">
      <n-text class="text-area" v-if="!isHovered">
        {{ text }}
      </n-text>
      <n-button class="btn-copy" text color="#8a2be2" v-if="isHovered">
        复制
      </n-button>
    </div>
  </n-card>
</template>

<style scoped>
.card:hover {
  box-shadow: rgb(255, 255, 255) -10px 20px 30px -12px, rgb(255, 255, 255) 20px 20px 20px -120px;
  border-radius: 1.5rem;
  border: 0.2rem solid #af80e7;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #9a49e5;
}

.card {
  box-shadow: rgb(255, 255, 255) 0px 10px 20px -12px, rgb(255, 255, 255) 0px 10px 20px -12px;
  border-radius: 1.5rem;
  border: 0.2rem solid #9a49e5;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) 0ms;
  height: 100%;
  background-color: #9a49e5;
}

.btn-copy {
  border: 0.2rem solid #9a49e5;
  border-radius: 1.5rem;
  width: 100%;
  height: 50%;
  padding: 10px;
  opacity: 1;
  color: black;
}

.btn-copy:hover {
  border: 0.2rem solid #661dab;
  border-radius: 1.5rem;
  width: 100%;
  height: 50%;
  color: black;
  opacity: 1;
}

.text-area {
  z-index: 3;
  width: 20%;
  height: 10%;
  position: relative;
  padding-left: 20px;
}
</style>