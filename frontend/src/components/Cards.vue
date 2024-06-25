<script setup>
import { useMessage, NAlert  } from 'naive-ui';
import {ref, h} from 'vue'

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})
const title = ref(props.item["name"]);
const pic = ref(props.item["pic"]);
const isHovered = ref(false);
const text = ref("************************************")
const msgInfo = useMessage()

const rendMessage = (props) => {
  const { type } = props;
  return h(
    NAlert,
    {
      closable: props.closable,
      onclose: props.onclose,
      type: type==="reloading"?"default":type,
      title: "获取激活码",
      style: {
        opacity: "0.8",
        backgroundColor: "pink",
      },
    },
    {
      default: () => props.content
    }
  );
};

function copyCode() {
  msgInfo.success("你已经获取激活码，现在可以直接粘贴啦！", {
    render: rendMessage,
    closable: true,
    duration: 1000
  });
}

</script>

<template>
  <n-card class="card" :title="title" @mouseleave="isHovered = false">
    <n-avatar :size="64" style="background-color: #9a49e5;">
      <div v-html="pic"></div>
    </n-avatar>
    <!-- :size="48"
    src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
/> -->
    <div @mouseenter="isHovered = true" @mouseleave="isHovered = false" style="border: 1em lightgreen; width: 100%; height: 50%; bottom: 20px;">
      <n-text class="text-area" v-if="!isHovered">
        {{ text }}
      </n-text>
      <n-button class="btn-copy" text @click="copyCode" color="#8a2be2" v-if="isHovered">
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
  height: 100%;
  padding: 10px;
  opacity: 1;
  color: black;
}

.btn-copy:hover {
  border: 0.2rem solid #661dab;
  border-radius: 1.5rem;
  width: 100%;
  height: 100%;
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