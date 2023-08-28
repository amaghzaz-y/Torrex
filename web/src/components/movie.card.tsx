import { A } from '@solidjs/router'
import { AiOutlineYoutube } from 'solid-icons/ai'
import { RiMediaPlayMiniLine } from 'solid-icons/ri'
import { Movie } from '~/types/movie'
export default function ({
    movie,
    stream,
}: {
    movie: Movie
    stream: string
}) {
    return (
        <div class="opacity-80 flex w-full h-60 bg-torrex-secondary rounded-3 hover:opacity-100 cursor-default">
            <img src={movie.bgimg} class="h-full bg-gray w-92 rounded-3" />
            <div class="flex justify-between flex-col p-5 flex-gap-2 w-full">
                <div class="flex items-center font-bold flex-gap-10 font-size-6">
                    {movie.title}
                    <text class="font-size-5 font-300">{movie.year}</text>
                </div>
                <text class="font-size-4 text-justify w-full line-clamp-3">
                    {movie.description}
                </text>
                <div class="flex flex-gap-2 place-self-end ">
                    <button
                        class="flex flex-gap-1 items-center p-2 font-size-4 rounded-2 text-torrex-text
                        bg-torrex-secondary outline-none 
                         border-none
                         fill-torrex-text
                         cursor-pointer
                         opacity-80
                         hover:opacity-100
                         ">
                        Trailer
                        <AiOutlineYoutube font-size="1.5rem" />
                    </button>
                    <A
                        href={'/watch/' + stream}
                        class="flex items-center p-2 font-size-4 rounded-2 text-torrex-text
                            bg-torrex-secondary outline-none 
                             border-none fill-torrex-text
                             hover:text-green hover:fill-green
                             cursor-pointer">
                        Join Stream
                        <RiMediaPlayMiniLine font-size="1.5rem" />
                    </A>
                </div>
            </div>
        </div>
    )
}
