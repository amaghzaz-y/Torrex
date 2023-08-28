import { A } from '@solidjs/router'
import { RiWeatherFireFill } from 'solid-icons/ri'
export default function () {
    return (
        <div
            class="my-1 place-self-center flex justify-between 
                    items-center  opacity-100 w-4xl 
                    rounded-3 cursor-default fixed z-12 bg-torrex-navbar">
            <A
                href="/"
                class="flex flex-gap-0.5 items-center text-torrex-accent font-size-8 font-900 px-5">
                <RiWeatherFireFill size={'1.8rem'} />
                Torrex
            </A>
            <div class="flex flex-gap-7 px-10">
                <A href="/" class="font-size-5 font-bold hover:text-white">
                    Home
                </A>
                <A href="/" class="font-size-5 font-bold hover:text-white">
                    Active
                </A>
                <A href="/" class="font-size-5 font-bold hover:text-white">
                    Upcoming
                </A>
                <A href="/" class="font-size-5 font-bold hover:text-white">
                    Calendar
                </A>
                <A href="/" class="font-size-5 font-bold hover:text-white">
                    Login
                </A>
            </div>
        </div>
    )
}
