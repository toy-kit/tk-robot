import { defineStore } from 'pinia'

export const useMainStore = defineStore('main', {
    state: () => ({
        runList: [] as any,
        metadata: {} as any,
    }),
    getters: {
    },
    actions: {
        setRunList(tasksList: any) {
            let events: any = [];
            for (let i = 0; i < tasksList.length; i++) {
                if (tasksList[i].name == "Run") {
                    events.push({ id: tasksList[i].id, label: tasksList[i].label })
                }
            }
            this.runList = events
        },
        setMetadata(data:any){
            this.metadata = data
        }
    },
})