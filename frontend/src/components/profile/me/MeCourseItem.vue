<script setup>
import { ref, defineProps, computed } from "vue";
import VueApexCharts from "vue3-apexcharts";
import { useUser } from '@/stores/user'

const userStore = useUser()

const props = defineProps({
    course: {
        type: Object,
        required: true
    }
});

const photoNumber = ref(getRandomNumberInRange(0, 3));

function getRandomNumberInRange(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

const progress = computed(() => {
    console.log(((parseInt(props.course.lecture.passed) + parseInt(props.course.tasks.passed)) / (parseInt(props.course.lecture.total) + parseInt(props.course.tasks.total))) * 100)
    return Math.floor(((parseInt(props.course.lecture.passed) + parseInt(props.course.tasks.passed)) / (parseInt(props.course.lecture.total) + parseInt(props.course.tasks.total))) * 100)
})
const chartOptions = ref({
    chart: {
        type: 'radialBar',
    },
    plotOptions: {
        radialBar: {
            dataLabels: {
                showOn: "always",
                name: {
                    show: false,
                },
                value: {
                    fontSize: "22px",
                    offsetY: 8,
                },
            },
            hollow: {
                size: '60%',
            },
        },
    },
    borderRadius: [20], // Устанавливайте здесь радиус скругления (border-radius)
    colors: ['#003791'], // Здесь указывайте свои цвета
});
console.log(userStore)


</script>


<template>
    <div class="me__courses-item">
    <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 0">
    <img class="pic" src="@/assets/images/program-disain.png" v-if="photoNumber === 1">
    <img class="pic" src="@/assets/images/card-steps.png" v-if="photoNumber === 2">
    <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 3">
        <div class="me__courses-item-body" style="display: flex; gap: 5px; flex-direction: row;">
            <div id="chart">
                <VueApexCharts class="test" type="radialBar" :options="chartOptions" :series="[progress]"></VueApexCharts>
            </div>
            <div class="me__courses-item-inside">
                <div class="me__courses-item-header">
                    <p>
                        <!-- Курс : Финансовая грамотность -->
                        {{ props.course.title }}
                    </p>
                    <span>
                        <!-- Вы прошли половину обучения. -->
                        {{ props.course.desc }}
                    </span>
                </div>
                <div class="me__courses-item-content">
                    <div class="me__courses-data">
                        <p>
                            Просмотрено лекций :  
                        </p>
                        <span>
                            <b>
                                <!-- 78 / -->
                                {{ props.course.lecture.passed }} /
                            </b>
                            <!-- 120 -->
                            {{ props.course.lecture.total }}
                        </span>
                    </div>
                    <div class="me__courses-data">
                        <p>
                            Выполнено ДЗ :
                        </p>
                        <span>
                            <b>
                                <!-- 65 /  -->
                                {{ props.course.tasks.passed }} /
                            </b>
                            <!-- 100 -->
                            {{ props.course.tasks.total }}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss">
.test {
    width: 200px;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    left: -33px;
    top: 0;
}
#chart {
}
.me__courses-item-inside {
    padding-left: 100px;
}
.me {

    &__courses {

    }

    &__courses-list {
        max-height: 600px;
        overflow: auto;
        border-radius: 20px;
    }

    &__courses-item {
        border-radius: 20px;
        background: var(--gray-50, #F9FAFB);
        position: relative;
        overflow: hidden;
        padding-left: 20% !important;
        @media (max-width: 767px) {
            padding-left: 20% !important;
        }
        &:not(:last-child) {
            margin-bottom: 30px;
            @media (max-width: 767px) {
                margin-bottom: 15px;
            }
        }
        & .pic {
            position: absolute;
            height: 100%;
            width: 100%;
            object-fit: cover;
            border-radius: var(--rounded-3xl, 24px);
            left: 0;
            top: 0;
            max-width: 32% !important;
        }
    }

    &__courses-item-body {
        padding: 30px;
        z-index: 2;
        position: relative;
        display: flex;
        flex-direction: column;
        gap: 20px;
        background: rgba(243, 244, 246, 0.80);
        backdrop-filter: blur(12.5px);
        border-radius: 30px 20px 20px 30px;
        @media (max-width: 767px) {
            gap: 10px;
            padding: 20px 30px !important;
        }
    }

    &__courses-item-header {
        & p {
            color: var(--gray-900, #111928);
            font-family: Inter;
            font-size: 18px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 27px */
            @media (max-width: 767px) {
                font-size: 15px;
            }
        }
        & span {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
    }

    &__courses-item-content {
        display: flex;
        column-gap: 40px !important;
        row-gap: 10px !important;
        flex-wrap: wrap;
    }

    &__courses-data {
        & p {
            color: var(--gray-500, #6B7280);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
        }
        & span {
            color: var(--unnamed, #7C8092);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%;
            @media (max-width: 767px) {
                font-size: 13px;
            }
            & b {
                color: var(--unnamed, #1F222E);
            }
        }
    }
}
[dark=true] {
    & .me__courses-item {
        background: #1A2537;
    }

    & .me__courses-item-body {
        background: rgba(55, 65, 81, 0.90);
    }

    & .me__courses-item-header {
        & p {
            color: var(--gray-100, #F3F4F6);

        }
        & span {
            color: var(--gray-100, #F3F4F6);

        }
    }

    & .me__courses-data {
        & p {
            color: var(--gray-400, #9CA3AF);
        }
        & span {
            color: var(--gray-400, #9CA3AF);
            & b {
                color: var(--gray-100, #F3F4F6);
            }
        }
    }
}
</style>