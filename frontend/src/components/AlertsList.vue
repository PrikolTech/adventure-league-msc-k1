<script setup>
import { useAlerts } from '@/stores/alerts'
const alertsStore = useAlerts()

</script>

<template>
    <div class="alerts">
        <div class="alert"
            v-for="alert of alertsStore.alerts" :key="alert.id"
            :class="[{ _error: alert.type === 'error', _success: alert.type === 'success'}]"
        >
        <span>
            {{ alert.text }}
        </span>
        <button
            @click="alertsStore.deleteAlert(alert.id)"
        >
            <v-icon
                v-if="alert.type === 'error'"
                end
                icon="mdi-close-circle"
            ></v-icon>
            <v-icon
                v-if="alert.type === 'success'"
                end
                icon="mdi-checkbox-marked-circle"
            ></v-icon>
        </button>
        </div>
    </div>
</template>



<style lang="scss" scoped>
.alerts {
    display: flex;
    flex-direction: column;
    position: fixed;
    right: 5px;
    bottom: 20px;
    gap: 5px;
}
.alert {
    padding: 10px;
    border-radius: 10px;
    color: #fff;
    display: flex;
    gap: 5px;
    align-items: flex-start;
    width: 280px;
    justify-content: space-between;
    &._error {
        background: #e14c4c;
    }
    &._success {
        background: green;
    }
}
</style>