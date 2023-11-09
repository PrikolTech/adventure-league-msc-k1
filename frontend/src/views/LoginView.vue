<script setup>
import { ref } from 'vue';
import TheButton from '@/components/layouts/TheButton.vue';
import { useUser } from '@/stores/user'
import EyeIcon from '../components/icons/EyeIcon.vue';
import router from '../router';
import { useAlerts } from '@/stores/alerts'
const userStore = useUser()
const alertsStore = useAlerts()
// import Cookies from 'js-cookie';
let email = ref('12345@mail.ru')
let password = ref('aA$123123')
let passwordIsHidden = ref('password')


const login = async () => {
    try {
        const grantType = 'password';
        const clientId = '000000';

        const queryParams = new URLSearchParams({
            grant_type: grantType,
            client_id: clientId,
            username: email.value,
            password: password.value
        });

        const url = `${import.meta.env.VITE_SERVICE_AUTH_URL}/token?${queryParams.toString()}`;

        const response = await fetch(url, {
            method: 'GET',
            mode: 'cors',
            credentials: 'include'
        });

        if (response.ok) {
            const data = await response.json();
            userStore.user = {}
            userStore.user.access_token = data.access_token
            userStore.user.token_type = data.token_type
            userStore.getUserInfo()
            router.push('/')
        } else {
            alertsStore.addAlert('Неверный логин или пароль', 'error')
            console.error('Ошибка при выполнении запроса.');
        }

    } catch (err) {
        console.error(err);
    }
}


</script>

<template>
    <main>
      <div class="login container">
        <div class="login__header">
            <div class="login__header-line line"></div>
            <div class="login__title">
                Авторизация
            </div>
            <div class="login__header-line line"></div>
        </div>
        <div class="login__form">
            <div class="field">
                <p>
                    Email *
                </p>
                <div class="input-w">
                    <input
                        placeholder=""
                        v-model="email"
                    >
                </div>
            </div>
            <div class="field">
                <p>
                    Пароль
                </p>
                <div class="input-w">
                    <input
                        placeholder=""
                        v-model="password"
                        :type="passwordIsHidden"
                    >
                    <eye-icon
                        :type="passwordIsHidden"
                        @click="passwordIsHidden = (passwordIsHidden === 'password') ? 'text' : 'password'"
                    />
                </div>
            </div>
            <!-- <div class="login__agree agree">
                <label class="b-contain">
                    <span>Запомните меня</span>
                    <input type="checkbox" />
                    <div class="b-input"></div>
                </label>
            </div> -->
            <the-button
                class="login__btn"
                :styles="['btn_red']"
                :type="'button'"
                @click="login()"
            >
                Войти
            </the-button>
        </div>
      </div>
    </main>
</template>

<style lang="scss">
.login {
    &__header {
        display: flex;
        align-items: center;
        margin-bottom: 40px;
        gap: 40px;
        @media (max-width: 539px) {
            margin-bottom: 10px;
        }
    }

    &__header-line {
        flex: 1;
    }

    &__title {
        color: var(--var-grey-lite-font);
        /* text-3xl/font-normal */
        font-family: Montserrat;
        font-size: 30px;
        font-style: normal;
        font-weight: 600;
        line-height: 150%; /* 45px */
    }

    &__form {
        max-width: 390px;
        margin: 0 auto;
    }

    &__agree {
    }

    &__btn {
        margin-top: 30px;
        @media (max-width: 539px) {
            width: 100%;
            margin-top: 20px;
        }
        width: 100%;
    }
}


</style>