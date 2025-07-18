<script lang="ts" setup>
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const emit = defineEmits(['ok'])
const model = defineModel({ type: Object })
const handleChange = (value: string) => {
    model.value.args.str = value
    model.value.label = `${t(model.value.name)} [text=${value}]`
    emit("ok")
}
handleChange("")
</script>

<template>
    <div class="flex flex-wrap gap-3">
        <div class="w-full">
            <a-textarea v-model="model.args.str" :auto-size="{ minRows: 3, maxRows: 6 }" placeholder="请输入文本内容" @change="handleChange" />
        </div>
    </div>
</template>

<style scoped></style>
