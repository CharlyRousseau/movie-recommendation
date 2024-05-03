<script lang="ts">
    import { onMount } from "svelte";
    import { link } from "svelte-routing";
    import { navigate } from "svelte-routing";
    import { cn } from "$lib/utils.js";

    let className: string | undefined | null = undefined;
    export { className as class };

    const API_URL = import.meta.env.VITE_API_URL;

    let currentPath = "";
    let username = "";

    onMount(() => {
        currentPath = window.location.pathname;
        window.addEventListener("popstate", () => {
            currentPath = window.location.pathname;
        });
    });

    onMount(async () => {
        const token = localStorage.getItem("jwt");
        console.log("oui");
        if (token) {
            console.log(token);
            try {
                const response = await fetch(`${API_URL}/me`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                if (response.ok) {
                    const data = await response.json();
                    username = data.Username;
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

<nav
    class={cn(
        "flex items-center justify-between space-x-4 lg:space-x-6 md:space-x-8 px-4 py-2 md:px-6 md:py-4 lg:px-8 lg:py-6 bg-white dark:bg-gray-800  dark:shadow-none rounded-lg dark:border dark:border-gray-700 transition-colors",
        className,
    )}
>
    <div class="space-x-4">
        <a
            href="/"
            use:link
            class={currentPath === "/"
                ? "text-sm font-medium transition-colors hover:text-primary"
                : "text-sm font-medium text-muted-foreground transition-colors hover:text-primary"}
        >
            Overview
        </a>

        <a
            href="/discover"
            use:link
            class={currentPath === "/discover"
                ? "text-sm font-medium transition-colors hover:text-primary"
                : "text-sm font-medium text-muted-foreground transition-colors hover:text-primary"}
        >
            Discover
        </a>
        {#if username}
            <a
                href="/favorites"
                use:link
                class={currentPath === "/favorites"
                    ? "text-sm font-medium transition-colors hover:text-primary"
                    : "text-sm font-medium text-muted-foreground transition-colors hover:text-primary"}
            >
                Favs
            </a>
        {/if}
    </div>
    <div>
        {#if username}
            <div>Welcome, {username}!</div>
        {:else}
            <a
                href="/login"
                use:link
                class={currentPath === "/login"
                    ? "text-sm font-medium transition-colors hover:text-primary"
                    : "text-sm font-medium text-muted-foreground transition-colors hover:text-primary"}
            >
                Login
            </a>
        {/if}
    </div>
</nav>
