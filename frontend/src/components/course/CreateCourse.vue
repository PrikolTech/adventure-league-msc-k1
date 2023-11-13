<script setup>
import TheModal from "../layouts/TheModal.vue";
import { computed, ref } from 'vue';
import { useAlerts } from '@/stores/alerts'
import { useUser } from '@/stores/user'
import { usePopups } from '@/stores/popups'

const emit = defineEmits(['createdLesoon'])

const popupStore = usePopups()
const userStore = useUser()
const alertsStore = useAlerts()
const namePopup = 'createCourse'

let nameCourse = ref('')
let descCourse = ref('')
let advantagesCourse = ref('')
let priceCourse = ref()
let tgCourse = ref('')
let prefixCourse = ref('')
let periodStartsAt = ref('')
let periodEndsAt = ref('')
let courseTypeName = ref('')
let educationFormName = ref('')

const createCourse = async () => {
    if(!nameCourse.value || !descCourse.value || !advantagesCourse.value || !priceCourse.value || !tgCourse.value || !prefixCourse.value || !periodStartsAt.value || !periodEndsAt.value || !courseTypeName.value || !educationFormName.value) {
        alertsStore.addAlert('Необходимо заполнить все поля', 'error')
        return
    }
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
        if(!response.ok) {
            alertsStore.addAlert('Курс создан', 'success')
            emit('createdCourse')
            popupStore.enableScroll(namePopup)
            nameCourse.value = ''
            descCourse.value = ''
            advantagesCourse.value = ''
            priceCourse.value = ''
            tgCourse.value = ''
            prefixCourse.value = ''
            periodStartsAt.value = ''
            periodEndsAt.value = ''
            courseTypeName.value = ''
            educationFormName.value = ''
        } else {
            alertsStore.addAlert('Курс создан', 'success')
        }


    } catch(err) {
        console.error(err)
    }
}

const currentDate = computed(() => {
    const today = new Date();
    today.setDate(today.getDate() + 1); // Добавляем один день
    const year = today.getFullYear();
    const month = `${today.getMonth() + 1}`.padStart(2, '0');
    const day = `${today.getDate()}`.padStart(2, '0');
    return `${year}-${month}-${day}`;
});
</script>

<template>
    <the-modal
        :name="namePopup"
        @send="createCourse()"
    >
        <template v-slot:title>
            Создание курса
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
                    :min="currentDate"
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
                <input v-model="courseTypeName">
            </div>
        </div>
        <div class="field">
            <p>
                Форма обучения
            </p>
            <div class="input-w">
                <select style="width: 100%; padding: 10px 5px;" v-model="educationFormName">
                    <option value="Очно">Очно</option>
                    <option value="Онлайн + Очно">Онлайн + Очно</option>
                </select>
            </div>
        </div>
        <template v-slot:btnSend>
            Создать
        </template>
    </the-modal>
</template>

<style lang="scss" scoped>
select {
    background: transparent;
}
.select-w {
    border-radius: 12px;
    border: 1px solid var(--gray-300, #D1D5DB);
    background: var(--gray-50, #F9FAFB);
}

[dark=true] {
    .select-w {
        border: 1px solid var(--gray-700, #374151);
        & select {
            background: var(--gray-700, #374151) !important;
        }

    }
}
</style>