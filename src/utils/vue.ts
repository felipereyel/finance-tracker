import { computed, ref } from "@vue/reactivity";
import { watch, type Ref } from "vue";

export const asyncComputed = <T>(
  asyncFunc: () => Promise<T>
): { result: Ref<T | null>; loading: Ref<boolean> } => {
  const loading = ref(true);
  const result: Ref<T | null> = ref(null);
  const funcPromise = computed(() => asyncFunc());

  const resolve = (res: T) => {
    result.value = res;
    loading.value = false;
  };

  watch(
    () => funcPromise.value,
    () => {
      result.value = null;
      loading.value = true;
      funcPromise.value.then(resolve);
    }
  );

  funcPromise.value.then(resolve);

  return { result, loading };
};
