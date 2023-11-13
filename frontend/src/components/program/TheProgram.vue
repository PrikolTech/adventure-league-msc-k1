<script setup>
import { onMounted, ref } from 'vue';
import TheButton from '@/components/layouts/TheButton.vue';
import ProgramItem from '@/components/program/ProgramItem.vue';
import { useUser } from '@/stores/user'

const userStore = useUser()
const emit = defineEmits(['openCoursePopup'])

const openPopup = (id) => {
    emit('openCoursePopup',id)
}
let filterProgram = ref([])

let currentFilterProgram = ref('all')
let quantityProgramsToShow = ref(7)
let programs = ref([])

let filteredPrograms = ref([])
const getPrograms = async () => {
    try {
        const url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses`;
        const response = await fetch(url, {
            method: 'GET',
            mode: 'cors',
        });
        const data = await response.json()
        console.log(data)
        
        programs.value = [...data];
        filteredPrograms.value = [...data];
        filterProgram.value = [
            { text: 'Все', value: 'all' },
            ...[...new Set(programs.value.map(program => program.course_type.name))].map(type => ({ text: type, value: type }))
        ];
    } catch (err) {
        console.error(err)
    }


}

const filterPrograms = (value) => {
    currentFilterProgram.value = value

    if(value === 'all') {
        filteredPrograms.value = [...programs.value]
        return
    }

    filteredPrograms.value = [...programs.value.filter(el => {
        return el.course_type.name === value;
    })];
}

onMounted(() => {
    getPrograms()
})

</script>

<template>
    <div class="program container">
        <div class="program__header">
            <div class="program__title title">
                Программы
            </div>
            <div class="program__line line"></div>
            <div class="program__filter">
                <div class="program__filter-item"
                    v-for="(item, index) of filterProgram" :key="index"
                >
                    <the-button
                        :styles="['btn_grey-lite', {active: currentFilterProgram === item.value}]"
                        :type="'button'"
                        @click="filterPrograms(item.value)"
                    >
                        {{ item.text }}
                    </the-button>
                </div>
            </div>
        </div>
        <div class="program__list">
            <program-item
                v-for="(program, index) of filteredPrograms" :key="index"
                :program="program"
                :number="index"
                :quantityProgramsToShow="quantityProgramsToShow"
                @openCoursePopup="(id) => openPopup(id)"
            />
        </div>
        <div class="program__footer"
            v-if="(quantityProgramsToShow < filteredPrograms.length - 1)"
        >
            <div class="program__footer-line line"></div>
            <the-button
                    :styles="['btn_grey']"
                    :type="'button'"
                    @click="quantityProgramsToShow += 8"
                >
                Показать ещё
            </the-button>
            <div class="program__footer-line line"></div>
        </div>
    </div>
</template>

<style lang="scss" >
.program {

    &__header {
        display: flex;
        gap: 40px;
        align-items: center;
        margin-bottom: 66px;
        @media (max-width: 1023px) {
            flex-wrap: wrap;
            gap: 20px;
            margin-bottom: 30px;
        }
        @media (max-width: 539px) {
            gap: 10px;
            margin-bottom: 15px;
        }
    }

    &__title {
    }

    &__filter {
        display: flex;
        gap: 18px;
        align-items: center;
        @media (max-width: 1023px) {
            flex: 0 0 100%;
            width: 100%;
            overflow: auto;
        }
    }

    &__filter-item {
        & button {
            @media (max-width: 1023px) {
                white-space: nowrap;
            }
            &.active {
                border: 1px solid var(--var-blue) !important;
                color: var(--var-blue) !important;
                &:hover {
                    background: var(--var-blue);
                    color: #fff !important;
                }
            }
        }
    }

    &__line {
        flex: 1;
        @media (max-width: 1023px) {
            display: none;
        }
    }

    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: 40px;
        margin-bottom: 60px;
        @media (max-width: 899px) {
            margin-bottom: 40px;
            gap: 20px;
        }
        @media (max-width: 539px) {
            margin-bottom: 25px;
        }
    }


    &__footer {
        display: flex;
        gap: 47px;
        align-items: center;
        justify-content: space-between;
        @media (max-width: 539px) {
            justify-content: center;
        }
    }

    &__footer-line {
        flex: 1;
        @media (max-width: 539px) {
            display: none;
        }
    }
}

[dark=true] {
    & .program__filter-item {
        & .btn_grey-lite {
            color: var(--gray-300, #D1D5DB);
            border: 1px solid var(--gray-700, #374151);
        }
    }
}

</style>