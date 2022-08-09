import { TodoId } from '../data/Todo';

export type DestroyTodoCommand = Readonly<{
  id: TodoId;
}>;
