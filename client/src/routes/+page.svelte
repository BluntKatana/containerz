<script lang='ts'>
    import { onMount } from "svelte";
    import type { Message } from "$lib/types";
    import MessageCard from "$lib/components/MessageCard.svelte";
    import { PUBLIC_API_URL } from '$env/static/public';
    import MessageForm from "@/components/MessageForm.svelte";
    import Separator from "@/components/ui/separator/separator.svelte";
    import { likedPostsStore, messageStore } from "@/stores";
    import { Button } from "@/components/ui/button";
    import { Moon, Sun } from "radix-icons-svelte";
    import { toggleMode } from "mode-watcher";

    let messages: Message[] = [];
    let messagesLoading = true;

  onMount(async () => {
    messages = await fetch(`${PUBLIC_API_URL}/messages`, {
      headers:{
      accept: 'application/json',
      'User-agent': 'learning app',
      'Content-Type': 'application/json',
    }
    })
    .then(response=>(response.json())).then(async (data) => {
        messageStore.set(data); 

        await new Promise((resolve) => setTimeout(resolve, 500));

        // messagesLoading = false;
        return data;
    });

    messageStore.subscribe((value) => {
        messages = value;
    })

    const localStorageLikedPosts = localStorage.getItem('y_liked_posts');

    likedPostsStore.set(localStorageLikedPosts ? JSON.parse(localStorageLikedPosts) : []);

    likedPostsStore.subscribe((value) => {
        localStorage.setItem('y_liked_posts', JSON.stringify(value));
    })
  })

</script>

<div class="flex justify-center flex-col">
    <div class="w-full flex items-center justify-between px-4 py-2">
        <span class="text-4xl">
            𝕐
        </span>
        <Button on:click={toggleMode} variant="outline" size="icon">
            <Sun
              class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
            />
            <Moon
              class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
            />
            <span class="sr-only">Toggle theme</span>
          </Button>

    </div>
    <div class="max-w-[600px] w-screen flex flex-col gap-4 mx-auto">
        <MessageForm />
        {#if messagesLoading}
        <div class="w-full flex-grow flex items-center justify-center">Loading</div>
        {:else}
            <Separator />
            <div class="flex flex-col gap-4">
            {#each messages as message (message.id)}
                <MessageCard {message} />
            {/each}
            </div>
        {/if}
        </div>
</div>