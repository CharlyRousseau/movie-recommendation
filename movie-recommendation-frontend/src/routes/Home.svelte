<script lang="ts">
    import MovieCard from "$lib/../components/movieCard.svelte";
    import { cn } from "$lib/utils.js";
    let className: string | undefined | null = undefined;
    export { className as class };
</script>

<div
    class="container space-x-4 lg:space-x-6 md:space-x-8 px-4 py-2 md:px-6 md:py-4 lg:px-8 lg:py-6"
>
    <nav
        class={cn(
            "flex items-center justify-between space-x-4 lg:space-x-6 md:space-x-8 px-4 py-2 md:px-6 md:py-4 lg:px-8 lg:py-6 bg-white dark:bg-gray-800  dark:shadow-none rounded-lg dark:border dark:border-gray-700 transition-colors",
            className,
        )}
    >
        <div class="space-x-4">
            <a
                href="/examples/dashboard"
                class="text-sm font-medium transition-colors hover:text-primary"
            >
                Overview
            </a>

            <a
                href="/examples/dashboard"
                class="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
            >
                Favorites
            </a>
        </div>
        <div>
            <a
                href="/login"
                class="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
            >
                Login
            </a>
        </div>
    </nav>
    <h1 class="text-2xl font-semibold tracking-tight">Trending</h1>
    <MovieCard
        title="Film Title"
        date="Année du film"
        imageSrc="path/to/image.jpg"
        grade="Your note here"
    />
</div>
