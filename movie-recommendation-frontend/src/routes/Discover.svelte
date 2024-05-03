<script lang="ts">
    import MovieCardList from "$components/movieCardList.svelte";
    import Navbar from "$components/navbar.svelte";
    import Pagination from "$components/pagination.svelte";
    import type { Movie } from "$types/movie";

    const TOKEN = import.meta.env.TMDB_TOKEN;

    let className: string | undefined | null = undefined;
    import { onMount } from "svelte";

    let currentPage = 1;
    let totalPages = 0;
    let movies: Movie[] = [];

    const fetchPage = async (page: number) => {
        const options = {
            method: "GET",
            headers: {
                accept: "application/json",
                Authorization: `Bearer ${TOKEN}`,
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
    <MovieCardList {movies} />
    <Pagination {fetchPage} {currentPage} {totalPages} />
</div>
