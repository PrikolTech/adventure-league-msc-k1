<script setup>
import { ref } from "vue"

const props = defineProps({
    selectors: {
        type: Array,
        required: true
    },
    modelValue: Boolean
})

const emit = defineEmits(['update:modelValue', 'select'])

let currentSelector = ref(props.selectors[0])

const toggleActive = () => {
    emit('update:modelValue', !props.modelValue);
}

</script>

<template>
    <div class="sort"
        :class="[{active: modelValue}]"
    >
        <div class="sort__title"
            @click="toggleActive()"
        >
            {{ currentSelector }}
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16" fill="none">
                <path d="M7.69622 10.0443L7.69628 10.0444C7.77575 10.1199 7.88872 10.1666 8.01153 10.1667L7.69622 10.0443ZM7.69622 10.0443L3.96289 6.49604L3.96294 6.49599M7.69622 10.0443L3.96294 6.49599M3.96294 6.49599L3.95655 6.49013M3.96294 6.49599L3.95655 6.49013M3.95655 6.49013C3.91618 6.45306 3.88547 6.41013 3.86489 6.3646C3.84435 6.31914 3.83395 6.27126 3.83351 6.22354C3.83307 6.17582 3.8426 6.12786 3.86225 6.0822C3.88194 6.03646 3.91176 5.99317 3.95132 5.95557C3.99094 5.91792 4.03944 5.88686 4.09462 5.8654C4.14983 5.84393 4.2098 5.83284 4.27081 5.83335C4.33182 5.83385 4.39151 5.84592 4.44622 5.86826C4.50088 5.89058 4.54867 5.92235 4.58746 5.96053L4.58741 5.96058M3.95655 6.49013L4.58741 5.96058M4.58741 5.96058L4.59371 5.96657M4.58741 5.96058L4.59371 5.96657M4.59371 5.96657L7.66717 8.88769L8.01163 9.21508M4.59371 5.96657L8.01163 9.21508M8.01163 9.21508L8.35609 8.88769M8.01163 9.21508L8.35609 8.88769M8.35609 8.88769L11.4258 5.97014C11.5059 5.89786 11.6173 5.85416 11.7375 5.85515C11.8589 5.85615 11.9704 5.90264 12.049 5.97737C12.1268 6.05132 12.166 6.14588 12.1668 6.23907C12.1677 6.33121 12.1311 6.42508 12.0566 6.49962L8.32704 10.0443L8.32699 10.0444M8.35609 8.88769L8.32699 10.0444M8.32699 10.0444C8.24751 10.1199 8.13454 10.1666 8.01173 10.1667L8.32699 10.0444Z" fill="#374151" stroke="#374151"/>
            </svg>
        </div>
        <div class="sort__list"
            v-show="modelValue === true"
        >
            <div class="sort__item"
                v-for="(select, index) of props.selectors" :key="index" v-show="select !== currentSelector"
                @click="currentSelector = select, toggleActive(), emit('select', select)"
            >
                {{ select }}
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.sort {
    position: relative;
    width: 300px;
    font-family: Roboto;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 125%; /* 17.5px */
    &.active {
        & .sort__title {
            & svg {
                transform: rotate(180deg);
            }
        }
    }
    &__title {
        border-radius: 12px;
        border: 1px solid var(--gray-300, #D1D5DB);
        padding: 12px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 5px;
        cursor: pointer;
        & svg {
            transition: .2s;
        }
    }

    &__list {
        border-radius: var(--rounded-lg, 8px);
        border: 1px solid var(--gray-200, #E5E7EB);
        background: var(--white, #FFF);
        position: absolute;
        z-index: 3;
        width: 100%;
        top: calc(100% + 12px);
        left: 0;
    }

    &__item {
        padding: 8px 12px;
        cursor: pointer;
        &:not(:last-child) {
            border-bottom: 1px solid var(--gray-200, #E5E7EB);
        }
    }
}

[dark=true] {
    & .sort__title {
        color: var(--gray-300, #D1D5DB);
        & svg {
            & path {
                stroke: #6B7280;
            }
        }
    }
    & .sort__list {
        background: #3c444e;
        color: #D1D5DB;
        border: 1px solid #76787a;
    }
    & .sort__item {
        &:not(:last-child) {
            border-bottom: 1px solid #76787a;
        }
    }
}

</style>