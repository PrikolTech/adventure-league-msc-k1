<script setup>
import TheModal from "../layouts/TheModal.vue";
import { ref } from 'vue';
import { useRoute } from 'vue-router'
import { useUser } from '@/stores/user'
import { useAlerts } from '@/stores/alerts'
import { usePopups } from '@/stores/popups'

const emit = defineEmits(['createdLesoon'])

const popupStore = usePopups()
const userStore = useUser()
const alertsStore = useAlerts()
const namePopup = 'createLesson'
const route = useRoute()

let nameLesson = ref('')
let descLesson = ref('')
let dateLesson = ref('')
let timeLesson = ref('')

const createLesson = async () => {
    if(!nameLesson.value || !descLesson.value || !dateLesson.value || !timeLesson.value ) {
        alertsStore.addAlert('Все поля должны быть заполнены!', 'error')
        return
    }
    const dateTime = `${dateLesson.value} - ${timeLesson.value}` 
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${route.params.id}/lectures`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${userStore.user.access}`,
            },
            body: JSON.stringify({
                name: nameLesson.value,
                description: descLesson.value,
                starts_at: dateTime,
            }),
        })

        if(response.ok) {
            alertsStore.addAlert('Лекция создана!', 'success')
            emit('createdLesoon')
            popupStore.enableScroll(namePopup)
            nameLesson.value = ''
            descLesson.value = ''
            dateLesson.value = ''
            timeLesson.value = ''
        }


    } catch(err) {
        console.error(err)
    }
}

</script>

<template>
    <the-modal
        :name="namePopup"
        @send="createLesson()"
    >
        <template v-slot:title>
            Создание занятия
        </template>
        <div class="field">
            <p>
                Название
            </p>
            <div class="input-w">
                <input
                    v-model="nameLesson"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Описание занятия
            </p>
            <div class="input-w">
                <textarea v-model="descLesson"></textarea>
            </div>
        </div>
        <div class="field">
            <p>
                Дата
            </p>
            <div class="input-w">
                <input
                    type="date"
                    v-model="dateLesson"
                >
            </div>
        </div>
        <div class="field">
            <div class="input-w">
                <input
                    type="time"
                    v-model="timeLesson"
                >
            </div>
        </div>
        <template v-slot:btnSend>
            Создать
        </template>
    </the-modal>
</template>

<style lang="scss" scoped>

</style>