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
            const response = await fetch(`${import.meta.env.VITE_SERVICE_AUTH_URL}/token`, {
                method: 'DELETE',
                mode: 'cors',
                credentials: 'include',
            })

            clearUserInfo()
            router.push('/login')

        } catch(err) {
            console.error(err)
        }
    }

    async function getUserInfo() {
        try {
            const response = await fetch(`${import.meta.env.VITE_SERVICE_USER_URL}/user/${user.value.user_id}`, {
                method: 'GET',
                mode: 'cors',
                headers: {
                    'Authorization': `Bearer ${user.value.access}`
                },
                credentials: 'include'
            })
            const data = await response.json()
            if(data.password) {
                delete data.password
            }
            Object.assign(user.value, data);
        } catch(err) {
            console.error(err)
        }

    }

    async function refreshTokens() {
        try {
            const response = await fetch(`${import.meta.env.VITE_SERVICE_AUTH_URL}/token?grant_type=refresh_token&client_id=000000`, {
                method: 'GET',
                headers: {'Content-Type': 'application/json'},
                mode: 'cors',
                credentials: 'include'
            });
            console.log('test',response)
            if(response.status === 401 || response.status === 400) {
                return
            }
            user.value = {}
            addCustomData()
            const data = await response.json()
            delete data.refresh_token
            Object.assign(user.value, data);
            await getUserInfo()
        } catch(err) {
            console.error(err);
        }
    }

    function clearUserInfo() {
        user.value = null
    }

    function checkRole(roleName) {
        return user.value.roles.some(role => role.title === roleName)
    }
    //временная функция для заполнения данных
    function addCustomData() {
        // user.value.id = '35e8a1e4-c0d9-4a79-be68-192603d0205f'
        // user.value.role = 'teacher'
    }


    return { user, logOut, getUserInfo, clearUserInfo, refreshTokens, checkRole, theme }
})