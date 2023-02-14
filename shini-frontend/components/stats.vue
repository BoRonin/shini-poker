<template>
    <div class="main-section">
        <p class="sectionTitle">Статистика</p>
        <ul class="stats">
            <p>Всего заработанных чипиков: </p>
            <li v-for="stat in stats" :key="stat.user">
             {{ stat.user }} - <strong>{{ stat.chips }}</strong>
            </li>
        </ul>
    </div>
</template>

<script setup lang="ts">
import {Stat} from '@/types/Stats'
const config = useAppConfig()
const {data: stats} = await useFetch<Stat[]>(config.BASE_URL + "admin/getstats",{
    method:'GET',
    credentials:'include',
    
})
</script>