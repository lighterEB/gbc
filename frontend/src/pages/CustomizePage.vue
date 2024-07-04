<script setup>
import { ref,reactive } from "vue"
import { parseTime } from "../utils/date";
const time = ref(null)
const value = ref(null)
const options= [
  {
    label: "是",
    value: "true"
  },
  {
    label: "否",
    value: "false"
  }
]
const tips = `如果要激活插件,
请在 https://plugins.jetbrains.com/搜索插件名称,
可以在插件首页看到URL中包含一串数字，
把数字替换到下面的链接'#'（注意是替换）\r\n https://plugins.jetbrains.com/api/plugins/#?filed=code ，
然后浏览器打开这个链接找到"purchaseInfo"中的"productCode",
并且选择是插件`

const listParams = reactive({
  "code": "",
  "name": "",
  "email": "",
  "extended": "",
  "expire": ""
})

const params = reactive({
  "licenseId": "",
  "licenseeName": "",
  "assigneeName": "",
  "assigneeEmail": "",
  "products": [{ "code": "PCWMP", "fallbackDate": "", "paidUpTo": "", "extended": "true" }, { "code": "PSI", "fallbackDate": "", "paidUpTo": "", "extended": "true" }]
})
const produts = reactive({ "code": "", "fallbackDate": "", "paidUpTo": "", "extended": "" })


function genKey() {
  listParams.expire = parseTime(time.value, "{y}-{m}-{d}")
  listParams.extended = value.value
  produts.code = listParams.code
  produts.fallbackDate = listParams.expire
  produts.paidUpTo = listParams.expire
  produts.extended = listParams.extended
  params.licenseeName = listParams.name
  params.assigneeName = listParams.name
  params.assigneeEmail = listParams.email
  params.products.forEach((item) => {
    item.fallbackDate = listParams.expire
    item.paidUpTo = listParams.expire
  })
  params.products.push(produts)
  console.log(params)
}
</script>

<template>
  <n-flex class="container" vertical>
    <n-list class="info-list">
      <n-list-item>
        <n-space>
          <n-thing class="n-thing">产品Code: </n-thing>
          <n-input placeholder="请输入产品Code" v-model:value="listParams.code"  style="width: 100%;"/>
          <n-tooltip width=200 placement="right-start" trigger="hover">
            <template #trigger>
              <n-tag type="info" style="border-radius: 1em; width: 20px; height: 20px; padding-left: 1px; padding-top: 1px; top: 5px;">
                <n-image src="../../assets/images/question.png" style="width: 18px; height: 18px;"/>
              </n-tag>
            </template>
              {{ tips }}
          </n-tooltip>
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space>
          <n-thing class="n-thing">是否是插件: </n-thing>
          <n-select placeholder="选择" v-model:value="value" :options="options" style="width: 100px"/>
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space>
          <n-thing class="n-thing">账号: </n-thing>
          <n-input placeholder="请输入账号" v-model:value="listParams.name" />
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space>
          <n-thing class="n-thing">邮箱: </n-thing>
          <n-input placeholder="请输入邮箱" v-model:value="listParams.email"/>
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space>
          <n-thing>过期时间: </n-thing>
          <n-date-picker v-model:value="time" size="medium" type="date" placeholder="请选择日期"/>
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space style="padding-left: 40%;">
          <n-button @click="genKey()">生成</n-button>
        </n-space>
      </n-list-item>
    </n-list>
  </n-flex>
</template>

<style scoped>
.container {
  position: relative;
  background-color: #f1f2e5;
  padding: 30px 0 50px 0;
  width: 100%;
  height: 100%;
  justify-content: center;
}

.info-list {
  margin: auto;
  width: 40%;
  padding: 13% 300px 265px 300px;
}
.n-thing {
  padding-top: 6px;
}
</style>