<script setup>
import { ref } from "vue";
let photoNumber = ref(getRandomNumberInRange(0,3))
function getRandomNumberInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}
</script>


<template>
    <div class="me__courses">
        <div class="me__courses-list">
            <div class="me__courses-item"
                v-for="(mark, index) of [1,2,3]" :key="index"
            >
            <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 0">
            <img class="pic" src="@/assets/images/program-disain.png" v-if="photoNumber === 1">
            <img class="pic" src="@/assets/images/card-steps.png" v-if="photoNumber === 2">
            <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 3">
                <div class="me__courses-item-body">
                    <div class="me__courses-item-header">
                        <p>
                            Курс : Финансовая грамотность
                        </p>
                        <span>
                            Вы прошли половину обучения.
                        </span>
                    </div>
                    <div class="me__courses-item-content">
                        <div class="me__courses-data">
                            <p>
                                Просмотрено лекций :  
                            </p>
                            <span>
                                <b>
                                    78 /
                                </b>
                                120
                            </span>
                        </div>
                        <div class="me__courses-data">
                            <p>
                                Выполнено ДЗ :
                            </p>
                            <span>
                                <b>
                                    65 / 
                                </b>
                                100
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
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
        padding-left: 150px;
        @media (max-width: 767px) {
            padding-left: 75px;
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
            max-width: 202px;
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
            padding: 15px;
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
        gap: 40px;
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