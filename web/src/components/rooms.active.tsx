import { For, Show, createResource } from 'solid-js'
import MovieCard from './movie.card'
import { Room } from '~/types/room'
import { HOSTPORT } from '~/config'

export default function () {
    const [data] = createResource<Room[]>(async () => {
        return await (await fetch(`${HOSTPORT}/rooms`)).json()
    })

    return (
        <div id="active" class="flex flex-col flex-gap-5">
            <text class="text-torrex-text font-size-3xl font-semibold">
                Active Sessions
            </text>
            <Show when={!data.loading} fallback={<>Loading...</>}>
                <For each={data()}>
                    {(room) => (
                        <MovieCard movie={room.movie} stream={room.id} />
                    )}
                </For>
            </Show>
        </div>
    )
}
