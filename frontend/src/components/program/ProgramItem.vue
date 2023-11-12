<script setup>
import TheButton from '@/components/layouts/TheButton.vue';
import { usePopups } from '@/stores/popups';
import { ref, onMounted, onBeforeUnmount } from 'vue';

import { useUser } from '@/stores/user'

const userStore = useUser()
const emit = defineEmits(['openCoursePopup'])

const openPopup = () => {
    if(!userStore.checkRole('teacher') && !userStore.checkRole('employee') && !userStore.checkRole('admin')) {
        popupStore.disableScroll('mainForm')
        emit('openCoursePopup',props.program.id)
    }
}

const popupStore = usePopups()

const props = defineProps({
    program: {
        type: Object,
        required: true
    },
    quantityProgramsToShow: {
        type: Number,
        required: true
    },
    number: {
        type: Number,
        required: true
    },
})

let photoNumber = ref(getRandomNumberInRange(0,3))

let hiddenBlock = ref(null)
let hiddenBlockHeight = ref(0)
let showHiddenText = ref(false)

const updateHiddenBlockHeight = () => {
    if (hiddenBlock.value) {
        hiddenBlockHeight.value = hiddenBlock.value.offsetHeight;
    }
}

function getRandomNumberInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}


onMounted(() => {
    updateHiddenBlockHeight();
    window.addEventListener('resize', updateHiddenBlockHeight);
});

onBeforeUnmount(() => {
    window.removeEventListener('resize', updateHiddenBlockHeight);
});


</script>


<template>
    <div class="program__item"
        v-if="number <= quantityProgramsToShow"
        @mouseenter="showHiddenText = true"
        @mouseleave="showHiddenText = false"
    >
        <div class="program__item-pic">
            <img src="@/assets/images/program-finance.png" v-if="photoNumber === 0">
            <img src="@/assets/images/program-disain.png" v-if="photoNumber === 1">
            <img src="@/assets/images/card-steps.png" v-if="photoNumber === 2">
            <img src="@/assets/images/program-finance.png" v-if="photoNumber === 3">
        </div>
        <div class="program__item-header">
            <div class="program__item-header-btn">
                {{ props.program.period.starts_at }}
            </div>
            <div class="program__item-header-btn">
                {{ props.program.education_form.name }}
            </div>
        </div>
        <div class="program__item-text"
            :class="{active: showHiddenText}"
            :style="{bottom: showHiddenText ? '0px' : `-212px`}"
            @resize="handleResize"
        >

            <div class="program__item-title">
                {{ props.program.name }}
            </div>
            <div class="hidden" ref="hiddenBlock"
            >
                <div class="program__item-description scroll-custom">
                    {{ props.program.description }}
                </div>
                <div class="program__item-btns">
                    <the-button
                        :styles="['btn btn_red']"
                        :type="'button'"
                        @click="openPopup()"
                    >
                        Записаться
                    </the-button>
                    <the-button
                        :styles="['btn btn_grey']"
                        :type="'button'"
                    >
                        Подробнее
                    </the-button>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.program {

    &__item {
        flex: 0 0 calc(25% - 30px);
        display: flex;
        flex-direction: column;
        position: relative;
        min-height: 400px;
        justify-content: space-between;
        border-radius: var(--rounded-3xl, 24px);
        overflow: hidden;
        @media (max-width: 1309px) {
            flex: 0 0 calc(33.33% - 27px);
        }
        @media (max-width: 1023px) {
            flex: 0 0 calc(50% - 20px);
        }
        @media (max-width: 899px) {
            flex: 0 0 calc(50% - 10px);
        }
        @media (max-width: 767px) {
            flex: 0 0 100%;
        }
    }

    &__item-pic {
        position: absolute;
        transition: 0.4s;
        bottom: 0;
        height: 100%;
        left: 0;
        max-width: 100%;
        position: absolute;
        right: 0;
        top: 0;
        border-radius: var(--rounded-3xl, 24px);
        width: 100%;
        &::after {
            content: "";
            width: 100%;
            height: 100%;
            left: 0;
            top: 0;
            background: rgba(0, 0, 0, 0.1);
            position: absolute;
            opacity: 0;
        }
        & img {
            height: 100%;
            width: 100%;
            object-fit: cover;
            border-radius: var(--rounded-3xl, 24px);
        }
    }

    &__item-header {
        display: flex;
        gap: 8px;
        padding: 20px 20px 20px 20px;
        flex-wrap: wrap;
    }

    &__item-header-btn {
        color: var( --var-grey-lite-font);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        display: flex;
        justify-content: center;
        align-items: center;
        text-align: center;
        padding: 10px 10px;
        border-radius: var(--rounded-xl, 12px);
        background: var(--var-grey-lite);
        backdrop-filter: blur(2px);
        transition: .2s;
    }

    &__item-text {
        // -webkit-line-clamp: 7; /* Число отображаемых строк */
        // display: -webkit-box; /* Включаем флексбоксы */
        // -webkit-box-orient: vertical; /* Вертикальная ориентация */
        // overflow: hidden; /* Обрезаем всё за пределами блока */
        z-index: 2;
        position: relative;
        border-radius: 20px;
        backdrop-filter: blur(4.5px);
        position: absolute;
        left: 0;
        width: 100%;
        background: linear-gradient(to bottom,  rgba(243,244,246,1) 52%,rgba(243,244,246,1) 60%,rgba(243,244,246,1) 63%); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
        transition: .5s;
        &.active {
            background: linear-gradient(to bottom,  rgba(243,244,246,0.7) 52%,rgba(243,244,246,0.92) 60%,rgba(243,244,246,1) 63%); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
        }
    }

    &__item-title {
        color: var(--var-grey-lite-font);
        font-family: Inter;
        font-size: 18px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 27px */
        padding: 30px 20px 30px 20px;
        border-radius: 24px 24px 0 0;
        max-height: 120px;
        overflow: hidden;
        text-overflow: ellipsis;
        -webkit-line-clamp: 2; /* Число отображаемых строк */
        display: -webkit-box; /* Включаем флексбоксы */
        -webkit-box-orient: vertical; /* Вертикальная ориентация */
        overflow: hidden; /* Обрезаем всё за пределами блока */
    }
    &__item-description {
        color: var(--var-grey-lite-font);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
        padding-top: 15px;
        padding-bottom: 30px;
        padding: 0px 20px 0px 20px;
        height: 126px;
        overflow: auto;
    }

    &__item-btns {
        display: flex;
        gap: 15px;
        padding: 15px 20px 25px 20px;
    }

    & .hidden {
    }
    
}

.fadeHeight-enter-active,
.fadeHeight-leave-active {
  transition: all 2s;
  height: 230px;
  
}
.fadeHeight-enter,
.fadeHeight-leave-to
{
  opacity: 0;
  height: 0px;
}

[dark=true] {
    & .program__item-text {
        background: linear-gradient(to bottom,  rgba(55,65,81,1) 52%,rgba(55,65,81,1) 60%,rgba(55,65,81,1) 63%); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
        &.active {
            background: linear-gradient(to bottom,  rgba(55,65,81,0.9) 52%,rgba(55,65,81,0.97) 60%,rgba(55,65,81,1) 63%); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
        }
    }

    & .program__item-pic {
        &::after {
            opacity: 1;
        }
    }
}
</style>