<script lang="ts" setup>
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const emit = defineEmits(['ok'])
import { OpenFileDialog } from "../../wailsjs/go/main/App"
const model = defineModel({ type: Object })
const handleChange = () => {
  let name = model.value.args.path.split(/[\\/]/).pop();
  if (model.value.name === "Run") {
    name = name?.split(".")[0]
  }
  model.value.args.name = name
  model.value.label = `${t(model.value.name)} [path=${name}]`
  emit("ok")
}
model.value.args.name = ""
model.value.args.path = ""
const openFileDialog = () => {
  OpenFileDialog().then((res) => {
    if (res) {
      model.value.args.path = res;
      handleChange()
    }
  })
}
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <div>
      <a-input v-model="model.args.path" placeholder="请输入或选择路径" @change="handleChange" class="input-append-p-0">
        <template #append>
          <a-button @click="openFileDialog">选择</a-button>
        </template>
      </a-input>
    </div>
  </div>
</template>

<style scoped></style>