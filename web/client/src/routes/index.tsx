import SessionActive from '~/components/session.active'
import Tagline from '~/components/tagline'

export default function Home() {
    return (
        <div class="flex flex-col w-4xl self-center flex-gap-8">
            <Tagline />
            <SessionActive />
        </div>
    )
}
