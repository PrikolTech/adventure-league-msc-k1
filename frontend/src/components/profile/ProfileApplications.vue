<script setup>
import { onMounted, ref } from "vue";
// import TheButton from '@/components/layouts/TheButton.vue';
import TheApplication from '@/components/profile/applications/TheApplication.vue'
import { useUser } from '@/stores/user'

const userStore = useUser()
const applications = ref([
    // {
    //     title: 'Финансовая грамотность',
    //     date: '06.11.23 - 10.01.24',
    //     format: 'Онлайн',
    //     status: 'Рассмотрение заявки : 10 дней'
    // },
])

const getApplications = async () => {
    let url = null
    if(userStore.checkRole('student')) {
        url = `${import.meta.env.VITE_SERVICE_FORM_URL}/registration?user_id=${userStore.user.user_id}`
    } else if(userStore.checkRole('employee')) {
        url = `${import.meta.env.VITE_SERVICE_FORM_URL}/registration?`
    }

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`,
                'Content-Type': 'application/json'
            },
            mode: 'cors',
        });

        const data = await response.json();
        applications.value = [...data]
        console.log('Заявки',data);
    } catch (err) {
        console.error(err);
    }
}

const getAllCourses = async () => {
    try {
        const url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses`

        const response = await fetch(url, {
            method: 'GET',
        });
        const data = await response.json()

        console.log('Все курсы',data)

        applications.value = [...data]
    } catch (err) {
        console.error(err)
    }
}


onMounted(() => {
    if(userStore.checkRole('employee')) {
        getAllCourses()
    } else if(userStore.checkRole('student')){
        getApplications()
    }
})
</script>

<template>
    <div class="profile__settings profile__block">
        <div class="profile__block-header">
            <div class="profile__title title">
                Заявки
            </div>
            <div class="profile__header-line line"></div>
        </div>
        <div class="applications__content">
            <div class="applications__header">

            </div>
            <div class="applications__list"
                v-if="applications.length"
            >
                <the-application
                    v-for="(course, index) of applications" :key="index" :course="course"
                />
            </div>
            <div class="text"
                v-else
            >
                Список заявок пуст
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>

[dark=true] {
}

</style>