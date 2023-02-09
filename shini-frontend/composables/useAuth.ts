import {User} from '@/types/User'

 export const useAuth = () =>{
    const config = useRuntimeConfig()
    const user = useState<User | null>('user', ()=> {
        return null})
    return{
        user,
    }
}