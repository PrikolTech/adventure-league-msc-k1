<script setup>
import { onMounted, ref, watch } from "vue";

import { useUser } from '@/stores/user'
import { useRoute, useRouter } from 'vue-router'
import AppliactionForModeration from "./AppliactionForModeration.vue";
const route = useRoute()
const router = useRouter()
const userStore = useUser()

const applications = ref([])
const getCourseApplications = async () => {
    const courseID = route.params.id 
    try {
        const url = `${import.meta.env.VITE_SERVICE_FORM_URL}/form/registration?course_id=${courseID}`
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
    console.log('Информация на курс для модерации Employee', data);
    
  } catch (err) {
    console.error(err);
  }
}

const courseInfo = ref({})
const getCourseInfo = async() => {
    const courseID = route.params.id 
    try {
        const url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${courseID}`
        const response = await fetch(url, {
        method: 'GET',
        mode: 'cors',
    });

    const data = await response.json();
    console.log('Курс с заявки', data);
    courseInfo.value = {...data}
    
  } catch (err) {
    console.error(err);
  }
}

watch(
    () => route.params, // Можно также использовать route.query для отслеживания изменений в query параметрах
    (to, from) => {
        getCourseApplications()
    }
);

onMounted(() => {
    getCourseApplications()
    getCourseInfo()
})

</script>

<template>
    <div class="list">
        <p class="text">
            {{ courseInfo.name }}
        </p>
        <appliaction-for-moderation
            v-for="application of applications" :key="application.id" :application=application
        />
    </div>
</template>

<style lang="scss" scoped>
.list {
    & p {
        margin-bottom: 20px;
    }
}
</style>