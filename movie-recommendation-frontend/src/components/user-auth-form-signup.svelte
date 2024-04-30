<script lang="ts">
    import Loader from "lucide-svelte/icons/loader";
    import { Button } from "$lib/components/ui/button/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import { cn } from "$lib/utils.js";
    import * as Alert from "$lib/components/ui/alert";
    import { navigate } from 'svelte-routing';

    const API_URL = import.meta.env.VITE_API_URL;

    let className: string | undefined | null = undefined;
    export { className as class };

    let email = "";
    let username = "";
    let password = "";
    let isLoading = false;
    let error = "";
    let success = ""; 

    async function onSubmit() {
        isLoading = true;
        error = "";

        try {
            const response = await fetch(`${API_URL}/users`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username,email, password }),
            });

            if (!response.ok) {
                const data = await response.json();
                error = data.message;
            }else{
                success = "User created successfully";
                      setTimeout(() => {
                        navigate('/login');
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

<div class={cn("grid gap-6", className)}>
    <form on:submit|preventDefault={onSubmit}>
        <div class="grid gap-2">
            <div class="grid gap-1">
                <Label class="sr-only" for="email">Email</Label>
                <Input
                    bind:value={email}
                    id="email"
                    placeholder="name@example.com"
                    type="email"
                    autocapitalize="none"
                    autocomplete="email"
                    autocorrect="off"
                    disabled={isLoading}
                />
                <Input
                    bind:value={username}
                    id="username"
                    placeholder="JohnDoe"
                    type="username"
                    autocapitalize="none"
                    autocorrect="off"
                    disabled={isLoading}
                />
                <Label class="sr-only" for="password">Password</Label>
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
                        <Alert.Title class="text-red-500" >Error</Alert.Title>
                            <Alert.Description>
                                {error}
                            </Alert.Description>
                    </Alert.Root>
                {/if}
                {#if success}
                    <Alert.Root>
                        <Alert.Title class="text-green-500">Success</Alert.Title>
                            <Alert.Description>
                            {success}
                            </Alert.Description>
                    </Alert.Root>
                {/if}
                <Button type="submit" disabled={isLoading}>
                    {#if isLoading}
                        <Loader />
                    {:else}
                        Sign Up
                    {/if}
                </Button>
            </div>
        </div>
    </form>
</div>
