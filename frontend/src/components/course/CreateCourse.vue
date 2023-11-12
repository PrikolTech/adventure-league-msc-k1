<script setup>
import TheModal from "../layouts/TheModal.vue";
import { ref } from 'vue';
import { useRoute } from 'vue-router'

import { useAlerts } from '@/stores/alerts'

const emit = defineEmits(['createdLesoon'])

const alertsStore = useAlerts()
const namePopup = 'createCourse'
const route = useRoute()

let nameCourse = ref('Курс по адекватности')
let descCourse = ref('Лекция по созданию фронта описание')
let advantagesCourse = ref('Какие то преимущества')
let priceCourse = ref(1000)
let periodStartsAt = ref('2023-08-10')
let periodEndAt = ref('2023-08-10')
let courseTypeName = ref('Тип курса')
let ceducationFormName = ref('Тип обучения на курса')

const createLesson = async () => {
    const dateTime = `${dateLesson.value} - ${timeLesson.value}` 
    // const dateTime = "2023-11-17T11:11:11"
    console.log(dateTime)
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json' // Specify content type as JSON
            },
            body: JSON.stringify({
                name: nameLesson.value,
                description: descLesson.value,
                starts_at: dateTime,
            }),
        })

        console.log(response)
        if(response.ok) {
            alertsStore.addAlert('Лекция создана!', 'success')
            emit('createdLesoon')
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