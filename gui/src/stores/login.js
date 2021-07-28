import { writable } from 'svelte/store';

// TODO: store values in localStorage
export const login = writable({
  url: "",
  username: "",
  password: "",
});
