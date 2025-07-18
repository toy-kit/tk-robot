<script lang="ts" setup>
const model = defineModel({ type: Object })
const emit = defineEmits(['ok'])
const handleChange = (value: number) => {
  model.value.args.tm = value || 1
  model.value.label = `暂停${model.value.args.tm}秒`
  emit("ok")
}
handleChange(1)
</script>

<template>
  <div class="flex flex-wrap gap-3">
    <div class="w-36">
      <a-input-number v-model="model.args.tm" :min="1" :hide-button="true" @change="handleChange" size="small">
        <template #prepend>时长</template>
        <template #suffix>秒</template>
      </a-input-number>
    </div>
  </div>
</template>

<style scoped></style>
