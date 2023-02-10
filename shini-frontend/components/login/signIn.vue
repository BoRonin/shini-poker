<template>
    <div class="flex min-h-full items-center text-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div class="w-full max-w-md space-y-8">
            <div>
                <Logo class="logo mx-auto hy-auto flex h-20 w-20" />
                <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Заходи. Не ссы.</h2>
                <a href="#" class="font-medium text-gray-600 hover:text-gray-400" @click="addText()">Регистрации нет,
                    так что если ты не в
                    клубе - ты лох. </a>
                    <p>{{ moreText }}</p>
                    <p class="text-sm text-red-600" v-if="alert"> {{ alert }}</p>
            </div>
            <form class="mt-8 space-y-6" action="#" method="POST" @submit.prevent="Login()">
                <input type="hidden" name="remember" value="true" />
                <div class="-space-y-px rounded-md shadow-sm">
                    <div>
                        <label for="login" class="sr-only">Логин</label>
                        <input id="login" name="login" type="text" required
                            class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-slate-700 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                            placeholder="Логин" v-model="login"/>
                    </div>
                    <div>
                        <label for="password" class="sr-only">Пароль</label>
                        <input id="password" name="password" type="password" autocomplete="current-password" required 
                            class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-slate-700 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                            placeholder="Пароль" v-model="password"/>
                    </div>
                </div>

                <div>
                    <button type="submit"
                        class="group flex w-full justify-center rounded-md border-transparent py-2 px-4 text-sm font-medium text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
                        В ЗАРУБУ!
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { User } from '@/types/User'
const alert = ref("")
const login = ref("")
const password = ref("")
const moreText = ref("")
const data = ref()
const {user} = useAuth()
const Login = async() => {
    
    await useFetch<User>('http://localhost:8080/login', {
        body: {
            login: login.value,
            password: password.value,
        },
        method: 'POST',
        credentials: 'include',
        onResponseError({ request, response, options }) {
            moreText.value = "АААА, хрен тебе!"
            if (response.status === 404){
                alert.value = "Нет такого кекса. Попробуй еще раз!"
            } else if (response.status === 401){
                alert.value = "Забыл пароль - без трусов король!"
            } else {
                alert.value = "Что-то пошло не так. Бей тревогу!"
            }   
  }
    }).then((response) => {
        user.value = response.data.value
        navigateTo('/')
    })
}

const addText = () => {
    moreText.value = "И не писька!"
}
</script>

<style lang="sass" scoped>
    button
        transition: all .3s ease-in-out
        background-image: linear-gradient(-45deg, #8303a0, #173977)

            
</style>

