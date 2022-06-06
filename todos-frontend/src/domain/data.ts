export type TodoId = number;

export type Todo = Readonly<{
  id: TodoId;
  title: string;
  completed: boolean;
}>;
