<script lang="ts" setup>
import { useI18n } from "vue-i18n"
const emit = defineEmits(['ok'])
const model: any = defineModel({ type: Object })
const buttons = [
  "left",
  "center",
  "right",
  "wheelDown",
  "wheelUp",
  "wheelLeft",
  "wheelRight",
]
const { t } = useI18n()
const handleChange = (value: string) => {
  model.value.args.button = value
  model.value.label = `${t(model.value.name)}${t(`MouseButtons.${value}`)}`
  emit("ok")
}
handleChange("left")
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <a-radio-group v-model="model.args.button" @change="handleChange" type="button" size="small">
      <a-radio v-for="item in buttons" :key="item" :value="item">{{ $t(`MouseButtons.${item}`) }}</a-radio>
    </a-radio-group>
  </div>
</template>

<style scoped></style>
