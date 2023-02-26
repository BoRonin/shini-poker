<template>
    <div class="games-container main-section">
        <h1 class="sectionTitle">Игры</h1>
        <p v-if="error !== null">
            {{ error.message }}</p>
        <ul v-if="games !== null">
            <li v-for="game in games" :key="game.id">
                <p class="section-subtitle">{{ game.title }}</p>
                <div class="game_info">
                    <div class="left">
                        <p> {{ getDate(game.created_at) }}<br />
                        <div v-if="!!game.is_finished" class="game-fit">
                            <p>Игра закончена.</p>
                            <p>{{ msToTime(game.duration) !== ' ' ? `${msToTime(game.duration)}` : '' }}</p>
                        </div>
                        <div v-else class="game-fit">
                            <NuxtLink :to="`/games/${game.id}`"><Button>Игра в процессе</Button></NuxtLink>
                        </div>
                        </p>
                    </div>
                    <div class="right" v-if="!!game.is_finished"><Button class="dark fit" @click="toggleModal(game)">Cтатистика
                            игры</Button></div>

                </div>
            </li>
        </ul>
        <Teleport to="body">
            <transition name="element">
                <Modalse v-if="!!modalOpen && gStat !== null && gStat.length > 0 && !!showgames" @close="closeModal()">
                    <p class="section-subtitle">{{ modalInfo.title }}</p>
                    <p>Множитель: {{ modalInfo.multiplier }}</p>
                    <p>{{ msToTime(modalInfo.duration) }}</p>
                    <div class="stats" >
                        <div class="stat" v-for="stat in gStat" :key="stat.login">
                            <Playerimage :source="`/images/${stat.login.toLowerCase()}.jpg`" :class="'small'" />
                            <p>Рубликов: <strong>{{ stat.money }}</strong></p>
                            <p>Побед: <strong>{{ stat.wins }}</strong></p>
                        </div>
                    </div>
                </Modalse>
            </transition>
        </Teleport>
        <NuxtLink to="/games"><Button class="dark left">Новая игра</Button></NuxtLink>
    </div>
</template>
<script setup lang="ts">
import { Game } from '@/types/Game'
interface player_stat {
    id: number,
    login: string,
    money: number,
    wins: number
}

const config = useAppConfig()
const modalInfo = ref<Game>({id:0, created_at: ""})
const modalOpen = ref(false)
const gStat = ref<player_stat[]>([])
const showgames = ref(false)
const { data: games, error } = useLazyFetch<Game[]>(config.BASE_URL + "admin/games", {
    credentials: 'include',
    method: 'GET',
})
const closeModal = () => {
    showgames.value = false
    modalOpen.value = false
}
function getDate(s: string) {
    var d = new Date(s)
    return d.toLocaleString()
}
function toggleModal(game: Game) {
    useFetch(config.BASE_URL + `admin/getstats/${game.id}`, {
        credentials:'include',
        method:'GET',
        onResponseError(){
            return
        },
        onResponse({response}){
            console.log(response._data);
            gStat.value = response._data.player_stats
        }
    }).then(()=>{
        showgames.value = true

    })
    modalInfo.value = game
    modalOpen.value = true
}
function msToTime(s: number | undefined) {
    if (s !== undefined) {
        // Pad to 2 or 3 digits, default is 2
        function pad(n: number, z: number = 2) {
            z = z || 2;
            return ('00' + n).slice(-z);
        }

        var ms = s % 1000;
        s = (s - ms) / 1000;
        var secs = s % 60;
        s = (s - secs) / 60;
        var mins = s % 60;
        var hrs = (s - mins) / 60;
        let hours = ""
        var minutes = ""
        if (pad(hrs) !== "00") {
            if (hrs.toString().endsWith("11") || hrs.toString().endsWith("12") || hrs.toString().endsWith("13") || hrs.toString().endsWith("14") || hrs.toString().endsWith("5") || hrs.toString().endsWith("6") || hrs.toString().endsWith("7") || hrs.toString().endsWith("8") || hrs.toString().endsWith("9") || hrs.toString().endsWith("0")) {
                hours = " часов "
            } else if (hrs.toString().endsWith("1")) {
                hours = " час "
            } else {
                hours = " часа "
            }
        }
        if (mins.toString().endsWith("11") || mins.toString().endsWith("12") || mins.toString().endsWith("13") || mins.toString().endsWith("14") || mins.toString().endsWith("5") || mins.toString().endsWith("6") || mins.toString().endsWith("7") || mins.toString().endsWith("8") || mins.toString().endsWith("9") || mins.toString().endsWith("0")) {
            minutes = " минут"
        } else if (mins.toString().endsWith("1")) {
            minutes = " минута"
        } else {
            minutes = " минуты"
        }
        return `${hrs > 0 ? hrs + hours : ''} ${mins > 0 ? mins + minutes : ''}`
    }

}
</script>
<style lang="sass">
.game_info
    display: flex
    flex-direction: row
    align-items: flex-end
    justify-content: space-between
    margin: .5rem 0
.game-fit
    width: fit-content
.stats
    display: flex
    flex-wrap: wrap
    justify-content: space-around
    gap: 2rem
</style>