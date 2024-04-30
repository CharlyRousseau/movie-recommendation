<script lang="ts">
    import * as Card from "$lib/components/ui/card";

    export let title: string;
    export let date: string;
    export let imageSrc: string;
    export let grade: string;
</script>

<Card.Root>
    <Card.Header>
        <Card.Title>{title}</Card.Title>
        <Card.Description>{date}</Card.Description>
    </Card.Header>
    <Card.Content>
        <img src={imageSrc} alt={"image"} />
    </Card.Content>
    <Card.Footer>
        <p>{grade}</p>
    </Card.Footer>
</Card.Root>
