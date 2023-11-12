<script setup>
import TheButton from '@/components/layouts/TheButton.vue';
import { ref } from "vue"
import { useUser } from '@/stores/user'
import { useAlerts } from '@/stores/alerts'

const alertsStore = useAlerts()
const userStore = useUser()

const props = defineProps({
    application: {
        type: Object,
        required: true,
    }
})

let isHidden = ref(true)

const accept = async() => {
    try {
        let status = null
        console.log(props.application.status)
        if(props.application.status === 'created') {
            status = 'accepted'
        }
        if(props.application.status === 'accepted') {
            status = 'approved'
        }
        if(props.application.status === 'approved') {
            return
        }
        const url = `${import.meta.env.VITE_SERVICE_FORM_URL}/form/registration/${props.application.id}`
        const response = await fetch(url, {
        method: 'PATCH',
        headers: {
            'Authorization': `Bearer ${userStore.user.access}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "status": status,
        }),
        mode: 'cors',
    });

    const data = await response.json();
    if(response.ok) {
        console.log('Заявка принята!', data)
        alertsStore.addAlert('Заявка принята!', 'success')
    }
    
  } catch (err) {
    console.error(err);
  }
}
</script>   

<template>
    <div class="appliaction">
        <div class="appliaction__preview"
            @click="isHidden = !isHidden"
        >
            <span>
                <svg class="created" xmlns="http://www.w3.org/2000/svg" width="24" height="25" viewBox="0 0 24 25" fill="none"
                    v-if="props.application.status === 'created'"
                >
                    <path d="M21.1927 10.3792L20.2947 9.48016C20.1047 9.29116 20.0007 9.03916 20.0007 8.77216V7.50016C20.0007 5.84616 18.6547 4.50016 17.0007 4.50016H15.7287C15.4657 4.50016 15.2077 4.39316 15.0217 4.20716L14.1227 3.30816C12.9527 2.13816 11.0507 2.13816 9.88069 3.30816L8.97969 4.20716C8.79369 4.39316 8.53569 4.50016 8.27269 4.50016H7.00069C5.34669 4.50016 4.00069 5.84616 4.00069 7.50016V8.77216C4.00069 9.03916 3.89669 9.29116 3.70769 9.48016L2.80869 10.3782C2.24169 10.9452 1.92969 11.6992 1.92969 12.5002C1.92969 13.3012 2.24269 14.0552 2.80869 14.6212L3.70669 15.5202C3.89669 15.7092 4.00069 15.9612 4.00069 16.2282V17.5002C4.00069 19.1542 5.34669 20.5002 7.00069 20.5002H8.27269C8.53569 20.5002 8.79369 20.6072 8.97969 20.7932L9.87869 21.6932C10.4637 22.2772 11.2317 22.5692 11.9997 22.5692C12.7677 22.5692 13.5357 22.2772 14.1207 21.6922L15.0197 20.7932C15.2077 20.6072 15.4657 20.5002 15.7287 20.5002H17.0007C18.6547 20.5002 20.0007 19.1542 20.0007 17.5002V16.2282C20.0007 15.9612 20.1047 15.7092 20.2947 15.5202L21.1927 14.6222C21.7587 14.0552 22.0717 13.3022 22.0717 12.5002C22.0717 11.6982 21.7597 10.9452 21.1927 10.3792Z" fill="#E5E7EB"/>
                </svg>
                <svg class="accepted" xmlns="http://www.w3.org/2000/svg" width="24" height="25" viewBox="0 0 24 25" fill="none"
                    v-if="props.application.status === 'accepted'"
                >
                    <path d="M21.1927 10.3792L20.2947 9.48016C20.1047 9.29116 20.0007 9.03916 20.0007 8.77216V7.50016C20.0007 5.84616 18.6547 4.50016 17.0007 4.50016H15.7287C15.4657 4.50016 15.2077 4.39316 15.0217 4.20716L14.1227 3.30816C12.9527 2.13816 11.0507 2.13816 9.88069 3.30816L8.97969 4.20716C8.79369 4.39316 8.53569 4.50016 8.27269 4.50016H7.00069C5.34669 4.50016 4.00069 5.84616 4.00069 7.50016V8.77216C4.00069 9.03916 3.89669 9.29116 3.70769 9.48016L2.80869 10.3782C2.24169 10.9452 1.92969 11.6992 1.92969 12.5002C1.92969 13.3012 2.24269 14.0552 2.80869 14.6212L3.70669 15.5202C3.89669 15.7092 4.00069 15.9612 4.00069 16.2282V17.5002C4.00069 19.1542 5.34669 20.5002 7.00069 20.5002H8.27269C8.53569 20.5002 8.79369 20.6072 8.97969 20.7932L9.87869 21.6932C10.4637 22.2772 11.2317 22.5692 11.9997 22.5692C12.7677 22.5692 13.5357 22.2772 14.1207 21.6922L15.0197 20.7932C15.2077 20.6072 15.4657 20.5002 15.7287 20.5002H17.0007C18.6547 20.5002 20.0007 19.1542 20.0007 17.5002V16.2282C20.0007 15.9612 20.1047 15.7092 20.2947 15.5202L21.1927 14.6222C21.7587 14.0552 22.0717 13.3022 22.0717 12.5002C22.0717 11.6982 21.7597 10.9452 21.1927 10.3792Z" fill="#C3DDFD"/>
                </svg>
                <svg class="approved" xmlns="http://www.w3.org/2000/svg" width="24" height="25" viewBox="0 0 24 25" fill="none"
                    v-if="props.application.status === 'approved'"
                >
                    <path d="M21.1927 10.3792L20.2947 9.48016C20.1047 9.29116 20.0007 9.03916 20.0007 8.77216V7.50016C20.0007 5.84616 18.6547 4.50016 17.0007 4.50016H15.7287C15.4657 4.50016 15.2077 4.39316 15.0217 4.20716L14.1227 3.30816C12.9527 2.13816 11.0507 2.13816 9.88069 3.30816L8.97969 4.20716C8.79369 4.39316 8.53569 4.50016 8.27269 4.50016H7.00069C5.34669 4.50016 4.00069 5.84616 4.00069 7.50016V8.77216C4.00069 9.03916 3.89669 9.29116 3.70769 9.48016L2.80869 10.3782C2.24169 10.9452 1.92969 11.6992 1.92969 12.5002C1.92969 13.3012 2.24269 14.0552 2.80869 14.6212L3.70669 15.5202C3.89669 15.7092 4.00069 15.9612 4.00069 16.2282V17.5002C4.00069 19.1542 5.34669 20.5002 7.00069 20.5002H8.27269C8.53569 20.5002 8.79369 20.6072 8.97969 20.7932L9.87869 21.6932C10.4637 22.2772 11.2317 22.5692 11.9997 22.5692C12.7677 22.5692 13.5357 22.2772 14.1207 21.6922L15.0197 20.7932C15.2077 20.6072 15.4657 20.5002 15.7287 20.5002H17.0007C18.6547 20.5002 20.0007 19.1542 20.0007 17.5002V16.2282C20.0007 15.9612 20.1047 15.7092 20.2947 15.5202L21.1927 14.6222C21.7587 14.0552 22.0717 13.3022 22.0717 12.5002C22.0717 11.6982 21.7597 10.9452 21.1927 10.3792ZM16.5557 11.3322L10.5557 15.3322C10.3867 15.4452 10.1927 15.5002 10.0007 15.5002C9.74269 15.5002 9.48669 15.4002 9.29369 15.2072L7.29369 13.2072C6.90269 12.8162 6.90269 12.1842 7.29369 11.7932C7.68469 11.4022 8.31669 11.4022 8.70769 11.7932L10.1277 13.2132L15.4457 9.66816C15.9067 9.36116 16.5267 9.48516 16.8327 9.94516C17.1397 10.4052 17.0157 11.0262 16.5557 11.3322Z" fill="#0E9F6E"/>
                </svg>
                <svg class="rejected" xmlns="http://www.w3.org/2000/svg" width="20" height="21" viewBox="0 0 20 21" fill="none"
                    v-if="props.application.status === 'rejected'"
                >
                    <path d="M18.2356 7.49999L19.1273 8.39265C19.6903 8.95466 20.0001 9.70235 20.0001 10.4987C20.0001 11.295 19.6893 12.0427 19.1273 12.6057L18.2356 13.4974C18.047 13.6851 17.9437 13.9353 17.9437 14.2004V15.4634C17.9437 17.1058 16.6072 18.4423 14.9648 18.4423H13.7018C13.4407 18.4423 13.1845 18.5485 12.9978 18.7332L12.1051 19.6259C11.5243 20.2068 10.7617 20.4967 9.9991 20.4967C9.23651 20.4967 8.47393 20.2068 7.89305 19.6269L7.00039 18.7332C6.8157 18.5485 6.55952 18.4423 6.29837 18.4423H5.03534C3.393 18.4423 2.05649 17.1058 2.05649 15.4634V14.2004C2.05649 13.9353 1.95322 13.6851 1.76456 13.4974L0.872894 12.6047C0.310884 12.0427 9.09073e-05 11.294 9.09073e-05 10.4987C9.09073e-05 9.70334 0.309891 8.95466 0.872894 8.39166L1.76556 7.49999C1.95322 7.31232 2.05649 7.0621 2.05649 6.79698V5.53395C2.05649 3.89161 3.393 2.55509 5.03534 2.55509H6.29837C6.55952 2.55509 6.8157 2.44885 7.00039 2.26416L7.89504 1.3715C9.05679 0.209746 10.9454 0.209746 12.1071 1.3715L12.9998 2.26416C13.1845 2.44885 13.4407 2.55509 13.7018 2.55509H14.9648C16.6072 2.55509 17.9437 3.89161 17.9437 5.53395V6.79698C17.9437 7.0621 18.047 7.31232 18.2356 7.49999Z" fill="#F98080"/>
                </svg>
                <p>
                    {{ application.initiator.last_name }} {{ application.initiator.first_name }}
                </p>
            </span>
            <svg width="8" height="13" viewBox="0 0 8 13" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M1.78852 12.3332C1.56924 12.3331 1.3549 12.2648 1.17259 12.1367C0.990275 12.0087 0.848183 11.8267 0.764274 11.6138C0.680365 11.4009 0.658406 11.1667 0.701173 10.9407C0.74394 10.7147 0.849512 10.5071 1.00454 10.3441L4.65608 6.507L1.00454 2.66988C0.898635 2.56239 0.814157 2.43381 0.756042 2.29165C0.697927 2.14948 0.667337 1.99658 0.666057 1.84186C0.664778 1.68714 0.692835 1.53371 0.748591 1.3905C0.804346 1.2473 0.886684 1.1172 0.9908 1.00779C1.09492 0.898386 1.21872 0.811864 1.355 0.753275C1.49128 0.694686 1.63729 0.665203 1.78453 0.666548C1.93177 0.667892 2.07727 0.700037 2.21256 0.761106C2.34785 0.822174 2.47021 0.910945 2.5725 1.02224L7.00802 5.68318C7.2159 5.90169 7.33268 6.19802 7.33268 6.507C7.33268 6.81598 7.2159 7.1123 7.00802 7.33082L2.5725 11.9918C2.36459 12.2103 2.08259 12.3331 1.78852 12.3332Z" fill="#9CA3AF"/>
            </svg>
        </div>
        <Transition name="fade">
            <div class="appliaction__hidden" v-if="!isHidden">
                <div class="field">
                    <p>
                        Email *
                    </p>
                    <div class="input-w">
                        <input
                            readonly
                            type="text"
                            v-model="props.application.email"
                        >
                    </div>
                </div>
                <div class="field">
                    <p>
                        Фамилия руководителя
                    </p>
                    <div class="input-w">
                        <input
                            readonly
                            type="text"
                            v-model="props.application.supervisor.last_name"
                        >
                    </div>
                </div>
                <div class="field">
                    <p>
                        Имя руководителя
                    </p>
                    <div class="input-w">
                        <input
                            readonly
                            type="text"
                            v-model="props.application.supervisor.first_name"
                        >
                    </div>
                </div>
                <div class="field">
                    <p>
                        Подразделение *
                    </p>
                    <div class="input-w">
                        <input
                            readonly
                            type="text"
                            v-model="props.application.department"
                        >
                    </div>
                </div>
                <div class="field">
                    <p>
                        Должность *
                    </p>
                    <div class="input-w">
                        <input
                            readonly
                            type="text"
                            v-model="props.application.post"
                        >
                    </div>
                </div>
                <div class="field">
                    <p>
                        Личные достижения в компании * (за последние 12 месяцев)
                    </p>
                    <div class="input-w">
                        <textarea readonly type="text" v-model="props.application.achievements"></textarea>
                    </div>
                </div>
                <div class="field">
                    <p>
                        Мотивационное письмо *
                    </p>
                    <div class="input-w">
                        <textarea readonly type="text" v-model="props.application.motivation"></textarea>
                    </div>
                </div>
                <div class="btns">
                <the-button btn_blue
                    :styles="['btn_blue']"
                    :type="'button'"
                    @click="accept()"
                    v-if="props.application.status !== 'approved'"
                >
                    <span v-if="props.application.status === 'created'">
                        Принять 
                    </span>
                    <span v-if="props.application.status === 'accepted'">
                        Одобрить 
                    </span>
                </the-button>
                <the-button
                    :styles="['btn_red-border']"
                    :type="'button'"

                    @click="Refuse()"
                >
                    Отказать
                </the-button>
                </div>
            </div>
        </Transition>
    </div>
</template>

<style lang="scss" scoped>
.btns {
    display: flex;
    gap: 20px;
    & .btn_blue {
        flex: 1;
    }
}
.appliaction {
    border-radius: 20px;
    background: var(--gray-50, #F9FAFB);
    & span {
        display: flex;
        gap: 15px;
        align-items: center;
    }
    &:not(:last-child) {
        margin-bottom: 20px;
    }
    &__preview {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 20px;
        cursor: pointer;
        padding: 20px;
    }

    &__hidden {
        padding: 0 20px 20px 20px;
    }
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>