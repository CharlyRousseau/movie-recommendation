<script lang="ts">
    import { Button } from "$shadcn/ui/button/index.js";

    export let currentPage: number;
    export let totalPages: number;
    export let fetchPage: (page: number) => void;

    const goToPage = (page: number) => {
        currentPage = page;
        fetchPage(currentPage);
    };
</script>

<div class="pagination mt-4">
    <Button
        variant="outline"
        on:click={() => goToPage(Math.max(1, currentPage - 1))}>Previous</Button
    >

    <Button variant="ghost" on:click={() => goToPage(currentPage)}
        >{currentPage}</Button
    >

    {#if currentPage + 1 <= totalPages}
        <Button variant="outline" on:click={() => goToPage(currentPage + 1)}
            >{currentPage + 1}</Button
        >
    {/if}

    {#if currentPage + 2 <= totalPages}
        <Button variant="outline" on:click={() => goToPage(currentPage + 2)}
            >{currentPage + 2}</Button
        >
    {/if}

    {#if currentPage < totalPages}
        <span>...</span>
        <Button variant="outline" on:click={() => goToPage(totalPages)}
            >{totalPages}</Button
        >
    {/if}

    <Button
        variant="outline"
        on:click={() => goToPage(Math.min(totalPages, currentPage + 1))}
        >Next</Button
    >
</div>
