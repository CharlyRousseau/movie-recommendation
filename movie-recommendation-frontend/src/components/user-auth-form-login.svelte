<script lang="ts">
    import Loader from "lucide-svelte/icons/loader";
    import { Button } from "$lib/components/ui/button/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import { cn } from "$lib/utils.js";
    import { navigate } from "svelte-routing";
    import * as Alert from "$lib/components/ui/alert";

    const API_URL = import.meta.env.VITE_API_URL;

    let className: string | undefined | null = undefined;
    export { className as class };

    let isLoading = false;
    let email = "";
    let password = "";
    let error = "";
    let success = "";

    async function onSubmit() {
        isLoading = true;
        error = "";

        try {
            const response = await fetch(`${API_URL}/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ email, password }),
            });

            if (!response.ok) {
                const data = await response.json();
                error = data.message;
            } else {
                success = "Successfully logged in! Redirecting...";
                setTimeout(() => {
                    navigate("/");
                }, 2000);
            }
        } catch (err) {
            error =
                err instanceof Error
                    ? err.message
                    : "An unknown error occurred";
        } finally {
            isLoading = false;
        }
    }
</script>

<div class={cn("grid gap-6", className)} {...$$restProps}>
    <form on:submit|preventDefault={onSubmit}>
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
                    <Alert.Root>
                        <Alert.Title class="text-red-500">Error</Alert.Title>
                        <Alert.Description>
                            {error}
                        </Alert.Description>
                    </Alert.Root>
                {/if}
                {#if success}
                    <Alert.Root>
                        <Alert.Title class="text-green-500">Success</Alert.Title
                        >
                        <Alert.Description>
                            {success}
                        </Alert.Description>
                    </Alert.Root>
                {/if}
            </div>
            <Button type="submit" disabled={isLoading}>
                {#if isLoading}
                    <Loader />
                {/if}
                Sign In with Email
            </Button>
        </div>
    </form>
    <div class="relative">
        <div class="absolute inset-0 flex items-center">
            <span class="w-full border-t" />
        </div>
    </div>
</div>
