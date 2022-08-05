<script lang="ts">
  // import { buildResult } from "./utils";

  let isOpen = false;

  type Result = {
    content: string;
    url: string;
    color: string;
  };
  export let results: Result[];

  function toggleCard() {
    isOpen = !isOpen;
  }

  function copyToClipboard(val: string) {
    navigator.clipboard.writeText(val);
  }
</script>

<div class="card">
  <div class="card_head">
    <h3>Profile Visits</h3>
    <span on:click={toggleCard}>
      {#if !isOpen}
        &plus;
      {:else}
        &minus;
      {/if}
    </span>
  </div>
  <div class:show={!isOpen} class="card_body">
    <slot />
    <div class="result_container">
      <ul class="results">
        {#each results as result}
          <li class="result">
            {result.content}
            <div class="actions">
              <span on:click={() => copyToClipboard(result.url)}>&boxbox;</span>
              <a
                target="_blank"
                referrerpolicy="noreferrer"
                rel="external"
                href={result.url}>&rightarrow;</a
              >
            </div>

            <!-- {@html buildResult({ label: result.content, color: result.color })} -->
          </li>
        {/each}
      </ul>
    </div>
  </div>
</div>

<style lang="postcss">
  .card {
    @apply w-full max-w-md;
  }

  .card_body {
    @apply bg-primary/50 p-2;
  }

  .card_head {
    @apply flex items-center justify-between;
    @apply ring-2 font-bold text-xl text-white bg-primary px-2 py-3 rounded-sm;
  }

  .card_head > span {
    @apply border rounded-sm w-8 text-center;
  }

  .show {
    @apply hidden;
  }

  .results {
    @apply mt-4;
  }

  .result {
    @apply flex items-center justify-between text-xs text-white m-2 px-2 py-3 bg-primary;
  }

  .result > .actions {
    @apply gap-2 text-2xl ml-2;
  }

  .actions > * {
    @apply text-accent hover:text-white;
  }
</style>
