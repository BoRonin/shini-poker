<template>
    <div class="wrapper">
        <div class="gamePlayers mt-10 flex ">
            <div class="user text-center" v-for="player, i in players" :key="player?.id">
                <div class="icon mx-auto">
                    <playerimage :src='`/images/${player.login?.toLowerCase()}.jpg`'
                        :alt='`${player.user_name} куканит в покерямбе`' />
                </div>
                <p>{{ player.user_name }}</p>
                <p>Чипиков: <span>{{ player.chips }}</span></p>
                <div class="add_chips">
                    <button @click="addChips(player.id)" class="addChips">Добавить чипики</button>
                </div>
                <input id="overall_chips" name="overall_chips" type="number"
                    class=" mt-5 relative block w-full appearance-none rounded-none border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-slate-700 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                    placeholder="Сколько всего фишек" v-model="players[i].chips_final" v-if="players !== null" />
                <div class="addChips win"><icons-plus height="2rem" width="2rem" @click="OpenWinOption(player.id)" />
                </div>
                <div class="Result" v-if="player.money !== undefined">
                    Результат: {{ player.money }}
                </div>
            </div>
            <Teleport to="body">
                <transition name="page">
                    <div class="addWin " v-if="combinations !== null && !!setWin">
                        <div class="ForTheWin flex" ref="FTW">
                            <div class="icon mx-auto mb-10">
                                <Playerimage :source='`/images/${PlayerForWin?.login?.toLowerCase()}.jpg`'
                                    :alt='`${PlayerForWin?.user_name} куканит в покерямбе`' />
                            </div>
                            <div class="combinationBlock">
                                <div class="combination" v-for="comb in combinations" :key="comb.id">
                                    <Button @click="Setwin(comb.id, PlayerForWin?.id)">{{ comb.name }}</Button>
                                </div>
                            </div>
                        </div>
                    </div>
                </transition>
            </Teleport>

        </div>
        <Button class="finishGameButton" @click="finishGame()">Закончить игру</Button>
        <div class="alert">{{ alert }}</div>
    </div>
</template>
<script setup lang="ts">
import { PlayerForGame, Combination } from '~~/types/Game';
import { onClickOutside } from '@vueuse/core'
import { Body } from 'nuxt/dist/head/runtime/components';
const setWin = ref(false)
const PlayerForWin = ref<PlayerForGame>()
const err = ref("")
const config = useAppConfig()
const route = useRoute()
const FTW = ref()
const alert = ref("")
const { data: players } = await useLazyFetch<PlayerForGame[]>(config.BASE_URL + `admin/getplayers/${route.params.id}`, {
    method: 'GET',
    credentials: 'include'
})
interface results {
    player: PlayerForGame,
    score: number,
}

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
    onRequestError({ error }) {
        err.value = error.message
    }
})
function OpenWinOption(id: number) {
    players.value?.forEach(e => {
        if (e.id === id) {
            PlayerForWin.value = e
        }
    })
    setWin.value = true
    onClickOutside(FTW, (e) => {
        setWin.value = false
    })
}
const finishGame = () => {
    let finalPlayers = new Map()
    let allGood = true
    players.value?.forEach(e => {
        if (e.chips_final === undefined) {
            alert.value = "Заполните фишки"
            allGood = false
            return
        }
        let id = e.id.toString()
        let chips = e.chips_final?.toString()
        finalPlayers.set(id, chips)

    });
    if (!!allGood) {
        const body = Object.fromEntries(finalPlayers)
        const { data: results } = useFetch<results[]>(config.BASE_URL + `admin/finishgame/${route.params.id}`, {
            credentials: 'include',
            body: body,
            method: 'POST',
            onResponse({response }) {
                console.log(toRaw(response._data));
                let results = toRaw(response._data)
                results.forEach((res:any) => {
                    players.value?.forEach(pl => {
                      
                        if (pl.id === res.id) {
                            pl.money = res.score
                            console.log(pl.money);
                            
                        }
                    });
                });
            }
        })

    }
}
function Setwin(cid: number, pid: number | undefined) {
    if ((typeof (pid) === 'number')) {
        setWin.value = false
        useFetch(config.BASE_URL + `admin/win/${pid}`, {
            body: {
                game_id: route.params.id,
                combination: cid.toString(),
            },
            credentials: 'include',
            method: 'POST',
        })
    }
}
</script>

<style lang="sass" scoped>
.gamePlayers
    justify-content: space-evenly
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
    top: 0
    display: grid
    place-items: center
    width: 100%
    // height: 100vh
    inset: 0
    background-color: rgba(#333, .9)
    .ForTheWin
        background-color: white
        padding: 5rem
        display: grid
        // grid-auto-columns: 1fr, 1fr
        
        img
            width: 10rem
            height: 10rem
            object-fit: cover
            border-radius: 50%
    .combinationBlock
        display: grid
        grid-template-columns: 1fr 1fr
            

</style>