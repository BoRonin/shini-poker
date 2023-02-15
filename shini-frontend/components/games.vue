<template>
    <div class="games-container main-section">
        <h1 class="sectionTitle">Игры</h1>
        <p v-if="error !== null">
        {{ error.message }}</p>
        <ul v-if="games !== null">
            <li v-for="game in games" :key="game.id">
            <p class="section-subtitle">{{ game.title }}</p>
            <p> {{ getDate(game.created_at) }}<br />
                {{ game.is_finished? "Игра закончена" : "Игра в процессе" }}</p>
            </li>
        </ul>
        <NuxtLink to="/games"><Button>Новая игра</Button></NuxtLink>
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