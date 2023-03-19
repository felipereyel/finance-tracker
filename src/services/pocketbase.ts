import PocketBase from "pocketbase";

export const pb = new PocketBase("/");

export const initPocketBase = async () => pb.health.check();
