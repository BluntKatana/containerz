<script lang="ts">
  import { onMount } from 'svelte';
  import type { Message } from '$lib/types';
  import MessageCard from '$lib/components/MessageCard.svelte';
  import { PUBLIC_API_URL } from '$env/static/public';
  import MessageForm from '@/components/MessageForm.svelte';
  import Separator from '@/components/ui/separator/separator.svelte';
  import { likedPostsStore, messageStore } from '@/stores';
  import { Button } from '@/components/ui/button';
  import { Moon, Reload, Sun } from 'radix-icons-svelte';
  import { toggleMode } from 'mode-watcher';

  // Store all the messages
  let messages: Message[] = [];

  // Whether the messages are being fetched
  let messagesLoading = true;

  onMount(async () => {
    // Get the messages from the server.
    messages = await fetch(`${PUBLIC_API_URL}/messages`, {
      headers: {
        accept: 'application/json',
        'User-agent': 'learning app',
        'Content-Type': 'application/json',
      },
    })
      // Parse the response as JSON (async operation)
      .then((response) => response.json())
      // Init the messageStore value and set the messagesLoading flag to false.
      .then(async (data) => {
        messageStore.set(data);

        await new Promise((resolve) => setTimeout(resolve, 500));

        messagesLoading = false;
        return data;
      });

    // Subscribe to the messageStore to keep the messages variable up to date.
    messageStore.subscribe((value) => {
      messages = value;
    });

    // Get the liked posts from the localStorage and init the likedPostsStore value.
    const localStorageLikedPosts = localStorage.getItem('y_liked_posts');

    likedPostsStore.set(localStorageLikedPosts ? JSON.parse(localStorageLikedPosts) : []);

    // When one of the posts gets liked, update the localStorage.
    likedPostsStore.subscribe((value) => {
      localStorage.setItem('y_liked_posts', JSON.stringify(value));
    });
  });
</script>

<svelte:head>
  <title>Home / Y</title>
</svelte:head>

<div class="flex items-center flex-col h-screen w-full">
  <!-- Header containing Y logo and dark/light mode toggle -->
  <nav class="fixed w-full flex items-center justify-between px-4 top-4">
    <span class="text-4xl">𝕐</span>
    <Button on:click={toggleMode} variant="outline" size="icon">
      <Sun
        class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
      />
      <Moon
        class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
      />
      <span class="sr-only">Toggle theme</span>
    </Button>
  </nav>
  <!-- Main content -->
  <main class="max-w-[600px] w-screen flex flex-col gap-4 mx-auto flex-grow py-4">
    <MessageForm />
    <!-- Show a loading spinner when the messages are being fetched. -->
    {#if messagesLoading}
      <div class="w-full flex-grow flex items-center flex-col">
        <Reload class="mr-2 h-4 w-4 animate-spin" />
        <span>Loading messages...</span>
      </div>
      <!-- Render a list of messages when they are defined and the list isnt empty. -->
    {:else if messages && messages.length > 0}
      <Separator />
      <div class="flex flex-col gap-4">
        {#each messages as message (message.id)}
          <MessageCard {message} />
        {/each}
      </div>
      <!-- If there are no messages, show a message to the user. -->
    {:else}
      <div class="w-full flex-grow flex items-center flex-col">
        <span>No messages yet. Be the first to post one!</span>
      </div>
    {/if}
  </main>
</div>
