<script lang="ts">
  import '../app.pcss';
  import { ModeWatcher } from 'mode-watcher';
  import { Toaster } from '$lib/components/ui/sonner';
  import gopherSvg from '$lib/assets/go-gopher.svg';
</script>

<ModeWatcher />
<Toaster />
<slot />
<!-- SVG of Gopher bottom left -->
<img alt="gopher golang icon" src={gopherSvg} class="fixed bottom-0 left-0 w-16 h-16" />
