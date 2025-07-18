<script lang="ts" setup>
import { ref } from 'vue'
import { components, events } from './events'
import { VueDraggable } from 'vue-draggable-plus'
import { Call, GetConfig } from "../wailsjs/go/main/App"
import { IconCheckCircle, IconCloseCircle, IconClockCircle, IconSync, IconRecordStop, IconPlayArrow, IconDelete, IconRefresh, IconSettings } from '@arco-design/web-vue/es/icon';
import { useI18n } from "vue-i18n"
import { useMainStore } from './store'
const runtime = (window as any).runtime
const store = useMainStore()
const tasks: any = ref([])
const config: any = ref({
  display_id: -1,
})
const eventSleep: any = ref(300);
const metadata: any = ref({});
const handleGetConfig = () => {
  GetConfig().then((res: any) => {
    metadata.value = res;
    store.setMetadata(res);
  })
}
handleGetConfig()
const { t } = useI18n()
const visible = ref(false);
const activeTask: any = ref({})
const handleDelete = (id: string) => {
  let i = tasks.value.findIndex((item: any) => item.id === id)
  if (i > -1) {
    tasks.value.splice(i, 1)
  }
}
const handleEdit = (item: any) => {
  if (visible.value && item.id === activeTask.value.id) {
    activeTask.value = {}
    visible.value = false;
  } else {
    activeTask.value = JSON.parse(JSON.stringify(item));
    visible.value = true;
  }
}
const handleOk = () => {
  let i = tasks.value.findIndex((item: any) => item.id === activeTask.value.id)
  if (i > -1) {
    tasks.value[i] = JSON.parse(JSON.stringify(activeTask.value))
  }
  store.setRunList(tasks.value)
};
const getTaskObj = (item: any) => {
  return {
    id: `T-${new Date().getTime()}`,
    name: item.label,
    label: t(item.label),
    component: item.component,
    args: item.args || {},
    results: {},
  }
}
const handleTaskAdd = (e: any) => {
  let item = e.data;
  let obj = getTaskObj(item)
  tasks.value[e.newIndex] = obj
  if (item.component) {
    handleEdit(obj)
  }
}
const handleEventsClick = (item: any) => {
  let obj = getTaskObj(item)
  if (item.component) {
    handleEdit(obj)
  }
  tasks.value.push(obj)
}
const results: any = ref({})
const handleCall = async (item: any) => {
  results.value[item.id].status = "running";
  try {
    if (item.results && item.results.Run) {
      item.args.pid = results.value[item.results.Run].data.pid
    }
  } catch { }
  results.value[item.id] = await Call(item.name, item.args, config.value)
}
const loading = ref(false)
const handleCallAll = async () => {
  if (tasks.value.length == 0) {
    return
  }
  runtime.WindowMinimise();
  loading.value = true;
  results.value = {};
  activeTask.value = {}
  visible.value = false;
  for (const item of tasks.value) {
    results.value[item.id] = { status: "wait" };
  }
  await new Promise(resolve => setTimeout(resolve, eventSleep.value))
  for (const item of tasks.value) {
    if (!loading.value) {
      break;
    }
    await handleCall(item)
    await new Promise(resolve => setTimeout(resolve, eventSleep.value))
  }
  loading.value = false;
  runtime.WindowUnminimise();
}
const handleStop = () => {
  loading.value = false;
}
const handleClear = () => {
  tasks.value = [];
}
const handleClearResults = () => {
  results.value = {};
}
</script>

<template>
  <div class="w-screen h-screen fixed">
    <div class="h-11 border-b border-color flex items-center px-2 gap-2">
      <a-button type="primary" @click="handleCallAll" :loading="loading" :disabled="tasks.length == 0" size="small">
        <template #icon><icon-play-arrow /></template>运行
      </a-button>
      <a-button type="primary" status="danger" @click="handleStop" title="停止运行" :disabled="!loading" size="small">
        <template #icon><icon-record-stop /></template>
      </a-button>
      <a-button type="primary" status="danger" @click="handleClear" title="清空任务" :disabled="loading || tasks.length == 0" size="small">
        <template #icon><icon-delete /></template>
      </a-button>
      <a-button type="primary" status="warning" @click="handleClearResults" title="清空结果" :disabled="loading || Object.keys(results).length === 0" size="small">
        <template #icon><i class="iconfont icon-clear" /></template>
      </a-button>
      <div class="w-28">
        <a-input-number v-model="eventSleep" class="input-center" :min="100" :hide-button="true" placeholder="间隔" title="事件执行间隔" size="small">
          <template #suffix>毫秒</template>
        </a-input-number>
      </div>
    </div>
    <div class="flex body-h">
      <a-scrollbar class="overflow-auto h-full" outer-class="h-full w-60">
        <a-collapse class="t-collapse" :bordered="false" :default-active-key="Object.keys(events)" expand-icon-position="right">
          <a-collapse-item v-for="(items, index) in events" :key="index" :header="$t(index.toString())">
            <VueDraggable class="grid grid-cols-2 gap-2 m-2" v-model="events[index]" :sort="false" :animation="150" :group="{ name: 'events', pull: 'clone', put: false }">
              <div v-for="item in items" :key="item.label">
                <a-button class="w-full !justify-start !px-2" type="dashed" status="normal" @click="handleEventsClick(item)" size="small">
                  <template #icon><i :class="item.icon"></i></template>
                  {{ $t(`${item.label}`) }}
                </a-button>
              </div>
            </VueDraggable>
          </a-collapse-item>
        </a-collapse>
      </a-scrollbar>
      <a-scrollbar class="overflow-auto h-full" outer-class="h-full border-l border-r border-color flex-1 w-0">
        <VueDraggable class="py-1 min-h-full" v-model="tasks" :animation="150" group="events" :onAdd="handleTaskAdd">
          <div v-for="(item, index) in tasks" :key="item.id" class="py-1 px-2">
            <a-button type="dashed" class="w-full !h-auto min-h-12">
              <div class="w-full py-2">
                <div class="w-full flex items-center gap-2">
                  <div class="flex-1 text-left">{{ index + 1 }}. {{ item.label }} </div>
                  <div v-if="results[item.id]">
                    <a-typography-text v-if="results[item.id].status === 'wait'" type="secondary">
                      <icon-clock-circle class="mr-1" />等待执行
                    </a-typography-text>
                    <a-typography-text v-else-if="results[item.id].status === 'running'" type="warning">
                      <icon-sync class="mr-1" />运行中
                    </a-typography-text>
                    <a-typography-text v-else-if="results[item.id].status === true" type="success">
                      <icon-check-circle class="mr-1" />执行成功
                    </a-typography-text>
                    <a-typography-text v-else type="danger">
                      <icon-close-circle class="mr-1" />{{ results[item.id].message }}
                    </a-typography-text>
                  </div>
                  <div v-else>
                    <a-button v-if="item.component" @click="handleEdit(item)" type="text" title="设置" size="small">
                      <template #icon><icon-settings /></template>
                    </a-button>
                    <a-button @click="handleDelete(item.id)" type="text" title="删除" status="danger" size="small">
                      <template #icon><icon-delete /></template>
                    </a-button>
                  </div>
                </div>
                <div v-if="activeTask.id === item.id && visible" class="border border-color rounded bg-white p-3 mt-2">
                  <component :is="components[activeTask.component]" v-model="activeTask" :key="activeTask.id" @ok="handleOk"></component>
                </div>
                <div v-if="results[item.id] && results[item.id].status === true && results[item.id].data" class="bg-slate-200 rounded text-xs p-2 mb-1 mt-2 text-wrap text-left">
                  {{ JSON.stringify(results[item.id].data, null, 2) }}
                </div>
              </div>
            </a-button>
          </div>
        </VueDraggable>
      </a-scrollbar>
    </div>
  </div>
</template>

<style></style>
