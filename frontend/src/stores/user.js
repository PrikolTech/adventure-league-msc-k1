import { ref } from 'vue'
import { defineStore } from "pinia";
// import Cookies from 'js-cookie';

export const useUser = defineStore('user', () => {
    const theme = ref(JSON.parse(localStorage.getItem('theme')))
    const user = ref({})
    // const router = useRouter()

    async function logOut() {
        // try {
        //     const { data } = await useFetch('/api/auth/logOut', {
        //         method: 'POST',
        //         headers: {
        //             'Content-Type': 'application/json',
        //         },
        //     })
    
        //     if(data.value) {
        //         Cookies.remove('sessionToken')
        //         clearUserInfo()
        //         router.push('/login')
        //     }

        // } catch(err) {
        //     console.error(err)
        // }
    }

    async function getUserInfo() {
        // try {
        //     const { data } = await useFetch('/api/user/getUser', {
        //         method: 'GET',
        //         headers: {
        //             'Content-Type': 'application/json',
        //         },
        //     })
        //     if(data.value) {
        //         if(data.value.user) {
        //             user.value = {...data.value.user}
        //         }
        //         if(!data.value.user.avatar || data.value.user.avatar === 'default') {
        //             user.value.avatar = 'uploads/avatar-default.svg'
        //         } else {
        //             user.value.avatar = `uploads/${data.value.user.avatar}`
        //         }
        //     }
        // } catch(err) {
        //     console.error(err)
        // }

    }

    function clearUserInfo() {
        user.value = {}
    }

    async function putUser(field) {
        console.log(field)
        // try {
        //     let url = '/api/user/putUser'
        //     let body = { ...user.value }
        //     if(field) {
        //         url = '/api/user/putFieldUser'
        //         body = { ...field }
        //         body.login = user.value.login
        //     }
        //     const { data } = await useFetch(url, {
        //         method: 'PUT',
        //         headers: {
        //             'Content-Type': 'application/json',
        //         },
        //         body: JSON.stringify(body),
        //     })
        //     if(data.value) {
        //         if(data.value.success) {
        //             return data.value.success
        //         }
        //     } else {
        //         throw new Error
        //     }

        // } catch(err) {
        //     console.error(err)
        // }
    }

    return { user, logOut, getUserInfo, putUser, clearUserInfo, theme }
})