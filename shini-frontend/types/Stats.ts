export interface Name_count_stats{
    combination: string,
    count: number,
}
export interface Stat {
    username: string,
    login: string,
    chips_after: number,
    chips_taken: number,
    name_count_stats: Name_count_stats[],

}