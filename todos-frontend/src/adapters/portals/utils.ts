export enum Filter {
  All = 'all',
  Active = 'active',
  Completed = 'completed',
}

export function pluralize(count: number, word: string) {
  return count > 1 ? `${word}s` : word;
}
