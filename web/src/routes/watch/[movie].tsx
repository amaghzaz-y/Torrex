import { useParams } from 'solid-start'
import MoviePlayer from '~/components/movie.player'

export default function Index() {
    const params = useParams()
    return (
        <div class="flex flex-col flex-gap-6 w-4xl self-center justify-center items-center p-10">
            <text class="w-full font-size-6 font-bold">
                Watching Red Notice - 2021
            </text>
            <div class="w-4xl">
                <MoviePlayer source="http://localhost:4000/stream/we8xY0eAVUWB/index.m3u8" />
            </div>
        </div>
    )
}
