import { ref } from 'vue'
import { defineStore } from "pinia";
import router from '../router';
// import Cookies from 'js-cookie';

export const useUser = defineStore('user', () => {
    const theme = ref(JSON.parse(localStorage.getItem('theme')))
    const user = ref(null)

    // const user = ref({
    //     first_name: 'Артем',
    //     last_name: 'Зимин'
    // })
    // const router = useRouter()

    async function logOut() {
        try {
            const response = await fetch('http://localhost:3001/token', {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
            })

            clearUserInfo()
            router.push('/login')
            console.log(response)

        } catch(err) {
            console.error(err)
        }
    }

    async function getUserInfo() {
        try {
            // const response = await fetch(`http://localhost:3002/user/${user.value.access_token}`, {
            //     method: 'GET',
            // })
            const response = await fetch('http://localhost:3002/user/943bc07e-37ce-49f3-8ed0-710cf980ba95', {
                method: 'GET',
            })
            const data = await response.json()
            Object.assign(user.value, data);
        } catch(err) {
            console.error(err)
        }

    }

    async function refreshTokens() {
        try {
            const response = await fetch('http://localhost:3001/token?grant_type=refresh_token&client_id=000000', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
            });
            console.log(response)
            console.log(response.status === 401)
            if(response.status === 401 || response.status === 401) {
                return
            }
            user.value = {}
            const data = await response.json()
            Object.assign(user.value, data);
            await getUserInfo()
        } catch(err) {
            console.error(err);
        }
    }

    function clearUserInfo() {
        user.value = null
    }


    return { user, logOut, getUserInfo, clearUserInfo, refreshTokens, theme }
})