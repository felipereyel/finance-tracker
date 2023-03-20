export const formatDate = (date: string) => {
  // date: YYY-MM-DD HH:MM:SS.SSSZ
  return date.split(" ")[0];
};

export const now = () => new Date().toISOString().split("T")[0];
