<script setup>
import TheButton from '@/components/layouts/TheButton.vue';
import { ref } from 'vue';
const props = defineProps({
    program: {
        type: Object,
        required: true
    },
})
let showHiddenText = ref(false)
</script>

<template>
    <div class="program__item"
        @mouseenter="showHiddenText = true"
        @mouseleave="showHiddenText = false"
    >
        <div class="program__item-pic">
            <img src="@/assets/images/program-finance.png" v-if="props.program.number === '1'">
            <img src="@/assets/images/program-disain.png" v-if="props.program.number === '2'">
            <img src="@/assets/images/card-steps.png" v-if="props.program.number === '3'">
            <img src="@/assets/images/program-finance.png" v-if="props.program.number === '4'">
        </div>
        <div class="program__item-header">
            <div class="program__item-header-btn">
                {{ props.program.date }}
            </div>
            <div class="program__item-header-btn">
                {{ props.program.type }}
            </div>
        </div>
        <div class="program__item-text" :class="{active: showHiddenText}">
            <div class="program__item-title">
                {{ props.program.title }}
            </div>
            <transition name="fadeHeight">
                <div class="hidden" v-if="showHiddenText">
                    <div class="program__item-description">
                        {{ props.program.description }}
                    </div>
                    <div class="program__item-btns">
                        <the-button
                            :styles="['swiper__slide-btn btn btn_red']"
                            :type="'button'"
                        >
                            Записаться
                        </the-button>
                        <the-button
                            :styles="['swiper__slide-btn btn btn_grey']"
                            :type="'button'"
                        >
                            Подробнее
                        </the-button>
                    </div>
                </div>
            </transition>
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
        overflow: hidden;
        min-height: 400px;
        justify-content: space-between;
        border-radius: var(--rounded-3xl, 24px);
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
        z-index: 2;
        position: relative;
        border-radius: 20px;
        backdrop-filter: blur(4.5px);
        position: absolute;
        left: 0;
        width: 100%;
        bottom: 0px;
        background: linear-gradient(to bottom,  rgba(243,244,246,1) 52%,rgba(243,244,246,1) 60%,rgba(243,244,246,1) 63%); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
        transition: .2s;
        &.active {
            background: linear-gradient(to bottom,  rgba(243,244,246,0.7) 52%,rgba(243,244,246,0.92) 60%,rgba(243,244,246,1) 63%); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
        }
    }

    &__item-title {
        // background: var(--gray-100, #F3F4F6);
        color: var(--var-grey-lite-font);
        font-family: Inter;
        font-size: 18px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 27px */
        padding: 30px 20px 30px 20px;
        border-radius: 24px 24px 0 0;
    }
    &__item-description {
        // background: var(--gray-100, #F3F4F6);
        color: var(--var-grey-lite-font);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
        padding-top: 15px;
        padding-bottom: 30px;
        padding: 0px 20px 0px 20px;

    }

    &__item-btns {
        // background: var(--gray-100, #F3F4F6);
        display: flex;
        justify-content: space-between;
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
}
</style>