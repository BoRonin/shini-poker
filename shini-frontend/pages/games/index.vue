<template>
    <div class="mt-10 text-center">

        <h1 class="title" ref="whos">Кто играет?</h1>
        <div class="pickPlayers mt-10 flex">
            <div class="user" :class="u.active ? `active` : ``" @click="toggleActive(i)" v-for="u, i in aUsers"
                :key=u?.user.id>
                <div class="icon">
                    <img :src="`/images/${u.user.login.toLowerCase()}.jpg`">
                    <transition name="page">
                        <IconsTick class="iconTick" v-if="u.active" />
                    </transition>
                </div>
                <h3>{{ u.user.name }}</h3>
            </div>
        </div>
        <form method="POST" @submit.prevent="CreateGame()">
            <div class="pickMulti flex">
                <div>
                    <input id="game_name" name="game_name" type="text" required
                        class="relative block appearance-none rounded-none border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-slate-700 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                        placeholder="Название игры" v-model="game" />
                </div>
                <div> <span>Множитель:</span>
                    <div class="multiplierBlock">
                        <div class="mult" @click="minus()">
                            <IconsMinus />
                        </div>
                        <div class="multiplier">{{ multiplier }}</div>
                        <div class="mult" @click="plus()">
                            <IconsPlus />
                        </div>
                    </div>
                </div>
                <button class="button" type="submit">Начать игру</button>
                <h1>{{ errorMessage }}</h1>
            </div>
        </form>
    </div>
</template>

<script setup lang="ts">
import { Game } from '@/types/Game'
import { User } from '@/types/User'
import { onClickOutside } from '@vueuse/core'
const config = useAppConfig()
const whos = ref()
const game = ref("")
const multiplier = ref(2)
const errorMessage = ref("")
interface activeUser {
    user: User,
    active: boolean,
}

const aUsers = ref<activeUser[]>([])
const minus = () => {
    if (multiplier.value === 0) {
        return
    }
    multiplier.value--
}
const plus = () => {
    multiplier.value++
}
function toggleActive(int: number) {
    aUsers.value[int].active = !aUsers.value[int].active
}

const { data: users } = await useFetch<User[]>(config.BASE_URL + "admin/users", {
    method: 'GET',
    credentials: 'include',
    transform: users => {
        users.forEach(element => {
            aUsers.value?.push({ user: element, active: true })
        })
        return users
    }
})

const CreateGame = () => {
    let gameid = 0
    useFetch(config.BASE_URL + "admin/creategame", {
        method: 'POST',
        credentials: 'include',
        body: {
            title: game.value,
            multiplier: multiplier.value.toString()
        },
        onResponse({ request, response, options }) {
            if (response.status === 403) {
                errorMessage.value = "Выбери название по-оригинальнее"
                return
            }
            aUsers.value.forEach(e => {
                if (!!e.active) {
                    useFetch(config.BASE_URL + `admin/addplayer/${response._data?.id}`, {
                        credentials: "include",
                        method: 'POST',
                        body: {
                            user_id: e.user.id.toString(),
                            chips: "100",
                        },
                    })
                }
                console.log('[fetch response]', response._data)

                gameid = response._data.id

            });
            setTimeout(()=>{
            navigateTo(`/games/${gameid}`)

        }, 200)
            
        },
    })
}

</script>

<style lang="sass">
.multiplierBlock
    display: flex
    justify-content: space-around
    gap: 1rem
    user-select: none
    .mult
        display: grid
        place-content: center
        padding: .1rem 1rem
        border: 1px solid black
        color: white
        background-color: #333
        cursor: pointer
        font-weight: bold
        text-transform: uppercase
.pickMulti
    justify-content: space-evenly
    align-items: flex-end

    button
        padding: .5rem 1rem
        border-radius: .3rem
        color: white
        background-color: #333
        transition: all .3s ease
        height: 3rem
        &:hover
            background-color: #555
    #miltiplier
        margin: 0 1rem
    
        
.pickPlayers
    justify-content: space-around
    width: 50rem
    margin: 5rem auto           
    .user
        .icon
            border-radius: 50%
            border: 1px solid black
            overflow: hidden
            width: 6rem
            cursor: pointer
            position: relative
            user-select: none
            .iconTick
                position: absolute
                top: 50%
                left: 50%
                transform: translateX(-50%) translateY(-50%)

            img
                aspect-ratio: 1/1
                object-fit: cover


</style>