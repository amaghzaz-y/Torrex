import { Title } from 'solid-start'
import { HttpStatusCode } from 'solid-start/server'

export default function NotFound() {
    return (
        <main class="flex flex-col h-full w-full">
            <Title>Not Found</Title>
            <HttpStatusCode code={404} />
            <text class="font-black font-size-26 text-center">
                Bruuh... <br /> Page Not Found
            </text>
        </main>
    )
}
