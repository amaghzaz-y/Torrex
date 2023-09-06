import MovieCard from '~/components/movie.card'

export default function Search() {
    return (
        <div class="flex flex-col w-4xl self-center flex-gap-8">
            <div class="flex flex-gap-3">
                <input
                    placeholder="Search Here..."
                    class=" text-white p-3 w-full
                     font-size-5 bg-torrex-secondary rounded-md outline-none border-none "
                />
                <button
                    class="text-torrex-text p-3 hover:bg-torrex-accent
                     font-size-4 bg-torrex-secondary rounded-md outline-none border-none">
                    Search
                </button>
            </div>
            {/* <MovieCard stream="/" /> */}
        </div>
    )
}
