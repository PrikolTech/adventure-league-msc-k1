<script setup>
import { computed, onMounted, ref } from "vue";
import TheStep from "./TheStep.vue";
import { useUser } from '@/stores/user'

const userStore = useUser()

const props = defineProps({
    course: {
        type: Object,
        required: true
    }
})

const steps = ref([
    {
        text: 'Рассмотрение',
        active: true
    },
    {
        text: 'Решение',
        active: false
    }
])

let photoNumber = ref(getRandomNumberInRange(0,3))

function getRandomNumberInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}


const courseInfo = ref({})
const getCourseInfo = async () => {
    try {
        const url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${props.course.course_id}`
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

const people = ref([])
const getCourseInfoByEmployee = async () => {
    try {
        const url = `${import.meta.env.VITE_SERVICE_FORM_URL}/form/registration?course_id=${props.course.id}`
        const response = await fetch(url, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${userStore.user.access}`,
            'Content-Type': 'application/json'
        },
        mode: 'cors',
    });

    const data = await response.json();
    console.log('Информация на курс для Employee', data);
    people.value = [...data]
    
  } catch (err) {
    console.error(err);
  }
}

const quantityOfPeople = computed(() => {
    console.log(people.value.length)
    if(userStore.checkRole('employee')) {
        return people.value.length
    } else {
        return ''
    }
})

onMounted(() => {
    if(userStore.checkRole('student')) {
        getCourseInfo()
    } else if(userStore.checkRole('employee')) {
        getCourseInfoByEmployee()
    }
})

</script>

<template>
    <div class="me__courses-item"
        v-if="Object.keys(courseInfo).length > 0 && userStore.checkRole('student')"
    >
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 0">
        <img class="pic" src="@/assets/images/program-disain.png" v-if="photoNumber === 1">
        <img class="pic" src="@/assets/images/card-steps.png" v-if="photoNumber === 2">
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 3">
        <div class="me__courses-item-body">
            <div class="me__courses-item-header">
                <div class="me__courses-item-info">
                    <p>
                        {{ courseInfo.period.starts_at }} - {{ courseInfo.period.ends_at }}
                    </p>
                    <p>
                        {{ courseInfo.education_form.name }}
                    </p>
                </div>
                <p>
                    {{ courseInfo.name }}
                </p>
                <span class="status">
                    {{ courseInfo.status }}
                </span>
                <div class="steps" v-if="userStore.checkRole('student')">
                    <the-step
                        v-for="(step, index) of steps" :key="index" :step="step"
                    />
                </div>
            </div>
        </div>
    </div>
    <router-link :to="`/profile/applications/ApplicationModeration/${props.course.id}`" class="me__courses-item"
        v-if="userStore.checkRole('employee')"
    >
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 0">
        <img class="pic" src="@/assets/images/program-disain.png" v-if="photoNumber === 1">
        <img class="pic" src="@/assets/images/card-steps.png" v-if="photoNumber === 2">
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 3">
        <div class="me__courses-item-body">
            <div class="me__courses-item-header">
                <div class="me__courses-item-info">
                    <p>
                        {{ props.course.period.starts_at }} - {{ props.course.period.ends_at }}
                    </p>
                    <p>
                        {{ props.course.education_form.name }}
                    </p>
                </div>
                <p>
                    {{ props.course.name }}
                </p>
                <div class="people">
                    Заявки: {{ quantityOfPeople }}
                </div>
            </div>
        </div>
    </router-link>
</template>

<style lang="scss" scoped>
.me__courses-item {
    padding-left: 40%;
    overflow: visible;
    display: block;
    @media (max-width: 1023px) {
        position: relative;
    }
    @media (max-width: 767px) {
        padding-left: 30%;
    }
    & .people{
        color: var(--gray-600, #4B5563);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
    }
}
.me__courses-item .pic {
    max-width: 50%;
}

.me__courses-item-info {
    display: flex;
    flex-wrap: wrap;
    gap: 12px;
    margin-bottom: 20px;
    & p {
        color: var(--gray-600, #4B5563);
        font-family: Roboto;
        font-size: 12px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 18px */
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 4px 8px;
        border-radius: var(--rounded-lg, 8px);
        background: var(--gray-200, #E5E7EB);
        transition: .2s;
    }
}

.me__courses-item-content {
    gap: 15px;
}
.me__courses-data {
    font-size: 14px;
}
.me__courses-progress-list {
    display: flex;
    gap: 6px;
    flex-wrap: wrap;
}

.status {
    color: var(--gray-400, #9CA3AF);
    font-family: Roboto;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%; /* 21px */
    margin-bottom: 15px;
    margin-top: 15px;
    display: block;
}
.steps {
    display: flex;
    align-items: center;
    gap: 5px;
    @media (max-width: 539px) {
        flex-direction: column;
        align-items: flex-start;
    }
    &__item {

    }

    &__line {

    }
}


[dark=true] {
    .me__courses-item-info {
        & p {
            color: var(--gray-200, #E5E7EB);
            background: var(--gray-500, #6B7280);
        }
    }
    & .me__courses-progress-item {
        background: var(--gray-300, #6B7280);
    }
    
    .me__courses-item {
            & .people{
                color: var(--gray-600, #F9FAFB);
            }
        }
}
</style>