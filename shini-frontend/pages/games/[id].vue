<template>
    <div class="wrapper main-section">
        <div class="gamePlayers mt-10 flex ">
            <div class="user text-center" v-for="player, i in players" :key="player?.id">
                <div class="icon mx-auto">
                    <playerimage :src='`/images/${player.login?.toLowerCase()}.jpg`'
                        :alt='`${player.user_name} куканит в покерямбе`' />
                </div>
                <p>{{ player.user_name }}</p>
                <p>Чипиков: <span>{{ player.chips }}</span></p>
                <div class="add_chips">
                    <button @click="OpenWinOption(player.id, 1)" class="addChips">Добавить чипики</button>
                </div>
                <input id="overall_chips" name="overall_chips" type="number" min="0"
                    class=" mt-5 relative block w-full appearance-none rounded-none border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-slate-700 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                    placeholder="Всего фишек" v-model="players[i].chips_final" v-if="players !== null" />
                <div class="addChips win"><icons-plus height="3rem" width="3rem" @click="OpenWinOption(player.id, 2)" />
                </div>
                <div class="Result" v-if="player.money !== undefined">
                    Результат: <strong>{{ player.money }}</strong>
                </div>
            </div>
            <Teleport to="body">
                <transition name="element">
                    <div class="addWin " v-if="combinations !== null && !!setWin">
                        <div class="ForTheWin flex" ref="FTW">
                            <div class="icon mx-auto mb-10">
                                <Playerimage :source='`/images/${PlayerForWin?.login?.toLowerCase()}.jpg`'
                                    :alt='`${PlayerForWin?.user_name} куканит в покерямбе`' />
                            </div>
                            <transition name="page" mode="out-in">
                                <div class="combinationBlock" v-if="!assertWinChoice">
                                    <div class="combination" v-for="comb in combinations" :key="comb.id">
                                        <Button class="dark" @click="openAssertWinChoice(comb.name, comb.id, PlayerForWin?.id)">{{
                                            comb.name.toLocaleUpperCase() }}</Button>
                                    </div>
                                </div>
                                <div class="assertWin" v-else>
                                    <p class="assert_question">{{ assertWinInfo.cname }}</p>
                                    <Button class="dark" @click="Setwin(assertWinInfo.cid, assertWinInfo.pid)">Да</Button>
                                    <Button class="dark" @click="assertWinChoice = !assertWinChoice">Назад</Button>
                                </div>
                            </transition>
                        </div>
                    </div>
                </transition>
            </Teleport>
            <Teleport to="body">
                <transition name="element">
                    <div class="addWin " v-if="!!getChips">
                        <div class="ForTheWin flex" ref="FTW">
                            <div class="icon mx-auto mb-10">
                                <Playerimage :source='`/images/${PlayerForWin?.login?.toLowerCase()}.jpg`'
                                    :alt='`${PlayerForWin?.user_name} куканит в покерямбе`' />
                            </div>
                            <transition name="page" mode="out-in">
                                <div class="combinationBlock" v-if="!assertGetChips">
                                    <p class="assert_question">Сколько добавляем чипиков?</p>
                                    <div class="combination">
                                        <Button class="dark" @click="addChips(PlayerForWin?.id, 50)">50</Button>
                                        <Button class="dark" @click="addChips(PlayerForWin?.id, 100)">100</Button>
                                    </div>
                                </div>
                            </transition>
                        </div>
                    </div>
                </transition>
            </Teleport>

        </div>
        <Button class="dark fit" @click="finishGame()">ЗАВЕРШИТЬ ИГРУ</Button>
        <div class="alert">{{ alert }}</div>
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
const alert = ref("")
const assertWinChoice = ref(false)
const getChips = ref(false)
const assertGetChips = ref(false)
const assertWinInfo = ref<{cname: string, cid: number, pid: number}>({cid: 0, cname: "", pid: 0})
const { data: players } = await useLazyFetch<PlayerForGame[]>(config.BASE_URL + `admin/getplayers/${route.params.id}`, {
    method: 'GET',
    credentials: 'include'
})
interface results {
    player: PlayerForGame,
    score: number,
}

const openAssertWinChoice = (cname: string, cid: number, pid: number | undefined) =>{
    if (pid != undefined){
        assertWinInfo.value.cid = cid
        assertWinInfo.value.pid = pid
        assertWinInfo.value.cname = cname
        console.log(assertWinInfo.value);
        assertWinChoice.value = true
    }
}
const addChips = (player_id: number | undefined, chips: number ) => {
    if (player_id !== undefined){
        useFetch<PlayerForGame[]>(config.BASE_URL + `admin/addchips/${player_id}`, {
            body: {
                chips: chips.toString(),
                game_id: route.params.id.toString(),
            },
            method: 'POST',
            credentials: 'include',
            async onResponse({ response }) {
                players.value = response._data
            }
        })
        playAudio(player_id, "lose")
        getChips.value = false

    }
}
const { data: combinations } = useLazyFetch<Combination[]>(config.BASE_URL + "admin/combinations", {
    credentials: 'include',
    method: 'GET',
    onRequestError({ error }) {
        err.value = error.message
    }
})
function OpenWinOption(id: number, action: number) {
    players.value?.forEach(e => {
        if (e.id === id) {
            PlayerForWin.value = e
        }
    })
    if (action === 2) {
        setWin.value = true
        window.scrollTo(0, 0)
        onClickOutside(FTW, (e) => {
            useHead({
                bodyAttrs: {
                    class: "",
                }
            })
            setWin.value = false
    
        })
        useHead({
            bodyAttrs: {
                class: !!setWin.value ? "locked" : "",
            }
        })
    }
    if (action === 1){
        getChips.value = true
        window.scrollTo(0, 0)
        onClickOutside(FTW, (e) => {
            useHead({
                bodyAttrs: {
                    class: "",
                }
            })
            getChips.value = false
        })
        useHead({
            bodyAttrs: {
                class: !!setWin.value ? "locked" : "",
            }
        })
    }
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
            onResponse({ response }) {
                console.log(toRaw(response._data));
                let results = toRaw(response._data)
                results.forEach((res: any) => {
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
        useHead({
            bodyAttrs: {
                class: "",
            }
        })
        playAudio(pid, "win")
        setWin.value = false
        assertWinChoice.value = false
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
function playAudio(pid:number, wl: string){
    players.value?.forEach(e => {
        if (e.id === pid) {
            const number = Math.floor(Math.random() * 4) + 1
            const audio = new Audio(`/sounds/${e.login?.toLowerCase()}-${number}-${wl}.mp3`)
            audio.play()
            return
        }
    });
}
</script>

<style lang="sass" scoped>
.gamePlayers
    justify-content: space-evenly
    flex-wrap: wrap
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
    overflow-x: hidden
    // height: 100vh
    inset: 0
    background-color: rgba(#333, .9)
    .ForTheWin
        background-color: white
        transition: all .2s ease-in
        padding: min(5rem, 5vw)
        display: grid
        // grid-auto-columns: 1fr, 1fr
        
        img
            width: 10rem
            height: 10rem
            object-fit: cover
            border-radius: 50%
    .combinationBlock
        display: grid
        grid-auto-rows: 1fr
        grid-template-columns: 1fr 1fr
        .combination
            display: grid
            
input
    max-width: 10rem
    &::placeholder
        font-size: .8rem
.user
    max-width: 45%
    margin-top: 2rem
</style>