<script lang="ts" setup>
import { ref } from "vue"
const emit = defineEmits(['ok'])
import { Process } from "../../wailsjs/go/main/App"
import { useMainStore } from '../store'
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const model = defineModel({ type: Object })
const process: any = ref([])
const store = useMainStore()
Process().then(res => {
  process.value = res
})
const handleChange = (value: number) => {
  model.value.results_type = "process"
  model.value.args.pid = value
  let name = process.value.find((item: any) => item.pid == value)?.name
  model.value.label = `${t(model.value.name)} [name=${name || '?'}]`
  emit("ok")
}
const handleChangeId = (value: string) => {
  model.value.results_type = "event"
  model.value.results.Run = value
  let name = store.runList.find((item: any) => item.id == value)?.label
  model.value.label = `${t(model.value.name)} [event=${name}]`
  emit("ok")
}
handleChangeId("")
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <div>
      <a-radio-group v-model="model.results_type" type="button" size="small">
        <a-radio value="event">事件</a-radio>
        <a-radio value="process">进程</a-radio>
      </a-radio-group>
    </div>
    <div v-if="model.results_type === 'event'">
      <a-select v-model="model.results.Run" @change="handleChangeId" placeholder="请选择事件" allow-search size="small">
        <a-option v-for="(item, index) of store.runList" :key="index" :value="item.id">{{ item.label }}</a-option>
      </a-select>
    </div>
    <div v-if="model.results_type === 'process'">
      <a-select v-model="model.args.pid" @change="handleChange" placeholder="请选择程序" allow-search size="small">
        <a-option v-for="(item, index) of process" :key="index" :value="item.pid">{{ item.pid }} - {{ item.name }}</a-option>
      </a-select>
    </div>
  </div>
</template>

<style scoped></style>
