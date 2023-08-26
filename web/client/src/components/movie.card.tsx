import { AiFillPlayCircle, AiOutlineYoutube } from 'solid-icons/ai'
import {
    RiMediaPlayFill,
    RiMediaPlayMiniFill,
    RiMediaPlayMiniLine,
} from 'solid-icons/ri'
export default function () {
    return (
        <div class="opacity-80 flex w-full h-60 bg-torrex-secondary rounded-3 hover:opacity-100 cursor-default">
            <img
                src="https://a.ltrbxd.com/resized/film-poster/8/4/0/8/6/3/840863-scrapper-0-500-0-750-crop.jpg?v=e9c47f6919"
                class="h-full bg-gray w-55 rounded-3"
            />
            <div class="flex justify-between flex-col p-5 flex-gap-2 w-full">
                <div class="flex items-center font-bold flex-gap-10 font-size-6">
                    Scrapper
                    <text class="font-size-5 font-300">2023</text>
                </div>
                <text class="font-size-4 text-justify w-full line-clamp-3">
                    Living alone since her beloved mum died, 12-year-old
                    Georgie fills the flat they shared with her own special
                    magic. But when her absent father Jason turns up out of
                    the blue, she’s forced to confront reality. Living
                    alone since her beloved mum died, 12-year-old Georgie
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
                    <button
                        class="flex items-center p-2 font-size-4 rounded-2 text-torrex-text
                            bg-torrex-secondary outline-none 
                             border-none fill-torrex-text
                             hover:text-green hover:fill-green
                             cursor-pointer">
                        Join Stream
                        <RiMediaPlayMiniLine font-size="1.5rem" />
                    </button>
                </div>
            </div>
        </div>
    )
}
