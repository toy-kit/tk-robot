<script lang="ts" setup>
import { useI18n } from "vue-i18n"
const emit = defineEmits(['ok'])
const model: any = defineModel({ type: Object })
const directions = ["up", "down", "left", "right"]
const { t } = useI18n()
const handleChange = () => {
  model.value.label = `${t("Mouse")}${t(`Directions.${model.value.args.direction}`)}${t("ScrollDir")} [x=${model.value.args.x} ]`
  emit("ok")
}
model.value.args.direction = "down"
model.value.args.x = 10
handleChange()
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <div>
      <a-radio-group v-model="model.args.direction" type="button" size="small" @change="handleChange">
        <a-radio v-for="item in directions" :key="item" :value="item">{{ $t(`Directions.${item}`) }}</a-radio>
      </a-radio-group>
    </div>
    <div class="w-36">
      <a-input-number v-model="model.args.x" :min="0" :hide-button="true" size="small" @change="handleChange">
        <template #prepend>位置</template>
      </a-input-number>
    </div>
  </div>
</template>

<style scoped></style>
