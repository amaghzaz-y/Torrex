import { RiWeatherFireFill } from 'solid-icons/ri'
export default function NavBar() {
    return (
        <div class="flex justify-between items-center bg-torrex-primary max-w-3xl px-4 py-2 rounded-3 cursor-default">
            <text class="flex flex-gap-0.5 items-center w-20 text-torrex-accent font-size-8 font-900">
                <RiWeatherFireFill size={'1.8rem'} />
                Torrex
            </text>
            <div class="flex flex-gap-7">
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
