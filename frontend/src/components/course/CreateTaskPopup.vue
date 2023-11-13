<script setup>
import TheModal from "../layouts/TheModal.vue";
import { ref } from "vue";
import { useUser } from '@/stores/user'
import { useAlerts } from '@/stores/alerts'
import { useRoute } from 'vue-router'
import { usePopups } from '@/stores/popups'

const emit = defineEmits(['createdTask'])

const popupStore = usePopups()
const alertsStore = useAlerts()
const route = useRoute()
const userStore = useUser()
const namePopup = 'createTask'
let nameTask = ref('')
let descTask = ref('')

const createTask = async () => {
    if(!nameTask.value || !descTask.value) {
        alertsStore.addAlert('Все поля должны быть заполнены!', 'error')
        return
    }
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/jobs/${route.query.lesson}/homeworks`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${userStore.user.access}`,
            },
            body: JSON.stringify({
                name: nameTask.value,
                description: descTask.value,
            }),
        })

        emit('createdTask')
        if(response.ok) {
            alertsStore.addAlert('Задание создано!', 'success')
            popupStore.enableScroll(namePopup)
        }


    } catch(err) {
        console.error(err)
    }
}


</script>

<template>
    <the-modal
        :name="namePopup"
        @send="createTask"
    >
        <template v-slot:title>
            Создание практического задания
        </template>
        <div class="field">
            <p>
                Название
            </p>
            <div class="input-w">
                <input
                    v-model="nameTask"
                    required
                >
            </div>
        </div>
        <div class="field">
            <p>
                Описание задания
            </p>
            <div class="input-w">
                <textarea v-model="descTask" required></textarea>
            </div>
        </div>
    </the-modal>
</template>

<style lang="scss" >

</style>