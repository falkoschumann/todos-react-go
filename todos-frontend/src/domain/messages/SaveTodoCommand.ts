import { TodoId } from '../data/Todo';

export type SaveTodoCommand = Readonly<{
  id: TodoId;
  newTitle: string;
}>;
