<script lang="ts">
    import Loader from "lucide-svelte/icons/loader";
    import { Button } from "$shadcn/ui/button/index.js";
    import { Input } from "$shadcn/ui/input/index.js";
    import { Label } from "$shadcn/ui/label/index.js";
    import Notification from "$components/notification.svelte";
    import { cn } from "$lib/utils.js";
    import { onSubmit } from "$lib/formHandler";

    let className: string | undefined | null = undefined;
    export { className as class };

    let isLoading = false;
    let email = "";
    let password = "";
    let error = "";
    let success = "";

    async function handleSubmit(
        url: string,
        body: any,
        successMessage: string,
        redirectUrl: string,
    ) {
        isLoading = true;
        error = "";
        const result = await onSubmit(url, body, successMessage, redirectUrl);
        isLoading = false;
        if (result.error) {
            error = result.error;
        } else {
            if (result.success) {
                success = result.success;
                localStorage.setItem("jwt", result.token);
            }
        }
    }
</script>

<div class={cn("grid gap-6", className)} {...$$restProps}>
    <form
        on:submit|preventDefault={() =>
            handleSubmit(
                "/login",
                { email, password },
                "Successfully logged in! Redirecting...",
                "/",
            )}
    >
        <div class="grid gap-2">
            <div class="grid gap-1">
                <Label class="sr-only" for="email">Email</Label>
                <Input
                    id="email"
                    bind:value={email}
                    placeholder="name@example.com"
                    type="email"
                    autocapitalize="none"
                    autocomplete="email"
                    autocorrect="off"
                    disabled={isLoading}
                />
                <Input
                    bind:value={password}
                    id="password"
                    type="password"
                    autocapitalize="none"
                    autocomplete="current-password"
                    autocorrect="off"
                    disabled={isLoading}
                />
                {#if error}
                    <Notification type="error" message={error} />
                {/if}
                {#if success}
                    <Notification type="success" message={success} />
                {/if}
            </div>
            <Button type="submit" disabled={isLoading}>
                {#if isLoading}
                    <Loader class="loader animate-spin" />
                {/if}
                Connect
            </Button>
        </div>
    </form>
</div>
