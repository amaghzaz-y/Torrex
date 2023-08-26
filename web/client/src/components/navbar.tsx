import { RiWeatherFireFill } from 'solid-icons/ri'
export default function () {
    return (
        <div
            class="my-1 place-self-center flex justify-between 
                    items-center  opacity-100 w-4xl 
                    rounded-3 cursor-default fixed z-12 bg-torrex-navbar">
            <text class="flex flex-gap-0.5 items-center text-torrex-accent font-size-8 font-900 px-5">
                <RiWeatherFireFill size={'1.8rem'} />
                Torrex
            </text>
            <div class="flex flex-gap-7 px-10">
                <text class="font-size-5 font-bold hover:text-white">
                    Home
                </text>
                <text class="font-size-5 font-bold hover:text-white">
                    Active
                </text>
                <text class="font-size-5 font-bold hover:text-white">
                    Upcoming
                </text>
                <text class="font-size-5 font-bold hover:text-white">
                    Calendar
                </text>
                <text class="font-size-5 font-bold hover:text-white">
                    Login
                </text>
            </div>
        </div>
    )
}
