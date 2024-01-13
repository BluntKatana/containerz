<script lang="ts">
  import { Button } from '$lib/components/ui/button';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Input } from '$lib/components/ui/input';
  import * as Card from '$lib/components/ui/card';
  import { PUBLIC_API_URL } from '$env/static/public';
  import { onMount } from 'svelte';
  import { Pencil1 } from 'radix-icons-svelte';
  import Skeleton from './ui/skeleton/skeleton.svelte';
  import { messageStore } from '@/stores';
  import type { Message } from '@/types';
  import { toast } from 'svelte-sonner';
  import { Reload } from 'radix-icons-svelte';
  import * as Dialog from '$lib/components/ui/dialog';
  import { Label } from './ui/label';

  // Constant IDs to use later for determining what form is being submitted.
  const TEXT_AREA_ID = 'post_textarea';
  const USERNAME_ID = 'username_input';

  // Content of the message
  let content = '';

  let username = '';
  let showUsernameInput = false;
  let loadingUsername = true;
  let usernameValue = username;

  let buttonLoading = false;

  onMount(() => {
    // Get the user's Y username from the localStorage.
    const localStorageUsername = localStorage.getItem('y_username');

    // If the username is not defined yet, we want to show the username input.
    if (localStorageUsername === null) {
      showUsernameInput = true;
    }
    // If the username is defined, we want to set the username variable to the localStorage value.
    else {
      username = localStorageUsername;
      usernameValue = username;
    }

    // At this point we are done loading the username.
    loadingUsername = false;

    // On mount, create an eventListener for the keydown event.
    function handleKeyDown(event: KeyboardEvent) {
      // If the user presses enter and there is an active element, we want to prevent the default behaviour
      // and determine which form needs to be submitted.
      if (event.key === 'Enter' && document.activeElement) {
        event.preventDefault();

        if (document.activeElement.id === TEXT_AREA_ID) {
          post();
          return;
        }

        if (document.activeElement.id === USERNAME_ID) {
          handleUsernameSave();
          return;
        }
      }
    }

    // Add the eventListener to the document.
    document.addEventListener('keydown', handleKeyDown);

    // On unmount, remove the eventListener from the document.
    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  });

  async function post() {
    // If there is no content, or when it is too long, or when there is no username, we want to return.
    if (content.length === 0 || content.length > 255 || !username) return;

    // Set the buttonLoading variable to true, so we can show a loading spinner.
    buttonLoading = true;

    // Send a request to the server to post the message.
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
      // If the response is not 201, we want to throw an error.
      if (response.status !== 201) {
        throw new Error('Something went wrong');
      }

      // Parse the response as JSON (async operation)
      const newMessage: Message = await response.json();

      // Add the new message to the front of the messageStore.
      messageStore.update((messages) => {
        return [newMessage, ...messages];
      });
    });

    // Reset the buttonLoading and content variables, and show a toast.
    // To give the user the feeling that something is happening, we wait 250ms before doing this.
    setTimeout(() => {
      buttonLoading = false;
      content = '';
      toast.success('Message posted!', {description: "Thanks for sharing your thoughts!"});
    }, 250);
  }

  function handleUsernameSave() {
    // If the username is empty, we want to return.
    if (usernameValue.length === 0) return;

    // Set the username in the localStorage and set the username variable.
    localStorage.setItem('y_username', usernameValue);

    // Set the username variable and hide the username input.
    username = usernameValue;

    // Hide the username input.
    showUsernameInput = false;
  }
</script>

<!-- Dialog for the user to change their username -->
<Dialog.Root bind:open={showUsernameInput}>
  <Dialog.Content class="sm:max-w-[425px]">
    <Dialog.Header>
      <Dialog.Title>Edit username</Dialog.Title>
      <Dialog.Description>
        Make changes to your profile here. Click save when you're done.
      </Dialog.Description>
    </Dialog.Header>

    <div class="grid gap-4 py-4">
      <div class="grid grid-cols-4 items-center gap-4">
        <Label for={USERNAME_ID} class="text-right">Username</Label>
        <Input
          id={USERNAME_ID}
          bind:value={usernameValue}
          class="col-span-3"
          placeholder="elonmusk"
        />
      </div>
    </div>
    <Dialog.Footer>
      <Button type="submit" on:click={handleUsernameSave}>Save changes</Button>
    </Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>

<!-- Form to post a new message to Y. -->
<Card.Root class="flex flex-col gap-2 justify-end z-10">
  <div class="min-h-8 max-h-8 flex items-center justify-between">
    <!-- If we are busy fetching the username from localstorage, show a skeleton loading. -->
    {#if loadingUsername}
      <Skeleton class="w-32 h-[18px]" />
      <!-- If the username has loaded, render it in a card title + a button to open the change username dialog. -->
    {:else}
      <div class="flex gap-2 items-center">
        <div class="flex items-center gap-2">
          <img
            src={`https://picsum.photos/seed/${username}/200`}
            alt="Avatar"
            class="h-8 aspect-square rounded-full"
          />
        </div>
        <Card.Title>
          {username}
        </Card.Title>
      </div>
      <Button
        variant="ghost"
        size="icon"
        on:click={() => (showUsernameInput = true)}
        class="w-8 h-8"
      >
        <Pencil1 class="h-4 w-4" />
      </Button>
    {/if}
  </div>

  <!-- Textarea to write your post -->
  <Textarea
    id={TEXT_AREA_ID}
    bind:value={content}
    placeholder="Why not share your thoughts?"
    cols={40}
    rows={3}
    class="resize-none border-none focus:border-none"
    maxlength={255}
  />

  <!-- Button to post the message -->
  <div class="flex items-end justify-between">
    <span class="text-sm text-gray-500">
      {content.length}/255
    </span>
    <Button
      type="submit"
      on:click={post}
      disabled={content.length === 0 || !username || buttonLoading}
      class="min-w-24"
    >
      <!-- If the button is loading, we show a spinning icon. -->
      {#if buttonLoading}
        <Reload class="mr-2 h-4 w-4 animate-spin" />
      {/if}
      Post
    </Button>
  </div>
</Card.Root>
