<script lang="ts">
    import Navbar from "$components/navbar.svelte";
    import { navigate } from "svelte-routing";
    import { onMount, getContext } from "svelte";

    const API_URL = import.meta.env.VITE_API_URL;
    let data: any = null;
    let className: string | undefined | null = undefined;

    // TO DO implements movie fetch based on ID
    onMount(async () => {
        const token = localStorage.getItem("jwt");
        console.log(token);
        if (token) {
            try {
                const options = {
                    method: "GET",
                    headers: {
                        accept: "application/json",
                        Authorization: `Bearer ${token}`,
                    },
                };

                const response = await fetch(`${API_URL}/favorites`, options);
                data = await response.json();
            } catch (err) {
                console.error(err);
            }
        } else {
            navigate("/");
        }
    });
</script>

<div
    class="container space-x-4 lg:space-x-6 md:space-x-8 px-4 py-2 md:px-6 md:py-4 lg:px-8 lg:py-6"
>
    <Navbar />
    <h1 class="text-2xl font-semibold tracking-tight">Favorites</h1>
    <pre>{data && JSON.stringify(data, null, 2)}</pre>
</div>
