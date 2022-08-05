import { writable } from "svelte/store";

const darkMode = writable(localStorage.getItem("mode") || false);
