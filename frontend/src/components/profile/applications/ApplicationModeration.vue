<script setup>
import { onMounted, ref, watch } from "vue";
import TheButton from '@/components/layouts/TheButton.vue';
import { useUser } from '@/stores/user'
import { useRoute, useRouter } from 'vue-router'
import AppliactionForModeration from "./AppliactionForModeration.vue";
const route = useRoute()
const router = useRouter()
const userStore = useUser()

const applications = ref([])
let filterByStatus = ref('Все')
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
    courseInfo.value = {...data}
    
  } catch (err) {
    console.error(err);
  }
}

const updateStatus = (application) => {
    const index = applications.value.findIndex(element => element.id === application.id);
    if (index !== -1) {
        applications.value[index] = { ...applications.value[index], ...application };
    }
}


watch(
    () => route.params,
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
        <div class="filter-btns">
            <the-button
                :styles="['btn_grey-lite']"
                :type="'button'"
                @click="filterByStatus = 'Все'"
                :class="{active: filterByStatus === 'Все'}"
            >
                Все
            </the-button>
            <the-button
                :styles="['btn_grey-lite']"
                :type="'button'"
                @click="filterByStatus = 'created'"
                :class="{active: filterByStatus === 'created'}"
            >
                Присланные
            </the-button>
            <the-button
                :styles="['btn_grey-lite']"
                :type="'button'"
                @click="filterByStatus = 'accepted'"
                :class="{active: filterByStatus === 'accepted'}"
            >
                Принятые
            </the-button>
            <the-button
                :styles="['btn_grey-lite']"
                :type="'button'"
                @click="filterByStatus = 'approved'"
                :class="{active: filterByStatus === 'approved'}"
            >
                Одобренные
            </the-button>
            <the-button
                :styles="['btn_grey-lite']"
                :type="'button'"
                @click="filterByStatus = 'rejected'"
                :class="{active: filterByStatus === 'rejected'}"
            >
                Отозванные
            </the-button>
        </div>
        <appliaction-for-moderation
            v-for="application of applications" :key="application.id" :application=application
            :filterStatus="filterByStatus"
            @updateStatus="((appliaction) => updateStatus(appliaction))"
        />
    </div>
</template>

<style lang="scss" scoped>
.list {
    & p {
        margin-bottom: 20px;
    }
}
.filter-btns {
    display: flex;
    flex-wrap: wrap;
    gap: 15px;
    margin-bottom: 20px;
    & .btn {
        &.active {
            border: 1px solid var(--var-blue) !important;
            color: var(--var-blue) !important;
            &:hover {
                background: var(--var-blue);
                color: #fff !important;
            }
        }
    }
    & .btn_grey-lite {
        color: var(--gray-300, #D1D5DB);
        border: 1px solid var(--gray-700, #374151);
        &:hover {
            
        }
    }
}
</style>