import { TodoId } from '../data/Todo';

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
