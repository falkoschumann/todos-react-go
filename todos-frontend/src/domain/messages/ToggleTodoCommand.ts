import { TodoId } from '../data/Todo';

export type ToggleTodoCommand = Readonly<{
  id: TodoId;
}>;
