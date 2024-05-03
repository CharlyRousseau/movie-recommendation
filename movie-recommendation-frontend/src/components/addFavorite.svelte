<script lang="ts">
    import Pagination from "$components/pagination.svelte";
    import type { Movie } from "$types/movie";
    import { Button } from "$shadcn/ui/button/index.js";
    import { onMount } from "svelte";

    const API_URL = import.meta.env.VITE_API_URL;

    export let itemid: number;
    let display = false;
    const token = localStorage.getItem("jwt");

    const addFavorite = async () => {
        const options = {
            method: "POST",
            headers: {
                accept: "application/json",
                Authorization: `Bearer ${token}`,
            },
        };

        const response = await fetch(
            `${API_URL}/favorites?itemId=${itemid}`,
            options,
        );
        const data = await response.json();
    };

    onMount(async () => {
        if (token) {
            try {
                const response = await fetch(`${API_URL}/me`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                if (response.ok) {
                    display = true;
                } else {
                    localStorage.removeItem("jwt");
                    navigate("/login");
                }
            } catch (err) {
                console.error(err);
            }
        }
    });
</script>

{#if display}
    <Button variant="outline" on:click={addFavorite}>Add to favs</Button>
{/if}
