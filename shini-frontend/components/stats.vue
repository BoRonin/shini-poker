<template>
    <div class="main-section stat-section">
        <p class="sectionTitle">Статистика</p>
        <ul class="stats">
            <li v-for="stat in stats" :key="stat.login">
                <Playerimage :source="`/images/${stat.login.toLowerCase()}.jpg`" :class="'small'" />

                <p class="namechips mb-5"><strong>{{ stat.chips_after - stat.chips_taken }}</strong> {{getChipString(stat.chips_after - stat.chips_taken)}}</p>

                <p class="section-subtitle mb-2">Комбинации</p>
                <p class="wins"> {{ getWins(stat.name_count_stats) }} </p>
                <ul>
                    <li v-for="count in stat.name_count_stats" :key="count.combination">
                        <strong>{{ count.count }}</strong>: {{ count.combination }} 
                    </li>
                </ul>

            </li>
        </ul>
    </div>
</template>

<script setup lang="ts">
import { Stat, Name_count_stats } from '@/types/Stats'
const config = useAppConfig()
const { data: stats } = await useLazyFetch<Stat[]>(config.BASE_URL + "admin/getstats", {
    method: 'GET',
    credentials: 'include',

})
const getWins = (all: Name_count_stats[]) => {
    let count = 0
    all.forEach(e => {
        count += e.count
    });
    let result = count.toString()
    if (result.endsWith("11") || result.endsWith("12") || result.endsWith("13") || result.endsWith("14")|| result.endsWith("5") || result.endsWith("6") || result.endsWith("7") || result.endsWith("8")|| result.endsWith("9") || result.endsWith("0")) {
        return count + " побед"
    } else if (result.endsWith("1")){
        return count + " победа"
    } else {
        return count + " победы"
    }
}

const getChipString = (chips: number )=>{
    let result = chips.toString()
    if (result.endsWith("11") || result.endsWith("12") || result.endsWith("13") || result.endsWith("14")|| result.endsWith("5") || result.endsWith("6") || result.endsWith("7") || result.endsWith("8")|| result.endsWith("9") || result.endsWith("0")) {
        return "чипиков"
    } else if (result.endsWith("1")){
        return "чипик"
    } else {
        return "чипика"
    }
}
</script>

<style lang="sass" scoped>
.stat-section
    .stats
        display: flex
        flex-wrap: wrap
        gap: 1rem
        justify-content: space-around
        .wins
            font-size: 0.9rem
            color: #999
</style>