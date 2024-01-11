<script lang='ts'>
    import type {Message} from '$lib/types';
    import * as Card from "$lib/components/ui/card";
    import { HeartFilled, Heart } from 'radix-icons-svelte';
    import { PUBLIC_API_URL } from '$env/static/public';
    import { onMount } from 'svelte';
    import { likedPostsStore } from '@/stores';
    import { get } from 'svelte/store';
    import Separator from './ui/separator/separator.svelte';

    export let message: Message;

    let liked = false;
    let yed = false;

    onMount(() => {
        liked = get(likedPostsStore).includes(message.id);
    })

    function like() {
        if (liked) {
            likeCount--;
        } else {
            likeCount++;
        }

        likedPostsStore.update((value) => {
            if (liked) {
                return value.filter((id) => id !== message.id);
            }

            return [...value, message.id];
        })

        fetch(`${PUBLIC_API_URL}/messages/${message.id}/${liked ? 'unlike' : 'like'}`, {
            method: 'PUT',
        }).then((response) => {
            if (response.status !== 200) {
                throw new Error('Something went wrong');
            }
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
    <div class="flex justify-between flex-col gap-1 sm:flex-row">
        <Card.Title>{message.username}</Card.Title>
        <Card.Description>{new Date(message.created_at).toLocaleString('en', {
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