<script lang="ts" setup>
import { Call } from "../../wailsjs/go/main/App"
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const emit = defineEmits(['ok'])
const model = defineModel({ type: Object })
const runtime = (window as any).runtime
const LocationByMouseDown = () => {
  runtime.WindowMinimise();
  Call("LocationByMouseDown", {} as any, {} as any).then((res: any) => {
    let data = res.data || {}
    model.value.args.x = data.x
    model.value.args.y = data.y
    handleChange()
    runtime.WindowUnminimise();
  })
}
const handleChange = () => {
  model.value.label = `${t('Mouse')}${t('Move')} [x=${model.value.args.x} y=${model.value.args.y}]`
  emit("ok")
}
model.value.args.x = 10
model.value.args.y = 10
handleChange()
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <div class="w-32">
      <a-input-number v-model="model.args.x" :min="1" :hide-button="true" @change="handleChange" size="small">
        <template #prepend>X</template>
      </a-input-number>
    </div>
    <div class="w-32">
      <a-input-number v-model="model.args.y" :min="1" :hide-button="true" @change="handleChange" size="small">
        <template #prepend>Y</template>
      </a-input-number>
    </div>
    <div>
      <a-button @click="LocationByMouseDown" title="提取坐标" size="small">
        <template #icon><i class="iconfont icon-coord" /></template>
      </a-button>
    </div>
  </div>
</template>

<style scoped></style>
