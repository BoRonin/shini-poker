<template>
    <div class="pickPlayers mt-10 flex ">
        <div class="user text-center" v-for="player in players" :key="player?.id">
            <div class="icon mx-auto">
                <img :src='`/images/${player.login?.toLowerCase()}.jpg`'
                    :alt='`${player.user_name} куканит в покерямбе`'>
            </div>
            <p>{{ player.user_name }}</p>
            <p>Чипиков: <span>{{ player.chips }}</span></p>

            <div class="add_chips">
                <button @click="addChips(player.id)" class="addChips">Добавить чипики</button>
            </div>
            <div class="addChips win"><icons-plus height="2rem" width="2rem" @click="OpenWinOption(player.id)"/></div>
        </div>
        <div class="addWin" v-if="combinations !== null && !!setWin">
            <div class="ForTheWin" ref="FTW">
                <div class="icon mx-auto">
                    <img :src='`/images/${PlayerForWin?.login?.toLowerCase()}.jpg`'
                        :alt='`${PlayerForWin?.user_name} куканит в покерямбе`'>
                </div>
                <h2>{{ PlayerForWin?.user_name }}, {{ PlayerForWin?.id }}</h2>
                <div class="combination" v-for="comb in combinations" :key="comb.id">
                
                <div class="button" @click="Setwin(comb.id, PlayerForWin?.id)">{{ comb.name }}</div>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { PlayerForGame, Combination } from '~~/types/Game';
import { onClickOutside } from '@vueuse/core'
const setWin = ref(false)
const PlayerForWin = ref<PlayerForGame>()
const err = ref("")
const config = useAppConfig()
const route = useRoute()
const FTW = ref()

const { data: players } = useFetch<PlayerForGame[]>(config.BASE_URL + `admin/getplayers/${route.params.id}`, {
    method: 'GET',
    credentials: 'include'
})
const addChips = (player_id: number) => {
    useFetch<PlayerForGame[]>(config.BASE_URL + `admin/addchips/${player_id}`, {
        body: {
            chips: "100",
            game_id: route.params.id.toString(),
        },
        method: 'POST',
        credentials: 'include',
        async onResponse({ response }) {
            players.value = response._data
        }
    })
}
const { data: combinations } = useLazyFetch<Combination[]>(config.BASE_URL + "admin/combinations", {
    credentials: 'include',
    method: 'GET',
    onRequestError({error}){
        err.value = error.message
    }
})
function OpenWinOption(id:number){
        players.value?.forEach(e => {
            if (e.id === id){
                PlayerForWin.value = e
            }
        })
        setWin.value = true
        onClickOutside(FTW, (e)=>{
            setWin.value = false
        })
}
function Setwin(cid: number, pid: number | undefined){
    if ( (typeof(pid) === 'number')){
        setWin.value = false
        useFetch(config.BASE_URL + `admin/win/${pid}`, {
            body:{
                game_id: route.params.id,
                combination: cid.toString(),
            },
            credentials:'include',
            method:'POST',
            onResponse({response}){
                console.log(response._data);
            }
        })
    }
}
</script>

<style lang="sass" scoped>
.user
    span
        font-weight: bold

.addChips
    color: white
    background-color: #333
    padding: .3rem 1rem
    border-radius: .2rem
.win
    padding: 0 0
    margin: 1rem auto
    width: fit-content
    cursor: pointer
.addWin
    position: absolute
    display: grid
    place-items: center
    width: 100%
    height: 100%
    inset: auto
    background-color: #333

</style>