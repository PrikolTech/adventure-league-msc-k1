<script setup>
import { onMounted, ref, watch } from "vue";
import CourseItem from "@/components/profile/courses/CourseItem.vue";
import TheSort from "@/components/layouts/TheSort.vue";
import TheSwitcher from "@/components/layouts/TheSwitcher.vue";
import { useUser } from '@/stores/user'
import TheButton from '@/components/layouts/TheButton.vue';
import { usePopups } from '@/stores/popups';
import CreateCourse from "../../course/CreateCourse.vue";

const popupStore = usePopups()

const userStore = useUser()
const sortSelectors = ref(['Все'])
let filterSelect = ref('Все')

let selectorIsActive = ref(false)
let showCompleted = ref(false)


let filteredCourses = ref([])
let courses = ref([])
const getUserCourse = async() => {
    if(userStore.checkRole('enrollee')) {
        if(!userStore.checkRole('student')) {
            return
        }
    }
    let url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses`
    if(userStore.checkRole('student')) {
        url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses?user_id=${userStore.user.id}`
    }
    try {
        const response = await fetch(url, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })
        
        const data = await response.json()


        filteredCourses.value.length = 0
        courses.value.length = 0
        filteredCourses.value = [...data]
        courses.value = [...data]

        filteredCourses.value.forEach(el => {
            if(!sortSelectors.value.includes(el.name)) {
                sortSelectors.value.push(el.name)
            }
        });
    } catch(err) {
        console.error(err)
    }
    // const response = [...courses.value]

    // filteredCourses.value = [...response]

    // filteredCourses.value.forEach(el => {
    //     if(!sortSelectors.value.includes(el.title)) {
    //         sortSelectors.value.push(el.title)
    //     }
    // });
}

const filterCourses = (filterValue) => {
    if(filterValue) {
        filterSelect.value = filterValue
    }

    filteredCourses.value = [...courses.value.filter(cource => {
        if(showCompleted.value) {
            if(filterSelect.value === 'Все') {
                return (cource.lecture.passed !== cource.lecture.total && cource.tasks.passed !== cource.tasks.total)
            } else {
                return ((cource.lecture.passed !== cource.lecture.total && cource.tasks.passed !== cource.tasks.total) && cource.title === filterSelect.value)
            }
        } else {
            if(filterSelect.value === 'Все') {
                return true
            } else {
                return cource.name === filterSelect.value
            }
        }
    })]

}

const openPopupCreateCourse = () => {
    popupStore.disableScroll('createCourse')

}

watch(showCompleted, () => {
    filterCourses()
});

onMounted(() => {
    getUserCourse()
})
</script>

<template>
    <div class="profile__courses profile__me profile__block">
        <div class="profile__me-body">
            <div class="profile__me-content">
                <div class="profile__block-header">
                    <div class="profile__title title">
                        Курсы
                    </div>
                    <div class="profile__header-line line"></div>
                </div>
                <div class="profile__courses-filter" v-if="courses.length" style="align-items: center; column-gap: 50px;">
                    <the-sort
                        :selectors="sortSelectors"
                        v-model="selectorIsActive"
                        @select="(modelValue) => filterCourses(modelValue)"
                        v-if="false"
                    />
                    <div class="field" style="flex: 1; margin-bottom: 0px;">
                        <div class="input-w">
                            <input placeholder="Поиск"/>
                        </div>
                    </div>
                    <the-switcher
                        :name="'showCompleted'"
                        v-model="showCompleted"
                        v-if="userStore.checkRole('student')"
                    >
                        Скрыть пройденные
                    </the-switcher>
                    <the-button
                        :styles="['btn_red']"
                        :type="'button'"
                        v-if="userStore.checkRole('tutor')"
                        @click="openPopupCreateCourse()"
                    >
                        Создать курс
                    </the-button>
                </div>
                <div class="text" v-else>
                    Список курсов пуст
                </div>
                <div class="profile__courses-header">
                    
                </div>
                <div class="profile__courses-list">
                    <course-item
                        v-for="(course, index) of filteredCourses" :key="index" :course="course"
                    />
                </div>
            </div>
        </div>
    </div>
    <create-course
        @createdCourse="getUserCourse()"
    />
</template>

<style lang="scss" scoped>

.profile__courses-filter {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    margin-bottom: 40px;
    flex-wrap: wrap;
    row-gap: 15px;
    @media (max-width: 767px) {
        margin-bottom: 20px;
    }
}

.me__courses-item {
    padding-left: 40%;
    @media (max-width: 767px) {
        padding-left: 30%;
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

[dark=true] {
    .me__courses-item-info {
    & p {
        color: var(--gray-200, #E5E7EB);
        background: var(--gray-500, #6B7280);
    }
}
}
</style>