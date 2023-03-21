export const formatCurrencyBRL = (value: number) =>
  new Intl.NumberFormat("en-US", { style: "currency", currency: "BRL" }).format(
    value
  );
