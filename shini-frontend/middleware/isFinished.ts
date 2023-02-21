
export default defineNuxtRouteMiddleware((to, from) => {
    const config = useAppConfig()
    useFetch(config.BASE_URL + `admin/games/isfinished/${to.params.id}`, {
        credentials:'include',
        method:'GET',
        onResponseError(){
            navigateTo('/')
        },
        onResponse({response}){
            if (response._data === true) {
                navigateTo('/')
            } 
        },
    })
})