<script setup>
import { useMessage, NAlert } from 'naive-ui';
import { ref, h, reactive } from 'vue'
import { KeyGenRequest } from '../../wailsjs/go/main/App'
import { clipboard } from 'clipboard'
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

const params = reactive({
  "licenseId": "DADDYHU726",
  "licenseeName": "lighterEB",
  "assigneeName": "daddyhu",
  "assigneeEmail": "hellohus726@126.com",
  "products": [{ "code": "PCWMP", "fallbackDate": "2099-12-31", "paidUpTo": "2099-12-31", "extended": "true" }, { "code": "PSI", "fallbackDate": "2099-12-31", "paidUpTo": "2099-12-31", "extended": "true" }]
})
const produts = reactive(
  { "code": "", "fallbackDate": "2099-12-31", "paidUpTo": "2099-12-31", "extended": "false" }
)
const rendMessage = (props) => {
  const { type } = props;
  return h(
    NAlert,
    {
      closable: props.closable,
      onclose: props.onclose,
      type: type === "reloading" ? "default" : type,
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

async function copyCode() {
  switch (props.item["name"]) {
    case "CLion":
      produts.code = "CL"
      break
    case "DataGrip":
      produts.code = "DG"
      break
    case "GoLand":
      produts.code = "GO"
      break
    case "IDEA":
      produts.code = "II"
      break
    case "PyCharm":
      produts.code = "PC"
      break
    case "Rider":
      produts.code = "RD"
      break
    case "RustRover":
      produts.code = "RR"
      break
    case "WebStorm":
      produts.code = "WS"
  }
  params.products.push(produts)
  console.log(params.products)


  await KeyGenRequest(JSON.stringify(params)).then(res => {
    if (res.code == 200) {
      navigator.clipboard.writeText(res.data)
      msgInfo.success("你已经获取激活码，现在可以直接粘贴啦！", {
        render: rendMessage,
        closable: true,
        duration: 1000
      });
    } else {
      msgInfo.error("糟糕~激活码丢了~~！", {
        render: rendMessage,
        closable: true,
        duration: 1000
      });
    }
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
    <div @mouseenter="isHovered = true" @mouseleave="isHovered = false"
      style="border: 1em lightgreen; width: 100%; height: 50%; bottom: 20px;">
      <n-text class="text-area" v-if="!isHovered">
        {{ text }}
      </n-text>
      <n-button class="btn-copy" text @click="copyCode()" color="#8a2be2" v-if="isHovered">
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