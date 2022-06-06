import { Todo, TodoId } from "./data";

export type CommandStatus = Readonly<{
  success: boolean;
  errorMessage?: string;
}>;

export type AddTodoCommand = Readonly<{
  title: string;
}>;

export type ClearCompletedCommand = Readonly<{}>;

export type DestroyTodoCommand = Readonly<{
  id: TodoId;
}>;

export type SaveTodoCommand = Readonly<{
  id: TodoId;
  newTitle: string;
}>;

export type ToggleAllCommand = Readonly<{
  checked: boolean;
}>;

export type ToggleTodoCommand = Readonly<{
  id: TodoId;
}>;

export type SelectTodosQuery = Readonly<{}>;

export type SelectTodosQueryResult = Readonly<{
  todos: readonly Todo[];
}>;
