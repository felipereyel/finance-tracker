import { computed, ref } from "@vue/reactivity";
import { watch, type Ref } from "vue";

export const asyncComputed = <T>(asyncFunc: () => Promise<T>) => {
  const loading = ref(true);
  const result: Ref<T | null> = ref(null);
  const funcPromise = computed(() => asyncFunc());

  const callBacks: Array<() => void | Promise<void>> = [];
  const onResult = (cb: () => void | Promise<void>) => {
    callBacks.push(cb);
  };

  const resolve = (res: T) => {
    result.value = res;
    loading.value = false;
    callBacks.forEach(cb => cb());
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

  return { result, loading, onResult };
};
