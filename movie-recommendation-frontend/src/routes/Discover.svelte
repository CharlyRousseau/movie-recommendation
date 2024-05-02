<script lang="ts">
    import MovieCard from "$lib/../components/movieCard.svelte";
    import Navbar from "$lib/../components/navbar.svelte";
    import { Button } from "$lib/components/ui/button/index.js";

    let className: string | undefined | null = undefined;
    import { onMount } from "svelte";
    type Movie = {
        overview: string;
        backdrop_path: string;
        id: number;
        original_title: string;
        release_date: string;
        genre_ids: number[];
        vote_average: number;
    };
    let currentPage = 1;
    let totalPages = 0;
    let movies: Movie[] = [];
    const fetchPage = async (page: number) => {
        const options = {
            method: "GET",
            headers: {
                accept: "application/json",
                Authorization:
                    "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIyYjc2YmJmNjVhOWUwZjJmOTJjMjVhZDAyMjM0NzVmMCIsInN1YiI6IjY2MzBlZmU2NmEzMDBiMDEyMjYwOWNiOSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.Emhv8C3hQseouTYE8SYFAqk5Ojrl5Aw2Zd_h4sp271Y",
            },
        };

        const response = await fetch(
            `https://api.themoviedb.org/3/discover/movie?page=${page}`,
            options,
        );
        const data = await response.json();
        movies = data.results;
        totalPages = data.total_pages;
    };

    onMount(() => {
        fetchPage(currentPage);
    });
</script>

<div
    class="container space-x-4 lg:space-x-6 md:space-x-8 px-4 py-2 md:px-6 md:py-4 lg:px-8 lg:py-6"
>
    <Navbar />
    <h1 class="text-2xl font-semibold tracking-tight">Filter</h1>
    <div class="grid grid-cols-3 gap-4">
        {#each movies as movie (movie.id)}
            <MovieCard
                title={movie.original_title}
                date={movie.release_date}
                imageSrc={movie.backdrop_path}
                grade={movie.vote_average}
            />
        {/each}
    </div>
    <div class="pagination mt-4">
        <Button
            variant="outline"
            on:click={() => {
                currentPage = Math.max(1, currentPage - 1);
                fetchPage(currentPage);
            }}
            >Previous
        </Button>

        <Button
            variant="ghost"
            on:click={() => {
                fetchPage(currentPage);
            }}>{currentPage}</Button
        >

        {#if currentPage + 1 <= totalPages}
            <Button
                variant="outline"
                on:click={() => {
                    currentPage = currentPage + 1;
                    fetchPage(currentPage);
                }}>{currentPage + 1}</Button
            >
        {/if}

        {#if currentPage + 2 <= totalPages}
            <Button
                variant="outline"
                on:click={() => {
                    currentPage = currentPage + 2;
                    fetchPage(currentPage);
                }}>{currentPage + 2}</Button
            >
        {/if}

        {#if currentPage < totalPages}
            <span>...</span>
            <Button
                variant="outline"
                on:click={() => {
                    currentPage = totalPages;
                    fetchPage(currentPage);
                }}>{totalPages}</Button
            >
        {/if}

        <Button
            variant="outline"
            on:click={() => {
                currentPage = Math.min(totalPages, currentPage + 1);
                fetchPage(currentPage);
            }}>Next</Button
        >
    </div>
</div>
