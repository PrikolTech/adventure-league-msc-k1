import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useAlerts = defineStore('alerts', () => {
    const alerts = ref([])
    let counterForId = 0

    function addAlert(text, type) {
        if(alerts.value.length >= 5) {
            alerts.value.shift()
        }
        counterForId = counterForId + 1
        const alert = {
            text: text,
            type: type,
            id: counterForId
        }

        alerts.value.push({...alert})

        setTimeout(() => {
            if(alerts.value) {
                alerts.value.shift()
            }
        }, 8000)
    }

    function deleteAlert(id) {
        const index = alerts.value.findIndex(alert => alert.id === id)
        if (index !== -1) {
            alerts.value.splice(index, 1)
        }
    }



    return { alerts, addAlert, deleteAlert}
})