<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import { likedPostsStore } from '@/stores';
  import type { Message } from '@/types';
  import { baseUrl } from '@/utils';
  import { Heart, HeartFilled } from 'radix-icons-svelte';
  import { onMount } from 'svelte';
  import { toast } from 'svelte-sonner';
  import { get } from 'svelte/store';

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
    fetch(`${baseUrl}/messages/${message.id}/${liked ? 'unlike' : 'like'}`, {
      method: 'PUT',
    }).then((response) => {
      if (response.status !== 200) {
        throw new Error('Something went wrong');
      }
    });

    // Toggle the liked variable.
    liked = !liked;

    // Show a toast message.
    toast.success(`Post ${liked ? 'liked' : 'unliked'}`);
  }

  function y() {
    // Up de y count
    yCount++;

    // Send a request to the server to like/unlike the post.
    // Note that we do not await the response, so we can show optimistic UI updates.
    fetch(`${baseUrl}/messages/${message.id}/y`, {
      method: 'PUT',
    });

    // Show a toast message.
    toast.success(`Post 𝕐'd`);
  }
</script>

<Card.Root>
  <!-- Card header containing the user name and post datetime -->
  <div class="flex justify-between flex-col gap-1 sm:flex-row">
    <div class="flex gap-2 items-center">
      <div class="flex items-center gap-2">
        <img
          src={`https://picsum.photos/seed/${message.username}/200`}
          alt="Avatar"
          class="h-8 aspect-square rounded-full"
        />
      </div>
      <Card.Title>
        {message.username}
      </Card.Title>
    </div>
    <Card.Description
      >{new Date(`${message.created_at} GMT`).toLocaleString('en', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric',
      })}</Card.Description
    >
  </div>

  <!-- Message content -->
  <Card.Content class="space-y-1">
    <span>
      {message.content}
    </span>
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
      <button class="flex gap-1 items-center" on:click={y}>
        <div>𝕐</div>
        {yCount}
      </button>
    </div>
  </Card.Content>
</Card.Root>
