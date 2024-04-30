<script lang="ts">
    import Loader from "lucide-svelte/icons/loader";
    import { Button } from "$lib/components/ui/button/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import { cn } from "$lib/utils.js";

    let className: string | undefined | null = undefined;
    export { className as class };

    let isLoading = false;
    async function onSubmit() {
        isLoading = true;

        setTimeout(() => {
            isLoading = false;
        }, 3000);
    }
</script>

<div class={cn("grid gap-6", className)} {...$$restProps}>
    <form on:submit|preventDefault={onSubmit}>
        <div class="grid gap-2">
            <div class="grid gap-1">
                <Label class="sr-only" for="email">Email</Label>
                <Input
                    id="email"
                    placeholder="name@example.com"
                    type="email"
                    autocapitalize="none"
                    autocomplete="email"
                    autocorrect="off"
                    disabled={isLoading}
                />
                <Input
                    id="password"
                    type="password"
                    autocapitalize="none"
                    autocorrect="off"
                    disabled={isLoading}
                />
            </div>
            <Button type="submit" disabled={isLoading}>
                {#if isLoading}
                    <Loader />
                {/if}
                Create Account
            </Button>
        </div>
    </form>
    <div class="relative">
        <div class="absolute inset-0 flex items-center">
            <span class="w-full border-t" />
        </div>
    </div>
</div>
