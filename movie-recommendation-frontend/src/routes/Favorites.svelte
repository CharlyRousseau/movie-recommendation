<script lang="ts">
    import MovieCard from "$lib/../components/movieCard.svelte";
    import Navbar from "$lib/../components/navbar.svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import { navigate } from "svelte-routing";
    import { onMount, getContext } from "svelte";

    const API_URL = import.meta.env.VITE_API_URL;

    type Movie = {
        overview: string;
        backdrop_path: string;
        id: number;
        original_title: string;
        release_date: string;
        genre_ids: number[];
        vote_average: number;
    };
    let className: string | undefined | null = undefined;
    onMount(() => {
        const token = localStorage.getItem("jwt");

        if (!token) {
            navigate("/");
        }
    });

    onMount(async () => {
        const token = localStorage.getItem("jwt");

        try {
            const options = {
                method: "GET",
                headers: {
                    accept: "application/json",
                    Authorization: `Bearer ${token}`,
                },
            };

            const response = await fetch(`${API_URL}/favorites`, options);
            const data = await response.json();
            console.log(data);
        } catch (err) {
            console.error(err);
        }
    });
</script>

<div
    class="container space-x-4 lg:space-x-6 md:space-x-8 px-4 py-2 md:px-6 md:py-4 lg:px-8 lg:py-6"
>
    <Navbar />
    <h1 class="text-2xl font-semibold tracking-tight">Favorites</h1>
</div>
