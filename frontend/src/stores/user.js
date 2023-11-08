import { ref } from 'vue'
import { defineStore } from "pinia";

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
        // await fetch('te', {
        //     method: "POST",
        //     headers: {
        //         'Content-Type': 'application/json',
        //         'Authorization': `Bearer ${}`
        //     },
        // })
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


    return { user, logOut, getUserInfo, clearUserInfo, theme }
})