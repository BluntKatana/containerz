<script lang='ts'>
    import { onMount } from "svelte";
    import type { Message } from "$lib/types";
    import MessageCard from "$lib/components/MessageCard.svelte";
    import { PUBLIC_API_URL } from '$env/static/public';

    let promise: Promise<Message[]>;

    let messages: Message[] = [];


  onMount(async () => {
    promise = await fetch(`${PUBLIC_API_URL}/messages`, {
      headers:{
      accept: 'application/json',
      'User-agent': 'learning app',
      'Content-Type': 'application/json',
    }
    })
    .then(response=>(response.json()))

    messages = await promise;
  })

</script>

<div class="flex justify-center">
        <div class="max-w-[600px] w-screen flex flex-col gap-4">
            {#each messages as message (message.id)}
                <MessageCard {message} />
            {/each}
        </div>
</div>