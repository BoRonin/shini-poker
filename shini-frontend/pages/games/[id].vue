<template>
    <div class="pickPlayers mt-10 flex ">
        <div class="user" v-for="player in players" :key="player?.id">
            <div class="icon">
                <img :src='`/images/${player.login?.toLowerCase()}.jpg`' :alt='`${player.name} куканит в покерямбе`'>
            </div>
            <p>Чипиков: {{ player.chips }}</p>
            <div class="add_chips">
                <button @click="addChips(player.id)">{{player.id}} Добавить чипики</button>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { PlayerForGame } from '~~/types/Game';
const config = useAppConfig()
const route = useRoute()

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

</script>

<style lang="sass" scoped>

</style>