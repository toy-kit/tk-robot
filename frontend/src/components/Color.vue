<script lang="ts" setup>
import { Call } from "../../wailsjs/go/main/App"
import { IconBgColors } from '@arco-design/web-vue/es/icon';
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const emit = defineEmits(['ok'])
const model = defineModel({ type: Object })
const runtime = (window as any).runtime
const colorByMouseDown = () => {
  runtime.WindowMinimise();
  Call("ColorByMouseDown", {} as any, {} as any).then((res: any) => {
    model.value.args.color = res.data || ""
    handleChange()
    runtime.WindowUnminimise();
  })
}
const handleChange = () => {
  model.value.label = `${t(model.value.name)} [color=${model.value.args.color}]`
  emit("ok")
}
model.value.args.color = ""
handleChange()
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <div>
      <a-color-picker v-model="model.args.color" format="hex" :show-text="true" :disabled-alpha="true" @change="handleChange" size="small" />
    </div>
    <div>
      <a-button @click="colorByMouseDown" title="提取颜色" size="small">
        <template #icon><icon-bg-colors /></template>
      </a-button>
    </div>
  </div>
</template>

<style scoped></style>
