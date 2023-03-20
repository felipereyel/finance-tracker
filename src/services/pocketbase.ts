import PocketBase from "pocketbase";

const POCKETBASE_URL = import.meta.env.VITE_POCKETBASE_URL || "/";

export const pb = new PocketBase(POCKETBASE_URL);

export const initPocketBase = async () => pb.health.check();
