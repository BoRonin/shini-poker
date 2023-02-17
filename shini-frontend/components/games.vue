<template>
    <div class="games-container main-section">
        <h1 class="sectionTitle">Игры</h1>
        <p v-if="error !== null">
        {{ error.message }}</p>
        <ul v-if="games !== null">
            <li v-for="game in games" :key="game.id">
            <p class="section-subtitle">{{ game.title }}</p>
            <p> {{ getDate(game.created_at) }}<br />
                <div v-if="!!game.is_finished" class="game-fit">Игра закончена</div>
                <div v-else class="game-fit"><NuxtLink :to="`/games/${game.id}`" ><Button>Игра в процессе</Button></NuxtLink></div>
                </p>
            </li>
        </ul>
        <NuxtLink to="/games"><Button class="dark left">Новая игра</Button></NuxtLink>
    </div>
</template>
<script setup lang="ts">
    import {Game} from '@/types/Game'
    const config = useAppConfig()

    const {data: games, error} = useLazyFetch<Game[]>(config.BASE_URL + "admin/games",{
        credentials:'include',
        method:'GET',
    })
    function getDate(s: string) {
    var d = new Date(s)
    return d.toLocaleString()
}
</script>
<style lang="sass">
.game-fit
    width: fit-content
</style>