<template>
    <div class="mt-10 text-center">
        
        <h1 class="title" ref="whos">Кто играет?</h1>
        <div class="pickPlayers mt-10 flex">
            <div class="user" :class="u.active ? `active` : ``" @click="toggleActive(i)" v-for="u, i in aUsers"
                :key=u?.user.id>
                <div class="icon">
                    <img :src="`/images/${u.user.login.toLowerCase()}.jpg`">
                    <IconsTick class="iconTick" v-if="u.active" />
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
            </div>
        </form>
    </div>
</template>

<script setup lang="ts">
import {Game} from '@/types/Game'
import { User } from '@/types/User'
import { onClickOutside } from '@vueuse/core'
const config = useAppConfig()
const whos = ref()
const game = ref("")
const multiplier = ref(2)
interface activeUser {
    user: User,
    active: boolean,
}
// onClickOutside(whos, (e) => {
//     console.log("Not in title block");
    
// })
const aUsers = ref<activeUser[]>([])
const minus = () =>{
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
            aUsers.value?.push({ user: element, active: false })
        })
        return users
    }
})
const CreateGame = async () => {
    
    useFetch(config.BASE_URL + "admin/creategame",{
        method:'POST',
        credentials:'include',
        body: {
            title: game.value,
            multiplier: multiplier.value.toString()
        },
        async onResponse({ request, response, options }) {
            aUsers.value.forEach(e => {
                if (!!e.active) {
                    useFetch(config.BASE_URL + `admin/addplayer/${response._data?.id}`,{
                        credentials:"include",
                        method:'POST',
                        body: {
                            user_id: e.user.id.toString(),
                            chips: "100",
                        },
                        async onResponseError({response}){
                            console.log('[fetch response error]', request, response.status, response._data)
                        }
                    })
                }
            });
            console.log('[fetch response]', response._data)
            navigateTo(`/games/${response._data.id}`)
        }
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