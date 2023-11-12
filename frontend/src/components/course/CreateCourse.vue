<script setup>
import TheModal from "../layouts/TheModal.vue";
import { ref } from 'vue';
import { useAlerts } from '@/stores/alerts'
import { useUser } from '@/stores/user'

const emit = defineEmits(['createdLesoon'])

const userStore = useUser()
const alertsStore = useAlerts()
const namePopup = 'createCourse'

let nameCourse = ref('Курс по адекватности')
let descCourse = ref('Лекция по созданию фронта описание')
let advantagesCourse = ref('Какие то преимущества')
let priceCourse = ref(1000)
let tgCourse = ref('link.com')
let prefixCourse = ref('ИДБ-20-09')
let periodStartsAt = ref('2023-08-10')
let periodEndsAt = ref('2023-08-10')
let courseTypeName = ref('Тип курса')
let educationFormName = ref('Test Form')

const createCourse = async () => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${userStore.user.access}`,
            },
            mode: 'cors',
            body: JSON.stringify({
                name: nameCourse.value,
                description: descCourse.value,
                price: priceCourse.value,
                tg_link: tgCourse.value,
                prefix: prefixCourse.value,
                period: {
                    "starts_at": periodStartsAt.value,
                    "ends_at": periodEndsAt.value
                },
                course_type: courseTypeName.value,
                education_form: educationFormName.value
            }),
        })

        console.log(response)
        if(response.ok) {
            alertsStore.addAlert('Курс создан', 'success')
            emit('createdCourse')
        }


    } catch(err) {
        console.error(err)
    }
}
</script>

<template>
    <the-modal
        :name="namePopup"
        @send="createCourse()"
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
                    v-model="nameCourse"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Описание курса
            </p>
            <div class="input-w">
                <textarea v-model="descCourse"></textarea>
            </div>
        </div>
        <div class="field">
            <p>
                Преимущества
            </p>
            <div class="input-w">
                <textarea v-model="advantagesCourse"></textarea>
            </div>
        </div>
        <div class="field">
            <p>
                Цена
            </p>
            <div class="input-w">
                <input
                    v-model="priceCourse"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Ссылка на телеграм
            </p>
            <div class="input-w">
                <input
                    v-model="tgCourse"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Название группы
            </p>
            <div class="input-w">
                <input
                    v-model="prefixCourse"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Дата начала
            </p>
            <div class="input-w">
                <input
                    type="date"
                    v-model="periodStartsAt"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Дата окончания
            </p>
            <div class="input-w">
                <input
                    type="date"
                    v-model="periodEndsAt"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Тип курса
            </p>
            <div class="input-w">
                <input
                    v-model="courseTypeName"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Тип обучения
            </p>
            <div class="input-w">
                <input
                    v-model="educationFormName"
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