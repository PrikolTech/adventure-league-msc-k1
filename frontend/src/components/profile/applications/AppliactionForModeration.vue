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
        const url = `${import.meta.env.VITE_SERVICE_FORM_URL}/form/registration/${props.application.id}`
        const response = await fetch(url, {
        method: 'PATCH',
        headers: {
            'Authorization': `Bearer ${userStore.user.access}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "status": "approved",
        }),
        mode: 'cors',
    });

    const data = await response.json();
    if(response.ok) {
        console.log('Заявка принята!', response)
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
                {{ application.initiator.last_name }} {{ application.initiator.first_name }}
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
                >
                    Одобрить
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