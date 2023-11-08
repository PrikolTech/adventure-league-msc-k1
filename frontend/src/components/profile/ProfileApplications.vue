<script setup>
import { ref } from "vue";
// import TheButton from '@/components/layouts/TheButton.vue';
import TheApplication from '@/components/profile/applications/TheApplication.vue'
const courses = ref([
    {
        title: 'Финансовая грамотность',
        date: '06.11.23 - 10.01.24',
        format: 'Онлайн',
        status: 'Рассмотрение заявки : 10 дней'
    },
])
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