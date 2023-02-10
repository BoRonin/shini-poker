export interface Game {
    id: number,
    title?: string,
    multiplier?: number,
    is_finished?: boolean
}
export interface Player{
    id: number,
    game_id?: number,
    chips?: number,
    chips_final?: number
}
export interface Combination{
    id: number,
    title: string
}
export interface Win {
    player: Player,
    combination: Combination
}