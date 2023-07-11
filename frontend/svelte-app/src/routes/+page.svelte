<script lang="ts">
  import {
    Theme,
    RadioButtonGroup,
    RadioButton,
  } from "carbon-components-svelte";

	import type {
    CarbonTheme
  } from "carbon-components-svelte/types/Theme/Theme.svelte";

  let theme: CarbonTheme = "g90";
</script>

<Theme bind:theme />

<RadioButtonGroup legendText="Carbon theme" bind:selected={theme}>
  {#each ["white", "g10", "g80", "g90", "g100"] as value}
    <RadioButton labelText={value} {value} />
  {/each}
</RadioButtonGroup>

<h1>Welcome to Rickenbacker !!</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>