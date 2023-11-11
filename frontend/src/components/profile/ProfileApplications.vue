<script setup>
import { onMounted, ref } from "vue";
// import TheButton from '@/components/layouts/TheButton.vue';
import TheApplication from '@/components/profile/applications/TheApplication.vue'
import { useUser } from '@/stores/user'

const userStore = useUser()
const courses = ref([
    {
        title: 'Финансовая грамотность',
        date: '06.11.23 - 10.01.24',
        format: 'Онлайн',
        status: 'Рассмотрение заявки : 10 дней'
    },
])

const getApplications = async () => {
    try {
        const url = `${import.meta.env.VITE_SERVICE_FORM_URL}/registration/${userStore.user.user_id}`;

        const response = await fetch(url, {
            method: 'GET',
            mode: 'cors',
            credentials: 'include'
        });

        console.log(response);
        const data = await response.json();
        console.log(data);
    } catch (err) {
        console.error(err);
    }
}



onMounted(() => {
    getApplications()
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
            <div class="applications__list">
                <the-application
                    v-for="(course, index) of courses" :key="index" :course="course"
                />
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>

[dark=true] {
}

</style>