<script setup>
import TheButton from './TheButton.vue';
import { usePopups } from '@/stores/popups'

const props = defineProps({
    name: String,
})

const emit = defineEmits(['send'])

const popupStore = usePopups()

</script>

<template>
    <div class="popup modal"
        :class="[{active: popupStore.popups.has(props.name)},props.name]"
        @click.stop.self="popupStore.enableScroll(props.name)"
    >
        <div class="popup__body">
            <div class="popup__content">
                <div class="popup__header">
                    <p class="popup__content-title title">
                        <slot name="title">
                            Заявка на программу
                        </slot>
                    </p>
                    <div class="popup__header-line line"></div>
                    <div class="popup__close"
                        @click="popupStore.enableScroll(props.name)"
                    >
                        <svg width="19" height="19" viewBox="0 0 19 19" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M11.2644 9.5L17.8691 2.89534C17.9882 2.78024 18.0833 2.64255 18.1487 2.49031C18.2141 2.33807 18.2485 2.17434 18.25 2.00865C18.2514 1.84297 18.2198 1.67866 18.1571 1.52531C18.0943 1.37196 18.0017 1.23263 17.8845 1.11547C17.7674 0.998314 17.628 0.905661 17.4747 0.84292C17.3213 0.780179 17.157 0.748607 16.9913 0.750047C16.8257 0.751487 16.6619 0.785909 16.5097 0.851306C16.3575 0.916702 16.2198 1.01176 16.1047 1.13094L9.5 7.7356L2.89534 1.13094C2.66 0.903643 2.3448 0.77787 2.01763 0.780714C1.69046 0.783557 1.3775 0.914788 1.14614 1.14614C0.914788 1.3775 0.783557 1.69046 0.780714 2.01763C0.77787 2.3448 0.903643 2.66 1.13094 2.89534L7.7356 9.5L1.13094 16.1047C1.01176 16.2198 0.916702 16.3575 0.851306 16.5097C0.785909 16.6619 0.751487 16.8257 0.750047 16.9913C0.748607 17.157 0.780179 17.3213 0.84292 17.4747C0.905661 17.628 0.998314 17.7674 1.11547 17.8845C1.23263 18.0017 1.37196 18.0943 1.52531 18.1571C1.67866 18.2198 1.84297 18.2514 2.00865 18.25C2.17434 18.2485 2.33807 18.2141 2.49031 18.1487C2.64255 18.0833 2.78024 17.9882 2.89534 17.8691L9.5 11.2644L16.1047 17.8691C16.34 18.0964 16.6552 18.2221 16.9824 18.2193C17.3095 18.2164 17.6225 18.0852 17.8539 17.8539C18.0852 17.6225 18.2164 17.3095 18.2193 16.9824C18.2221 16.6552 18.0964 16.34 17.8691 16.1047L11.2644 9.5Z" fill="#1F2A37"/>
                        </svg>
                    </div>
                </div>
                <div class="content">
                    <slot>

                    </slot>
                </div>
                <div class="popup__btns">
                    <the-button
                        :styles="['btn_red-border']"
                        :type="'button'"
                        @click="emit('send')"
                    >
                        <slot name="btnSend">
                            Подать заявку
                        </slot>
                    </the-button>
                    <the-button
                        :styles="['btn btn_grey']"
                        :type="'button'"
                        @click="popupStore.enableScroll(props.name)"
                    >
                        Отмена
                    </the-button>
                </div>
            </div>	
        </div>
    </div>
</template>

<style lang="scss" >
.popup {
    width: 100%;
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    overflow-y: auto;
    overflow-x: hidden;
    display: none;
    z-index: 10;
    &.active {
      display: block;
    }
    &::before {
        content: "";
        background: #000000;
        opacity: 0.7;
        position: fixed;
        left: 0px;
        top: 0px;
        width: 100%;
        height: 100%;
        z-index: 104;
    }
}
.popup__content {
    position: relative;
    background: var(--var-popup-bg);
    border-radius: 20px;
    padding: 30px;
    width: 95%;
    max-width: 715px;
    z-index: 105;
    @media (max-width: 539px) {
        padding: 35px 15px 38px;
    }
}
.popup__body {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 30px 10px;
    min-height: 100%;
}

.popup {

    &__header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 20px;
        margin-bottom: 45px;
        @media (max-width: 539px) {
            margin-bottom: 15px;
        }
    }

    &__content-title {
    }

    &__header-line {
        flex: 1;
        @media (max-width: 539px) {
            display: none;
        }
    }

    &__list {
        margin-bottom: 30px;
    }

    &__close {
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
    }

    &__btns {
        display: flex;
        justify-content: center;
        column-gap: 30px;
        row-gap: 10px;
        margin-top: 50px;
        flex-wrap: wrap;
        @media (max-width: 539px) {
            column-gap: 15px;
        }
        & .btn {
            @media (max-width: 539px) {
                white-space: nowrap;
            }
        }
    }

    &__agree {
        display: flex;
        align-items: flex-start;

    }
}

.field {
    width: 100%;
    &:not(:last-child) {
        margin-bottom: 30px;
        @media (max-width: 539px) {
            margin-bottom: 20px;
        }
    }
    & p {
        color: var(--gray-800, #1F2A37);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%;
        @media (max-width: 539px) {
            font-size: 14px;
        }
    }

}

.input-w {
    border-radius: 12px;
    border: 1px solid var(--gray-300, #D1D5DB);
    background: var(--gray-50, #F9FAFB);
    & input, textarea {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
        border-radius: 12px;
        width: 100%;
        background: rgba(255,255,255,0);
        padding: 10px 30px;
        @media (max-width: 539px) {
            padding: 10px 15px;
        }
    }
    & textarea {
        height: 180px;
        resize: none;
    }
}

[dark=true] {
    & .popup__content-title {
        color: var(--gray-50, #F9FAFB);
    }
    & .field {
        & p {
            color: var(--gray-50, #F9FAFB);
        }
    }
    & .input-w {
        border: 1px solid var(--gray-700, #374151);
        background: var(--gray-700, #374151);
        & input, textarea {
            color: #fff;
            &::placeholder {
                color: var(--gray-500, #6B7280);
            }
        }
    }

    & .popup__close {
        & svg {
            & path {
                fill: #fff;
            }
        }
    }
}

</style>