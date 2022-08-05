<script lang="ts">
  import Card from "./lib/Card.svelte";
  import Input from "./lib/Input.svelte";
  import Navbar from "./lib/Navbar.svelte";
  const APP_URL = location.origin;
  let username = "";
  let customLabel = "";

  const buildURL = (cLabel: string, uname: string, cColor: string) => {
    return `${APP_URL}/profile/${uname}?${
      cLabel.length > 0 ? `label=${cLabel}&` : ``
    }${cColor.length > 0 ? `color=${cColor}` : ``}`;
  };

  $: results = [
    {
      content: "Default label, no color.",
      url: buildURL("", username, ""),
      color: "",
    },
    {
      content: "Label + yellow template.",
      url: buildURL(customLabel, username, "yellow"),
      color: "yellow",
    },
    {
      content: "Label + green template.",
      url: buildURL(customLabel, username, "green"),
      color: "green",
    },
  ];
</script>

<Navbar />
<main>
  <div class="head">
    <h1 class="text-gray-400">Github Status for your README.md</h1>
    <Input
      required={true}
      bind:input_value={username}
      pholder="Your Github username"
    />
  </div>
  <div class="card_list">
    <Card {results}>
      <Input
        required={false}
        bind:input_value={customLabel}
        pholder="If you want a custom label"
      />
    </Card>
  </div>
</main>

<style lang="postcss">
  main {
    @apply max-w-4xl m-auto px-2 py-4;
  }

  main > .head {
    @apply text-primary text-center;
  }

  .card_list {
    @apply mt-8;
  }
</style>
