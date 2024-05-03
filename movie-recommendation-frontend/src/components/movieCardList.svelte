<script lang="ts">
    import MovieCard from "$components/movieCard.svelte";
    import type { Movie } from "$types/movie";
    export let movies: Movie[];
</script>

<div
    class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4"
>
    {#each movies as movie (movie.id)}
        <MovieCard
            title={movie.original_title}
            date={movie.release_date}
            imageSrc={movie.backdrop_path}
            grade={movie.vote_average}
        />
    {/each}
</div>
