<template>
    <div class="main-section stat-section">
        <p class="sectionTitle">Статистика</p>
        <ul class="stats">
            <li v-for="stat in stats" :key="stat.login">
                <Playerimage :source="`/images/${stat.login.toLowerCase()}.jpg`" :class="'small'"/>
             <p class="namechips"><strong>{{ stat.chips_after - stat.chips_taken }}</strong> чипиков</p>
             <ul>
                <p class="section-subtitle">Комбинации</p>
                <li v-for="count in stat.name_count_stats" :key="count.combination">
                    {{ count.combination }} - <strong>{{ count.count }}</strong>
                
                </li>
             </ul>
            </li>
        </ul>
    </div>
</template>

<script setup lang="ts">
import {Stat} from '@/types/Stats'
const config = useAppConfig()
const {data: stats} = await useLazyFetch<Stat[]>(config.BASE_URL + "admin/getstats",{
    method:'GET',
    credentials:'include',
    
})
</script>

<style lang="sass" scoped>
.stat-section
    .stats
        display: flex
        gap: 1rem
</style>