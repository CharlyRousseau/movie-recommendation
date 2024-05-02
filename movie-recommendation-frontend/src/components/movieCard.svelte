<script lang="ts">
    import * as Card from "$lib/components/ui/card";

    export let title: string;
    export let date: string;
    export let imageSrc: string;
    export let grade: number;
    const baseimg = "https://image.tmdb.org/t/p/w500";
</script>

<Card.Root class="w-[200px] mt-4">
    <Card.Header>
        <Card.Title class="line-clamp-1">{title}</Card.Title>
        <Card.Description>{date}</Card.Description>
    </Card.Header>
    <Card.Content class="relative h-[300px] m-4">
        <img
            class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-full h-full object-cover rounded shadow"
            src={baseimg + imageSrc}
            alt={"image"}
        />
    </Card.Content>
    <Card.Footer>
        <p>{grade}</p>
    </Card.Footer>
</Card.Root>
