<script lang='ts'>
    import type {Message} from '$lib/types';
    import * as Card from "$lib/components/ui/card";
    import { HeartFilled, Heart } from 'radix-icons-svelte';
    import { PUBLIC_API_URL } from '$env/static/public';

    export let message: Message;

    let liked = false;
    let yed = false;

    function like() {
        if (liked) {
            likeCount--;
        } else {
            likeCount++;
        }

        fetch(`${PUBLIC_API_URL}/messages/${message.id}/${liked ? 'unlike' : 'like'}`, {
            method: 'PUT',
        })

        liked = !liked;
    }

    function y() {
        yCount++;

        fetch(`${PUBLIC_API_URL}/messages/${message.id}/y`, {
            method: 'PUT',
        })

        yed = !yed;
    }

    let likeCount = message.likes;
    let yCount = message.ys;
</script>

<Card.Root>
    <div class="flex justify-between">
        <Card.Title>{message.username}</Card.Title>
        <Card.Description>{new Date(message.created_at).toLocaleString('nl-NL', {
            weekday: 'long',
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric'
        })}</Card.Description>
    </div>
    <Card.Content>
        {message.content}
    </Card.Content>
    <div class="flex gap-4">
        <button on:click={like} class="flex gap-1 items-center">
            {#if liked}
                <HeartFilled/>
            {:else}
                <Heart/>
            {/if}

            {likeCount}
        </button>
        <button class="flex gap-1 items-center" on:click={y}><div>𝕐</div> {yCount}</button>
    </div>
</Card.Root>