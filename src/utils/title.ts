import { RouteLocationNormalizedLoaded } from "vue-router";

export const setTitle = (router: RouteLocationNormalizedLoaded, title?: string) => {
  let docTitle = `${router.meta.title} | Finance Tracker`;
  if (title) docTitle = `${title} | ${docTitle}`;
  window.document.title = docTitle;
}