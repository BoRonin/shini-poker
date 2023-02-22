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
                    <div class="right" v-if="!!game.is_finished"><Button class="dark fit" @click="toggleModal()">Cтатистика
                            игры</Button></div>

                </div>
            </li>
        </ul>
        <Teleport to="body">
            <transition name="element">
                <Modalse v-if="modalOpen" @close="modalOpen = !modalOpen">
                    <p>sd</p>
                </Modalse>
            </transition>
        </Teleport>
        <NuxtLink to="/games"><Button class="dark left">Новая игра</Button></NuxtLink>
    </div>
</template>
<script setup lang="ts">
import { Game } from '@/types/Game'
import { onClickOutside } from '@vueuse/core'
const config = useAppConfig()
const modalOpen = ref(false)
const { data: games, error } = useLazyFetch<Game[]>(config.BASE_URL + "admin/games", {
    credentials: 'include',
    method: 'GET',
})
function getDate(s: string) {
    var d = new Date(s)
    return d.toLocaleString()
}
function toggleModal() {
    modalOpen.value = !modalOpen.value
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
    flex-direction: column
    justify-content: space-between
    margin: .5rem 0
.game-fit
    width: fit-content
</style>