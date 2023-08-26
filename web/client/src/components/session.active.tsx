import MovieCard from './movie.card'

export default function () {
    return (
        <div class="flex flex-col flex-gap-5">
            <text class="text-torrex-text font-size-3xl font-semibold">
                Active Sessions
            </text>
            <MovieCard />
            <MovieCard />
            <MovieCard />
            <MovieCard />
            <MovieCard />
            <MovieCard />
            <MovieCard />
        </div>
    )
}
