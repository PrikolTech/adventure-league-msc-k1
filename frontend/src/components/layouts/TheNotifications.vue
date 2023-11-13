<script setup>
import { onMounted } from "vue"
import { useUser } from '@/stores/user'

const userStore = useUser()
const props = defineProps({
    notificationIsActive: {
        type: Boolean,
        required: true
    }
})

const emit = defineEmits('update:modelValue')

const getNotifications = async () => {
    try {
        const response = await fetch(`http://localhost:3006/api/notifications/`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })
        console.log(response,'уведомления')
    } catch(err) {
        console.error(err)
    }
}

onMounted(() => {
    // getNotifications()
})
</script>

<template>
    <div class="notification">
        <div class="notification__btn">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"
                @click="emit('toggleNotification')"
            >
                <path d="M19.079 14.829C18.4881 14.2477 18.1488 13.4578 18.134 12.629V10.829C18.133 9.35216 17.6212 7.92113 16.6855 6.77854C15.7497 5.63596 14.4477 4.85214 13 4.56001V3.10001C13 2.83479 12.8946 2.58044 12.7071 2.3929C12.5196 2.20536 12.2652 2.10001 12 2.10001C11.7348 2.10001 11.4804 2.20536 11.2929 2.3929C11.1054 2.58044 11 2.83479 11 3.10001V4.56001C9.55183 4.85225 8.24937 5.6365 7.31359 6.7797C6.37781 7.9229 5.86635 9.35465 5.866 10.832V12.418C5.86607 12.4886 5.87378 12.559 5.889 12.628H5.866C5.85125 13.4568 5.51194 14.2467 4.921 14.828C4.38423 15.3542 4.05692 16.0575 4 16.807C4 17.35 4 19 5.538 19H7.746C8.03584 19.9003 8.60385 20.6855 9.36829 21.2425C10.1327 21.7995 11.0542 22.0995 12 22.0995C12.9458 22.0995 13.8673 21.7995 14.6317 21.2425C15.3961 20.6855 15.9642 19.9003 16.254 19H18.462C20 19 20 17.35 20 16.807C19.9428 16.0578 19.6156 15.3549 19.079 14.829ZM12 20.1C11.5932 20.0995 11.1928 19.9989 10.834 19.8071C10.4753 19.6152 10.1693 19.3381 9.943 19H14.057C13.8307 19.3381 13.5247 19.6152 13.166 19.8071C12.8072 19.9989 12.4068 20.0995 12 20.1ZM18 17H6C6 16.933 6 16.864 6 16.807C6.1063 16.54 6.26627 16.2977 6.47 16.095C7.35569 15.1579 7.85559 13.9214 7.87 12.632V10.832C7.8681 9.71355 8.2954 8.637 9.06382 7.8243C9.83224 7.01159 10.8832 6.5247 12 6.46401C13.1175 6.52372 14.1695 7.01017 14.9387 7.82298C15.708 8.63578 16.1358 9.7129 16.134 10.832V12.418C16.1342 12.4886 16.1416 12.5589 16.156 12.628H16.134C16.1484 13.9174 16.6483 15.1539 17.534 16.091C17.7377 16.2937 17.8977 16.536 18.004 16.803C18 16.864 18 16.933 18 17Z" fill="#9CA3AF"/>
            </svg>

        </div>
        <div class="notification__content" v-if="props.notificationIsActive">
            <svg class="arrow" width="52" height="31" viewBox="0 0 52 31" fill="none" xmlns="http://www.w3.org/2000/svg">
                <g filter="url(#filter0_d_2016_11845)">
                <path d="M26 9L45.0526 26.25H6.94744L26 9Z" fill="white"/>
                </g>
                <defs>
                <filter id="filter0_d_2016_11845" x="0.447266" y="0.5" width="51.1055" height="30.25" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
                <feFlood flood-opacity="0" result="BackgroundImageFix"/>
                <feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
                <feOffset dy="-2"/>
                <feGaussianBlur stdDeviation="3.25"/>
                <feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.1 0"/>
                <feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_2016_11845"/>
                <feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_2016_11845" result="shape"/>
                </filter>
                </defs>
            </svg>
            <div class="notification__title">
                <p>Уведомления</p>
                <div class="notification__line line"></div>
                <svg width="30" height="30" viewBox="0 0 30 30" fill="none" xmlns="http://www.w3.org/2000/svg"
                    @click="emit('toggleNotification')"
                >
                    <path d="M16.7644 15L23.3691 8.39534C23.4882 8.28024 23.5833 8.14255 23.6487 7.99031C23.7141 7.83807 23.7485 7.67434 23.75 7.50865C23.7514 7.34297 23.7198 7.17866 23.6571 7.02531C23.5943 6.87196 23.5017 6.73263 23.3845 6.61547C23.2674 6.49831 23.128 6.40566 22.9747 6.34292C22.8213 6.28018 22.657 6.24861 22.4913 6.25005C22.3257 6.25149 22.1619 6.28591 22.0097 6.35131C21.8575 6.4167 21.7198 6.51176 21.6047 6.63094L15 13.2356L8.39534 6.63094C8.16 6.40364 7.8448 6.27787 7.51763 6.28071C7.19046 6.28356 6.8775 6.41479 6.64614 6.64614C6.41479 6.8775 6.28356 7.19046 6.28071 7.51763C6.27787 7.8448 6.40364 8.16 6.63094 8.39534L13.2356 15L6.63094 21.6047C6.51176 21.7198 6.4167 21.8575 6.35131 22.0097C6.28591 22.1619 6.25149 22.3257 6.25005 22.4913C6.24861 22.657 6.28018 22.8213 6.34292 22.9747C6.40566 23.128 6.49831 23.2674 6.61547 23.3845C6.73263 23.5017 6.87196 23.5943 7.02531 23.6571C7.17866 23.7198 7.34297 23.7514 7.50865 23.75C7.67434 23.7485 7.83807 23.7141 7.99031 23.6487C8.14255 23.5833 8.28024 23.4882 8.39534 23.3691L15 16.7644L21.6047 23.3691C21.84 23.5964 22.1552 23.7221 22.4824 23.7193C22.8095 23.7164 23.1225 23.5852 23.3539 23.3539C23.5852 23.1225 23.7164 22.8095 23.7193 22.4824C23.7221 22.1552 23.5964 21.84 23.3691 21.6047L16.7644 15Z" fill="#1F2A37"/>
                </svg>
            </div>
            <div class="notification__list">
                <div class="notification__item"
                    v-for="(notification, index) of [1]" :key="index"
                >
                    <p>
                        Ваше практическое задание по курсу “Финансовая грамотность” к уроку 2 проверено
                    </p>
                    <span>
                        21:00
                    </span>
                </div>
            </div>
        </div>
    </div>
</template>



<style lang="scss" scoped>
.notification {
    position: relative;
    &__btn {
        display: flex;
        cursor: pointer;
    }

    &__content {
        top: 53px;
        right: -59px;
        position: absolute;
        border-radius: 30px;
        background: #FFF;
        box-shadow: 0px 0px 25.8px 0px rgba(0, 0, 0, 0.10);
        width: 570px;
        padding: 30px 40px;
        & .arrow {
            position: absolute;
            top: -26px;
            right: 45px;
        }
    }

    &__title {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 20px;
        margin-bottom: 40px;
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 24px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 36px */
        & svg {
            cursor: pointer;
        }
    }

    &__line {
        flex: 1;
    }

    &__list {
        max-height: 300px;
        overflow: auto;
    }

    &__item {
        &:not(:last-child) {
            margin-bottom: 20px;
        }
        & p {
            color: #000;
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
            margin-bottom: 5px;
        }
        & span {
            color: var(--gray-400, #9CA3AF);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
    }
}

[dark=true] {
    .notification {
    &__btn {
        & svg {
            & path {
                fill: #9CA3AF;
            }
        }
    }

    &__content {
        background: var(--gray-800, #1F2A37);
        box-shadow: 0px 0px 25.8px 0px rgba(0, 0, 0, 0.10);
        & .arrow {
            & path {
                fill: #1F2A37;
            }
        }
    }

    &__title {
        color: var(--gray-50, #F9FAFB);
        & svg {
            & path {
                fill: #D1D5DB;
            }
        }
    }

    &__line {
    }

    &__list {
    }

    &__item {

        & p {
            color: var(--gray-50, #F9FAFB);
        }
        & span {
            color: var(--gray-400, #9CA3AF);
        }
    }
}
}
</style>