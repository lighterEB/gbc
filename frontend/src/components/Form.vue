<script setup>
import { ref, reactive } from 'vue'
import { parseTime } from "../utils/date";
import { useMessage } from 'naive-ui';
import { KeyGenRequest } from '../../wailsjs/go/main/App'

const tips = `如果要激活插件,
请在 https://plugins.jetbrains.com/搜索插件名称,
可以在插件首页看到URL中包含一串数字，
把数字替换到下面的链接'#'（注意是替换）\r\n https://plugins.jetbrains.com/api/plugins/#?filed=code ，
然后浏览器打开这个链接找到"purchaseInfo"中的"productCode",
并且选择是插件`

const formRef = ref(null)
const message = useMessage();
const listParams = ref({
    productCode: null,
    nameInput: null,
    emailInput: null,
    extendSelect: null,
    datetimeSelect: null
})

const rules = {
    productCode: {
        required: true,
        trigger: ["blur", "input"],
        message: "请输入产品Code"
    },
    extendSelect: {
        required: true,
        trigger: ["blur", "change"],
        message: "请选择是否为插件"
    },
    nameInput: {
        required: true,
        trigger: ["blur", "input"],
        message: "请输入用户名"
    },
    emailInput: {
        required: true,
        trigger: ["blur", "input"],
        message: "请输入邮箱"
    },
    datetimeSelect: {
        type: "number",
        required: true,
        trigger: ["blur", "change"],
        message: "请选择过期时间"
    },
}
const generalOpt = ["是", "否"].map(
    (v) => ({
        label: v,
        value: v === "是" ? "true" : "false"
    })
)

const params = reactive({
    "licenseId": "",
    "licenseeName": "",
    "assigneeName": "",
    "assigneeEmail": "",
    "products": [{ "code": "PCWMP", "fallbackDate": "", "paidUpTo": "", "extended": "true" }, { "code": "PSI", "fallbackDate": "", "paidUpTo": "", "extended": "true" }]
})
const produts = reactive({ "code": "", "fallbackDate": "", "paidUpTo": "", "extended": "" })


async function genKey() {
    const expired = parseTime(listParams.value.datetimeSelect, "{y}-{m}-{d}")
    produts.code = listParams.value.productCode
    produts.fallbackDate = expired
    produts.paidUpTo = expired
    produts.extended = listParams.value.extendSelect
    params.licenseeName = listParams.value.nameInput
    params.assigneeName = listParams.value.nameInput
    params.assigneeEmail = listParams.value.emailInput
    params.products.forEach((item) => {
        item.fallbackDate = expired
        item.paidUpTo = expired
    })
    params.products.push(produts)
    console.log(params)
    await KeyGenRequest(JSON.stringify(params)).then(res => {
        try {
            if (res.code == 200) {
                navigator.clipboard.writeText(res.data)
                message.success("激活码已成功生成，可以直接粘贴激活。")
            } else {
                message.error("激活码生成失败")
            }
        }catch(e) {
            console.log("激活码生成出错：",e)
            message.error("出错了，错误原因：",e)
        }
        // console.log(res.code)
        // console.log(res.data)
    })
}

function handleValidateButtonClick(e) {
    e.preventDefault();
    formRef.value?.validate((errors) => {
        if (!errors){
            genKey()
        }else {
            message.error("请补全信息！");
        }
    });
}

function clearForm() {
    for (const i in listParams.value) {
        listParams.value[i] = null
    }
}
</script>
<template>
    <n-form ref="formRef" :model="listParams" :rules="rules" label-placement="left" label-width="auto"
        require-mark-placement="right-hanging" size="medium" :style="{
            maxWidth: '640px'
        }">
        <n-form-item label="产品Code:" path="productCode">
            <n-input placeholder="请输入产品Code" v-model:value="listParams.productCode" />
            <n-tooltip width=200 placement="right-start" trigger="hover">
                <template #trigger>
                    <n-tag type="info"
                        style="border-radius: 1em; width: 20px; height: 20px; padding-left: 1px; padding-top: 1px; top: 5px;">
                        <n-image src="../../assets/images/question.png" style="width: 18px; height: 18px;" />
                    </n-tag>
                </template>
                {{ tips }}
            </n-tooltip>
        </n-form-item>
        <n-form-item label="账号:" path="nameInput">
            <n-input placeholder="请输入账号" v-model:value="listParams.nameInput" />
        </n-form-item>
        <n-form-item label="是否为插件" path="extendSelect">
            <n-select v-model:value="listParams.extendSelect" :options="generalOpt" />
        </n-form-item>
        <n-form-item label="邮箱:" path="emailInput">
            <n-input placeholder="请输入邮箱" v-model:value="listParams.emailInput" />
        </n-form-item>
        <n-form-item label="过期时间:" path="datetimeSelect">
            <n-date-picker placeholder="请选择日期" v-model:value="listParams.datetimeSelect" type="date" />
        </n-form-item>
        <div style="display: flex; justify-content: center; gap: 20px">
            <n-button round type="primary" @click="handleValidateButtonClick">
                生成
            </n-button>
            <n-button round type="primary" @click="clearForm" >
                重置
            </n-button>
        </div>
    </n-form>
    <!-- <pre>{{ JSON.stringify(listParams) }}</pre> -->
</template>
<style scoped></style>