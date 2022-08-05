<script>
  import githubIcon from "../assets/github-svgrepo-com.svg";
  const GITHUB_URL = "https://github.com/mystique09/gh-stats";
</script>

<div class="wrap">
  <nav>
    <h2><a href="/">{"<Github Status />"}</a></h2>
    <ul>
      <li>
        <a class="github_img" href={GITHUB_URL}>
          <img
            rel="opener"
            target="_blank"
            referrerpolicy="noreferrer"
            class="github_icon"
            src={githubIcon}
            alt="Github Icon"
          />
          <span>Star the github.</span>
        </a>
      </li>
    </ul>
  </nav>
</div>

<style lang="postcss">
  .wrap {
    @apply w-full shadow-md bg-gray/10;
  }
  nav {
    @apply p-2 h-16 m-auto max-w-4xl flex items-center justify-between;
  }

  .github_icon {
    @apply h-8;
  }

  h2 {
    @apply text-secondary/80;
  }

  a.github_img {
    @apply flex flex-col items-center text-xs;
  }
  a.github_img > span {
    @apply text-gray/70;
  }
</style>
