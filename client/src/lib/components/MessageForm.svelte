<script lang="ts">
  import { Button } from '$lib/components/ui/button';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Input } from '$lib/components/ui/input';
  import * as Card from '$lib/components/ui/card';
  import { PUBLIC_API_URL } from '$env/static/public';
  import { onMount } from 'svelte';
  import { Check } from 'radix-icons-svelte';
  import Skeleton from './ui/skeleton/skeleton.svelte';
  import { messageStore } from '@/stores';
  import type { Message } from '@/types';
  import { toast } from 'svelte-sonner';
  import { Reload } from 'radix-icons-svelte';

  let content = '';
  let username = '';
  let showUsernameInput = false;

  let loadingUsername = true;

  let usernameValue = username;

  let buttonLoading = false;

  onMount(() => {
    const localStorageUsername = localStorage.getItem('y_username');

    loadingUsername = false;

    if (localStorageUsername === null) {
      showUsernameInput = true;
      return;
    }

    username = localStorageUsername;

    function handleKeyDown(event: KeyboardEvent) {
      if (event.key === 'Enter') {
        event.preventDefault();
        post();
      }
    }

    document.addEventListener('keydown', handleKeyDown);

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  });

  async function post() {
    if (content.length === 0 || content.length > 255 || !username) return;

    buttonLoading = true;
    await fetch(`${PUBLIC_API_URL}/messages`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        content,
        username,
      }),
    }).then(async (response) => {
      if (response.status !== 201) {
        throw new Error('Something went wrong');
      }

      const newMessage: Message = await response.json();

      messageStore.update((messages) => {
        return [newMessage, ...messages];
      });
    });

    setTimeout(() => {
      buttonLoading = false;
      content = '';
      toast('Message posted!');
    }, 500);
  }

  function handleUsernameSave() {
    if (usernameValue.length === 0) return;

    localStorage.setItem('y_username', usernameValue);

    username = usernameValue;
    showUsernameInput = false;
  }
</script>

<Card.Root class="flex flex-col gap-2 justify-end">
  {#if loadingUsername}
    <Skeleton class="w-32 h-5" />
  {:else if showUsernameInput}
    <div class="flex gap-2">
      <Input
        bind:value={usernameValue}
        placeholder={username || 'What should you be called on Y?'}
      />
      <Button size="icon" disabled={!usernameValue} on:click={handleUsernameSave}>
        <Check />
      </Button>
    </div>
  {:else}
    <Card.Title>{username}</Card.Title>
  {/if}
  <Textarea
    bind:value={content}
    placeholder="What's happening?"
    cols={40}
    rows={3}
    class="resize-none border-none focus:border-none"
    maxlength={255}
  />
  <Button
    type="submit"
    on:click={post}
    disabled={content.length === 0 || !username || buttonLoading}
    class="max-w-36 w-full ml-auto"
  >
    {#if buttonLoading}
      <Reload class="mr-2 h-4 w-4 animate-spin" />
    {/if}
    Plaatsen
  </Button>
</Card.Root>
