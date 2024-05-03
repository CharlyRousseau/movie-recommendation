<script>
    import * as Alert from "$shadcn/ui/alert";
    export let type = "";
    export let message = "";
</script>

<Alert.Root>
    <Alert.Title class={type === "error" ? "text-red-500" : "text-green-500"}>
        {type.charAt(0).toUpperCase() + type.slice(1)}
    </Alert.Title>
    <Alert.Description>
        {message}
    </Alert.Description>
</Alert.Root>
