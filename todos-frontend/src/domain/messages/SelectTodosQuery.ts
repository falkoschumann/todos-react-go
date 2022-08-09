import { Todo } from '../data/Todo';

export type SelectTodosQuery = Readonly<{}>;

export type SelectTodosQueryResult = Readonly<{
  todos: readonly Todo[];
}>;
