import RoomsActive from '~/components/rooms.active'
import RoomsNext from '~/components/rooms.next'
import Tagline from '~/components/tagline'

export default function Home() {
    return (
        <div class="flex flex-col w-4xl self-center flex-gap-8">
            <Tagline />
            <RoomsActive />
            <RoomsNext />
        </div>
    )
}
