<script lang="ts" setup>
import { Call } from "../../wailsjs/go/main/App"
import { useI18n } from "vue-i18n"
import { ref } from 'vue'
import keys from '../keys'
const { t } = useI18n()
const emit = defineEmits(['ok'])
const model = defineModel({ type: Object })
const multiple = model.value.name === 'KeyTap'
const loading = ref(false)
const colorByMouseDown = () => {
  loading.value = true;
  Call("KeyByListen", {} as any, {} as any).then((res: any) => {
    if (res.data) {
      if (multiple) {
        model.value.args.keys.push(res.data)
      } else {
        model.value.args.keys = [res.data]
      }
    }
    handleChange()
    loading.value = false;
  })
}
const handleChange = () => {
  model.value.label = `${t(model.value.name)} [key=${model.value.args.keys.join('+')}]`
  emit("ok")
}
model.value.args.keys = []
handleChange()
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <div>
      <a-select v-model="model.args.keys" placeholder="请选择按键" @change="handleChange" size="small" :multiple="multiple" allow-search>
        <a-option v-for="(item, index) of keys" :key="index" :value="item">{{ item }}</a-option>
      </a-select>
    </div>
    <div>
      <a-button @click="colorByMouseDown" :loading="loading" title="提取按键" size="small">
        <template #icon><i class="iconfont icon-keyboard" /></template>
      </a-button>
    </div>
  </div>
</template>

<style scoped></style>
