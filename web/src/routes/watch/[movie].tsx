import { Show, createResource } from 'solid-js'
import { useParams } from 'solid-start'
import MoviePlayer from '~/components/movie.player'
import { HOSTPORT } from '~/config'
import { Room } from '~/types/room'

export default function Index() {
    const params = useParams()
    let path = `${HOSTPORT}/stream/${params['movie']}/index.m3u8`
    const [data] = createResource<Room>(async () => {
        let x = await (
            await fetch(`${HOSTPORT}/rooms/${params['movie']}`)
        ).json()
        console.log(x)
        return x
    })

    return (
        <div class="flex flex-col flex-gap-6 w-4xl self-center justify-center items-center p-10">
            <text class="w-full font-size-6 font-bold">
                <Show when={!data.loading} fallback={<>Loading...</>}>
                    {data()?.movie.title} {data()?.movie.year}
                </Show>
            </text>
            <div class="w-4xl">
                <MoviePlayer source={path} />
            </div>
        </div>
    )
}
