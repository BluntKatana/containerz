<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import { HeartFilled, Heart } from 'radix-icons-svelte';
  import { PUBLIC_API_URL } from '$env/static/public';
  import { onMount } from 'svelte';
  import { likedPostsStore } from '@/stores';
  import { get } from 'svelte/store';
  import Separator from './ui/separator/separator.svelte';
  import type { Message } from '@/types';

  export let message: Message;

  // Store whether the post is liked or not.
  let liked = false;

  // Initialize the like and y counts based on the message prop.
  let likeCount = message.likes;
  let yCount = message.ys;

  onMount(() => {
    // When the component mounts, check if the current user has already liked it.
    liked = get(likedPostsStore).includes(message.id);
  });

  function like() {
    // If the post is liked, unlike it, otherwise like it.
    if (liked) {
      likeCount--;
    } else {
      likeCount++;
    }

    // Update the likedPostsStore value accordingly, depending on whether the post is liked or not.
    likedPostsStore.update((value) => {
      if (liked) {
        return value.filter((id) => id !== message.id);
      }

      return [...value, message.id];
    });

    // Send a request to the server to like/unlike the post.
    // Note that we do not await the response, so we can show optimistic UI updates.
    fetch(`${PUBLIC_API_URL}/messages/${message.id}/${liked ? 'unlike' : 'like'}`, {
      method: 'PUT',
    }).then((response) => {
      if (response.status !== 200) {
        throw new Error('Something went wrong');
      }
    });

    // Toggle the liked variable.
    liked = !liked;
  }

  function y() {
    // Up de y count
    yCount++;

    // Send a request to the server to like/unlike the post.
    // Note that we do not await the response, so we can show optimistic UI updates.
    fetch(`${PUBLIC_API_URL}/messages/${message.id}/y`, {
      method: 'PUT',
    });
  }
</script>

<Card.Root>
  <!-- Card header containing the user name and post datetime -->
  <div class="flex justify-between flex-col gap-1 sm:flex-row">
    <Card.Title>{message.username}</Card.Title>
    <Card.Description
      >{new Date(message.created_at).toLocaleString('en', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric',
      })}</Card.Description
    >
  </div>

  <!-- Message content -->
  <Card.Content>
    {message.content}
  </Card.Content>

  <!-- Footer of the card, containing buttons to (un)like or y the post. -->
  <div class="flex gap-4">
    <!-- Like/Unlike button, also showing the count -->
    <button on:click={like} class="flex gap-1 items-center">
      {#if liked}
        <HeartFilled />
      {:else}
        <Heart />
      {/if}

      {likeCount}
    </button>

    <!-- Y button, also showing the count. -->
    <button class="flex gap-1 items-center" on:click={y}
      ><div>𝕐</div>
      {yCount}</button
    >
  </div>
</Card.Root>
