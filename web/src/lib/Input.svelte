<script lang="ts">
  export let input_value: string;
  export let pholder: string;
  export let required: boolean;
</script>

<input
  type="text"
  name="value"
  id="value"
  bind:value={input_value}
  placeholder={pholder}
  required={required && required}
/>

<style lang="postcss">
  input[type="text"] {
    @apply mt-2 rounded-md px-3 py-2 border-none outline-none ring ring-secondary/80 hover:ring-2 hover:ring-accent;
    @apply font-light;
  }
</style>
