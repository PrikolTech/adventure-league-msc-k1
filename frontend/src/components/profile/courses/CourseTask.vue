<script setup>
import { ref } from "vue";

const props = defineProps({
    job: {
        type: Object,
        required: true
    }
})

let showHiddenText = ref(false)
</script>

<template>
    <div class="me__courses-progress-item"
        :class="[{ complited: props.job.is_viewed === true }, { home: props.job.type == 'home' }]"
        @click.stop.prevent=""
        @mouseenter="showHiddenText = true"
        @mouseleave="showHiddenText = false"
    >
        <span class="hidden-info"
            :class="{active: showHiddenText}"
        >
            <!-- Урок 2. Знакомство с компонентами. -->
            {{ props.job.name }}
            <svg xmlns="http://www.w3.org/2000/svg" width="17" height="8" viewBox="0 0 17 8" fill="none">
                <path d="M9.20711 0.707107L16.2071 7.70711C16.3946 7.89464 16.649 8 16.9142 8L0.0857883 8C0.351005 8 0.605358 7.89464 0.792894 7.70711L7.79289 0.707107C8.18342 0.316583 8.81658 0.316582 9.20711 0.707107Z" fill="#1F2A37"/>
            </svg>
        </span>
    </div>
</template>

<style lang="scss" scoped>

.me__courses-progress-item {
    // width: 44px;
    flex: 1;
    max-width: 20px;
    height: 5px;
    border-radius: 30px;
    background: var(--gray-300, #c0c5cd);
    position: relative;
    cursor: pointer;
    @media (max-width: 1023px) {
        // width: 25px;
        position: static;
    }
    &.complited {
        background:#0E9F6E;
    }
    // &.home {
    //     width: 5px;
    //     height: 5px;
    //     border-radius: 50%;
    // }
}
.hidden-info {
    border-radius: var(--rounded-lg, 8px);
    background: var(--white, #FFF);
    box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.08);
    color: var(--gray-900, #111928);
    text-align: center;
    font-family: Roboto;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%; /* 21px */
    position: absolute;
    white-space: nowrap;
    left: 50%;
    transform: translate(-50%,0);
    bottom: -50px;
    padding: 8px 12px;
    opacity: 0;
    pointer-events: none;
    transition: .2s;
    z-index: 2;
    @media (max-width: 1023px) {
        padding: 0px 5px;
        font-size: 12px;
        right: 11px;
        bottom: 2px;
        transform: translate(0, 0);
        left: auto;
        width: fit-content;
    }
    & svg {
        position: absolute;
        left: 50%;
        transform: translate(-50%, 0);
        top: -8px;
        @media (max-width: 1023px) {
            display: none;
        }
        & path {
            fill: #FFF;
        }
    }
    &.active {
        opacity: 1;
        pointer-events: all;
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
        &.complited {
            background:#0E9F6E;
        }
    }
    
    & .hidden-info {
        color: var(--white, #FFF);
        background: var(--gray-800, #1F2A37);
        & svg {
            & path {
                fill: #1F2A37;
            }
        }
    }
}
</style>