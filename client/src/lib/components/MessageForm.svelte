<script>
    import  {Button} from "$lib/components/ui/button";
    import  {Textarea} from "$lib/components/ui/textarea";
    import  {Input} from "$lib/components/ui/input";
    import * as Card from "$lib/components/ui/card";
    import { PUBLIC_API_URL } from "$env/static/public";
    import { onMount } from "svelte";
    import { Check } from "radix-icons-svelte";

    let content = "";
    let username = "";
    let showUsernameInput = false;

    let usernameValue = username;

    onMount(() => {
        const localStorageUsername = localStorage.getItem('y_username');

        if (localStorageUsername === null) {
            showUsernameInput = true;
            return;
        }

        username = localStorageUsername;
    })

    async function post() {
        if (content.length === 0 || content.length > 255) return;

        await fetch(`${PUBLIC_API_URL}/messages`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                content: content,
                username: 'test',
            })
        })
    }
</script>

<Card.Root class='flex flex-col gap-2 justify-end'>
    {#if showUsernameInput}
        <div class="flex gap-2">
            <Input placeholder={username || 'What should you be called on Y?'}/>
            <Button variant="outline" size="icon">
                <Check  />
            </Button>
        </div>
        {:else} 
        <Card.Title>{username}</Card.Title>
    {/if}
    <Textarea bind:value={content} placeholder="What's happening?" cols={40} rows={3} class='resize-none border-none focus:border-none' maxlength={255}/>
    <Button on:click={post} disabled={content.length === 0} class='max-w-36 w-full ml-auto'>Plaatsen</Button>
</Card.Root>
