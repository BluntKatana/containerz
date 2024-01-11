import { writable, type Writable } from "svelte/store";
import type { Message } from "./types";

export const messageStore: Writable<Message[]> = writable([]);