<template>
    <div class="mt-10 text-center">
        <h1 class="title">Кто играет?</h1>
        <div class="pickPlayers mt-10 flex" >
            <div class="user" :class="u.active? `active`:``" @click="toggleActive(i)" v-for="u, i in aUsers" :key=u?.user.id >
                <div class="icon">
                    <img :src="`/images/${u.user.login.toLowerCase()}.jpg`">
                    <IconsTick class="iconTick" v-if="u.active"/>
                </div>
                <h3>{{ u.user.name }}</h3>
        </div>
        </div>
        <div class="pickMulti flex">
            <div> <span>Множитель:</span>  
                <input type="number" id="miltiplier" v-model="multiplier">
            </div>
            <button class="button">Начать игру</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { User } from '@/types/User'
const config = useAppConfig()
const multiplier = ref ()
interface activeUser {
    user: User,
    active: boolean,
}

const aUsers = ref<activeUser[]>([])
function toggleActive(int: number){
    aUsers.value[int].active = !aUsers.value[int].active
}

const { data: users } = await useFetch<User[]>(config.BASE_URL + "admin/users", {
    method: 'GET',
    credentials: 'include',
    transform: users=> {
        users.forEach(element => {
            aUsers.value?.push({user: element, active: false})
        })
        return users
    }
})

</script>

<style lang="sass" scoped>
.pickMulti
    justify-content: space-evenly

    button
        padding: .5rem 1rem
        border-radius: .3rem
        color: white
        background-color: #333
        transition: all .3s ease
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